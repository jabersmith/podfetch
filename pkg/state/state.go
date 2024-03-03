package state

import (
	"fmt"
	"os"
	"time"

	"github.com/goccy/go-yaml"
)

type State struct {
	filename string
	s        map[string]FeedState
}

type FeedState struct {
	last time.Time
}

func stateFromYaml(contents []byte) (map[string]FeedState, error) {
	tmp := map[string]int64{}
	cooked := map[string]FeedState{}

	if err := yaml.Unmarshal(contents, &tmp); err != nil {
		return cooked, err
	}

	for url, epoch := range tmp {
		cooked[url] = FeedState{last: time.Unix(epoch, 0)}
	}
	return cooked, nil
}

func yamlFromState(s map[string]FeedState) ([]byte, error) {
	tmp := map[string]int64{}

	for url, fs := range s {
		tmp[url] = fs.last.Unix()
	}

	b, err := yaml.Marshal(tmp)
	if err != nil {
		return []byte{}, err
	}
	return b, nil
}

func LoadState(filename string) (*State, error) {
	stateYaml, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Failed to read state file: %v", err)
	}

	s, err := stateFromYaml(stateYaml)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse state file: %v", err)
	}
	return &State{filename: filename, s: s}, nil
}

func (s *State) Flush() error {
	y, err := yamlFromState(s.s)
	if err != nil {
		return fmt.Errorf("failed to marshal state: %v", err)
	}

	tmpfile := s.filename + ".tmp"
	err = os.WriteFile(tmpfile, y, 0666)
	if err != nil {
		return fmt.Errorf("failed to write tmp state file %s: %v", tmpfile, err)
	}

	err = os.Rename(tmpfile, s.filename)
	if err != nil {
		return fmt.Errorf("failed to rename tmp state file %s to %s: %v",
			tmpfile, s.filename, err)
	}

	return nil
}

func (s *State) Last(url string) time.Time {
	return s.s[url].last
}

func (s *State) Update(url string, last time.Time) {
	s.s[url] = FeedState{last: last}
}
