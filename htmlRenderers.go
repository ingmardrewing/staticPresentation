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

/* Global Renderer */

type renderer struct {
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

func (c *renderer) SiteDto() staticIntf.Site {
	return c.site
}

func (c *renderer) GetPagesObsolete() []staticIntf.Page {
	return c.pages
}

func (c *renderer) Render() []fs.FileContainer {
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

func (c *renderer) FsSetOff(fsSetOff ...string) string {
	if len(fsSetOff) > 0 {
		c.fsSetOff = fsSetOff[0]
	}
	return c.fsSetOff
}

func (c *renderer) SetPages(pages []staticIntf.Page) {
	c.pages = pages
}

func (c *renderer) GetComponents() []staticIntf.Component {
	return c.components
}

func (c *renderer) GetPages() []staticIntf.Page {
	return c.pages
}

func (c *renderer) AddPage(p staticIntf.Page) {
	c.pages = append(c.pages, p)
}

func (c *renderer) AddComponent(comp staticIntf.Component) {
	c.components = append(c.components, comp)
}

func (c *renderer) AddComponents(comps ...staticIntf.Component) {
	for _, comp := range comps {
		comp.Renderer(c)
		c.AddComponent(comp)
	}
}

func (c *renderer) GetDisqusShortname() string {
	return c.disqusShortname
}

func (c *renderer) GetMainNavigationLocations() []staticIntf.Location {
	return c.site.Main()
}

func (c *renderer) GetFooterNavigationLocations() []staticIntf.Location {
	return c.site.Marginal()
}

func (c *renderer) GetCssUrl() string {
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

func (c *renderer) GetTwitterPage() string {
	return c.twitterPageUrl
}

func (c *renderer) GetFBPageUrl() string {
	return c.fbPageUrl
}

func (c *renderer) GetOGType() string {
	return c.ogType
}

func (c *renderer) GetTwitterCardType() string {
	return c.twitterCardType
}

func (c *renderer) GetTwitterHandle() string {
	return c.twitterHandle
}

func (c *renderer) GetContentSection() string {
	return c.contentSection
}

func (c *renderer) GetContentTags() string {
	return c.tags
}

func (c *renderer) GetSiteName() string {
	return c.siteName
}

func (c *renderer) GetCss() string {
	css := ""
	for _, c := range c.components {
		css += c.GetCss()
	}
	return c.minifyCss(css)
}

func (c *renderer) GetJs() string {
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

func (c *renderer) minifyCss(txt string) string {
	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	s, err := m.String("text/css", txt)
	if err != nil {
		panic(err)
	}
	return s
}

func (c *renderer) GetReadNavigationLocations() []staticIntf.Location {
	return nil
}

func NewRenderer(site staticIntf.Site) staticIntf.Renderer {
	c := new(renderer)
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

// Create Narrrative Margina Renderer
// used for marginal pages of graphic novels
func NewNarrativeMarginalRenderer(cd staticIntf.Site) staticIntf.Renderer {

	c := NewRenderer(cd)

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

	c := NewRenderer(cd)

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

	c := NewRenderer(cd)

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

	c := NewRenderer(cd)

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

	c := NewRenderer(cd)

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

	c := NewRenderer(cd)

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

	c := NewRenderer(cd)

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

	c := NewRenderer(cd)

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
