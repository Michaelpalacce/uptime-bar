package services

import (
	"log/slog"

	"github.com/Michaelpalacce/uptime-bar/internal/configuration"
	"github.com/Michaelpalacce/uptime-bar/pkgs/monitors"
)

type Storable interface {
	Store() any
}

type Retrievable interface{}

type StorableRetrievable interface {
	Storable
	Retrievable
}

// StatusService is responsible for polling the status of the different services
type StatusService struct {
	Configuration *configuration.Configuration
}

func NewStatusService(configuration *configuration.Configuration) *StatusService {
	service := &StatusService{
		Configuration: configuration,
	}

	go service.WatchStatus()

	return service
}

func (s *StatusService) GetStatusForAll() {
}

// WatchStatus is meant to be executed in a goroutine
// It will poll the configured services and get their state
func (s *StatusService) WatchStatus() {
	for _, conf := range s.Configuration.HttpStatuses {
		go func() {
			monitor := monitors.HttpMonitor{
				HttpStatus: conf,
			}

			statusChan := make(chan monitors.Change)

			go monitor.Watch(statusChan)

			for {
				change := <-statusChan
				slog.Info("Change detected", "address", conf.Address, "state", change.Status, "reason", change.Reason)
				conf.State = change.Status
				conf.Reason = change.Reason
			}
		}()
	}
}
