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

	nav := htmlDoc.NewNode("nav", "", "class", "prevnextnavigation")
	pn := b.prevNode(p)
	nn := b.nextNode(p)
	if pn != nil {
		nav.AddChild(pn)
	}

	if nn != nil {
		nav.AddChild(nn)
	}

	wn := b.wrap(nav, "prevnextnavigation__wrapper")
	p.AddBodyNodes([]*htmlDoc.Node{wn})
}

func (b *BlogPrevNextNaviComponent) nextNode(p staticIntf.Page) *htmlDoc.Node {
	nextPage := b.nextPage(p)
	if nextPage == nil {
		return nil
	}
	nextLabel := htmlDoc.NewNode("span", "next post", "class", "label")
	nextIcon := htmlDoc.NewNode("span", "&gt;", "class", "icon")

	classes := "prevnextnavigation__next prevnextnavigation__item prevnextnavigation__placeholder"
	nextNode := b.getOtherNode(nextPage, classes)
	nextNode.AddChild(nextLabel)
	nextNode.AddChild(nextIcon)
	return nextNode
}

func (b *BlogPrevNextNaviComponent) prevNode(p staticIntf.Page) *htmlDoc.Node {
	prevPage := b.previousPage(p)
	if prevPage == nil {
		return nil
	}
	prevIcon := htmlDoc.NewNode("span", "&lt;", "class", "icon")
	prevLabel := htmlDoc.NewNode("span", "previous post", "class", "label")

	classes := "prevnextnavigation__previous prevnextnavigation__item prevnextnavigation__placeholder"
	prevNode := b.getOtherNode(prevPage, classes)
	prevNode.AddChild(prevIcon)
	prevNode.AddChild(prevLabel)
	return prevNode
}

func (b *BlogPrevNextNaviComponent) getOtherNode(otherPage staticIntf.Page, classes string) *htmlDoc.Node {
	if otherPage == nil {
		span := htmlDoc.NewNode("span", "", "class", classes)
		return span
	}

	a := htmlDoc.NewNode(
		"a", "",
		"class", classes,
		"href", otherPage.PathFromDocRootWithName())

	imageObj := otherPage.Images()[0]
	imgUrl := imageObj.W800Square()
	if len(imageObj.W200Square()) > 0 {
		imgUrl = imageObj.W200Square()
	} else if len(imageObj.W390Square()) > 0 {
		imgUrl = imageObj.W390Square()
	}

	img := htmlDoc.NewNode(
		"img", "",
		"src", imgUrl,
		"width", "200",
		"height", "200",
		"alt", "",
		"class", "prevnextnavigation__image")
	a.AddChild(img)
	return a
}

func (b *BlogPrevNextNaviComponent) GetCss() string {
	return `
.prevnextnavigation__image {
	display: none;
}
.prevnextnavigation__item {
	text-transform: uppercase;
	font-weight: 700;
	font-size: 16px;
}
.prevnextnavigation__item .icon{
	color: lightgrey;
	transition: 0.5s;
}
.prevnextnavigation__item .label{
	color: lightgrey;
	display: inline-block;
	margin-left: 5px;
	margin-right: 5px;
	transition: 0.5s;
}
.prevnextnavigation__item:hover {
	text-decoration: none;
}
.prevnextnavigation__item:hover .icon,
.prevnextnavigation__item:hover .label {
	color: grey;
}
.prevnextnavigation__item + .prevnextnavigation__item {
	margin-left: 10px;
}

@media only screen and (max-width: 1240px) {
	.prevnextnavigation{
		text-align: right;
		margin-bottom: 50px;
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
@media only screen and (min-width: 1240px) {
	.prevnextnavigation__image {
		height: 200px;
		width: 200px;
		display: block;
		opacity: 0.3;
		transition: 0.5s;
	}
	.prevnextnavigation__wrapper {
		position: absolute;
		top: -200px;
		height: 0;
	}
	.prevnextnavigation__item {
		position: fixed;
		top: 213px;
		font-weight: 700;
	}
	.prevnextnavigation__item:hover {
		color: black;
		text-decoration: none;
	}
	.prevnextnavigation__item:hover .prevnextnavigation__image{
		opacity: 1;
	}
	.prevnextnavigation__previous {
		left: 0;
	}
	.prevnextnavigation__next{
		right: 0;
	}
}
`
}
