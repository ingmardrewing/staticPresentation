package staticPresentation

import (
	"fmt"
	"math/rand"
	"path"
	"time"

	"github.com/ingmardrewing/fs"
	"github.com/ingmardrewing/staticIntf"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
)

var headerComponents []staticIntf.Component = []staticIntf.Component{
	NewGeneralMetaComponent(),
	NewFaviconComponent(),
	NewGlobalCssComponent(),
	NewGoogleComponent(),
	NewTwitterComponent(),
	NewFBComponent(),
	NewCssLinkComponent()}

/* Global Renderer */

type renderer struct {
	rendererName    string
	twitterHandle   string
	contentSection  string
	tags            string
	siteName        string
	twitterCardType string
	ogType          string
	fbPageUrl       string
	twitterPageUrl  string
	cssUrl          string
	disqusShortname string
	pages           []staticIntf.Page
	components      []staticIntf.Component
	site            staticIntf.Site
}

func (c *renderer) Render() []fs.FileContainer {
	fmt.Printf("Called Render() on %s\n", c.rendererName)
	targetDir := c.site.TargetDir()
	fcs := []fs.FileContainer{}
	for _, p := range c.pages {

		for _, comp := range c.components {
			p.AcceptVisitor(comp)
		}
		doc := p.GetDoc()
		doc.AddRootAttr("itemscope")
		doc.AddRootAttr("lang", "en")
		html := doc.Render()
		path := path.Join(targetDir, p.PathFromDocRoot())

		fc := fs.NewFileContainer()
		fc.SetPath(path)
		fc.SetFilename(p.HtmlFilename())

		fc.SetDataAsString(html)
		fcs = append(fcs, fc)
	}
	return fcs
}

func (c *renderer) Components() []staticIntf.Component {
	return c.components
}

func (c *renderer) Pages(ps ...staticIntf.Page) []staticIntf.Page {
	if len(ps) > 0 {
		c.pages = ps
	}
	return c.pages
}

func (c *renderer) AddPage(p staticIntf.Page) {
	c.pages = append(c.pages, p)
}

func (c *renderer) addComponent(comp staticIntf.Component) {
	c.components = append(c.components, comp)
}

func (c *renderer) AddComponents(comps ...staticIntf.Component) {
	for _, comp := range comps {
		comp.Renderer(c)
		c.addComponent(comp)
	}
}

func (c *renderer) DisqusShortname() string {
	return c.disqusShortname
}

func (c *renderer) MainNavigationLocations() []staticIntf.Location {
	return c.site.Main()
}

func (c *renderer) FooterNavigationLocations() []staticIntf.Location {
	return c.site.Marginal()
}

func (c *renderer) CssUrl() string {
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

func (c *renderer) TwitterPage() string {
	return c.twitterPageUrl
}

func (c *renderer) FBPageUrl() string {
	return c.fbPageUrl
}

func (c *renderer) OGType() string {
	return c.ogType
}

func (c *renderer) TwitterCardType() string {
	return c.twitterCardType
}

func (c *renderer) TwitterHandle() string {
	return c.twitterHandle
}

func (c *renderer) ContentSection() string {
	return c.contentSection
}

func (c *renderer) ContentTags() string {
	return c.tags
}

func (c *renderer) SiteName() string {
	return c.siteName
}

func (c *renderer) Css() string {
	cssString := ""
	for _, c := range c.components {
		cssString += c.GetCss()
	}
	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	s, err := m.String("text/css", cssString)
	if err != nil {
		panic(err)
	}
	return s
}

func NewRenderer(site staticIntf.Site, rendererName string) staticIntf.Renderer {
	c := new(renderer)
	c.site = site
	c.rendererName = rendererName

	c.twitterHandle = site.TwitterHandle()
	c.contentSection = site.Topic()
	c.tags = site.Tags()
	c.siteName = site.Site()
	c.twitterCardType = site.CardType()
	c.ogType = site.Section()
	c.fbPageUrl = site.FBPage()
	c.twitterPageUrl = site.TwitterPage()
	c.cssUrl = site.Css()
	c.disqusShortname = site.DisqusId()

	return c
}

// Create Narrrative Margina Renderer
// used for marginal pages of graphic novels
func NewNarrativeMarginalRenderer(cd staticIntf.Site) staticIntf.Renderer {

	c := NewRenderer(cd, "Narrative Marginal Renderer")

	c.AddComponents(headerComponents...)
	c.AddComponents(
		NewTitleComponent(),
		NewNarrativeHeaderComponent(),
		NewPlainContentComponent(),
		NewNarrativeCopyRightComponent(),
		NewFooterNaviComponent())

	return c
}

// Create Narrrative Renderer
// used for graphic novels
func NewNarrativeRenderer(cd staticIntf.Site) staticIntf.Renderer {

	c := NewRenderer(cd, "Narrative Renderer")

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

// Create Narrrative Renderer
// used for graphic novels
func NewNarrativeArchiveRename(cd staticIntf.Site) staticIntf.Renderer {

	c := NewRenderer(cd, "Narrative Archive Renderer")

	c.AddComponents(headerComponents...)
	c.AddComponents(
		NewTitleComponent(),
		NewNarrativeHeaderComponent(),
		NewNarrativeArchiveComponent(),
		NewNarrativeCopyRightComponent(),
		NewFooterNaviComponent())

	return c
}

// Pages context, used for static pages
// of a site, featuring separate subjects
func NewPagesRenderer(cd staticIntf.Site) staticIntf.Renderer {

	c := NewRenderer(cd, "Pages Renderer")

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
func NewBlogRenderer(cd staticIntf.Site) staticIntf.Renderer {

	c := NewRenderer(cd, "Blog Renderer")

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
func NewBlogNaviRenderer(cd staticIntf.Site) staticIntf.Renderer {

	c := NewRenderer(cd, "Blog Navi Renderer")

	c.AddComponents(headerComponents...)
	c.AddComponents(
		NewTitleComponent(),
		NewBlogNaviPageContentComponent(),
		NewBlogNaviComponent(),
		NewMainHeaderComponent(),
		NewMainNaviComponent(),
		NewCopyRightComponent(),
		NewFooterNaviComponent())

	return c
}

// Marginal context use for pages contained
// within the marginal navigation (imprint, terms of use, etc.)
func NewMarginalRenderer(cd staticIntf.Site) staticIntf.Renderer {

	c := NewRenderer(cd, "Marginal Renderer")

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

// Entry page renderer
func NewEntryPageRenderer(cd staticIntf.Site) staticIntf.Renderer {

	c := NewRenderer(cd, "Entry Page Renderer")

	c.AddComponents(headerComponents...)
	c.AddComponents(
		NewTitleComponent(),
		NewContentComponent(),
		NewMainHeaderComponent(),
		NewMainNaviComponent(),
		NewEntryPageComponent(),
		NewCopyRightComponent(),
		NewFooterNaviComponent())

	return c
}
