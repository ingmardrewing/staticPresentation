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

type subContext struct {
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
	fsSetOff        string
	pages           []staticIntf.Page
	components      []staticIntf.Component
	site            staticIntf.Site
}

func (c *subContext) SiteDto() staticIntf.Site {
	return c.site
}

func (c *subContext) GetPages() []staticIntf.Page {
	return c.pages
}

func (c *subContext) RenderPages() []fs.FileContainer {
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

func (c *subContext) FsSetOff(fsSetOff ...string) string {
	if len(fsSetOff) > 0 {
		c.fsSetOff = fsSetOff[0]
	}
	return c.fsSetOff
}

func (c *subContext) SetElements(pages []staticIntf.Page) {
	c.pages = pages
}

func (c *subContext) GetComponents() []staticIntf.Component {
	return c.components
}

func (c *subContext) GetElements() []staticIntf.Page {
	return c.pages
}

func (c *subContext) AddPage(p staticIntf.Page) {
	c.pages = append(c.pages, p)
}

func (c *subContext) AddComponent(comp staticIntf.Component) {
	c.components = append(c.components, comp)
}

func (c *subContext) AddComponents(comps ...staticIntf.Component) {
	for _, comp := range comps {
		comp.SetContext(c)
		c.AddComponent(comp)
	}
}

func (c *subContext) GetDisqusShortname() string {
	return c.disqusShortname
}

func (c *subContext) GetMainNavigationLocations() []staticIntf.Location {
	return c.site.Main()
}

func (c *subContext) GetFooterNavigationLocations() []staticIntf.Location {
	return c.site.Marginal()
}

func (c *subContext) GetCssUrl() string {
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

func (c *subContext) GetTwitterPage() string {
	return c.twitterPageUrl
}

func (c *subContext) GetFBPageUrl() string {
	return c.fbPageUrl
}

func (c *subContext) GetOGType() string {
	return c.ogType
}

func (c *subContext) GetTwitterCardType() string {
	return c.twitterCardType
}

func (c *subContext) GetTwitterHandle() string {
	return c.twitterHandle
}

func (c *subContext) GetContentSection() string {
	return c.contentSection
}

func (c *subContext) GetContentTags() string {
	return c.tags
}

func (c *subContext) GetSiteName() string {
	return c.siteName
}

func (c *subContext) GetCss() string {
	css := ""
	for _, c := range c.components {
		css += c.GetCss()
	}
	return c.minifyCss(css)
}

func (c *subContext) GetJs() string {
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

func (c *subContext) minifyCss(txt string) string {
	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	s, err := m.String("text/css", txt)
	if err != nil {
		panic(err)
	}
	return s
}

func (c *subContext) GetReadNavigationLocations() []staticIntf.Location {
	return nil
}

func NewContext(site staticIntf.Site) staticIntf.SubContext {
	c := new(subContext)
	c.site = site

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

// Create Narrrative Context
// used for graphic novels
func NewNarrativeMarginalContext(cd staticIntf.Site) staticIntf.SubContext {

	c := NewContext(cd)

	c.AddComponents(headerComponents...)
	c.AddComponents(
		NewTitleComponent(),
		NewNarrativeHeaderComponent(),
		NewPlainContentComponent(),
		NewNarrativeCopyRightComponent(),
		NewFooterNaviComponent())

	return c
}

// Create Narrrative Context
// used for graphic novels
func NewNarrativeContext(cd staticIntf.Site) staticIntf.SubContext {

	c := NewContext(cd)

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

// Create Narrrative Context
// used for graphic novels
func NewNarrativeArchiveContext(cd staticIntf.Site) staticIntf.SubContext {

	c := NewContext(cd)

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
func NewPagesContext(cd staticIntf.Site) staticIntf.SubContext {

	c := NewContext(cd)

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
func NewBlogContext(cd staticIntf.Site) staticIntf.SubContext {

	c := NewContext(cd)

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
func NewBlogNaviContext(cd staticIntf.Site) staticIntf.SubContext {

	c := NewContext(cd)

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
func NewMarginalContext(cd staticIntf.Site) staticIntf.SubContext {

	c := NewContext(cd)

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
