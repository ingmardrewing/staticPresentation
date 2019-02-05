package staticPresentation

import (
	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
	"github.com/ingmardrewing/staticUtil"
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
	for _, navigated := range p.NavigatedPages() {
		n.AddChild(b.renderNavigatedPage(navigated))
	}
	n.AddChild(htmlDoc.NewNode("div", "", "style", "clear: both"))

	wn := b.wrap(n)
	p.AddBodyNodes([]*htmlDoc.Node{wn})
}

func (b *BlogNaviPageContentComponent) renderNavigatedPage(n staticIntf.Page) *htmlDoc.Node {
	a := htmlDoc.NewNode(
		"a", " ",
		"href", n.Link(),
		"class", "blognavientry__tile")

	a.AddChild(htmlDoc.NewNode(
		"h2", n.Title()))
	if len(staticUtil.MakeSrcSet(n)) > 5 {
		a.AddChild(htmlDoc.NewNode(
			"img", "",
			"src", n.ThumbnailUrl(),
			"srcset", staticUtil.MakeSrcSet(n),
			"class", "blognavientry__image"))
	} else {
		a.AddChild(htmlDoc.NewNode(
			"img", "",
			"src", n.ThumbnailUrl(),
			"class", "blognavientry__image"))
	}
	return a
}

func (b *BlogNaviPageContentComponent) GetCss() string {
	return `
.blognavientry__tile {
	box-sizing: border-box;
	display: block;
	overflow: hidden;
	width: 390px;
	height: 390px;
	max-height: 390px;
	position: relative;
}

.blognavientry__tile h2 {
	display: none;
}
.blognavientry__tile:hover h2 {
	display: absolute;
	bottom: 0;
	right: 0;
	font-weight: 700;
	text-transform: uppercase;
	color: black;
	margin-bottom: 4px;
	line-height: 24px;
	background-color: white;
	text-align: right;
}

@media only screen and (max-width: 768px) {
	.blognavientry__tile  {
		width: 100%;
		height: auto;
		max-height: none;
		text-align: center;
	}
	.blognavientry__image {
		display: block;
		margin: 0 auto;
	}
	.blognavientry__tile h2 {
		text-align: left;
		padding-left: 10px;
		font-size: 1.2em;
		text-transform: uppercase;
		display: block;
	}
}

@media only screen and (min-width: 769px) {
	.blognavipagecomponent {
		padding-top: 123px;
	}
	a.blognavientry__tile {
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
		/* height: 390px; */
	}
	.blognavientry__tile h2 {
		font-size: 1.2em;
		text-align: left;
	}
}
`
}
