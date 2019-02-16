package staticPresentation

import (
	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
)

// Creates anew BlogPrevNextNaviComponent
func NewBlogPrevNextNaviComponent(r staticIntf.Renderer) *BlogPrevNextNaviComponent {
	nc := new(BlogPrevNextNaviComponent)
	nc.abstractComponent.Renderer(r)
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
		"prevnextnavigation__previous prevnextnavigation__item prevnextnavigation__placeholder")
	prevNode.AddChild(prevIcon)
	prevNode.AddChild(prevLabel)

	nextLabel := htmlDoc.NewNode(
		"span", "next post",
		"class", "label")
	nextIcon := htmlDoc.NewNode(
		"span", "▶",
		"class", "icon")
	nextNode := b.nextFromDocRoot(p, "",
		"prevnextnavigation__next prevnextnavigation__item prevnextnavigation__placeholder")
	nextNode.AddChild(nextLabel)
	nextNode.AddChild(nextIcon)

	nav := htmlDoc.NewNode("nav", "", "class", "prevnextnavigation")
	nav.AddChild(prevNode)
	nav.AddChild(nextNode)

	wn := b.wrap(nav, "prevnextnavigation__wrapper")
	p.AddBodyNodes([]*htmlDoc.Node{wn})
}

func (b *BlogPrevNextNaviComponent) GetCss() string {
	return `
@media only screen and (max-width: 1080px) {
	.prevnextnavigation{
		text-align: right;
		margin-bottom: 50px;
	}
	.prevnextnavigation__item {
		color: grey;
		text-transform: uppercase;
		font-weight: 700;
		font-size: 16px;
	}
	.prevnextnavigation__item .label{
		display: inline-block;
		margin-left: 5px;
		margin-right: 5px;
	}
	.prevnextnavigation__item:hover {
		text-decoration: none;
	}
	span.prevnextnavigation__item {
		color: lightgrey;
	}
	.prevnextnavigation__item + .prevnextnavigation__item {
		margin-left: 10px;
	}
}
@media only screen and (max-width: 768px) {
	.prevnextnavigation{
		padding-top: 10px;
		padding-bottom: 10px;
		border-top: 1px solid black;
		border-bottom: 1px solid black;
		text-align: center;
	}
}
@media only screen and (min-width: 1081px) {
	.prevnextnavigation__wrapper {
		position: absolute;
		top: -200px;
		height: 0;
	}
	.prevnextnavigation__item {
		position: fixed;
		top: calc(50vh - 50px);
		color: lightgrey;
		font-weight: 700;
		font-size: 100px;
	}
	.prevnextnavigation__item .label {
		display: none;
	}
	span.prevnextnavigation__item {
		display: none;
	}
	.prevnextnavigation__item:hover {
		color: black;
		text-decoration: none;
	}
	.prevnextnavigation__previous {
		left: calc(50vw - 520px);
	}
	.prevnextnavigation__next{
		right: calc(50vw - 520px);
	}
}
`
}
