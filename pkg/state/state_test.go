package state

import (
	"bytes"
	"testing"
	"time"
)

func TestEmptyYaml(t *testing.T) {
	s, err := stateFromYaml([]byte("\n "))
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	if s == nil {
		t.Fatalf("nil map")
	}

	if len(s) != 0 {
		t.Fatalf("non-empty map: %+v", s)
	}
}

func TestYamlParse(t *testing.T) {
	s, err := stateFromYaml([]byte(`
http://wtfpod.libsyn.com/rss: 111111
https://www.patreon.com/rss/theflagrantones?auth=PYkre__74n16LEDkBSkLAk4dkdRmZANq: 3123
`))
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	if s == nil {
		t.Fatalf("nil map")
	}

	if len(s) != 2 {
		t.Fatalf("expected 2 entries, got %d: %+v", len(s), s)
	}

	if s["https://www.patreon.com/rss/theflagrantones?auth=PYkre__74n16LEDkBSkLAk4dkdRmZANq"].last != time.Unix(3123, 0) {
		t.Fatalf("bad last time for HH: expected 3123, got %v: %+v",
			s["https://www.patreon.com/rss/theflagrantones?auth=PYkre__74n16LEDkBSkLAk4dkdRmZANq"].last, s)
	}
}

func TestYamlFromState(t *testing.T) {
	in := map[string]FeedState{
		"http://wtfpod.libsyn.com/rss": FeedState{last: time.Unix(111111, 0)},
		"https://www.patreon.com/rss/theflagrantones?auth=PYkre__74n16LEDkBSkLAk4dkdRmZANq": FeedState{last: time.Unix(3123, 0)},
	}

	out := []byte(`http://wtfpod.libsyn.com/rss: 111111
https://www.patreon.com/rss/theflagrantones?auth=PYkre__74n16LEDkBSkLAk4dkdRmZANq: 3123
`)

	b, err := yamlFromState(in)
	if err != nil {
		t.Fatalf("marshal error: %v", err)
	}

	if !bytes.Equal(b, out) {
		t.Fatalf("bad marshal results: expected %+v, got %+v", string(out), string(b))
	}

}
