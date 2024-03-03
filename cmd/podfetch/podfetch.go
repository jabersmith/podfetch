package main

import (
	"flag"
	"fmt"
	"os"

	"jaypod/pkg/engine"
	"jaypod/pkg/state"
	"jaypod/pkg/subscription"
)

func main() {

	var subscriptionFile = flag.String("f", "", "subscriptions yaml file")
	var stateFile = flag.String("s", "", "subscriptions state file")
	var dir = flag.String("d", "", "directory into which podcasts should be saved")
	var testmode = flag.Bool("t", false, "log output without downloading files")

	flag.Parse()

	if *subscriptionFile == "" {
		fmt.Fprintf(os.Stderr, "missing required feeds.yaml file\n")
		os.Exit(1)
	}

	if *stateFile == "" {
		fmt.Fprintf(os.Stderr, "missing required state file\n")
		os.Exit(1)
	}

	if *dir == "" {
		fmt.Fprintf(os.Stderr, "missing required output dir\n")
		os.Exit(1)
	}

	feedsYaml, err := os.ReadFile(*subscriptionFile)
	if err != nil {
		fmt.Printf("Bad file %s: %v\n", *subscriptionFile, err)
		os.Exit(1)
	}

	feeds, err := subscription.ParseFeeds(feedsYaml)
	if err != nil {
		fmt.Printf("parse error on feeds.yaml: %v\n", err)
		os.Exit(1)
	}

	state, err := state.LoadState(*stateFile)
	if err != nil {
		fmt.Printf("error loading state file %s: %v\n", *stateFile, err)
		os.Exit(1)
	}

	err = engine.Fetch(feeds, state, *dir, *testmode)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}
