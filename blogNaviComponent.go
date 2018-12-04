package staticPresentation

import (
	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
)

// Creates a new BlogNaviComponent
func NewBlogNaviComponent(r staticIntf.Renderer) *BlogNaviComponent {
	bnc := new(BlogNaviComponent)
	bnc.abstractComponent.Renderer(r)
	return bnc
}

type BlogNaviComponent struct {
	wrapper
	abstractComponent
}

func (b *BlogNaviComponent) addBodyNodes(p staticIntf.Page) {

	prevLabel := htmlDoc.NewNode(
		"span", "previous posts",
		"class", "label")
	prevIcon := htmlDoc.NewNode(
		"span", "◀",
		"class", "icon")
	prev := b.previous(p, "", "blognavicomponent__previous blognavicomponent__item")
	prev.AddChild(prevIcon)
	prev.AddChild(prevLabel)

	nextLabel := htmlDoc.NewNode(
		"span", "next posts",
		"class", "label")
	nextIcon := htmlDoc.NewNode(
		"span", "▶",
		"class", "icon")
	next := b.next(p, " ", "blognavicomponent__next blognavicomponent__item")
	next.AddChild(nextLabel)
	next.AddChild(nextIcon)

	nav := htmlDoc.NewNode("nav", "", "class", "blognavicomponent__nav")
	nav.AddChild(prev)
	nav.AddChild(next)

	d := htmlDoc.NewNode("div", "", "class", "blognavicomponent meta")
	d.AddChild(htmlDoc.NewNode("div", p.Content()))
	d.AddChild(nav)

	wn := b.wrap(d, "blognavi__wrapper")
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
@media only screen and (max-width: 1080px) {
	.blognavi__wrapper {
		padding-top: 10px;
		padding-bottom: 12px;
		margin-top: 30px;
	}
	.blognavicomponent {
		padding-top: 123px;
	}
	.blognavicomponent.meta {
		padding-top: 0;
	}
	.blognavicomponent__nav {
		text-align: right;
		color: lightgrey;
	}
	.blognavicomponent__nav span.blognavicomponent__item {
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
	.blognavientry__tile h2{
		padding: 10px;
		margin: 0;
	}
	.blognavicomponent__item .label {
		display: inline-block ;
		margin-left: 5px;
		margin-right: 5px;
	}
}
@media only screen and (max-width: 768px) {
	.blognavi__wrapper {
		border-top: 1px solid black;
		border-bottom: 1px solid black;
	}
	.blognavientry__tile + .blognavientry__tile h2{
		border-top: 1px solid black;
	}
	.blognavicomponent__nav {
		text-align: center;
	}
}
@media only screen and (min-width: 1081px) {
	.blognavicomponent__item {
		position: fixed;
		top: calc(50vh - 50px);
		font-size: 100px;
		font-weight: 900;
		color: lightgrey;
	}
	.blognavicomponent__item:hover {
		color: black;
		text-decoration: none;
	}
	span.blognavicomponent__item {
		display: none;
	}
	.blognavicomponent__next {
		right: calc(50vw - 520px);
	}
	.blognavicomponent__previous {
		left: calc(50vw - 520px);
	}
	.blognavicomponent__item .label {
		display: none;
	}
}
`
}
