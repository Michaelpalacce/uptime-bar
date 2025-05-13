package run

import (
	"flag"
	"os"

	"github.com/Michaelpalacce/uptime-bar/pkgs/args"
)

var usage = `uptime-bar is a utility tool used to provide a local rest endpoint that will give you details about uptime of different systems.
`

var examples = `
# Basic Usage
`

// Args will parse the CLI arguments once and return the parsed options from then on
// This will panic if there are any validation issues
func (c *RunCommand) Args() {
	// if runOptions.Parsed {
	// 	return runOptions
	// }

	args, err := args.NewArgs(
		os.Args[2:],
		args.WithExamples(examples),
		args.WithUsage(usage),
		args.WithFs(flag.NewFlagSet("run", flag.ExitOnError)),
	)
	if err != nil {
		panic(err)
	}

	// args.AddVar(&runOptions.Software.JavaVersion, "javaVersion", "jv", "17", "Which version of java to install? If not set, will skip installation.")

	if err := args.Parse(); err != nil {
		panic(err)
	}

	// runOptions.Parsed = true
	//
	// return runOptions
}

