package subscription

import (
	"testing"

	"jaypod/pkg/rss"
)

const subYaml = `
feeds:
  - name: "Comedy/WTF"
    url: http://wtfpod.libsyn.com/rss
    filters:
      - title_regex: "Episode (?P<epnum>\d*) - (?P<eptitle>.*)"
        filename: "{{.epnum}} {{.eptitle}}"
        incoming: true
      - filename: "{{.title}}"
        incoming: true
  - name: "Comedy/TheBestShow"
    url: "https://www.patreon.com/rss/TheBestShow?auth=mxJC8pCEXWZ7LYrtS9vX7ek3vPHWB6LV"
    filters:
      - title_regex: "Meet My Friends.*"
        subdir: "MMFTF"
        incoming: true
      - title_regex: "Takin' The Leap.*"
        subdir: "MMFTF"
        incoming: true
      - title_regex: "Meet My Friends.*"
        subdir: "MMFTF"
        incoming: true
      - title_regex: ".*John Gentle .*"
        subdir: "JohnGentle"
        incoming: true
      - title_regex: "Ask Tom.*"
        subdir: "AskTom"
        incoming: true
      - title_regex: "Make Mike.*"
        subdir: "MakeMikeMarvel"
        incoming: true
      - title_regex: "MAKE DC.*"
        subdir: "MakeMikeMarvel"
        incoming: true
      - title_regex: "Rubinesque.*"
        subdir: "Rubinesque"
        incoming: true
      - title_regex: "Four Horsemen.*"
        subdir: "Horsemen"
        incoming: true
      - title_regex: "Four Horseman.*"
        subdir: "Horsemen"
        incoming: true
      - title_regex: "BEST SHOW: FOUR HORSEMEN.*"
        subdir: "Horsemen"
        incoming: true
      - title_regex: "S&W.*"
        subdir: "SandW"
        incoming: true
      - title_regex: "Ahoy.*"
        subdir: "Ahoy"
        incoming: true
      - title_regex: "So Far.*"
        subdir: "SoFar"
        incoming: true
      - description_regex: ".*Originally.*"
        subdir: "BestShowBests"
        incoming: true
      - subdir: "Main"
        incoming: true
`

func TestParseFeeds(t *testing.T) {

	feeds, err := ParseFeeds([]byte(subYaml))
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	if feeds[0].Url != "http://wtfpod.libsyn.com/rss" {
		t.Fatalf("wrong podcast subscription: expected wtf, got %s", feeds[0].Url)
	}

	if len(feeds[0].Filters) != 2 {
		t.Fatalf("wrong number of filters: expected 2, got %d", len(feeds[0].Filters))
	}

	var wtfExpected = []struct {
		title        string
		description  string
		match        bool
		dest         string
		filebasename string
		incoming     bool
	}{
		{
			title:        "Episode 1512 - Da'Vine Joy Randolph",
			description:  "",
			match:        true,
			dest:         "Comedy/WTF",
			filebasename: "1512 Da'Vine Joy Randolph",
			incoming:     true,
		},
		{
			title:        "Wayne Kramer from 2014",
			description:  "",
			match:        true,
			dest:         "Comedy/WTF",
			filebasename: "Wayne Kramer from 2014",
			incoming:     true,
		},
	}

	for i, x := range wtfExpected {
		match, dest, filebasename, incoming := feeds[0].MatchAndMap(
			makeRssItem(x.title, x.description))
		if match != x.match {
			t.Errorf("wtfExpected[%d] - expected match %v, got %v", i, x.match, match)
		}
		if dest != x.dest {
			t.Errorf("wtfExpected[%d] - expected dest %v, got %v", i, x.dest, dest)
		}
		if filebasename != x.filebasename {
			t.Errorf("wtfExpected[%d] - expected filebasename %v, got %v", i, x.filebasename, filebasename)
		}
		if incoming != x.incoming {
			t.Errorf("wtfExpected[%d] - expected incoming %v, got %v", i, x.incoming, incoming)
		}
	}

	var bestShowExpected = []struct {
		title        string
		description  string
		match        bool
		dest         string
		filebasename string
		incoming     bool
	}{
		{
			title:        "LIVE MUSIC FROM DAIISTAR! WHAT DO YOU LOVE? WE FOOLED PETER FUNT! WE ARE THE WORLD DOCUMENTARY!",
			description:  "&lt;p&gt;Phones ring on the topic: WHAT DO YOU LOVE? Austin TX noise pop band DAIISTAR plays a live set! Candid Camera's Peter Funt fooled by our Billy Joel song! The Horseman discuss the new documentary about We Are The World! And so much more!!&lt;/p&gt;",
			match:        true,
			dest:         "Comedy/TheBestShow/Main",
			filebasename: "",
			incoming:     true,
		},
		{
			title:        "Darren From Work",
			description:  "&lt;p&gt;&lt;strong&gt;BEST SHOW BESTS! In this classic clip, Tom gets a call from DARREN FROM WORK! (Originally Aired August 23rd, 2022)&lt;/strong&gt;&lt;/p&gt;",
			match:        true,
			dest:         "Comedy/TheBestShow/BestShowBests",
			filebasename: "",
			incoming:     true,
		},
	}

	for i, x := range bestShowExpected {
		match, dest, filebasename, incoming := feeds[1].MatchAndMap(
			makeRssItem(x.title, x.description))
		if match != x.match {
			t.Errorf("bestShowExpected[%d] - expected match %v, got %v", i, x.match, match)
		}
		if dest != x.dest {
			t.Errorf("bestShowExpected[%d] - expected dest %v, got %v", i, x.dest, dest)
		}
		if filebasename != x.filebasename {
			t.Errorf("bestShowExpected[%d] - expected filebasename %v, got %v", i, x.filebasename, filebasename)
		}
		if incoming != x.incoming {
			t.Errorf("bestShowExpected[%d] - expected incoming %v, got %v", i, x.incoming, incoming)
		}
	}
}

func makeRssItem(title, description string) *rss.RssItem {
	return &rss.RssItem{
		MyTitle:       rss.NonNamespaceString(title),
		MyDescription: description,
	}
}
