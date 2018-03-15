package staticPresentation

import (
	"fmt"
	"strings"

	"github.com/ingmardrewing/fs"
	"github.com/ingmardrewing/staticIntf"
)

func NewRssRenderer(pages []staticIntf.Page) *rssRenderer {
	r := new(rssRenderer)
	r.pages = pages
	return r
}

type rssRenderer struct {
	pages []staticIntf.Page
}

func (r *rssRenderer) Render(rssPath, rssFilename string) fs.FileContainer {
	fc := fs.NewFileContainer()
	fc.SetPath(rssPath)
	fc.SetFilename(rssFilename)
	fc.SetDataAsString(r.rssContent())
	return fc
}

func (r *rssRenderer) rssContent() string {
	rss := []string{}
	for _, p := range r.pages {
		rss = append(rss, r.getSingleRssEntry(p))
	}
	itemsRss := strings.Join(rss, "\n")

	rssTemplate := `<?xml version="1.0" encoding="UTF-8"?><rss version="2.0"
	xmlns:content="http://purl.org/rss/1.0/modules/content/"
	xmlns:wfw="http://wellformedweb.org/CommentAPI/"
	xmlns:dc="http://purl.org/dc/elements/1.1/"
	xmlns:atom="http://www.w3.org/2005/Atom"
	xmlns:sy="http://purl.org/rss/1.0/modules/syndication/"
	xmlns:slash="http://purl.org/rss/1.0/modules/slash/"
	xmlns:media="http://search.yahoo.com/mrss/"
	>

<channel>
	<title>%s</title>
    <image>
      <url>https://%s/favicon-32x32.png</url>
      <title>%s</title>
      <link>https://%s</link>
      <width>32</width>
      <height>32</height>
      <description>A science-fiction webcomic about the lives of software developers in the far, funny and dystopian future</description>
    </image>
	<atom:link href="https://%s/%s" rel="self" type="application/rss+xml" />
	<link>https://%s</link>
	<description>%s</description>
	<lastBuildDate>%s</lastBuildDate>
	<language>en-US</language>
	<sy:updatePeriod>weekly</sy:updatePeriod>
	<sy:updateFrequency>1</sy:updateFrequency>
	<generator>https://github.com/ingmardrewing/static</generator>
%s
	</channel>
</rss>
`
	domain := r.pages[0].Domain()
	date := r.pages[len(r.pages)-1].PublishedTime()
	return fmt.Sprintf(rssTemplate, domain, domain, domain, domain, domain, "rss.xml", date, "", itemsRss, "")
}

func (r *rssRenderer) getSingleRssEntry(p staticIntf.Page) string {
	rssItem := `  <item>
	<title>%s</title>
	<link>%s</link>
	<pubDate>%s</pubDate>
	<dc:creator><![CDATA[Ingmar Drewing]]></dc:creator>
	<guid>%s/index.html</guid>
	<description><![CDATA[%s]]></description>
	<content:encoded><![CDATA[%s]]></content:encoded>

	<media:thumbnail url="%s" />
	<media:content url="%s" medium="image">
	  <media:title type="html">%s</media:title>
	  <media:thumbnail url="%s" />
	</media:content>
  </item>`
	return fmt.Sprintf(rssItem, p.Title(), p.Url(), p.PublishedTime(), p.Url(), p.Description(), p.Content(), p.ImageUrl(), p.ImageUrl(), p.ImageUrl(), p.Title())

}
