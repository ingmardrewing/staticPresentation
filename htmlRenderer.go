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

func (r *renderer) getIndexOfPage(p staticIntf.Page) int {
	for i, l := range r.pages {
		lurl := l.PathFromDocRoot() + l.HtmlFilename()
		purl := p.PathFromDocRoot() + p.HtmlFilename()
		if lurl == purl {
			return i
		}
	}
	return -1
}

func (r *renderer) GetLastPage() staticIntf.Page {
	if len(r.pages) > 0 {
		return r.pages[len(r.pages)-1]
	}
	return nil
}

func (r *renderer) GetFirstPage() staticIntf.Page {
	if len(r.pages) > 0 {
		return r.pages[0]
	}
	return nil
}

func (r *renderer) GetPageBefore(p staticIntf.Page) staticIntf.Page {
	i := r.getIndexOfPage(p)
	if i > 0 {
		return r.pages[i-1]
	}
	return nil
}

func (r *renderer) GetPageAfter(p staticIntf.Page) staticIntf.Page {
	i := r.getIndexOfPage(p)
	if i+1 < len(r.pages) {
		return r.pages[i+1]
	}
	return nil
}

func (r *renderer) Render() []fs.FileContainer {
	targetDir := r.site.TargetDir()
	fcs := []fs.FileContainer{}

	for _, p := range r.pages {
		html := r.RenderHtml(p)
		path := path.Join(targetDir, p.PathFromDocRoot())
		fc := r.CreateFileContainer(html, path, p.HtmlFilename())
		fcs = append(fcs, fc)
	}

	return fcs
}

func (r *renderer) RenderHtml(p staticIntf.Page) string {
	for _, comp := range r.components {
		p.AcceptVisitor(comp)
	}
	doc := p.GetDoc()
	doc.AddRootAttr("itemscope")
	doc.AddRootAttr("lang", "en")
	return doc.Render()
}

func (r *renderer) CreateFileContainer(html, path, filename string) fs.FileContainer {
	fc := fs.NewFileContainer()
	fc.SetPath(path)
	fc.SetFilename(filename)
	fc.SetDataAsString(html)
	return fc
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
	return r
}

// Create Narrrative Margina Renderer
// used for marginal pages of graphic novels
func NewNarrativeMarginalRenderer(site staticIntf.Site) staticIntf.Renderer {

	r := NewRenderer(site, "Narrative Marginal Renderer")

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
func NewNarrativeRenderer(site staticIntf.Site) staticIntf.Renderer {

	r := NewRenderer(site, "Narrative Renderer")

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
func NewNarrativeArchiveRenderer(site staticIntf.Site) staticIntf.Renderer {
	r := NewRenderer(site, "Narrative Archive Renderer")

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
func NewPagesRenderer(site staticIntf.Site) staticIntf.Renderer {

	r := NewRenderer(site, "Pages Renderer")

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
func NewBlogRenderer(site staticIntf.Site) staticIntf.Renderer {

	r := NewRenderer(site, "Blog Renderer")

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
func NewPortfolioRenderer(site staticIntf.Site) staticIntf.Renderer {

	r := NewRenderer(site, "Portfolio Renderer")

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
func NewBlogNaviRenderer(site staticIntf.Site) staticIntf.Renderer {

	r := NewRenderer(site, "Blog Navi Renderer")

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
func NewMarginalRenderer(site staticIntf.Site) staticIntf.Renderer {

	r := NewRenderer(site, "Marginal Renderer")

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
func NewHomePageRenderer(site staticIntf.Site) staticIntf.Renderer {

	r := NewRenderer(site, "Entry Page Renderer")

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
