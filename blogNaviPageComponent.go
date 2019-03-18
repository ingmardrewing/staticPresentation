package staticPresentation

import (
	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
	"github.com/ingmardrewing/staticUtil"
)

// Generates navigational overview in the
// content area of the page. The Component
// places thumbnail elements into a grid and
// links them to the pages or posts represented
// by them.
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
		"class", "blogNaviPageComponent__grid")
	for _, navigated := range p.NavigatedPages() {
		n.AddChild(b.renderNavigatedPage(navigated))
	}

	wn := b.wrap(n)
	p.AddBodyNodes([]*htmlDoc.Node{wn})
}

func (b *BlogNaviPageContentComponent) renderNavigatedPage(n staticIntf.Page) *htmlDoc.Node {
	a := htmlDoc.NewNode(
		"a", " ",
		"title", n.Title(),
		"href", n.Link(),
		"class", "blogNaviPageComponent__gridItem")
	a.AddChild(b.createItemLabel(n))
	a.AddChild(b.createImage(n))
	return a
}

func (b *BlogNaviPageContentComponent) createItemLabel(n staticIntf.Page) *htmlDoc.Node {
	return htmlDoc.NewNode(
		"h2", n.Title(),
		"class", "blogNaviPageComponent__gridItemTitle")
}

func (b *BlogNaviPageContentComponent) createImage(n staticIntf.Page) *htmlDoc.Node {
	if len(staticUtil.MakeSrcSet(n)) > 5 {
		return htmlDoc.NewNode(
			"img", "",
			"src", n.ThumbnailUrl(),
			"alt", n.Title(),
			"width", "390",
			"height", "390",
			"srcset", staticUtil.MakeSrcSet(n),
			"class", "blogNaviPageComponent__image")
	}
	return htmlDoc.NewNode(
		"img", "",
		"src", n.ThumbnailUrl(),
		"alt", n.Title(),
		"width", "390",
		"height", "390",
		"class", "blogNaviPageComponent__image")
}

func (b *BlogNaviPageContentComponent) GetCss() string {
	return `
/* BlogNaviPageContentComponent start */
.blogNaviPageComponent__grid {
	display: grid;
	grid-gap: 20px;
}

.blogNaviPageComponent__image {
	max-height: 390px;
	max-width: 390px;
	-webkit-transition: opacity 0.5s;
    -moz-transition: opacity 0.5s;
    -o-transition: opacity 0.5s;
    transition: opacity 0.5s;
}

.blogNaviPageComponent__gridItem:hover .blogNaviPageComponent__image {
	opacity: 0.3;
}
@media only screen and (max-width: 768px) {
	.blogNaviPageComponent__grid {
		grid-template-columns: 100%;
		margin-top: 20px;
	}
}

@media only screen and (min-width: 769px) {
	.blogNaviPageComponent__grid {
		grid-template-columns: 390px 390px;
	}
}

.blogNaviPageComponent__gridItem:hover,
.blogNaviPageComponent__gridItem {
	text-decoration: none;
	display: block;
	position: relative;
	overflow: hidden;
	max-height: 390px;
}

@media only screen and (max-width: 768px) {
	.blogNaviPageComponent__gridItem {
		margin-bottom: 0;
	}
}

.blogNaviPageComponent__gridItemTitle{
	background-color:rgba(255, 255, 255, 0.8);
	text-align: left;
	font-size: 18px;
	font-weight: 700;
	text-transform: uppercase;
	color: rgb(0,0,0);
	margin-top: 0;
	margin-bottom: 0;
	border-bottom: 1px solid black;
	text-decoration: none;
}
@media only screen and (min-width: 411px) and (max-width: 768px) {
	.blogNaviPageComponent__gridItemTitle{
		max-width: 390px;
		margin-left: auto;
		margin-right: auto;
	}
	.blogNaviPageComponent__image {
		max-width: 390px;
	}
}

@media only screen and (max-width: 410px) {
	.blogNaviPageComponent__image {
		max-height: none;
		max-width: calc(100% - 20px);
		height: auto;
	}
	.blogNaviPageComponent__gridItem:hover,
	.blogNaviPageComponent__gridItem {
		max-height: none;
	}
	.blogNaviPageComponent__gridItemTitle {
		margin-left: 10px;
		margin-right: 10px;
	}
}
/* BlogNaviPageContentComponent end */
`
}
