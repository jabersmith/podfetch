package rss

import (
	"encoding/xml"
	"fmt"
	"net/url"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// Uggh.  Cribbing from https://github.com/golang/go/issues/8535#issuecomment-312429764,
// a workaround for a years-old stdlib bug that keeps us from consuming both the podcast
// <title> and <itunes:title> tags
type NonNamespaceString string

func (s *NonNamespaceString) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Space != "" {
		// We do not want a namespace, so we need to consume it
		d.Skip()
		return nil
	} else {
		str := ""
		d.DecodeElement(&str, &start)
		*s = NonNamespaceString(str)
	}

	return nil
}

type RssContainer struct {
	XMLName xml.Name   `xml:"rss"`
	Version string     `xml:"version,attr"`
	Feed    RssChannel `xml:channel"`
}

type RssChannel struct {
	XMLName xml.Name   `xml:"channel"`
	Title   string     `xml:"title"`
	Items   []*RssItem `xml:"item"`
}

type RssItem struct {
	XMLName       xml.Name           `xml:"item"`
	MyTitle       NonNamespaceString `xml:"title"`
	MyDescription string             `xml:"description"`
	// guid?
	PubDateString string       `xml:"pubDate"`
	PubDate       time.Time    `xml:"-"`
	ItunesTitle   string       `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd title"`
	Duration      string       `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd duration"`
	Episode       string       `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd episode"`
	Season        string       `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd season"`
	Enclosure     RssEnclosure `xml:"enclosure"`
}

type RssEnclosure struct {
	XMLName       xml.Name `xml:"enclosure"`
	Url           string   `xml:"url,attr"`
	Length        string   `xml:"length,attr"`
	EnclosureType string   `xml:"type,attr"`
}

func ParseRss(doc []byte) (RssContainer, error) {
	var rc RssContainer

	if err := xml.Unmarshal(doc, &rc); err != nil {
		return rc, err
	}

	for _, item := range rc.Feed.Items {
		var err error
		item.PubDate, err = parsePodcastDate(item.PubDateString)
		if err != nil {
			return rc, fmt.Errorf("error parsing date %s: %v", item.PubDateString, err)
		}
	}
	return rc, nil
}

func parsePodcastDate(date string) (time.Time, error) {
	// rfc822 specifies support for the *ST USA timezone abbreviations, and
	// they've been grandfathered into the podcast "standards".  Hopefully
	// most feeds use UTC or timezone offsets.
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		return time.Time{}, err
	}

	formats := []string{
		time.RFC1123,
		time.RFC1123Z,
		"Mon, 2 Jan 2006 15:04:05 MST",   // rfc1123 without date padding
		"Mon, 2 Jan 2006 15:04:05 -0700", // rfc1123Z without date padding
	}

	var firstErr error
	for i, format := range formats {
		t, err := time.ParseInLocation(format, date, loc)
		if err == nil {
			return t, nil
		}
		if i == 0 {
			firstErr = err
		}
	}

	return time.Time{}, firstErr
}

func (rc RssContainer) Podcasts() []*RssItem {
	var ret []*RssItem
	for _, item := range rc.Feed.Items {
		if strings.HasPrefix(item.Enclosure.EnclosureType, "audio/") {
			ret = append(ret, item)
		}
	}
	return ret
}

func (i *RssItem) Attrs() map[string]string {
	m := map[string]string{}

	if string(i.MyTitle) != "" {
		m["title"] = string(i.MyTitle)
	}

	if i.ItunesTitle != "" {
		m["itunestitle"] = i.ItunesTitle
	}

	if i.Episode != "" {
		if epno, err := strconv.Atoi(i.Episode); err != nil {
			m["episode"] = fmt.Sprintf("%3d", epno)
		} else {
			m["episode"] = i.Episode
		}
	}

	if i.Season != "" {
		if sno, err := strconv.Atoi(i.Season); err != nil {
			m["season"] = fmt.Sprintf("%3d", sno)
		} else {
			m["season"] = i.Season
		}
	}

	m["date"] = i.PubDate.Format("2006-01-02")

	return m
}

func (i *RssItem) Date() time.Time {
	return i.PubDate
}

func (i *RssItem) Url() string {
	return i.Enclosure.Url
}

func (i *RssItem) Title() string {
	return string(i.MyTitle)
}

func (i *RssItem) Description() string {
	return i.MyDescription
}

func (i *RssItem) Type() string {
	return i.Enclosure.EnclosureType
}

func (i *RssItem) FileBaseName() string {
	u, err := url.Parse(i.Enclosure.Url)
	if err != nil {
		return ""
	}
	ubase := filepath.Base(u.Path)
	sep := strings.LastIndex(ubase, ".")
	if sep < 0 {
		return ubase
	}

	return ubase[:sep]
}

func (i *RssItem) ExtensionFromMimeType() string {
	extensions := map[string]string{
		"audio/mpeg":  "mp3",
		"audio/mp3":   "mp3",
		"audio/mp4":   "mp4",
		"audio/ogg":   "ogg",
		"audio/x-wav": "wav",
	}
	return extensions[i.Enclosure.EnclosureType]
}
