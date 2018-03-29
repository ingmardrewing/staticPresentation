package staticPresentation

import (
	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
)

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

	for _, page := range p.NavigatedPages() {

		ta := page.ThumbnailUrl()
		if ta == "" {
			ta = page.ImageUrl()
		}
		a := htmlDoc.NewNode("a", " ",
			"href", page.PathFromDocRootWithName(),
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
