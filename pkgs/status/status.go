package status

const (
	STATE_UP   = 0
	STATE_DOWN = 1
)

type Status struct {
	State  int
	Reason string
	Name   string `mapstructure:"name"`
}

type HttpStatus struct {
	Status `mapstructure:",squash"`

	Address             string `mapstructure:"address"`
	Port                int    `mapstructure:"port"`
	ExpectedStatusCodes int    `mapstructure:"expectedStatusCodes"`
}
