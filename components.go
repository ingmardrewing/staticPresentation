package staticPresentation

import (
	"fmt"
	"path"
	"strings"

	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
)

// abstractComponent implementing default functions
// for implementing components
type abstractComponent struct {
	context staticIntf.Context
}

func (ac *abstractComponent) SetContext(context staticIntf.Context) {
	ac.context = context
}

func (ac *abstractComponent) GetCss() string { return "" }

func (ac *abstractComponent) GetJs() string { return "" }

func (ac *abstractComponent) VisitPage(p staticIntf.Page) {}

func (b *abstractComponent) getIndexOfPage(p staticIntf.Page) int {
	for i, l := range b.context.GetElements() {
		lurl := l.PathFromDocRoot() + l.HtmlFilename()
		purl := p.PathFromDocRoot() + p.HtmlFilename()
		if lurl == purl {
			return i
		}
	}
	return -1
}

func (b *abstractComponent) getFirstPage() staticIntf.Page {
	pages := b.context.GetElements()
	if len(pages) > 0 {
		return pages[0]
	}
	return nil
}

func (b *abstractComponent) getLastPage() staticIntf.Page {
	pages := b.context.GetElements()
	if len(pages) > 0 {
		return pages[len(pages)-1]
	}
	return nil
}

func (b *abstractComponent) getPageBefore(p staticIntf.Page) staticIntf.Page {
	index := b.getIndexOfPage(p)
	pages := b.context.GetElements()
	if index > 0 {
		return pages[index-1]
	}
	return nil
}

func (b *abstractComponent) getPageAfter(p staticIntf.Page) staticIntf.Page {
	index := b.getIndexOfPage(p)
	pages := b.context.GetElements()
	if index+1 < len(pages) {
		return pages[index+1]
	}
	return nil
}

// wrapper
type wrapper struct{}

func (cw *wrapper) wrap(n *htmlDoc.Node, addedclasses ...string) *htmlDoc.Node {
	inner := htmlDoc.NewNode("div", "", "class", "wrapperInner")
	inner.AddChild(n)
	classes := "wrapperOuter " + strings.Join(addedclasses, " ")
	wrapperNode := htmlDoc.NewNode("div", "", "class", classes)
	wrapperNode.AddChild(inner)
	return wrapperNode
}

/* global css component */

func NewGlobalCssComponent() *GlobalCssComponent {
	return new(GlobalCssComponent)
}

type GlobalCssComponent struct {
	abstractComponent
}

func (gcc *GlobalCssComponent) GetCss() string {
	return `
body, p, span {
	margin: 0;
	padding: 0;
	font-family: Arial, Helvetica, sans-serif;
}
a {
	color: grey;
	text-decoration: none;
}
a:hover {
	text-decoration: underline;
}
.wrapperOuter {
	text-align: center;
}

.wrapperInner {
	margin: 0 auto;
	width: 800px;
}
p + p {
	margin-top: 10px;
}
`
}

/* GeneralMetaComponent */

func NewGeneralMetaComponent() *GeneralMetaComponent {
	return new(GeneralMetaComponent)
}

type GeneralMetaComponent struct {
	abstractComponent
}

func (g *GeneralMetaComponent) VisitPage(p staticIntf.Page) {
	m := []*htmlDoc.Node{
		htmlDoc.NewNode("meta", "", "name", "viewport", "content", "width=device-width, initial-scale=1.0"),
		htmlDoc.NewNode("meta", "", "name", "robots", "content", "index,follow"),
		htmlDoc.NewNode("meta", "", "name", "author", "content", "Ingmar Drewing"),
		htmlDoc.NewNode("meta", "", "name", "publisher", "content", "Ingmar Drewing"),
		htmlDoc.NewNode("meta", "", "name", "keywords", "content", "storytelling, illustration, drawing, web comic, comic, cartoon, caricatures"),
		htmlDoc.NewNode("meta", "", "name", "DC.subject", "content", "storytelling, illustration, drawing, web comic, comic, cartoon, caricatures"),
		htmlDoc.NewNode("meta", "", "name", "page-topic", "content", "art"),
		htmlDoc.NewNode("meta", "", "charset", "UTF-8")}
	p.AddHeaderNodes(m)
}

/* favicon component */

func NewFaviconComponent() *FaviconComponent {
	return new(FaviconComponent)
}

type FaviconComponent struct {
	abstractComponent
}

func (f *FaviconComponent) VisitPage(p staticIntf.Page) {
	m := []*htmlDoc.Node{
		htmlDoc.NewNode("link", "", "rel", "icon", "type", "image/png", "sizes", "16x16", "href", "/icons/favicon-16x16.png"),
		htmlDoc.NewNode("link", "", "rel", "icon", "type", "image/png", "sizes", "32x32", "href", "/icons/favicon-32x32.png"),
		htmlDoc.NewNode("link", "", "rel", "icon", "type", "image/png", "sizes", "192x192", "href", "/icons/android-192x192.png"),
		htmlDoc.NewNode("link", "", "rel", "apple-touch-icon", "type", "image/png", "sizes", "180x180", "href", "/icons/apple-touch-icon-180x180.png"),
		htmlDoc.NewNode("meta", "", "name", "msapplication-config", "content", "/icons/browserconfig.xml")}
	p.AddHeaderNodes(m)
}

/* fb component */
type FBComponent struct {
	abstractComponent
}

func NewFBComponent() *FBComponent {
	fb := new(FBComponent)
	return fb
}

func (fbc *FBComponent) VisitPage(p staticIntf.Page) {
	m := []*htmlDoc.Node{
		htmlDoc.NewNode("meta", "", "property", "og:title", "content", p.Title()),
		htmlDoc.NewNode("meta", "", "property", "og:url", "content", p.PathFromDocRoot()+p.HtmlFilename()),
		htmlDoc.NewNode("meta", "", "property", "og:image", "content", p.ImageUrl()),
		htmlDoc.NewNode("meta", "", "property", "og:description", "content", p.Description()),
		htmlDoc.NewNode("meta", "", "property", "og:site_name", "content", fbc.abstractComponent.context.GetSiteName()),
		htmlDoc.NewNode("meta", "", "property", "og:type", "content", fbc.abstractComponent.context.GetOGType()),
		htmlDoc.NewNode("meta", "", "property", "article:published_time", "content", p.PublishedTime()),
		htmlDoc.NewNode("meta", "", "property", "article:modified_time", "content", p.PublishedTime()),
		htmlDoc.NewNode("meta", "", "property", "article:section", "content", fbc.abstractComponent.context.GetContentSection()),
		htmlDoc.NewNode("meta", "", "property", "article:tag", "content", fbc.abstractComponent.context.GetContentTags())}

	p.AddHeaderNodes(m)
}

/* google component */

type GoogleComponent struct {
	abstractComponent
}

func NewGoogleComponent() *GoogleComponent {
	gc := new(GoogleComponent)
	return gc
}

func (goo *GoogleComponent) VisitPage(p staticIntf.Page) {
	m := []*htmlDoc.Node{
		htmlDoc.NewNode("meta", "", "itemprop", "name", "content", p.Title()),
		htmlDoc.NewNode("meta", "", "itemprop", "description", "content", p.Description()),
		htmlDoc.NewNode("meta", "", "itemprop", "image", "content", p.ImageUrl())}
	p.AddHeaderNodes(m)
}

/* twitter component */

type TwitterComponent struct {
	abstractComponent
}

func NewTwitterComponent() *TwitterComponent {
	t := new(TwitterComponent)
	return t
}

func (tw *TwitterComponent) VisitPage(p staticIntf.Page) {
	m := []*htmlDoc.Node{
		htmlDoc.NewNode("meta", "",
			"name", "t:card",
			"content", tw.abstractComponent.context.GetTwitterCardType()),
		htmlDoc.NewNode("meta", "",
			"name", "t:site",
			"content", tw.abstractComponent.context.GetTwitterHandle()),
		htmlDoc.NewNode("meta", "",
			"name", "t:title",
			"content", p.Title()),
		htmlDoc.NewNode("meta", "",
			"name", "t:text:description",
			"content", p.Description()),
		htmlDoc.NewNode("meta", "",
			"name", "t:creator",
			"content", tw.abstractComponent.context.GetTwitterHandle()),
		htmlDoc.NewNode("meta", "",
			"name", "t:image",
			"content", p.ImageUrl())}
	p.AddHeaderNodes(m)
}

/* title component */
type TitleComponent struct {
	abstractComponent
}

func NewTitleComponent() *TitleComponent {
	return new(TitleComponent)
}

func (tc *TitleComponent) VisitPage(p staticIntf.Page) {
	title := htmlDoc.NewNode("title", p.Title())
	p.AddHeaderNodes([]*htmlDoc.Node{title})
}

/* css link component */

type CssLinkComponent struct {
	abstractComponent
}

func NewCssLinkComponent() *CssLinkComponent {
	clc := new(CssLinkComponent)
	return clc
}

func (clc *CssLinkComponent) VisitPage(p staticIntf.Page) {
	link := htmlDoc.NewNode("link", "", "href", clc.abstractComponent.context.GetCssUrl(), "rel", "stylesheet", "type", "text/css")
	p.AddHeaderNodes([]*htmlDoc.Node{link})
}

/**/

type BlogNaviComponent struct {
	wrapper
	abstractComponent
}

func NewBlogNaviContextComponent() *BlogNaviComponent {
	bnc := new(BlogNaviComponent)
	return bnc
}

func (b *BlogNaviComponent) addPrevious(p staticIntf.Page) *htmlDoc.Node {
	pageBefore := b.getPageBefore(p)
	if pageBefore == nil {
		return htmlDoc.NewNode("span", "< previous posts", "class", "blognavicomponent__previous")
	}
	return htmlDoc.NewNode("a", "< previous posts", "href", pageBefore.PathFromDocRoot()+pageBefore.HtmlFilename(), "rel", "prev", "class", "blognavicomponent__previous")
}

func (b *BlogNaviComponent) addNext(p staticIntf.Page) *htmlDoc.Node {
	pageAfter := b.getPageAfter(p)
	if pageAfter == nil {
		return htmlDoc.NewNode("span", "next posts >", "class", "blognavicomponent__next")
	}
	return htmlDoc.NewNode("a", "next posts >", "href", pageAfter.PathFromDocRoot()+pageAfter.HtmlFilename(), "rel", "next", "class", "blognavicomponent__next")
}

func (b *BlogNaviComponent) addBodyNodes(p staticIntf.Page) {
	nav := htmlDoc.NewNode("nav", "", "class", "blognavicomponent__nav")

	prev := b.addPrevious(p)
	nav.AddChild(prev)

	next := b.addNext(p)
	nav.AddChild(next)

	d := htmlDoc.NewNode("div", "", "class", "blognavicomponent meta")
	d.AddChild(htmlDoc.NewNode("div", p.Content()))
	d.AddChild(nav)
	wn := b.wrap(d)
	p.AddBodyNodes([]*htmlDoc.Node{wn})
}

func (b *BlogNaviComponent) VisitPage(p staticIntf.Page) {
	if len(b.abstractComponent.context.GetElements()) < 3 {
		return
	}
	b.addBodyNodes(p)
}

func (b *BlogNaviComponent) GetCss() string {
	return `
.blognavicomponent {
	text-align: left;
	padding-top: 123px;
}
.blognavicomponent.meta {
	padding-top: 0;
}
.blognavicomponent__nav {
	text-align: center;
	color: lightgrey;
	margin-bottom: 50px;
}
.blognavicomponent__nav span {
	font-family: Arial Black, Arial, Helvetica, sans-serif;
	color: lightgrey;
	font-weight: 900;
}
.blognavicomponent__next {
	margin-left: 10px;
}
.blognavicomponent__previous,
.blognavicomponent__next {
	font-family: Arial Black, Arial, Helvetica, sans-serif;
	color: grey;
	text-transform: uppercase;
	font-weight: 900;
	font-size: 16px;
}
`
}

/* NarrativeNaviComponent */
func NewNarrativeNaviComponent() *NarrativeNaviComponent {
	nc := new(NarrativeNaviComponent)
	return nc
}

type NarrativeNaviComponent struct {
	abstractComponent
	wrapper
	cssClass string
}

func (nv *NarrativeNaviComponent) VisitPage(p staticIntf.Page) {
	firstNode := nv.first(p)
	prevNode := nv.previous(p)
	nextNode := nv.next(p)
	lastNode := nv.last(p)

	nav := htmlDoc.NewNode("nav", "", "class", "narrativenavigation")
	nav.AddChild(firstNode)
	nav.AddChild(prevNode)
	nav.AddChild(nextNode)
	nav.AddChild(lastNode)

	wn := nv.wrap(nav, "narrativenavi__wrapper")
	p.AddBodyNodes([]*htmlDoc.Node{wn})
}

func (nv *NarrativeNaviComponent) first(p staticIntf.Page) *htmlDoc.Node {
	fPage := nv.getFirstPage()
	if fPage == nil || fPage.Id() == p.Id() {
		return htmlDoc.NewNode("span", "&lt;&lt; first page", "class", "narrativenavigation__first narrativenavigation__item narrativenavigation__placeholder")
	}
	href := path.Join(fPage.PathFromDocRoot(), fPage.HtmlFilename())
	return htmlDoc.NewNode("a", "&lt;&lt; first page", "href", href, "rel", "first", "class", "narrativenavigation__first narrativenavigation__item")
}

func (nv *NarrativeNaviComponent) last(p staticIntf.Page) *htmlDoc.Node {
	lPage := nv.getLastPage()
	if lPage == nil || lPage.Id() == p.Id() {
		return htmlDoc.NewNode("span", "last page &gt;&gt;", "class", "narrativenavigation__last narrativenavigation__item narrativenavigation__placeholder")
	}
	href := path.Join(lPage.PathFromDocRoot(), lPage.HtmlFilename())
	return htmlDoc.NewNode("a", "last page &gt;&gt;", "href", href, "rel", "last", "class", "narrativenavigation__last narrativenavigation__item")
}

func (nv *NarrativeNaviComponent) previous(p staticIntf.Page) *htmlDoc.Node {
	pageB := nv.getPageBefore(p)
	if pageB == nil {
		return htmlDoc.NewNode("span", "&lt; previous page", "class", "narrativenavigation__previous narrativenavigation__item narrativenavigation__placeholder")
	}
	href := path.Join(pageB.PathFromDocRoot(), pageB.HtmlFilename())
	return htmlDoc.NewNode("a", "&lt; previous page", "href", href, "rel", "prev", "class", "narrativenavigation__previous narrativenavigation__item")
}

func (nv *NarrativeNaviComponent) next(p staticIntf.Page) *htmlDoc.Node {
	pageA := nv.getPageAfter(p)
	if pageA == nil {
		return htmlDoc.NewNode("span", "next page &gt;", "class", "narrativenavigation__next narrativenavigation__item narrativenavigation__placeholder")
	}
	href := path.Join(pageA.PathFromDocRoot(), pageA.HtmlFilename())
	return htmlDoc.NewNode("a", "next page &gt;", "href", href, "rel", "next", "class", "narrativenavigation__next narrativenavigation__item")
}

func (mhc *NarrativeNaviComponent) GetCss() string {
	return `
.narrativenavigation{
	text-align: right;
	margin-bottom: 50px;
}
.narrativenavigation__item {
	font-family: Arial Black, Arial, Helvetica, sans-serif;
	color: grey;
	text-transform: uppercase;
	font-weight: 900;
	font-size: 16px;
}
.narrativenavigation__item.narrativenavigation__placeholder {
	color: lightgrey;
}
.narrativenavigation__item + .narrativenavigation__item {
	margin-left: 10px;
}
`
}

/* NarrativeHeaderComponent */
func NewNarrativeHeaderComponent() *NarrativeHeaderComponent {
	nc := new(NarrativeHeaderComponent)
	return nc
}

type NarrativeHeaderComponent struct {
	abstractComponent
	wrapper
	cssClass string
}

func (nv *NarrativeHeaderComponent) VisitPage(p staticIntf.Page) {
	a1 := htmlDoc.NewNode("a", "<!-- Devabo.de-->", "href", "https://devabo.de", "class", "home")
	a2 := htmlDoc.NewNode("a", "New Reader? Start here!", "href", "https://devabo.de/2013/08/01/a-step-in-the-dark/", "class", "orange")
	h1 := htmlDoc.NewNode("h1", p.Title(), "class", "maincontent__h1")

	n := htmlDoc.NewNode("header", "")
	n.AddChild(a1)
	n.AddChild(a2)
	n.AddChild(h1)

	wn := nv.wrap(n, "header__wrapper")
	p.AddBodyNodes([]*htmlDoc.Node{wn})
}

func (mhc *NarrativeHeaderComponent) GetCss() string {
	return `header .home {
    display: block;
    line-height: 80px;
    height: 30px;
    width: 800px;
    text-align: left;
    color: rgb(0, 0, 0);
    margin-bottom: 0px;
    margin-top: 0px;
    background: url(https://devabo.de/imgs/header_devabo_de.png) 0px 0px no-repeat transparent;
}

header .orange {
    display: block;
    height: 2.2em;
    background-color: rgb(255, 136, 0);
    color: rgb(255, 255, 255);
    line-height: 1em;
    box-sizing: border-box;
    width: 100%;
    font-size: 24px;
    font-family: "Arial Black";
    text-transform: uppercase;
    margin-bottom: 1rem;
    padding: 0.5em;
    text-decoration: underline;
}

header {
	text-align: left;
}
`
}

/* MainNaviComponent */
func NewMainNaviComponent() *MainNaviComponent {
	nc := new(MainNaviComponent)
	return nc
}

type MainNaviComponent struct {
	abstractComponent
	wrapper
	cssClass string
}

func (nv *MainNaviComponent) VisitPage(p staticIntf.Page) {
	nav := htmlDoc.NewNode("nav", "",
		"class", "mainnavi")
	for _, l := range nv.abstractComponent.context.GetMainNavigationLocations() {
		if p.Url() == l.Url() {
			span := htmlDoc.NewNode("span", l.Title(),
				"class", "mainnavi__navelement--current")
			nav.AddChild(span)
		} else {
			a := htmlDoc.NewNode("a", l.Title(),
				"href", l.Url(),
				"class", "mainnavi__navelement")
			nav.AddChild(a)
		}
	}
	node := htmlDoc.NewNode("div", "", "class", nv.cssClass)
	node.AddChild(nav)
	wn := nv.wrap(node, "mainnavi__wrapper")
	p.AddBodyNodes([]*htmlDoc.Node{wn})
}

func (mhc *MainNaviComponent) GetCss() string {
	return `
.mainnavi {
	border-top: 1px solid black;
	border-bottom: 2px solid black;
}
.mainnavi__wrapper {
	position: fixed;
	width: 100%;
	top: 80px;
	background-color: white;
}
.mainnavi__navelement--current,
a.mainnavi__navelement {
	display: inline-block;
	font-family: Arial Black, Arial, Helvetica, sans-serif;
	font-weight: 900;
	font-size: 18px;
	line-height: 20px;
	text-transform: uppercase;
	color: black;
	padding: 10px 20px;
}
.mainnavi__navelement--current,
a.mainnavi__navelement:hover {
	text-decoration: none;
	color: gray;
}
.mainnavi__nav {
	border-bottom: 2px solid black;
}
`
}

/* FooterNaviComponent */

func NewFooterNaviComponent() *FooterNaviComponent {
	nc := new(FooterNaviComponent)
	return nc
}

type FooterNaviComponent struct {
	abstractComponent
	wrapper
	cssClass string
}

func (f *FooterNaviComponent) VisitPage(p staticIntf.Page) {
	nav := htmlDoc.NewNode("nav", "",
		"class", "footernavi")
	for _, l := range f.abstractComponent.context.GetFooterNavigationLocations() {
		if p.Url() == l.Url() {
			span := htmlDoc.NewNode("span", l.Title(),
				"class", "footernavi__navelement--current")
			nav.AddChild(span)
		} else {
			a := htmlDoc.NewNode("a", l.Title(),
				"href", l.Url(),
				"class", "footernavi__navelement")
			nav.AddChild(a)
		}
	}
	node := htmlDoc.NewNode("div", "", "class", f.cssClass)
	node.AddChild(nav)
	wn := f.wrap(node, "footernavi__wrapper")
	p.AddBodyNodes([]*htmlDoc.Node{wn})
}

func (f *FooterNaviComponent) GetCss() string {
	return `
.footernavi {
	border-top: 1px solid black;
}
.footernavi__wrapper {
	position: fixed;
	width: 100%;
	bottom: 0;
	background-color: white;
}
.footernavi__navelement--current ,
a.footernavi__navelement {
	display: inline-block;
	font-family: Arial Black, Arial, Helvetica, sans-serif;
	font-weight: 900;
	font-size: 16px;
	line-height: 20px;
	text-transform: uppercase;
	text-decoration: none;
	color: black;
	padding: 10px 15px;
}
a.footernavi__navelement:hover,
.footernavi__navelement--current {
	color: gray;
}
`
}

/* disqus component */

type DisqusComponent struct {
	abstractComponent
	wrapper
	configuredJs string
}

func NewDisqusComponent() *DisqusComponent {
	d := new(DisqusComponent)
	return d
}

func (dc *DisqusComponent) GetCss() string {
	return `
.disqus,
.disqus p {
	font-family: Arial, Helvetica, sans-serif;
}
`
}

func (dc *DisqusComponent) GetJs() string {
	return dc.configuredJs
}

func (dc *DisqusComponent) VisitPage(p staticIntf.Page) {
	dc.configuredJs = fmt.Sprintf(`var disqus_config = function () { this.page.title= "%s"; this.page.url = '%s'; this.page.identifier =  '%s'; }; (function() { var d = document, s = d.createElement('script'); s.src = 'https://%s.disqus.com/embed.js'; s.setAttribute('data-timestamp', +new Date()); (d.head || d.body).appendChild(s); })();`, p.Title(), p.Domain()+p.PathFromDocRoot()+p.HtmlFilename(), p.DisqusId(), dc.abstractComponent.context.GetDisqusShortname())
	n := htmlDoc.NewNode("div", " ", "id", "disqus_thread", "class", "disqus")
	js := htmlDoc.NewNode("script", dc.configuredJs)
	wn := dc.wrap(n)
	p.AddBodyNodes([]*htmlDoc.Node{wn, js})
}

/* main  header component */

type MainHeaderComponent struct {
	abstractComponent
	wrapper
}

func NewMainHeaderComponent() *MainHeaderComponent {
	mhc := new(MainHeaderComponent)
	return mhc
}

func (mhc *MainHeaderComponent) VisitPage(p staticIntf.Page) {
	logo := htmlDoc.NewNode("a", "<!-- logo -->",
		"href", mhc.abstractComponent.context.GetHomeUrl(),
		"class", "headerbar__logo")
	logocontainer := htmlDoc.NewNode("div", "",
		"class", "headerbar__logocontainer")
	logocontainer.AddChild(logo)

	header := htmlDoc.NewNode("header", "", "class", "headerbar")
	header.AddChild(logocontainer)

	wn := mhc.wrap(header, "headerbar__wrapper")
	p.AddBodyNodes([]*htmlDoc.Node{wn})
}

func (mhc *MainHeaderComponent) GetCss() string {
	return `
.headerbar__wrapper {
	position: fixed;
	width: 100%;
	top: 0;
	background-color: white;
}
.headerbar__logo {
	background-image: url(https://s3.amazonaws.com/drewingdeblog/drewing_de_logo.png);
	background-repeat: no-repeat;
	background-position: center center;
	display: block;
	width: 100%;
	height: 80px;
}
.headerbar__navelement {
	display: inline-block;
	font-family: Arial Black, Arial, Helvetica, sans-serif;
	font-weight: 900;
	font-size: 18px;
	line-height: 20px;
	text-transform: uppercase;
	text-decoration: none;
	color: black;
	padding: 10px 20px;
}
`
}

/* start page component */
type NarrativeComponent struct {
	abstractComponent
	wrapper
}

func NewNarrativeComponent() *NarrativeComponent {
	return new(NarrativeComponent)
}

func (cc *NarrativeComponent) VisitPage(p staticIntf.Page) {
	img := htmlDoc.NewNode("img", "", "src", p.ImageUrl(), "width", "800")
	n := htmlDoc.NewNode("main", "", "class", "mainnarrativecontent")
	n.AddChild(img)

	wn := cc.wrap(n)
	p.AddBodyNodes([]*htmlDoc.Node{wn})
}

/* start page component */
type StartPageComponent struct {
	abstractComponent
	wrapper
}

func NewStartPageComponent() *StartPageComponent {
	return new(StartPageComponent)
}

func (cc *StartPageComponent) VisitPage(p staticIntf.Page) {
	c1 := htmlDoc.NewNode("div", "portfoliocontent", "class", "home__portfolio")
	c2 := htmlDoc.NewNode("div", "devabode", "class", "home__devabode")
	c3 := htmlDoc.NewNode("div", "blog", "class", "home__blog")

	n := htmlDoc.NewNode("main", "", "class", "maincontent")
	n.AddChild(c1)
	n.AddChild(c2)
	n.AddChild(c3)

	wn := cc.wrap(n)
	p.AddBodyNodes([]*htmlDoc.Node{wn})
}

/* content component */
type ContentComponent struct {
	abstractComponent
	wrapper
}

func NewContentComponent() *ContentComponent {
	return new(ContentComponent)
}

func (cc *ContentComponent) VisitPage(p staticIntf.Page) {
	h1 := htmlDoc.NewNode("h1", p.Title(),
		"class", "maincontent__h1")
	h2 := htmlDoc.NewNode("h2", p.PublishedTime(),
		"class", "maincontent__h2")
	n := htmlDoc.NewNode("main", p.Content(),
		"class", "maincontent")
	n.AddChild(h1)
	n.AddChild(h2)
	wn := cc.wrap(n)
	p.AddBodyNodes([]*htmlDoc.Node{wn})
}

func (cc *ContentComponent) GetCss() string {
	return `
.maincontent{
	padding-top: 126px;
	padding-bottom: 50px;
	text-align: left;
	line-height: 20px;
}
.maincontent li,
.maincontent p {
	line-height: 30px;
}
.maincontent h1,
.maincontent h2 {
	text-transform: uppercase;
}
.maincontent__h1,
.maincontent__h2 {
	display: inline-block;
	font-family: Arial Black, Arial, Helvetica, sans-serif;
	text-transform: uppercase;
}
.maincontent__h1 ,
.maincontent__h2 {
	font-size: 18px;
	line-height: 20px;
	text-transform: uppercase;
}
.maincontent__h2 {
	color: grey;
	margin-left: 10px;
}
`
}

/* gallery component */

type GalleryComponent struct {
	wrapper
}

func NewGalleryComponent() *GalleryComponent {
	gc := new(GalleryComponent)
	return gc
}

func (gal *GalleryComponent) VisitPage(p staticIntf.Page) {
	inner := htmlDoc.NewNode("div", "", "class", "maincontent__inner")
	for i := 0; i < 5; i++ {
		title := htmlDoc.NewNode("span", "At The Zoo", "class", "portfoliothumb__title")
		subtitle := htmlDoc.NewNode("span", "Digital drawing", "class", "portfoliothumb__details")

		label := htmlDoc.NewNode("div", "", "class", "portfoliothumb__label")
		label.AddChild(title)
		label.AddChild(subtitle)

		img := htmlDoc.NewNode("img", "", "class", "portfoliothumb__image", "src", "https://s3.amazonaws.com/drewingdeblog/blog/wp-content/uploads/2017/12/02152842/atthezoo-400x400.png")

		div := htmlDoc.NewNode("a", "", "class", "portfoliothumb", "href", "https://drewing.de")
		div.AddChild(img)
		div.AddChild(label)

		inner.AddChild(div)
	}

	m := htmlDoc.NewNode("main", "", "class", "maincontent")
	m.AddChild(inner)
	wn := gal.wrap(m)
	p.AddBodyNodes([]*htmlDoc.Node{wn})
}

/* copyright component */
type CopyRightComponent struct {
	abstractComponent
	wrapper
}

func NewCopyRightComponent() *CopyRightComponent {
	c := new(CopyRightComponent)
	return c
}

func (crc *CopyRightComponent) VisitPage(p staticIntf.Page) {
	n := htmlDoc.NewNode("div", `<a rel="license" class="copyright__cc" href="https://creativecommons.org/licenses/by-nc-nd/3.0/"></a><p class="copyright__license">&copy; 2017 by Ingmar Drewing </p><p class="copyright__license">Except where otherwise noted, content on this site is licensed under a <a rel="license" href="https://creativecommons.org/licenses/by-nc-nd/3.0/">Creative Commons Attribution-NonCommercial-NoDerivs 3.0 Unported (CC BY-NC-ND 3.0) license</a>.</p><p class="copyright__license">Soweit nicht anders explizit ausgewiesen, stehen die Inhalte auf dieser Website unter der <a rel="license" href="https://creativecommons.org/licenses/by-nc-nd/3.0/">Creative Commons Namensnennung-NichtKommerziell-KeineBearbeitung (CC BY-NC-ND 3.0)</a> Lizenz. Unless otherwise noted the author of the content on this page is <a href="https://plus.google.com/113943655600557711368?rel=author">Ingmar Drewing</a></p>`, "class", "copyright")
	wn := crc.wrap(n)
	p.AddBodyNodes([]*htmlDoc.Node{wn})
}

func (crc *CopyRightComponent) GetCss() string {
	return `
.copyright {
	background-color: black;
	color: white;
	text-align: left;
	font-family: Arial, Helvetica, sans-serif;
	font-size: 14px;
	padding: 20px 20px 50px;
	margin-top: 20px;
}
.copyright__license {
	margin-top: 20px;
	margin-bottom: 20px;
}
.copyright__cc {
    display: block;
    border-width: 0;
    background-image: url(https://i.creativecommons.org/l/by-nc-nd/3.0/88x31.png);
    width: 88px;
    height: 31px;
    margin-right: 15px;
    margin-bottom: 5px;
}
`
}

// NarrativeCopyRightComponent
type NarrativeCopyRightComponent struct {
	abstractComponent
	wrapper
}

func NewNarrativeCopyRightComponent() *NarrativeCopyRightComponent {
	c := new(NarrativeCopyRightComponent)
	return c
}

func (crc *NarrativeCopyRightComponent) VisitPage(p staticIntf.Page) {
	n := htmlDoc.NewNode("div", `All content including but not limited to the art, characters, story, website design & graphics are © copyright 2013-2018 Ingmar Drewing unless otherwise stated. All rights reserved. Do not copy, alter or reuse without expressed written permission.`, "class", "copyright")
	wn := crc.wrap(n)
	p.AddBodyNodes([]*htmlDoc.Node{wn})
}

func (crc *NarrativeCopyRightComponent) GetCss() string {
	return `
.copyright {
	background-color: black;
	color: white;
	text-align: left;
	font-family: Arial, Helvetica, sans-serif;
	font-size: 14px;
	padding: 20px 20px 50px;
	margin-top: 20px;
}
.copyright__license {
	margin-top: 20px;
	margin-bottom: 20px;
}
.copyright__cc {
    display: block;
    border-width: 0;
    background-image: url(https://i.creativecommons.org/l/by-nc-nd/3.0/88x31.png);
    width: 88px;
    height: 31px;
    margin-right: 15px;
    margin-bottom: 5px;
}
`
}

/* cookie notifier component */

type CookieNotifierComponent struct{}

func (cnc *CookieNotifierComponent) getJs() string {
	return `
function cli_show_cookiebar(p) {
	var Cookie = {
		set: function(name,value,days) {
			if (days) {
				var date = new Date();
				date.setTime(date.getTime()+(days*24*60*60*1000));
				var expires = "; expires="+date.toGMTString();
			}
			else var expires = "";
			document.cookie = name+"="+value+expires+"; path=/";
		},
		read: function(name) {
			var nameEQ = name + "=";
			var ca = document.cookie.split(';');
			for(var i=0;i < ca.length;i++) {
				var c = ca[i];
				while (c.charAt(0)==' ') {
					c = c.substring(1,c.length);
				}
				if (c.indexOf(nameEQ) === 0) {
					return c.substring(nameEQ.length,c.length);
				}
			}
			return null;
		},
		erase: function(name) {
			this.set(name,"",-1);
		},
		exists: function(name) {
			return (this.read(name) !== null);
		}
	};

	var ACCEPT_COOKIE_NAME = 'viewed_cookie_policy',
		ACCEPT_COOKIE_EXPIRE = 365,
		json_payload = p.settings;

	if (typeof JSON.parse !== "function") {
		console.log("CookieLawInfo requires JSON.parse but your browser doesn't support it");
		return;
	}
	var settings = JSON.parse(json_payload);

	var cached_header = jQuery(settings.notify_div_id),
		cached_showagain_tab = jQuery(settings.showagain_div_id),
		btn_accept = jQuery('#cookie_hdr_accept'),
		btn_decline = jQuery('#cookie_hdr_decline'),
		btn_moreinfo = jQuery('#cookie_hdr_moreinfo'),
		btn_settings = jQuery('#cookie_hdr_settings');

	cached_header.hide();
	if ( !settings.showagain_tab ) {
		cached_showagain_tab.hide();
	}

	var hdr_args = { };

	var showagain_args = { };
	cached_header.css( hdr_args );
	cached_showagain_tab.css( showagain_args );

	if (!Cookie.exists(ACCEPT_COOKIE_NAME)) {
		displayHeader();
	}
	else {
		cached_header.hide();
	}

	if ( settings.show_once_yn ) {
		setTimeout(close_header, settings.show_once);
	}
	function close_header() {
		Cookie.set(ACCEPT_COOKIE_NAME, 'yes', ACCEPT_COOKIE_EXPIRE);
		hideHeader();
	}

	var main_button = jQuery('.cli-plugin-main-button');
	main_button.css( 'color', settings.button_1_link_colour );

	if ( settings.button_1_as_button ) {
		main_button.css('background-color', settings.button_1_button_colour);

		main_button.hover(function() {
			jQuery(this).css('background-color', settings.button_1_button_hover);
		},
		function() {
			jQuery(this).css('background-color', settings.button_1_button_colour);
		});
	}
	var main_link = jQuery('.cli-plugin-main-link');
	main_link.css( 'color', settings.button_2_link_colour );

	if ( settings.button_2_as_button ) {
		main_link.css('background-color', settings.button_2_button_colour);

		main_link.hover(function() {
			jQuery(this).css('background-color', settings.button_2_button_hover);
		},
		function() {
			jQuery(this).css('background-color', settings.button_2_button_colour);
		});
	}

	cached_showagain_tab.click(function(e) {
		e.preventDefault();
		cached_showagain_tab.slideUp(settings.animate_speed_hide, function slideShow() {
			cached_header.slideDown(settings.animate_speed_show);
		});
	});

	jQuery("#cookielawinfo-cookie-delete").click(function() {
		Cookie.erase(ACCEPT_COOKIE_NAME);
		return false;
	});
	jQuery("#cookie_action_close_header").click(function(e) {
		e.preventDefault();
		accept_close();
	});

	function accept_close() {
		Cookie.set(ACCEPT_COOKIE_NAME, 'yes', ACCEPT_COOKIE_EXPIRE);

		if (settings.notify_animate_hide) {
			cached_header.slideUp(settings.animate_speed_hide);
		}
		else {
			cached_header.hide();
		}
		cached_showagain_tab.slideDown(settings.animate_speed_show);
		return false;
	}

	function closeOnScroll() {
		if (window.pageYOffset > 100 && !Cookie.read(ACCEPT_COOKIE_NAME)) {
			accept_close();
			if (settings.scroll_close_reload === true) {
				location.reload();
			}
			window.removeEventListener("scroll", closeOnScroll, false);
		}
	}
	if (settings.scroll_close === true) {
		window.addEventListener("scroll", closeOnScroll, false);
	}

	function displayHeader() {
		if (settings.notify_animate_show) {
			cached_header.slideDown(settings.animate_speed_show);
		}
		else {
			cached_header.show();
		}
		cached_showagain_tab.hide();
	}
	function hideHeader() {
		if (settings.notify_animate_show) {
			cached_showagain_tab.slideDown(settings.animate_speed_show);
		}
		else {
			cached_showagain_tab.show();
		}
		cached_header.slideUp(settings.animate_speed_show);
	}
};

function l1hs(str){if(str.charAt(0)=="#"){str=str.substring(1,str.length);}else{return "#"+str;}return l1hs(str);}

cli_show_cookiebar({
					settings: '{"animate_speed_hide":"500","animate_speed_show":"500","background":"#fff","border":"#444","border_on":true,"button_1_button_colour":"#000","button_1_button_hover":"#000000","button_1_link_colour":"#fff","button_1_as_button":true,"button_2_button_colour":"#333","button_2_button_hover":"#292929","button_2_link_colour":"#444","button_2_as_button":false,"font_family":"inherit","header_fix":false,"notify_animate_hide":true,"notify_animate_show":false,"notify_div_id":"#cookie-law-info-bar","notify_position_horizontal":"right","notify_position_vertical":"bottom","scroll_close":false,"scroll_close_reload":false,"showagain_tab":false,"showagain_background":"#fff","showagain_border":"#000","showagain_div_id":"#cookie-law-info-again","showagain_x_position":"100px","text":"#000","show_once_yn":false,"show_once":"10000"}'
});

`
}

// Generates navigational overview pages filled
// with thumbnails

func NewBlogNaviPageContentComponent() *BlogNaviPageContentComponent {
	bnpc := new(BlogNaviPageContentComponent)
	return bnpc
}

type BlogNaviPageContentComponent struct {
	abstractComponent
	wrapper
}

func (b *BlogNaviPageContentComponent) VisitPage(p staticIntf.Page) {
	n := htmlDoc.NewNode("div", "", "class", "blognavicomponent")

	for _, page := range p.(staticIntf.NaviPage).NavigatedPages() {

		ta := page.ThumbnailUrl()
		if ta == "" {
			ta = page.ImageUrl()
		}

		a := htmlDoc.NewNode("a", " ",
			"href", "/"+page.PathFromDocRoot()+page.HtmlFilename(),
			"class", "blognavientry__tile")
		span := htmlDoc.NewNode("span", " ",
			"style", "background-image: url("+page.ThumbnailUrl()+")",
			"class", "blognavientry__image")
		h2 := htmlDoc.NewNode("h2", page.Title())
		a.AddChild(span)
		a.AddChild(h2)
		n.AddChild(a)
	}
	n.AddChild(htmlDoc.NewNode("div", "", "style", "clear: both"))
	wn := b.wrap(n)
	p.AddBodyNodes([]*htmlDoc.Node{wn})
}

func (b *BlogNaviPageContentComponent) GetCss() string {
	return `
a.blognavientry__tile {
	display: block;
	position: relative;
	width: 390px;
	height: 470px;
	margin-bottom: 20px;
	float: left;
	text-decoration: none;
}

.blognavientry__tile:nth-child(odd) {
	margin-right: 20px;
}

.blognavientry__image {
	display: block;
	width: 390px;
	height: 390px;
	background-size: cover;
}
.blognavientry__tile h2 {
	font-family: Arial Black, Arial, Helvetica, sans-serif;
	text-transform: uppercase;
	color: black;
	margin-top: 4px;
	line-height: 24px;
}
.blognavientry__tile:hover h2 {
	color: grey;
}
`
}
