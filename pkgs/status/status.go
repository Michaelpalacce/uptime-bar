package status

import "time"

const (
	STATE_UNKNOWN = iota
	STATE_UP
	STATE_DOWN
)

type Status struct {
	State  int
	Reason string
	Name   string `mapstructure:"name"`
}

type PatchFunc func(h *Status) (bool, error)

// Patch can be used to update the state of the Status struct.
// Use provided helper methods
func (h *Status) Patch(patches ...PatchFunc) (bool, error) {
	hadUpdate := false
	for _, apply := range patches {
		didUpdate, err := apply(h)
		if didUpdate {
			hadUpdate = true
		}

		if err != nil {
			return hadUpdate, err
		}
	}

	return hadUpdate, nil
}

func SetToUp(h *Status) (bool, error) {
	if h.State != STATE_UP {
		h.State = STATE_UP
		return true, nil
	}

	return false, nil
}

func SetToDown(h *Status) (bool, error) {
	if h.State != STATE_DOWN {
		h.State = STATE_DOWN
		return true, nil
	}

	return false, nil
}

func SetReason(reason string) PatchFunc {
	return func(h *Status) (bool, error) {
		if h.Reason != reason {
			h.Reason = reason
			return true, nil
		}

		return false, nil
	}
}

type HttpStatus struct {
	Status `mapstructure:",squash"`

	Address             string        `mapstructure:"address"`
	ExpectedStatusCodes []int         `mapstructure:"expectedStatusCodes"`
	Interval            time.Duration `mapstructure:"interval"`
	Timeout             time.Duration `mapstructure:"timeout"`
}

// time.Duration has a base type of int64, so 0 should be the empty/nil value. In this case, default to something else
func (h HttpStatus) GetTimeout() time.Duration {
	if h.Timeout == 0 {
		return time.Second * 30
	}

	return h.Timeout
}

// TODO: Add a compare function
