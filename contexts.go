package staticPresentation

import (
	"fmt"
	"math/rand"
	"path"
	"strings"
	"time"

	"github.com/ingmardrewing/fs"
	"github.com/ingmardrewing/staticIntf"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/js"
)

var headerComponents []staticIntf.Component = []staticIntf.Component{
	NewGeneralMetaComponent(),
	NewFaviconComponent(),
	NewGlobalCssComponent(),
	NewGoogleComponent(),
	NewTwitterComponent(),
	NewFBComponent(),
	NewCssLinkComponent()}

/* Global Context */

type ContextImpl struct {
	twitterHandle             string
	contentSection            string
	tags                      string
	siteName                  string
	twitterCardType           string
	ogType                    string
	fbPageUrl                 string
	twitterPageUrl            string
	cssUrl                    string
	rssUrl                    string
	homeUrl                   string
	disqusShortname           string
	fsSetOff                  string
	addRss                    bool
	mainNavigationLocations   []staticIntf.Location
	footerNavigationLocations []staticIntf.Location
	pages                     []staticIntf.Page
	components                []staticIntf.Component
}

func (c *ContextImpl) SetContextDto(dto staticIntf.ContextDto) {
	c.twitterHandle = dto.TwitterHandle()
	c.contentSection = dto.Topic()
	c.tags = dto.Tags()
	c.siteName = dto.Site()
	c.twitterCardType = dto.CardType()
	c.ogType = dto.Section()
	c.fbPageUrl = dto.FBPage()
	c.twitterPageUrl = dto.TwitterPage()
	c.cssUrl = dto.Css()
	c.rssUrl = dto.Rss()
	c.homeUrl = dto.Domain()
	c.disqusShortname = dto.DisqusId()
}

func (c *ContextImpl) getSingleRssEntry(p staticIntf.Page) string {
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

func (c *ContextImpl) AddRss() {
	c.addRss = true
}

func (c *ContextImpl) rss() string {

	last10 := c.getLastPages(10)
	rss := []string{}
	for _, p := range last10 {
		rss = append(rss, c.getSingleRssEntry(p))
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

func (c *ContextImpl) getLastPages(nr int) []staticIntf.Page {
	if len(c.pages) > nr {
		return c.pages[len(c.pages)-nr:]
	}
	return c.pages
}

func (c *ContextImpl) RenderPages(targetDir string) []fs.FileContainer {
	fcs := []fs.FileContainer{}
	for _, p := range c.pages {

		for _, comp := range c.components {
			p.AcceptVisitor(comp)
		}
		doc := p.GetDoc()
		doc.AddRootAttr("itemscope")
		doc.AddRootAttr("lang", "en")
		html := doc.Render()
		path := path.Join(targetDir, c.FsSetOff(), p.PathFromDocRoot())

		fc := fs.NewFileContainer()
		fc.SetPath(path)
		fc.SetFilename(p.HtmlFilename())

		fc.SetDataAsString(html)
		fcs = append(fcs, fc)
	}
	if c.addRss && len(c.pages) > 0 {
		path := path.Join(targetDir, c.FsSetOff())
		fc := fs.NewFileContainer()
		fc.SetPath(path)
		fc.SetFilename("rss.xml")
		fc.SetDataAsString(c.rss())
		fcs = append(fcs, fc)
	}
	return fcs
}

func (c *ContextImpl) FsSetOff(fsSetOff ...string) string {
	if len(fsSetOff) > 0 {
		c.fsSetOff = fsSetOff[0]
	}
	return c.fsSetOff
}

func (c *ContextImpl) SetElements(pages []staticIntf.Page) {
	c.pages = pages
}

func (c *ContextImpl) GetComponents() []staticIntf.Component {
	return c.components
}

func (c *ContextImpl) GetElements() []staticIntf.Page {
	return c.pages
}

func (c *ContextImpl) AddPage(p staticIntf.Page) {
	c.pages = append(c.pages, p)
}

func (c *ContextImpl) AddComponent(comp staticIntf.Component) {
	c.components = append(c.components, comp)
}

func (c *ContextImpl) AddComponents(comps ...staticIntf.Component) {
	for _, comp := range comps {
		comp.SetContext(c)
		c.AddComponent(comp)
	}
}

func (c *ContextImpl) GetHomeUrl() string {
	return c.homeUrl
}

func (c *ContextImpl) GetRssUrl() string {
	return c.rssUrl
}

func (c *ContextImpl) GetDisqusShortname() string {
	return c.disqusShortname
}

func (c *ContextImpl) GetMainNavigationLocations() []staticIntf.Location {
	return c.mainNavigationLocations
}

func (c *ContextImpl) GetFooterNavigationLocations() []staticIntf.Location {
	return c.footerNavigationLocations
}

func (c *ContextImpl) GetCssUrl() string {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	n := 10
	b := make([]byte, n)

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	for i := range b {
		b[i] = letterBytes[r.Intn(len(letterBytes))]
	}
	pseudoParam := string(b)
	return c.cssUrl + "?" + pseudoParam
}

func (c *ContextImpl) GetTwitterPage() string {
	return c.twitterPageUrl
}

func (c *ContextImpl) GetFBPageUrl() string {
	return c.fbPageUrl
}

func (c *ContextImpl) GetOGType() string {
	return c.ogType
}

func (c *ContextImpl) GetTwitterCardType() string {
	return c.twitterCardType
}

func (c *ContextImpl) GetTwitterHandle() string {
	return c.twitterHandle
}

func (c *ContextImpl) GetContentSection() string {
	return c.contentSection
}

func (c *ContextImpl) GetContentTags() string {
	return c.tags
}

func (c *ContextImpl) GetSiteName() string {
	return c.siteName
}

func (c *ContextImpl) GetCss() string {
	css := ""
	for _, c := range c.components {
		css += c.GetCss()
	}
	return c.minifyCss(css)
}

func (c *ContextImpl) GetJs() string {
	jsCode := ""
	for _, c := range c.components {
		jsCode += c.GetJs()
	}

	m := minify.New()
	m.AddFunc("text/javascript", js.Minify)
	s, err := m.String("text/javascript", jsCode)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf(`<script>%s</script>`, s)
}

func (c *ContextImpl) minifyCss(txt string) string {
	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	s, err := m.String("text/css", txt)
	if err != nil {
		panic(err)
	}
	return s
}

func (c *ContextImpl) GetReadNavigationLocations() []staticIntf.Location {
	return nil
}

// Create Narrrative Context
// used for graphic novels
func NewNarrativeContext(mainnavi, footernavi []staticIntf.Location) staticIntf.Context {

	c := new(ContextImpl)
	c.mainNavigationLocations = mainnavi
	c.footerNavigationLocations = footernavi
	c.fsSetOff = "devabo.de/"

	c.AddComponents(headerComponents...)
	c.AddComponents(
		NewTitleComponent(),
		NewNarrativeHeaderComponent(),
		NewNarrativeComponent(),
		NewNarrativeNaviComponent(),
		NewDisqusComponent(),
		NewNarrativeCopyRightComponent(),
		NewFooterNaviComponent())

	return c
}

// Pages context, used for static pages
// of a site, featuring separate subjects
func NewPagesContext(mainnavi, footernavi []staticIntf.Location) staticIntf.Context {

	c := new(ContextImpl)
	c.mainNavigationLocations = mainnavi
	c.footerNavigationLocations = footernavi

	c.AddComponents(headerComponents...)
	c.AddComponents(
		NewTitleComponent(),
		NewStartPageComponent(),
		NewMainHeaderComponent(),
		NewMainNaviComponent(),
		NewCopyRightComponent(),
		NewFooterNaviComponent())

	return c
}

// Blog context, used for blog pages
func NewBlogContext(mainnavi, footernavi []staticIntf.Location) staticIntf.Context {

	c := new(ContextImpl)
	c.mainNavigationLocations = mainnavi
	c.footerNavigationLocations = footernavi

	c.AddComponents(headerComponents...)
	c.AddComponents(
		NewTitleComponent(),
		NewContentComponent(),
		NewDisqusComponent(),
		NewMainHeaderComponent(),
		NewMainNaviComponent(),
		NewCopyRightComponent(),
		NewFooterNaviComponent())

	return c
}

// Blog navigation context
// creates pages containing a navigations overview
// of blog pages
func NewBlogNaviContext(mainnavi, footernavi []staticIntf.Location) staticIntf.Context {

	c := new(ContextImpl)
	c.mainNavigationLocations = mainnavi
	c.footerNavigationLocations = footernavi

	c.AddComponents(headerComponents...)
	c.AddComponents(
		NewTitleComponent(),
		NewBlogNaviPageContentComponent(),
		NewBlogNaviContextComponent(),
		NewMainHeaderComponent(),
		NewMainNaviComponent(),
		NewCopyRightComponent(),
		NewFooterNaviComponent())

	return c
}

// Marginal context use for pages contained
// within the marginal navigation (imprint, terms of use, etc.)
func NewMarginalContext(mainnavi, footernavi []staticIntf.Location) staticIntf.Context {

	c := new(ContextImpl)
	c.mainNavigationLocations = mainnavi
	c.footerNavigationLocations = footernavi

	c.AddComponents(headerComponents...)
	c.AddComponents(
		NewTitleComponent(),
		NewContentComponent(),
		NewMainHeaderComponent(),
		NewMainNaviComponent(),
		NewCopyRightComponent(),
		NewFooterNaviComponent())

	return c
}
