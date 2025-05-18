package monitors

import (
	"fmt"
	"log/slog"
	"net/http"
	"slices"
	"time"

	"github.com/Michaelpalacce/uptime-bar/pkgs/status"
)

type HttpMonitor struct {
	HttpStatus *status.HttpStatus
}

// Watch will watch the configured HttpStatus.
// Make sure to call this in a goroutine
// The chan will notify if there is an update. if false is given back, then the watch failed
func (m *HttpMonitor) Watch(statusChan chan bool) {
	if m.HttpStatus == nil {
		slog.Error("You created a monitor that does not have a provied HttpStatus")
		statusChan <- false
		return
	}

	slog.Info("Starting a new http monitor", "name", m.HttpStatus.Name, "address", m.HttpStatus.Address, "interval", m.HttpStatus.Interval)
	t := time.NewTicker(m.HttpStatus.Interval)

	for {
		<-t.C
		var (
			statusCode int
			err        error
		)

		if statusCode, err = m.checkURLStatus(); err != nil {
			if update, _ := m.HttpStatus.Patch(
				status.SetToDown,
				status.SetReason(fmt.Errorf("%s returned %d. Err was: %w", m.HttpStatus.Address, statusCode, err).Error()),
			); update {
				statusChan <- true
			}

			continue
		}

		if !slices.Contains(m.HttpStatus.ExpectedStatusCodes, statusCode) {
			if update, _ := m.HttpStatus.Patch(
				status.SetToDown,
				status.SetReason(fmt.Errorf("%s returned %d, which was not expected", m.HttpStatus.Address, statusCode).Error()),
			); update {
				statusChan <- true
			}

			continue
		}

		if upddate, _ := m.HttpStatus.Patch(status.SetToUp, status.SetReason(fmt.Sprintf("Status is: %d", statusCode))); upddate {
			statusChan <- true
		}
	}
}

func (m *HttpMonitor) checkURLStatus() (int, error) {
	// Create a custom HTTP client with a timeout
	client := http.Client{Timeout: m.HttpStatus.GetTimeout()}

	// Send an HTTP GET request
	resp, err := client.Get(m.HttpStatus.Address)
	if err != nil {
		return -1, fmt.Errorf("failed to check %s: %w", m.HttpStatus.Address, err)
	}

	return resp.StatusCode, nil
}
