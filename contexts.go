package staticPresentation

import (
	"fmt"
	"math/rand"

	"github.com/ingmardrewing/fs"
	"github.com/ingmardrewing/staticIntf"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/js"
)

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
	mainNavigationLocations   []staticIntf.Location
	footerNavigationLocations []staticIntf.Location
	pages                     []staticIntf.Page
	components                []staticIntf.Component
}

func (c *ContextImpl) SetGlobalFields(
	twitterHandle,
	topic,
	tags,
	site,
	cardType,
	section,
	fbPage,
	twitterPage,
	cssUrl,
	rssUrl,
	home,
	disqusShortname string) {
	c.twitterHandle = twitterHandle
	c.contentSection = topic
	c.tags = tags
	c.siteName = site
	c.twitterCardType = cardType
	c.ogType = section
	c.fbPageUrl = fbPage
	c.twitterPageUrl = twitterPage
	c.cssUrl = cssUrl
	c.rssUrl = rssUrl
	c.homeUrl = home
	c.disqusShortname = disqusShortname
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
		path := targetDir + p.PathFromDocRoot()

		fc := fs.NewFileContainer()
		fc.SetPath(path)
		fc.SetFilename(p.HtmlFilename())

		fmt.Println(path, p.HtmlFilename())

		fc.SetDataAsString(html)
		fcs = append(fcs, fc)
	}
	return fcs
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
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
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

func fillContextWithComponents(context staticIntf.Context, components ...staticIntf.Component) {
	for _, compo := range components {
		compo.SetContext(context)
		context.AddComponent(compo)
	}
}

func newContext(mainnavi, footernavi []staticIntf.Location, contentComponents []staticIntf.Component) staticIntf.Context {
	c := new(ContextImpl)
	c.mainNavigationLocations = mainnavi
	c.footerNavigationLocations = footernavi

	fillContextWithComponents(c,
		NewGeneralMetaComponent(),
		NewFaviconComponent(),
		NewGlobalCssComponent(),
		NewGoogleComponent(),
		NewTwitterComponent(),
		NewFBComponent(),
		NewCssLinkComponent(),
		NewTitleComponent())

	fillContextWithComponents(c, contentComponents...)

	fillContextWithComponents(c,
		NewMainHeaderComponent(),
		NewMainNaviComponent(),
		NewCopyRightComponent(),
		NewFooterNaviComponent())

	return c
}

/* Pages Context */

func NewNarrativeContext(mainnavi, footernavi []staticIntf.Location) staticIntf.Context {

	c := new(ContextImpl)
	c.mainNavigationLocations = mainnavi
	c.footerNavigationLocations = footernavi

	fillContextWithComponents(c,
		NewGeneralMetaComponent(),
		NewFaviconComponent(),
		NewGlobalCssComponent(),
		NewGoogleComponent(),
		NewTwitterComponent(),
		NewFBComponent(),
		NewCssLinkComponent(),
		NewTitleComponent(),
		NewNarrativeHeaderComponent(),
		NewNarrativeComponent(),
		NewNarrativeNaviComponent(),
		NewDisqusComponent())

	return c
}

/* Pages Context */

func NewPagesContext(mainnavi, footernavi []staticIntf.Location) staticIntf.Context {

	c := new(ContextImpl)
	c.mainNavigationLocations = mainnavi
	c.footerNavigationLocations = footernavi

	fillContextWithComponents(c,
		NewGeneralMetaComponent(),
		NewFaviconComponent(),
		NewGlobalCssComponent(),
		NewGoogleComponent(),
		NewTwitterComponent(),
		NewFBComponent(),
		NewCssLinkComponent(),
		NewTitleComponent(),
		NewStartPageComponent(),
		NewMainHeaderComponent(),
		NewMainNaviComponent(),
		NewCopyRightComponent(),
		NewFooterNaviComponent())

	return c
}

/* Blog Context */

func NewBlogContext(mainnavi, footernavi []staticIntf.Location) staticIntf.Context {
	contentComponents := []staticIntf.Component{
		NewContentComponent(),
		NewDisqusComponent()}
	c := newContext(mainnavi, footernavi, contentComponents)
	return c
}

/* Footer Context */

func NewFooterContext(mainnavi, footernavi []staticIntf.Location) staticIntf.Context {
	contentComponents := []staticIntf.Component{
		NewContentComponent()}
	c := newContext(mainnavi, footernavi, contentComponents)
	return c
}

/* Blog Navi Context */

func NewBlogNaviContext(mainnavi, footernavi []staticIntf.Location) staticIntf.Context {
	contentComponents := []staticIntf.Component{
		NewBlogNaviContextComponent()}
	c := newContext(mainnavi, footernavi, contentComponents)
	return c
}
