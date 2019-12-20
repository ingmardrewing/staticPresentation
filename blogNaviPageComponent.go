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

func (c *BlogNaviPageContentComponent) VisitPage(page staticIntf.Page) {
	div := htmlDoc.NewNode(
		"div", "",
		"class", "blogNaviPageComponent__grid")

	for _, navigated := range page.NavigatedPages() {
		nn := c.renderNavigatedPage(navigated)
		div.AddChild(nn)
	}

	wrapper := c.wrap(div)
	page.AddBodyNodes([]*htmlDoc.Node{wrapper})
}

func (c *BlogNaviPageContentComponent) renderNavigatedPage(page staticIntf.Page) *htmlDoc.Node {
	a := c.createLink(page)
	a.AddChild(c.createItemLabel(page))
	a.AddChild(c.createImage(page))
	if page.Description() != page.Site().Description() && page.Description() != "" {
		a.AddChild(c.createDescription(page))
	}
	return a
}

func (c *BlogNaviPageContentComponent) createLink(page staticIntf.Page) *htmlDoc.Node {
	return htmlDoc.NewNode(
		"a", " ",
		"href", page.Link(),
		"class", "blogNaviPageComponent__gridItem")
}

func (c *BlogNaviPageContentComponent) createItemLabel(page staticIntf.Page) *htmlDoc.Node {
	return htmlDoc.NewNode(
		"h2", page.Title(),
		"class", "blogNaviPageComponent__gridItemTitle")
}

func (c *BlogNaviPageContentComponent) createImage(page staticIntf.Page) *htmlDoc.Node {
	if len(staticUtil.MakeSrcSet(page)) > 5 {
		return htmlDoc.NewNode(
			"img", "",
			"alt", page.Title(),
			"src", page.ThumbnailUrl(),
			"width", "390",
			"height", "390",
			"srcset", staticUtil.MakeSrcSet(page),
			"class", "blogNaviPageComponent__image")
	}
	return htmlDoc.NewNode(
		"img", "",
		"alt", page.Title(),
		"src", page.ThumbnailUrl(),
		"width", "390",
		"height", "390",
		"class", "blogNaviPageComponent__image")
}

func (c *BlogNaviPageContentComponent) createDescription(page staticIntf.Page) *htmlDoc.Node {
	descriptionContainer := htmlDoc.NewNode(
		"div", "",
		"class", "blogNaviPageComponent__descriptionContainer")
	descriptionContainer.AddChild(htmlDoc.NewNode(
		"div", page.Description(),
		"class", "blogNaviPageComponent__description"))
	return descriptionContainer
}

func (c *BlogNaviPageContentComponent) GetCss() string {
	return `
/* BlogNaviPageContentComponent start */
.blogNaviPageComponent__descriptionContainer {
	position: absolute;
	bottom: 0;
	left: 0;
	right: 0;
	overflow: hidden;
	width: 100%;
	height: 0;

	background-color: #FFFFFFCF;

	-webkit-transition: 0.4s ease;
    -moz-transition: 0.4s ease;
    -o-transition: 0.4s ease;
    transition: 0.4s ease;

}
@media (hover) {
	.blogNaviPageComponent__gridItem:hover .blogNaviPageComponent__descriptionContainer {
		height: 33%;
	}
}
.blogNaviPageComponent__description {
	box-sizing: border-box;
	padding: 20px;
	width: 100%;
	height: 100%;
	text-align: left;
	color: #000;
}
.blogNaviPageComponent__grid {
	display: grid;
	grid-gap: 20px;
}
.blogNaviPageComponent__image {
	max-height: 390px;
	max-width: 390px;
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
