package rss

import (
	"testing"
	"time"
)

const waitWhatFragment = `<?xml version="1.0" encoding="UTF-8"?>
<rss xmlns:itunes="http://www.itunes.com/dtds/podcast-1.0.dtd" version="2.0">

  <channel>

    <title>Wait, What?</title>
    <description>A fortnightly comics and pop culture podcast</description>
    <link>http://www.waitwhatpodcast.com</link>
    <language>en-us</language>
    <copyright>Copyright 2010</copyright>
    <lastBuildDate>Mon, 16 Aug 2010 11:30:00 -0500</lastBuildDate>
    <pubDate>Tue, 10 Aug 2010 11:30:00 -0500</pubDate>
    <docs>http://blogs.law.harvard.edu/tech/rss</docs>
    <webMaster>pig.latin@gmail.com (Jeff Lester)</webMaster>
    <itunes:author>Jeff and Graeme @ waitwhatpodcast.com</itunes:author>
    <itunes:subtitle>Comic books, graphic novels, and pop culture, past and present.</itunes:subtitle>
    <itunes:summary>Ostensibly about comic books and graphic novels, "Wait, What?" is a podcast in which two friends swap stories, theories, and jokes about all aspects of pop culture...but especially comics.
	For our early episodes, our theme music is Track 18 from Nine Inch Nails&apos; wonderful ambient album Ghosts, which Trent Reznor graciously made available under a Creative Commons BY-NC-SA license.
        But for some time now, it has been an original composition by our very own Graeme McMillan: this update is overdue!</itunes:summary>

    <itunes:owner>
           <itunes:name>Jeff L</itunes:name>
           <itunes:email>pig.latin@gmail.com</itunes:email>
    </itunes:owner>

<itunes:explicit>No</itunes:explicit>

<itunes:image href="http://www.theworkingdraft.com/img/WaitWhatXLarge.jpg"/>
   
<itunes:category text="Arts">
     <itunes:category text="Books"/>
</itunes:category>

<item>
<title>Wait, What: Episode 1.1 </title>
<link>http://www.waitwhatpodcast.com</link>
<guid>http://theworkingdraft.com/media/podcasts/WaitWhat1point1.mp3</guid>
<description>In the very first installment of our first episode, Graeme McMillan and Jeff Lester talk about the genius that is Marvel Comics&apos; Dazzler.

However, because we are still learning the ropes of how to record our conversations, nobody will be talking about the genius of our sound mixing.

We promise you it improves greatly in future podcasts and hope you hang around.  Thanks for listening!</description>
<enclosure url="http://theworkingdraft.com/media/podcasts/WaitWhat1point1.mp3" length="13753600" type="audio/mpeg"/>
<category>Books</category>
<pubDate>Mon, 08 Jun 2009 11:30:00 -0500</pubDate>

<itunes:author>Jeff and Graeme @ waitwhatpodcast</itunes:author>

<itunes:explicit>No</itunes:explicit>
<itunes:subtitle>Welcome to Wait, What, a comics podcast for the chronically digressive!</itunes:subtitle>
<itunes:summary> In the very first installment of our first episode, Graeme McMillan and Jeff Lester talk about the genius that is Marvel Comics&apos; Dazzler and Skull the Slayer.

However, because we are still learning the ropes of how to record our conversations, nobody will be talking about the genius of our sound mixing.

We promise you it improves greatly in future podcasts and hope you hang around.  Thanks for listening!</itunes:summary>
<itunes:duration>21:41</itunes:duration>
<itunes:keywords>Wait What, Savage Critic, comic books, graphic novels, Graeme McMillan, Jeff Lester, Marvel Comics, Dazzler, Negative Zone, Skull The Slayer</itunes:keywords>
</item>

<item>
<title>Wait, What? - The April Fools Edition</title>
<link>http://www.waitwhatpodcast.com</link>
<guid>http://theworkingdraft.com/media/podcasts5/WaitWhatAprilFools.mp3</guid>
<description>Fools Rush In When Ike Perlmutter Gets Rushed Out!</description>
<enclosure url="http://theworkingdraft.com/media/podcasts5/WaitWhatAprilFools.mp3" length="36965224" type="audio/mpeg"/>
<category>Books</category>
<pubDate>Sat, 01 Apr 2023 18:58:00 EST</pubDate>

<itunes:author>Jeff and Graeme @ waitwhatpodcast</itunes:author>

<itunes:explicit>No</itunes:explicit>
<itunes:subtitle>Welcome to Wait, What, a comics podcast!</itunes:subtitle>
<itunes:summary>Fools Rush In When Ike Perlmutter Gets Rushed Out!</itunes:summary>
<itunes:duration>01:17:01</itunes:duration>
<itunes:keywords>Wait What, Savage Critic, comic books, graphic novels, Graeme McMillan, Jeff Lester, Graham, Geoff, Steve Englehart</itunes:keywords>
</item>


<item>
  <title>1: No Such Thing As A Pilot Fish</title>
  <link>https://audioboom.com/posts/4960884</link>
  <itunes:episode>1</itunes:episode>
  <itunes:title>No Such Thing As A Pilot Fish</itunes:title>
  <enclosure url="https://pscrb.fm/rss/p/pdst.fm/e/arttrk.com/p/ABMA5/dts.podtrac.com/redirect.mp3/audioboom.com/posts/4960884.mp3?modified=1599215998&amp;sid=2399216&amp;source=rss" length="0" type="audio/mpeg" />
  <media:content url="https://pscrb.fm/rss/p/pdst.fm/e/arttrk.com/p/ABMA5/dts.podtrac.com/redirect.mp3/audioboom.com/posts/4960884.mp3?modified=1599215998&amp;sid=2399216&amp;source=rss" type="audio/mpeg" medium="audio" duration="1980" />
  <itunes:image href="https://audioboom.com/i/36345807/s=1400x1400/el=1/rt=fill.jpg" />
  <media:content url="https://audioboom.com/i/36345807/s=1400x1400/el=1/rt=fill.jpg" type="image/jpg" medium="image"  />
  <itunes:duration>1980</itunes:duration>
  <itunes:explicit>no</itunes:explicit>
  <itunes:episodeType>full</itunes:episodeType>
  <description><![CDATA[<div>A new podcast from the writers of QI, who discuss the best facts they've found that week. The pilot episode features Dan Schreiber (@schreiberland), James Harkin (@eggshaped), Anna Ptaszynski (@nosuchthing) &amp; Andrew Hunter Murray (@andrewhunterm)<br>
<br>

</div>
<div>For more check out <a href="http://www.qi.com/podcast">www.qi.com/podcast</a><br>
<br>

</div>
]]></description>
  <pubDate>Sat, 08 Mar 2014 00:00:00 +0000</pubDate>
  <guid isPermaLink="false">tag:soundcloud,2010:tracks/138526614</guid>
  <itunes:author>No Such Thing As A Fish</itunes:author>
  <dc:creator>No Such Thing As A Fish</dc:creator>
  <media:rights status="userCreated" />
</item>


</channel>
</rss>
`

var waitWhatExpected = []struct {
	name string
	date time.Time
	url  string
}{
	{
		name: "Wait, What: Episode 1.1 ",
		date: time.Date(2009, 6, 8, 11, 30, 00, 00, time.FixedZone("UTC-5", -5*60*60)),
		url:  "http://theworkingdraft.com/media/podcasts/WaitWhat1point1.mp3",
	},
	{
		name: "Wait, What? - The April Fools Edition",
		date: time.Date(2023, 4, 1, 18, 58, 00, 00, time.FixedZone("UTC-5", -5*60*60)),
		url:  "http://theworkingdraft.com/media/podcasts5/WaitWhatAprilFools.mp3",
	},
	{
		name: "1: No Such Thing As A Pilot Fish",
		date: time.Date(2014, 3, 8, 00, 00, 00, 00, time.FixedZone("UTC", 0)),
		url:  "https://pscrb.fm/rss/p/pdst.fm/e/arttrk.com/p/ABMA5/dts.podtrac.com/redirect.mp3/audioboom.com/posts/4960884.mp3?modified=1599215998&sid=2399216&source=rss",
	},
}

func TestParseRss(t *testing.T) {

	rc, err := ParseRss([]byte(waitWhatFragment))
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	podcasts := rc.Podcasts()

	if len(podcasts) != len(waitWhatExpected) {
		t.Fatalf("wrong number of podcasts: expected %d, got %d", len(waitWhatExpected), len(podcasts))
	}

	for i, p := range podcasts {
		exp := waitWhatExpected[i]
		if exp.date.Unix() != p.Date().Unix() {
			t.Errorf("wrong date for podcast %d: expected %v (%v), got %v (%v)", i, exp.date, exp.date.Unix(), p.PubDate, p.Date().Unix())
		}
		if p.Title() != exp.name {
			t.Errorf("wrong name for podcast %d: expected %v, got %v", i, exp.name, p.Title())
		}
		if p.Url() != exp.url {
			t.Errorf("wrong url for podcast %d: expected %v, got %v", i, exp.url, p.Url())
		}

	}

}
