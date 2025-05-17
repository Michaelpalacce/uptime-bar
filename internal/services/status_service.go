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

	go service.StartStatusWatch()

	return service
}

func (s *StatusService) GetStatusForAll() {
}

// StartStatusWatch is meant to be executed in a goroutine
// It will poll the configured services and get their state
// @TODO: On change, update the store
func (s *StatusService) StartStatusWatch() {
	for _, conf := range s.Configuration.HttpStatuses {
		go func() {
			monitor := monitors.HttpMonitor{
				HttpStatus: conf,
			}

			changeChan := make(chan bool)
			go monitor.Watch(changeChan)

			for {
				<-changeChan

				slog.Info("Change detected", "name", conf.Name, "address", conf.Address, "state", conf.State, "reason", conf.Reason)
			}
		}()
	}
}
