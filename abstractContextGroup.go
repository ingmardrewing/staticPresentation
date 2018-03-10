package staticPresentation

import (
	"fmt"
	"path"
	"strings"

	"github.com/ingmardrewing/fs"
	"github.com/ingmardrewing/staticIntf"
)

type abstractContextGroup struct {
	pages   []staticIntf.Page
	context staticIntf.Context
	site    staticIntf.Site
}

func (a *abstractContextGroup) GetComponents() []staticIntf.Component {
	return a.context.GetComponents()
}

func (a *abstractContextGroup) RenderPages() []fs.FileContainer {
	return a.context.RenderPages()
}

func (a *abstractContextGroup) Domain() string { return a.site.Domain() }

func (a *abstractContextGroup) naviPageDescription() string { return "" }

func (a *abstractContextGroup) naviPageTitle() string { return "" }

func (a *abstractContextGroup) naviPagePathFromDocRoot() string { return "" }

func (a *abstractContextGroup) Init() {}

func (a *abstractContextGroup) rss(targetDir string) fs.FileContainer {
	if len(a.context.GetPages()) > 0 {

		rssPath := a.context.SiteDto().Rss()
		rssFilename := "rss.xml"

		fc := fs.NewFileContainer()
		fc.SetPath(path.Join(targetDir, rssPath))
		fc.SetFilename(rssFilename)
		fc.SetDataAsString(a.rssContent())

		return fc
	}
	return nil
}

func (a *abstractContextGroup) getSingleRssEntry(p staticIntf.Page) string {
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

func (a *abstractContextGroup) rssContent() string {

	last10 := a.getLastPages(10)
	rss := []string{}
	for _, p := range last10 {
		rss = append(rss, a.getSingleRssEntry(p))
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
	domain := last10[0].Domain()
	date := last10[len(last10)-1].PublishedTime()
	return fmt.Sprintf(rssTemplate, domain, domain, domain, domain, domain, "rss.xml", date, itemsRss)
}

func (a *abstractContextGroup) getLastPages(nr int) []staticIntf.Page {
	pgs := a.context.GetPages()
	if len(pgs) > nr {
		return pgs[len(pgs)-nr:]
	}
	return pgs
}
