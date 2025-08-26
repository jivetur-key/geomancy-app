package main

import (
	"flag"
	"fmt"
	"geomancy-app/cli"
	"geomancy-app/cmd"
	"os"
)

func init() {
	cli.InitFlags()
}

func main() {
	flag.Parse()
	var (
		planet string
		err    error
	)

	if cli.IsFlagSet("planet") {
		planet, err = cli.GetPlanetNumber()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
	} else {
		planet = "None"
	}
	err = cmd.Cmd(planet, cli.Rubeus)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
