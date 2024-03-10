package engine

import (
	"fmt"
	"io"
	"net/http"
	"slices"
	"time"

	"jaypod/pkg/rss"
	"jaypod/pkg/state"
	"jaypod/pkg/subscription"
)

func Fetch(feeds []*subscription.Feed, state *state.State, rootdir string, testmode bool) error {
	for _, feed := range feeds {
		resp, err := http.Get(feed.Url)
		if err != nil {
			return fmt.Errorf("failed getting %s: %v", feed.Url, err)
		}

		contents, err := io.ReadAll(resp.Body)
		resp.Body.Close()

		rc, err := rss.ParseRss([]byte(contents))
		if err != nil {
			return fmt.Errorf("parse error on %s: %v", feed.Url, err)
		}

		last := state.Last(feed.Url)
		newLast, err := fetchNewFromFeed(rc, feed, rootdir, last, testmode)
		if err != nil {
			return err
		}
		state.Update(feed.Url, newLast)

		err = state.Flush()
		if err != nil {
			return fmt.Errorf("error flushing state: %v\n", err)
		}

	}

	return nil
}

func fetchNewFromFeed(rc rss.RssContainer, feed *subscription.Feed, rootdir string, last time.Time, testmode bool) (time.Time, error) {

	podcasts := rc.Podcasts()

	newPodcasts := make([]*rss.RssItem, 0, len(podcasts))
	for _, p := range podcasts {
		if p.PubDate.After(last) {
			newPodcasts = append(newPodcasts, p)
		}
	}

	// sort oldest to newest, so we can advance the last even if we
	// don't get all the podcasts this time
	slices.SortFunc(newPodcasts, func(a, b *rss.RssItem) int {
		return a.PubDate.Compare(b.PubDate)
	})

	newLast := last

	for _, p := range newPodcasts {
		match, dest, basename, incoming := feed.MatchAndMap(p)
		if match && dest != "" {
			err := act(testmode, p, rootdir, dest, basename, incoming)
			if err != nil {
				fmt.Printf("%v\n", err)
				break
			}
			if p.PubDate.After(newLast) {
				newLast = p.PubDate
			}
		}
	}

	return newLast, nil
}

func act(testmode bool, podcast *rss.RssItem, rootdir string, dest string, basename string, incoming bool) error {
	if testmode {
		return trialRun(podcast, rootdir, dest, basename, incoming)
	} else {
		return download(podcast, rootdir, dest, basename, incoming)
	}
}

func trialRun(podcast *rss.RssItem, rootdir string, dest string, basename string, incoming bool) error {

	fmt.Printf("%s: would download to %s/%s", podcast.Title(), rootdir, dest)
	if basename != "" {
		fmt.Printf(" and rename to %s", basename)
	}
	if incoming {
		fmt.Printf(" and copy to %s/Incoming", rootdir)
	}
	fmt.Printf("\n")

	return nil
}
