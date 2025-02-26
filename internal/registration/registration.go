package registration

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/project-flotta/flotta-device-worker/internal/metrics"

	"git.sr.ht/~spc/go-log"
	"github.com/google/uuid"
	"github.com/hashicorp/go-multierror"
	"github.com/project-flotta/flotta-device-worker/internal/configuration"
	"github.com/project-flotta/flotta-device-worker/internal/datatransfer"
	hardware2 "github.com/project-flotta/flotta-device-worker/internal/hardware"
	"github.com/project-flotta/flotta-device-worker/internal/heartbeat"
	os2 "github.com/project-flotta/flotta-device-worker/internal/os"
	"github.com/project-flotta/flotta-device-worker/internal/workload"
	"github.com/project-flotta/flotta-operator/models"
	pb "github.com/redhatinsights/yggdrasil/protocol"
)

const (
	retryAfter = 10
)

type Registration struct {
	hardware         *hardware2.Hardware
	os               *os2.OS
	dispatcherClient pb.DispatcherClient
	config           *configuration.Manager
	RetryAfter       int64
	heartbeat        *heartbeat.Heartbeat
	workloads        *workload.WorkloadManager
	monitor          *datatransfer.Monitor
	registered       bool
	metricsStorage   metrics.API
	systemMetrics    *metrics.SystemMetrics
	lock             sync.RWMutex
}

func NewRegistration(hardware *hardware2.Hardware, os *os2.OS, dispatcherClient DispatcherClient,
	config *configuration.Manager, heartbeatManager *heartbeat.Heartbeat, workloadsManager *workload.WorkloadManager,
	monitorManager *datatransfer.Monitor, metricsStorage metrics.API, systemMetrics *metrics.SystemMetrics) *Registration {
	return &Registration{
		hardware:         hardware,
		os:               os,
		dispatcherClient: dispatcherClient,
		config:           config,
		RetryAfter:       retryAfter,
		heartbeat:        heartbeatManager,
		workloads:        workloadsManager,
		monitor:          monitorManager,
		metricsStorage:   metricsStorage,
		systemMetrics:    systemMetrics,
		lock:             sync.RWMutex{},
	}
}

func (r *Registration) RegisterDevice() {
	err := r.registerDeviceOnce()

	if err != nil {
		log.Errorf("cannot register device. DeviceID: %s; err: %v", r.workloads.GetDeviceID(), err)
	}

	go r.registerDeviceWithRetries(r.RetryAfter)
}

func (r *Registration) registerDeviceWithRetries(interval int64) {
	ticker := time.NewTicker(time.Second * time.Duration(interval))
	for range ticker.C {
		if !r.config.IsInitialConfig() {
			ticker.Stop()
			break
		}
		log.Infof("configuration has not been initialized yet. Sending registration request. DeviceID: %s;", r.workloads.GetDeviceID())
		err := r.registerDeviceOnce()
		if err != nil {
			log.Errorf("cannot register device. DeviceID: %s; err: %v", r.workloads.GetDeviceID(), err)
		}
	}
}

func (r *Registration) registerDeviceOnce() error {
	hardwareInformation, err := r.hardware.GetHardwareInformation()
	if err != nil {
		return err
	}
	registrationInfo := models.RegistrationInfo{
		Hardware: hardwareInformation,
	}
	content, err := json.Marshal(registrationInfo)
	if err != nil {
		return err
	}
	data := &pb.Data{
		MessageId: uuid.New().String(),
		Content:   content,
		Directive: "registration",
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Call "Send"
	if _, err := r.dispatcherClient.Send(ctx, data); err != nil {
		return err
	}
	r.lock.Lock()
	r.registered = true
	r.lock.Unlock()
	return nil
}

func (r *Registration) IsRegistered() bool {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return r.registered
}

func (r *Registration) Deregister() error {

	var errors error
	err := r.workloads.Deregister()
	if err != nil {
		errors = multierror.Append(errors, fmt.Errorf("failed to deregister workloads: %v", err))
		log.Errorf("failed to deregister workloads. DeviceID: %s; err: %v", r.workloads.GetDeviceID(), err)
	}

	err = r.config.Deregister()
	if err != nil {
		errors = multierror.Append(errors, fmt.Errorf("failed to deregister configuration: %v", err))
		log.Errorf("failed to deregister configuration. DeviceID: %s; err: %v", r.workloads.GetDeviceID(), err)
	}

	err = r.heartbeat.Deregister()
	if err != nil {
		errors = multierror.Append(errors, fmt.Errorf("failed to deregister heartbeat: %v", err))
		log.Errorf("failed to deregister heartbeat. DeviceID: %s; err: %v", r.workloads.GetDeviceID(), err)
	}

	err = r.monitor.Deregister()
	if err != nil {
		errors = multierror.Append(errors, fmt.Errorf("failed to deregister monitor: %v", err))
		log.Errorf("failed to deregister monitor. DeviceID: %s; err: %v", r.workloads.GetDeviceID(), err)
	}

	err = r.systemMetrics.Deregister()
	if err != nil {
		errors = multierror.Append(errors, fmt.Errorf("failed to deregister system metrics: %v", err))
		log.Errorf("failed to deregister system metrics. DeviceID: %s; err: %v", r.workloads.GetDeviceID(), err)
	}

	err = r.metricsStorage.Deregister()
	if err != nil {
		errors = multierror.Append(errors, fmt.Errorf("failed to deregister metrics storage: %v", err))
		log.Errorf("failed to deregister metrics storage. DeviceID: %s; err: %v", r.workloads.GetDeviceID(), err)
	}

	r.registered = false
	return errors
}
