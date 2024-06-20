package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"time"

	"jaypod/pkg/engine"
	"jaypod/pkg/state"
	"jaypod/pkg/subscription"
)

func main() {

	var subscriptionDir = flag.String("f", "", "subscriptions directory")
	var stateFile = flag.String("s", "", "subscriptions state file")
	var dir = flag.String("d", "", "directory into which podcasts should be saved")
	var testmode = flag.Bool("t", false, "log output without downloading files")
	var wakeInterval = flag.Int("w", 0, "if > 0, number of minutes to wait between rss pulls")

	flag.Parse()

	if *subscriptionDir == "" {
		fmt.Fprintf(os.Stderr, "missing required feeds directory\n")
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

	//	slog.SetDefault(

	if *wakeInterval > 0 {
		tick := time.NewTicker(time.Duration(*wakeInterval) * time.Minute)
		for ; ; <-tick.C {
			pull(*subscriptionDir, *stateFile, *dir, *testmode)
		}
	} else {
		pull(*subscriptionDir, *stateFile, *dir, *testmode)
	}

}

func pull(subscriptionDir, stateFile, dir string, testmode bool) {

	start := time.Now()

	feeds, err := subscription.ParseDir(subscriptionDir)
	if err != nil {
		slog.Error("error loading feeds",
			"error", err)
		return
	}

	state, err := state.LoadState(stateFile)
	if err != nil {
		slog.Error("error loading state file",
			"filename", stateFile,
			"error", err)
		return
	}

	downloads, err := engine.Fetch(feeds, state, dir, testmode)
	if err != nil {
		slog.Error("error during fetch",
			"error", err)
		return
	}

	slog.Info("wakeup",
		"elapsed", time.Now().Sub(start),
		"downloads", downloads)
	return
}
