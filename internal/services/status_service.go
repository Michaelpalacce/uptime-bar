package services

import (
	"log/slog"

	"github.com/Michaelpalacce/uptime-bar/internal/configuration"
	"github.com/Michaelpalacce/uptime-bar/pkgs/monitors"
	"github.com/Michaelpalacce/uptime-bar/pkgs/status"
)

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

// GetStatusForAll will return up and down counter
// While there is a race condition of up and down status being misreported, because in a split second we may
// get the old status, this is not a worry due to the nature of our program.
func (s *StatusService) GetStatusForAll() GetAllStatusResponseBody {
	body := GetAllStatusResponseBody{}
	for _, conf := range s.Configuration.HttpStatuses {
		if conf.State == status.STATE_UP {
			body.Up++
		}

		if conf.State == status.STATE_DOWN {
			body.Down++
		}
	}

	return body
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
				isOperational := <-changeChan
				if !isOperational {
					break
				}

				slog.Info("Change detected", "name", conf.Name, "address", conf.Address, "state", conf.State, "reason", conf.Reason)
			}
		}()
	}
}
