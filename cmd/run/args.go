package run

import (
	"flag"
	"os"

	"github.com/Michaelpalacce/uptime-bar/internal/options"
	"github.com/Michaelpalacce/uptime-bar/pkgs/args"
)

var runOptions = &options.RunOptions{
	Parsed: false,
}

var usage = `uptime-bar is a utility tool used to provide a local rest endpoint that will give you details about uptime of different systems.
`

var examples = `
# Basic Usage

Due to the nature of this tool, the server should not be exposed outside of localhost. To start it, you can run:

uptime-bar run --address="http://127.0.0.1"
`

// Args will parse the CLI arguments once and return the parsed options from then on
// This will panic if there are any validation issues
func (c *RunCommand) Args() *options.RunOptions {
	if runOptions.Parsed {
		return runOptions
	}

	args, err := args.NewArgs(
		os.Args[2:],
		args.WithExamples(examples),
		args.WithUsage(usage),
		args.WithFs(flag.NewFlagSet("run", flag.ExitOnError)),
	)
	if err != nil {
		panic(err)
	}

	args.AddVar(&runOptions.RouterOptions.Address, "address", "a", "127.0.0.1", "Address to listen on. Do not specify the schema")
	args.AddVar(&runOptions.RouterOptions.Port, "port", "p", "9876", "Port to listen on.")

	if err := args.Parse(); err != nil {
		panic(err)
	}

	runOptions.Parsed = true

	return runOptions
}
