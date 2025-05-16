package monitors

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Michaelpalacce/uptime-bar/pkgs/status"
)

type HttpMonitor struct {
	HttpStatus status.HttpStatus
}

// Watch will watch the configured HttpStatus.
// Make sure to call this in a goroutine
func (m *HttpMonitor) Watch(statusChan chan Change) {
	t := time.NewTicker(10 * time.Second)

	for {
		<-t.C

		// @TODO: Save the status too
		if _, err := m.checkURLStatus(); err != nil {
			statusChan <- Change{
				Status: status.STATE_DOWN,
				Reason: err.Error(),
			}
		}

		statusChan <- Change{
			Status: status.STATE_UP,
			Reason: "",
		}
	}
}

func (m *HttpMonitor) checkURLStatus() (int, error) {
	// Create a custom HTTP client with a timeout
	client := &http.Client{
		Timeout: 5 * time.Second, // 5-second timeout for the request
	}

	// Send an HTTP GET request
	resp, err := client.Get(m.HttpStatus.Address)
	if err != nil {
		return -1, fmt.Errorf("failed to check %s: %w", m.HttpStatus.Address, err)
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}
