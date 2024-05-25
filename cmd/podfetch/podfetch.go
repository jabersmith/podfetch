package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"jaypod/pkg/engine"
	"jaypod/pkg/state"
	"jaypod/pkg/subscription"
)

func main() {

	var subscriptionFile = flag.String("f", "", "subscriptions yaml file")
	var stateFile = flag.String("s", "", "subscriptions state file")
	var dir = flag.String("d", "", "directory into which podcasts should be saved")
	var testmode = flag.Bool("t", false, "log output without downloading files")
	var wakeInterval = flag.Int("w", 0, "if > 0, number of minutes to wait between rss pulls")

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs

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

	if *wakeInterval > 0 {
		tick := time.NewTicker(time.Duration(*wakeInterval) * time.Minute)
		for ; ; <-tick.C {
			pull(*subscriptionFile, *stateFile, *dir, *testmode)
		}
	} else {
		pull(*subscriptionFile, *stateFile, *dir, *testmode)
	}

}

func pull(subscriptionFile, stateFile, dir string, testmode bool) {

	start := time.Now()

	feedsYaml, err := os.ReadFile(subscriptionFile)
	if err != nil {
		log.Error().
			Str("filename", subscriptionFile).
			Err(err).
			Msg("Bad file")
		return
	}

	feeds, err := subscription.ParseFeeds(feedsYaml)
	if err != nil {
		log.Error().
			Err(err).
			Msg("parse error on feeds.yaml")
		return
	}

	state, err := state.LoadState(stateFile)
	if err != nil {
		log.Error().
			Err(err).
			Str("file", stateFile).
			Msg("error loading state file")
		return
	}

	downloads, err := engine.Fetch(feeds, state, dir, testmode)
	if err != nil {
		log.Error().
			Err(err).
			Msg("error during fetch")
		return
	}

	log.Info().
		Dur("elapsed", time.Now().Sub(start)).
		Int("downloads", downloads).
		Msg("wakeup")
	return
}
