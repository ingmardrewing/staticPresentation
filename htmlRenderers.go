package staticPresentation

import (
	"math/rand"
	"path"
	"time"

	"github.com/ingmardrewing/fs"
	"github.com/ingmardrewing/staticIntf"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
)

func getHeaderComponents(r staticIntf.Renderer) []staticIntf.Component {
	return []staticIntf.Component{
		NewGeneralMetaComponent(r),
		NewFaviconComponent(r),
		NewGlobalCssComponent(r),
		NewGoogleComponent(r),
		NewTwitterComponent(r),
		NewFBComponent(r),
		NewCssLinkComponent(r)}
}

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

func (r *renderer) Render() []fs.FileContainer {
	targetDir := r.site.TargetDir()
	fcs := []fs.FileContainer{}
	for _, p := range r.pages {

		for _, comp := range r.components {
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

func (r *renderer) Components() []staticIntf.Component {
	return r.components
}

func (r *renderer) Pages(ps ...staticIntf.Page) []staticIntf.Page {
	if len(ps) > 0 {
		r.pages = ps
	}
	return r.pages
}

func (r *renderer) AddPage(p staticIntf.Page) {
	r.pages = append(r.pages, p)
}

func (r *renderer) addComponent(comp staticIntf.Component) {
	r.components = append(r.components, comp)
}

func (r *renderer) AddComponents(comps ...staticIntf.Component) {
	for _, comp := range comps {
		r.addComponent(comp)
	}
}

func (r *renderer) DisqusShortname() string {
	return r.disqusShortname
}

func (r *renderer) MainNavigationLocations() []staticIntf.Location {
	return r.site.Main()
}

func (r *renderer) FooterNavigationLocations() []staticIntf.Location {
	return r.site.Marginal()
}

func (r *renderer) CssUrl() string {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	n := 10
	b := make([]byte, n)

	s := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(s)
	for i := range b {
		b[i] = letterBytes[rnd.Intn(len(letterBytes))]
	}
	pseudoParam := string(b)
	return r.cssUrl + "?" + pseudoParam
}

func (r *renderer) TwitterPage() string {
	return r.twitterPageUrl
}

func (r *renderer) FBPageUrl() string {
	return r.fbPageUrl
}

func (r *renderer) OGType() string {
	return r.ogType
}

func (r *renderer) TwitterCardType() string {
	return r.twitterCardType
}

func (r *renderer) TwitterHandle() string {
	return r.twitterHandle
}

func (r *renderer) ContentSection() string {
	return r.contentSection
}

func (r *renderer) ContentTags() string {
	return r.tags
}

func (r *renderer) Site() staticIntf.Site {
	return r.site
}

func (r *renderer) SiteName() string {
	return r.siteName
}

func (r *renderer) Css() string {
	cssString := ""
	for _, r := range r.components {
		cssString += r.GetCss()
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
	r := new(renderer)
	r.site = site
	r.rendererName = rendererName

	r.twitterHandle = site.TwitterHandle()
	r.contentSection = site.Topic()
	r.tags = site.Tags()
	r.siteName = site.Site()
	r.twitterCardType = site.CardType()
	r.ogType = site.Section()
	r.fbPageUrl = site.FBPage()
	r.twitterPageUrl = site.TwitterPage()
	r.cssUrl = site.Css()
	r.disqusShortname = site.DisqusId()

	return r
}

// Create Narrrative Margina Renderer
// used for marginal pages of graphic novels
func NewNarrativeMarginalRenderer(cd staticIntf.Site) staticIntf.Renderer {

	r := NewRenderer(cd, "Narrative Marginal Renderer")

	r.AddComponents(getHeaderComponents(r)...)
	r.AddComponents(
		NewTitleComponent(r),
		NewNarrativeHeaderComponent(r),
		NewPlainContentComponent(r),
		NewNarrativeCopyRightComponent(r),
		NewFooterNaviComponent(r))

	return r
}

// Create Narrrative Renderer
// used for graphic novels
func NewNarrativeRenderer(cd staticIntf.Site) staticIntf.Renderer {

	r := NewRenderer(cd, "Narrative Renderer")

	r.AddComponents(getHeaderComponents(r)...)
	r.AddComponents(
		NewTitleComponent(r),
		NewNarrativeHeaderComponent(r),
		NewNarrativeComponent(r),
		NewNarrativeNaviComponent(r),
		NewNarrativeCopyRightComponent(r),
		NewFooterNaviComponent(r))

	return r
}

// Create Narrrative Renderer
// used for graphic novels
func NewNarrativeArchiveRename(cd staticIntf.Site) staticIntf.Renderer {

	r := NewRenderer(cd, "Narrative Archive Renderer")

	r.AddComponents(getHeaderComponents(r)...)
	r.AddComponents(
		NewTitleComponent(r),
		NewNarrativeHeaderComponent(r),
		NewNarrativeArchiveComponent(r),
		NewNarrativeCopyRightComponent(r),
		NewFooterNaviComponent(r))

	return r
}

// Pages context, used for static pages
// of a site, featuring separate subjects
func NewPagesRenderer(cd staticIntf.Site) staticIntf.Renderer {

	r := NewRenderer(cd, "Pages Renderer")

	r.AddComponents(getHeaderComponents(r)...)
	r.AddComponents(
		NewTitleComponent(r),
		NewStartPageComponent(r),
		NewMainHeaderComponent(r),
		NewMainNaviComponent(r),
		NewCopyRightComponent(r),
		NewFooterNaviComponent(r))

	return r
}

// Blog context, used for blog pages
func NewBlogRenderer(cd staticIntf.Site) staticIntf.Renderer {

	r := NewRenderer(cd, "Blog Renderer")

	r.AddComponents(getHeaderComponents(r)...)
	r.AddComponents(
		NewTitleComponent(r),
		NewMainHeaderComponent(r),
		NewMainNaviComponent(r),
		NewContentComponent(r),
		NewBlogPrevNextNaviComponent(r),
		NewCopyRightComponent(r),
		NewFooterNaviComponent(r))
	return r
}

// Blog context, used for blog pages
func NewPortfolioRenderer(cd staticIntf.Site) staticIntf.Renderer {

	r := NewRenderer(cd, "Portfolio Renderer")

	r.AddComponents(getHeaderComponents(r)...)
	r.AddComponents(
		NewTitleComponent(r),
		NewMainHeaderComponent(r),
		NewMainNaviComponent(r),
		NewContentComponent(r),
		NewCopyRightComponent(r),
		NewFooterNaviComponent(r))
	return r
}

// Blog navigation context
// creates pages containing a navigations overview
// of blog pages
func NewBlogNaviRenderer(cd staticIntf.Site) staticIntf.Renderer {

	r := NewRenderer(cd, "Blog Navi Renderer")

	r.AddComponents(getHeaderComponents(r)...)
	r.AddComponents(
		NewTitleComponent(r),
		NewMainHeaderComponent(r),
		NewMainNaviComponent(r),
		NewBlogNaviPageContentComponent(r),
		NewBlogNaviComponent(r),
		NewCopyRightComponent(r),
		NewFooterNaviComponent(r))

	return r
}

// Marginal context use for pages contained
// within the marginal navigation (imprint, terms of use, etc.)
func NewMarginalRenderer(cd staticIntf.Site) staticIntf.Renderer {

	r := NewRenderer(cd, "Marginal Renderer")

	r.AddComponents(getHeaderComponents(r)...)
	r.AddComponents(
		NewTitleComponent(r),
		NewMainHeaderComponent(r),
		NewMainNaviComponent(r),
		NewContentComponent(r),
		NewCopyRightComponent(r),
		NewFooterNaviComponent(r))

	return r
}

// Entry page renderer
func NewHomePageRenderer(cd staticIntf.Site) staticIntf.Renderer {

	r := NewRenderer(cd, "Entry Page Renderer")

	r.AddComponents(getHeaderComponents(r)...)
	r.AddComponents(
		NewTitleComponent(r),
		NewMainHeaderComponent(r),
		NewMainNaviComponent(r),
		NewHomePageComponent(r),
		NewCopyRightComponent(r),
		NewFooterNaviComponent(r))

	return r
}
