package staticPresentation

import (
	"fmt"
	"strings"

	"github.com/ingmardrewing/fs"
	"github.com/ingmardrewing/staticIntf"
)

func NewRssRenderer(
	pages []staticIntf.Page,
	fsPath string,
	pathFromDocRoot string,
	rssFilename string) *rssRenderer {

	r := new(rssRenderer)
	r.pages = pages

	r.fsPath = fsPath
	r.pathFromDocRoot = pathFromDocRoot

	r.rssFilename = rssFilename
	return r
}

type rssRenderer struct {
	pages     []staticIntf.Page
	targetDir string

	fsPath          string
	pathFromDocRoot string

	rssFilename string
}

func (r *rssRenderer) Render() fs.FileContainer {
	if len(r.pages) > 0 {
		fc := fs.NewFileContainer()
		fc.SetPath(r.fsPath)
		fc.SetFilename(r.rssFilename)
		fc.SetDataAsString(r.renderRssContent())
		return fc
	}
	return nil
}

func (r *rssRenderer) renderRssContent() string {
	rss := []string{}
	for _, p := range r.pages {
		rss = append(rss, r.renderSingleRssEntry(p))
	}
	itemsRss := strings.Join(rss, "\n")
	domain := r.pages[0].Domain()
	title := domain
	favUrl := "https://" + domain + "/favicon-32x32.png"
	favTitle := domain
	favLink := "https://" + domain
	atomLink := "https://" + domain + r.pathFromDocRoot + r.rssFilename
	link := "https://" + domain + r.pathFromDocRoot + r.rssFilename
	description := ""

	lastbuilddate := r.pages[0].PublishedTime()

	return fmt.Sprintf(
		rssTemplate,
		title,
		favUrl,
		favTitle,
		favLink,
		atomLink,
		link,
		description,
		lastbuilddate,
		itemsRss)
}

func (r *rssRenderer) renderSingleRssEntry(p staticIntf.Page) string {
	return fmt.Sprintf(rssItemTemplate,
		p.Title(),
		p.Url(),
		p.PublishedTime(),
		p.Url(),
		p.Description(),
		p.Content(),
		p.ImageUrl(),
		p.ImageUrl(),
		p.ImageUrl(),
		p.ImageUrl())
}

var rssItemTemplate string = `  <item>
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

var rssTemplate string = `<?xml version="1.0" encoding="UTF-8"?><rss version="2.0"
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
      <url>%s</url>
      <title>%s</title>
      <link>%s</link>
      <width>32</width>
      <height>32</height>
      <description>A science-fiction webcomic about the lives of software developers in the far, funny and dystopian future</description>
    </image>
	<atom:link href="%s" rel="self" type="application/rss+xml" />
	<link>%s</link>
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
