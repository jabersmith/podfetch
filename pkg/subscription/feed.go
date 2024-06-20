package subscription

import (
	"bytes"
	"fmt"
	"log/slog"
	"os"
	"regexp"
	"strings"
	"text/template"

	"github.com/goccy/go-yaml"

	"jaypod/pkg/rss"
)

type Wrapper struct {
	Feeds []*Feed
}

type Feed struct {
	Name    string
	Url     string
	Filters []*Filter
}

type Filter struct {
	TitleExpression       string `yaml:"title_regex"`
	TitleRegexp           *regexp.Regexp
	DescriptionExpression string `yaml:"description_regex"`
	DescriptionRegexp     *regexp.Regexp
	FilenameExpression    string `yaml:"filename_regex"`
	FilenameRegexp        *regexp.Regexp
	Subdir                string
	Skip                  bool
	Filename              string
	FilenameTemplate      *template.Template
	Incoming              bool
	dest                  string
}

func ParseDir(dir string) ([]*Feed, error) {

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	feeds := []*Feed{}
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if !strings.HasSuffix(file.Name(), ".yaml") {
			continue
		}

		path := dir + "/" + file.Name()

		feedsYaml, err := os.ReadFile(path)
		if err != nil {
			slog.Error("Unreadable file",
				"filename", path,
				"error", err)
		}

		newFeeds, err := ParseFeeds(feedsYaml)
		if err != nil {
			slog.Error("Error parsing file",
				"filename", path,
				"error", err)
		}
		feeds = append(feeds, newFeeds...)
	}
	return feeds, nil
}

func ParseFeeds(doc []byte) ([]*Feed, error) {
	var w Wrapper

	if err := yaml.Unmarshal(doc, &w); err != nil {
		return []*Feed{}, err
	}

	for _, feed := range w.Feeds {
		for _, filter := range feed.Filters {
			//			fmt.Printf("filter: %+v\n", filter)

			if filter.TitleExpression != "" {
				re, err := regexp.Compile("^" + filter.TitleExpression + "$")
				if err != nil {
					return []*Feed{}, fmt.Errorf("error parsing regexp %s for feed %s", filter.TitleExpression, feed.Url)
				}
				filter.TitleRegexp = re
			}

			if filter.DescriptionExpression != "" {
				re, err := regexp.Compile("^" + filter.DescriptionExpression + "$")
				if err != nil {
					return []*Feed{}, fmt.Errorf("error parsing regexp %s for feed %s", filter.DescriptionExpression, feed.Url)
				}
				filter.DescriptionRegexp = re
			}

			if filter.FilenameExpression != "" {
				re, err := regexp.Compile("^" + filter.FilenameExpression + "$")
				if err != nil {
					return []*Feed{}, fmt.Errorf("error parsing regexp %s for feed %s", filter.FilenameExpression, feed.Url)
				}
				filter.FilenameRegexp = re
			}

			if filter.Subdir != "" {
				filter.dest = fmt.Sprintf("%s/%s", feed.Name, filter.Subdir)
			} else {
				filter.dest = feed.Name
			}

			if filter.Filename != "" {
				t, err := template.New(filter.dest).Option("missingkey=zero").Parse(filter.Filename)
				if err != nil {
					return []*Feed{}, fmt.Errorf("error parsing filename %s for feed %s", filter.Filename, feed.Url)
				}
				filter.FilenameTemplate = t
			}

		}
	}

	return w.Feeds, nil
}

func (f *Feed) MatchAndMap(podcast *rss.RssItem) (bool, string, string, bool) {
	for _, filter := range f.Filters {
		match, dest, filebasename, incoming := filter.matchAndMap(podcast)
		if match {
			return match, dest, filebasename, incoming
		}
	}
	return false, "", "", false
}

func (f *Filter) matchAndMap(podcast *rss.RssItem) (bool, string, string, bool) {

	//	fmt.Printf("comparing %+v to {%s, %s}\n", f, title, description)

	subst := map[string]string{}
	for k, v := range podcast.Attrs() {
		subst[k] = v
	}

	if f.TitleRegexp != nil {
		matches := f.TitleRegexp.FindStringSubmatch(podcast.Title())
		if matches == nil {
			return false, "", "", false
		}

		for i, m := range matches {
			if i > 0 {
				subst[f.TitleRegexp.SubexpNames()[i]] = m
			}
		}

	}

	if f.DescriptionRegexp != nil {
		matches := f.DescriptionRegexp.FindStringSubmatch(podcast.Description())
		if matches == nil {
			return false, "", "", false
		}

		for i, m := range matches {
			if i > 0 {
				subst[f.DescriptionRegexp.SubexpNames()[i]] = m
			}
		}

	}

	if f.FilenameRegexp != nil {
		matches := f.FilenameRegexp.FindStringSubmatch(podcast.FileBaseName())
		if matches == nil {
			return false, "", "", false
		}

		for i, m := range matches {
			if i > 0 {
				subst[f.FilenameRegexp.SubexpNames()[i]] = m
			}
		}

	}

	if f.Skip {
		return true, "", "", false
	}

	if f.FilenameTemplate == nil {
		return true, f.dest, "", f.Incoming
	}

	var b bytes.Buffer
	f.FilenameTemplate.Execute(&b, subst)
	return true, f.dest, b.String(), f.Incoming
}
