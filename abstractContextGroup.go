package staticPresentation

import (
	"fmt"
	"path"
	"strconv"
	"strings"

	"github.com/ingmardrewing/fs"
	"github.com/ingmardrewing/staticIntf"
	"github.com/ingmardrewing/staticModel"
)

type abstractContextGroup struct {
	pages        []staticIntf.Page
	pagesContext staticIntf.Context
	commonData   staticIntf.CommonData
}

func (a *abstractContextGroup) GetComponents() []staticIntf.Component {
	return a.pagesContext.GetComponents()
}

func (a *abstractContextGroup) RenderPages(dir string) []fs.FileContainer {
	return a.pagesContext.RenderPages(dir)
}

func (a *abstractContextGroup) Domain() string {
	return ""
}

func (a *abstractContextGroup) naviPageDescription() string {
	return ""
}

func (a *abstractContextGroup) naviPageTitle() string {
	return ""
}

func (a *abstractContextGroup) naviPagePathFromDocRoot() string {
	return ""
}

//
type navigationalContextGroup struct {
	abstractContextGroup
	naviContext staticIntf.Context
}

func (n *navigationalContextGroup) GetComponents() []staticIntf.Component {
	components := n.pagesContext.GetComponents()
	return append(components, n.naviContext.GetComponents()...)
}

func (n *navigationalContextGroup) RenderPages(dir string) []fs.FileContainer {
	pages := n.pagesContext.RenderPages(dir)
	return append(pages, n.naviContext.RenderPages(dir)...)
}

func (n *navigationalContextGroup) Init() {
	bundles := n.generateBundles()
	last := len(bundles) - 1
	naviPages := []staticIntf.Page{}
	for i, bundle := range bundles {
		filename := "index" + strconv.Itoa(i) + ".html"
		if i == last {
			filename = "index.html"
		}

		np := staticModel.NewNaviPage()
		np.NavigatedPages(bundle...)
		np.Domain(n.Domain())
		np.Title(n.naviPageTitle())
		np.Description(n.naviPageDescription())
		np.HtmlFilename(filename)
		np.PathFromDocRoot(n.naviPagePathFromDocRoot())

		naviPages = append(naviPages, np)
	}
	n.naviContext.SetElements(naviPages)
}

func (n *navigationalContextGroup) getReversedPages() []staticIntf.Page {
	pages := n.pagesContext.GetElements()
	length := len(pages)
	reversed := []staticIntf.Page{}
	for i := length - 1; i >= 0; i-- {
		reversed = append(reversed, pages[i])
	}
	return reversed
}

func (n *navigationalContextGroup) generateReversedBundles() []*elementBundle {
	reversed := n.getReversedPages()
	b := newElementBundle()
	bundles := []*elementBundle{}
	for _, p := range reversed {
		b.addElement(p)
		if b.full() {
			bundles = append(bundles, b)
			b = newElementBundle()
		}
	}
	if !b.full() {
		bundles = append(bundles, b)
	}
	return bundles
}

func (n *navigationalContextGroup) generateBundles() [][]staticIntf.Page {
	reversedBundles := n.generateReversedBundles()
	length := len(reversedBundles)
	pageBundles := [][]staticIntf.Page{}
	for i := length - 1; i >= 0; i-- {
		pageBundles = append(pageBundles, reversedBundles[i].getElements())
	}
	return pageBundles
}

func (b *navigationalContextGroup) rss(targetDir string) fs.FileContainer {
	if len(b.pagesContext.GetPages()) > 0 {
		cd := b.pagesContext.CommonData()

		rssPath := cd.ContextDto().Rss()
		rssFilename := "rss.xml"
		rssLabel := "rss"

		url := path.Join(rssPath, rssFilename)
		l := staticModel.NewLocation(url, "", rssLabel, "", "", "")
		cd.AddMarginal(l)

		fc := fs.NewFileContainer()
		fc.SetPath(path.Join(targetDir, rssPath))
		fc.SetFilename(rssFilename)
		fc.SetDataAsString(b.rssContent())

		return fc
	}
	return nil
}

func (b *navigationalContextGroup) getSingleRssEntry(p staticIntf.Page) string {
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
  </item>
`
	return fmt.Sprintf(rssItem, p.Title(), p.Url(), p.PublishedTime(), p.Content(), p.Url(), p.Description(), p.ImageUrl(), p.ImageUrl(), p.ImageUrl(), p.Title())

}

func (b *navigationalContextGroup) rssContent() string {

	last10 := b.getLastPages(10)
	rss := []string{}
	for _, p := range last10 {
		rss = append(rss, b.getSingleRssEntry(p))
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

func (b *navigationalContextGroup) getLastPages(nr int) []staticIntf.Page {
	pgs := b.pagesContext.GetPages()
	if len(pgs) > nr {
		return pgs[len(pgs)-nr:]
	}
	return pgs
}
