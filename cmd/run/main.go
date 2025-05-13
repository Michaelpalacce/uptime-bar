package run

type RunCommand struct{}

func (c *RunCommand) Name() string {
	return "run"
}

func (c *RunCommand) Run() error {
	return nil
}
