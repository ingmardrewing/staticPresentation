package staticPresentation

import (
	"path"

	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
)

// Creates a new BlogNaviComponent
func NewBlogNaviComponent() *BlogNaviComponent {
	bnc := new(BlogNaviComponent)
	return bnc
}

type BlogNaviComponent struct {
	wrapper
	abstractComponent
}

func (b *BlogNaviComponent) addPrevious(p staticIntf.Page) *htmlDoc.Node {
	pageBefore := b.getPageBefore(p)
	if pageBefore == nil {
		return htmlDoc.NewNode("span", "< previous posts", "class", "blognavicomponent__previous")
	}
	href := path.Join(pageBefore.PathFromDocRoot(), pageBefore.HtmlFilename())
	return htmlDoc.NewNode("a", "< previous posts", "href", href, "rel", "prev", "class", "blognavicomponent__previous")
}

func (b *BlogNaviComponent) addNext(p staticIntf.Page) *htmlDoc.Node {
	pageAfter := b.getPageAfter(p)
	if pageAfter == nil {
		return htmlDoc.NewNode("span", "next posts >", "class", "blognavicomponent__next")
	}

	href := path.Join(pageAfter.PathFromDocRoot(), pageAfter.HtmlFilename())
	return htmlDoc.NewNode("a", "next posts >", "href", href, "rel", "next", "class", "blognavicomponent__next")
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
	if len(b.abstractComponent.renderer.Pages()) < 3 {
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
