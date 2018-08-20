package staticPresentation

import (
	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
)

// Creates anew BlogPrevNextNaviComponent
func NewBlogPrevNextNaviComponent() *BlogPrevNextNaviComponent {
	nc := new(BlogPrevNextNaviComponent)
	return nc
}

type BlogPrevNextNaviComponent struct {
	abstractComponent
	wrapper
	cssClass string
}

func (b *BlogPrevNextNaviComponent) VisitPage(p staticIntf.Page) {
	prevNode := b.previousFromDocRoot(p,
		"&lt; previous page",
		"blogprevnextnavigation__previous blogprevnextnavigation__item blogprevnextnavigation__placeholder")
	nextNode := b.nextFromDocRoot(p,
		"next page &gt;",
		"blogprevnextnavigation__next blogprevnextnavigation__item blogprevnextnavigation__placeholder")

	nav := htmlDoc.NewNode("nav", "", "class", "blogprevnextnavigation")
	nav.AddChild(prevNode)
	nav.AddChild(nextNode)

	wn := b.wrap(nav, "blogprevnextnavigation__wrapper")
	p.AddBodyNodes([]*htmlDoc.Node{wn})
}

func (b *BlogPrevNextNaviComponent) GetCss() string {
	return `
.blogprevnextnavigation{
	text-align: right;
	margin-bottom: 50px;
}
.blogprevnextnavigation__item {
	font-family: Arial Black, Arial, Helvetica, sans-serif;
	color: grey;
	text-transform: uppercase;
	font-weight: 900;
	font-size: 16px;
}
span.blogprevnextnavigation__item {
	color: lightgrey;
}
.blogprevnextnavigation__item + .blogprevnextnavigation__item {
	margin-left: 10px;
}
`
}
