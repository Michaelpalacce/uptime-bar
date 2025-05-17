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

type HttpStatus struct {
	Status `mapstructure:",squash"`

	Address             string        `mapstructure:"address"`
	Port                int           `mapstructure:"port"`
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
