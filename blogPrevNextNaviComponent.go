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
	prevLabel := htmlDoc.NewNode(
		"span", "previous post",
		"class", "label")
	prevIcon := htmlDoc.NewNode(
		"span", "◀",
		"class", "icon")
	prevNode := b.previousFromDocRoot(p, "",
		"blogprevnextnavigation__previous blogprevnextnavigation__item blogprevnextnavigation__placeholder")
	prevNode.AddChild(prevIcon)
	prevNode.AddChild(prevLabel)

	nextLabel := htmlDoc.NewNode(
		"span", "next post",
		"class", "label")
	nextIcon := htmlDoc.NewNode(
		"span", "▶",
		"class", "icon")
	nextNode := b.nextFromDocRoot(p, "",
		"blogprevnextnavigation__next blogprevnextnavigation__item blogprevnextnavigation__placeholder")
	nextNode.AddChild(nextLabel)
	nextNode.AddChild(nextIcon)

	nav := htmlDoc.NewNode("nav", "", "class", "blogprevnextnavigation")
	nav.AddChild(prevNode)
	nav.AddChild(nextNode)

	wn := b.wrap(nav, "blogprevnextnavigation__wrapper")
	p.AddBodyNodes([]*htmlDoc.Node{wn})
}

func (b *BlogPrevNextNaviComponent) GetCss() string {
	return `
@media only screen and (max-width: 1080px) {
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
	.blogprevnextnavigation__item .label{
		display: inline-block;
		margin-left: 5px;
		margin-right: 5px;
	}
	.blogprevnextnavigation__item:hover {
		text-decoration: none;
	}
	span.blogprevnextnavigation__item {
		color: lightgrey;
	}
	.blogprevnextnavigation__item + .blogprevnextnavigation__item {
		margin-left: 10px;
	}
}
@media only screen and (min-width: 1081px) {
	.blogprevnextnavigation__wrapper {
		position: absolute;
		top: -200px;
		height: 0;
	}
	.blogprevnextnavigation__item {
		position: fixed;
		top: calc(50vh - 50px);
		font-family: Arial Black, Arial, Helvetica, sans-serif;
		color: lightgrey;
		font-weight: 900;
		font-size: 100px;
	}
	.blogprevnextnavigation__item .label {
		display: none;
	}
	span.blogprevnextnavigation__item {
		display: none;
	}
	.blogprevnextnavigation__item:hover {
		color: grey;
		text-decoration: none;
	}
	.blogprevnextnavigation__previous {
		left: 10px;
	}
	.blogprevnextnavigation__next{
		right : 10px;
	}
}
`
}
