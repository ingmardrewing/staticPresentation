package staticPresentation

import (
	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
)

// Generates navigational overview pages filled
// with thumbnails
func NewBlogNaviPageContentComponent(r staticIntf.Renderer) *BlogNaviPageContentComponent {
	bnpc := new(BlogNaviPageContentComponent)
	bnpc.abstractComponent.Renderer(r)
	return bnpc
}

type BlogNaviPageContentComponent struct {
	abstractComponent
	wrapper
}

func (b *BlogNaviPageContentComponent) VisitPage(p staticIntf.Page) {
	n := htmlDoc.NewNode(
		"div", "",
		"class", "blognavipagecomponent")

	for _, page := range p.NavigatedPages() {

		ta := page.ThumbnailUrl()
		if ta == "" {
			ta = page.ImageUrl()
		}
		a := htmlDoc.NewNode(
			"a", " ",
			"href", page.Link(),
			"class", "blognavientry__tile")

		a.AddChild(htmlDoc.NewNode(
			"h2", page.Title()))
		a.AddChild(htmlDoc.NewNode(
			"img", "",
			"src", page.ThumbnailUrl(),
			"class", "blognavientry__image"))

		n.AddChild(a)
	}
	n.AddChild(htmlDoc.NewNode("div", "", "style", "clear: both"))
	wn := b.wrap(n)
	p.AddBodyNodes([]*htmlDoc.Node{wn})
}

func (b *BlogNaviPageContentComponent) GetCss() string {
	return `
.blognavientry__tile {
	display: block;
}
.blognavientry__tile h2 {
	font-family: Arial Black, Arial, Helvetica, sans-serif;
	text-transform: uppercase;
	color: black;
	margin-bottom: 4px;
	line-height: 24px;
}
.blognavientry__tile:hover h2 {
	color: grey;
}
@media only screen and (max-width: 768px) {
	.blognavientry__image {
		max-width: 100%;
		height: auto;
	}
	.blognavientry__tile h2 {
		text-align: left;
		padding-left: 10px;
		font-size: 1.2em;
	}
}
@media only screen and (min-width: 769px) {
	.blognavipagecomponent {
		padding-top: 140px;
	}
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
		width: 390px;
		height: 390px;
	}
	.blognavientry__tile h2 {
		font-size: 1.2em;
		text-align: left;
	}
}
`
}
