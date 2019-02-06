package staticPresentation

import (
	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
	"github.com/ingmardrewing/staticUtil"

	log "github.com/sirupsen/logrus"
)

// Creates a new EntryPageComponent
func NewHomePageComponent(r staticIntf.Renderer) *HomePageComponent {
	h := new(HomePageComponent)
	h.abstractComponent.Renderer(r)
	return h
}

type HomePageComponent struct {
	abstractComponent
	wrapper
	mainDiv *htmlDoc.Node
}

func (e *HomePageComponent) VisitPage(p staticIntf.Page) {
	e.mainDiv = htmlDoc.NewNode("div", "", "class", "homepage__content")

	textBlock := e.getHomeTextBlock(p.Site())
	e.mainDiv.AddChild(textBlock)

	containerBlocks := e.getBlocksFromContainers(p.Site())
	for _, cb := range containerBlocks {
		e.mainDiv.AddChild(cb)
	}

	w := e.wrap(e.mainDiv, "homepage__wrapperouter")
	p.AddBodyNodes([]*htmlDoc.Node{w})
}

func (e *HomePageComponent) getHomeTextBlock(site staticIntf.Site) *htmlDoc.Node {
	hl := site.HomeHeadline()
	txt := site.HomeText()
	return e.createBlockFromTexts(hl, txt)
}

func (e *HomePageComponent) getBlocksFromContainers(site staticIntf.Site) []*htmlDoc.Node {
	tool := staticUtil.NewPagesContainerCollectionTool(site)
	containers := tool.ContainersOrderedByVariants("blog", "portfolio")
	return e.createBlocksFrom(containers)
}

func (e *HomePageComponent) createBlocksFrom(containers []staticIntf.PagesContainer) []*htmlDoc.Node {
	blocks := []*htmlDoc.Node{}
	for _, c := range containers {
		block := e.createBlockFrom(c)
		if block != nil {
			blocks = append(blocks, block)
		}
	}
	return blocks
}

func (e *HomePageComponent) createBlockFrom(c staticIntf.PagesContainer) *htmlDoc.Node {
	pages := c.Representationals()
	log.Debugf("HomePageComponent.createBlockFrom(), found %d representational pages\n", len(pages))
	if len(pages) > 0 {
		block := htmlDoc.NewNode(
			"div", "",
			"class", "homepageblock")
		block.AddChild(htmlDoc.NewNode(
			"h2", c.Headline(),
			"class", "homepageblock__headline"))
		block.AddChild(e.createGridWithLinksFrom(pages))
		return block
	}
	return nil
}

func (e *HomePageComponent) createBlockFromTexts(headlineTxt, bodyCopy string) *htmlDoc.Node {
	block := htmlDoc.NewNode(
		"div", "",
		"class", "homepageblock")
	block.AddChild(htmlDoc.NewNode(
		"h2", headlineTxt,
		"class", "homepageblock__headline"))
	block.AddChild(htmlDoc.NewNode(
		"p", bodyCopy,
		"class", "homepageblock__paragraph"))
	return block
}

func (e *HomePageComponent) createGridWithLinksFrom(pages []staticIntf.Page) *htmlDoc.Node {
	grid := htmlDoc.NewNode(
		"div", " ",
		"class", "homepage__grid")
	for i := len(pages) - 1; i >= 0; i-- {
		grid.AddChild(e.getElementLinkingToPages(pages[i]))
	}
	return grid
}

func (e *HomePageComponent) getElementLinkingToPages(page staticIntf.Page) *htmlDoc.Node {
	a := htmlDoc.NewNode(
		"a", " ",
		"href", page.Link(),
		"title", page.Title(),
		"class", "homepage__tile")
	a.AddChild(htmlDoc.NewNode(
		"img", " ",
		"src", page.MicroThumbnailUrl(),
		"srcset", staticUtil.MakeMicroSrcSet(page),
		"alt", page.Title(),
		"class", "homepage__tileImg"))
	return a
}

func (b *HomePageComponent) GetCss() string {
	return `
/* HomePageComponent start */
.homepageblock {
	text-align: center;
}
.homepage__grid {
	display: grid;
	grid-template-columns: 190px 190px 190px 190px;
	grid-gap: 20px;
}
.homepage__tile {
	display: block;
	overflow: hidden;
	max-height: 190px;
}
.homepage__tileImg {
	max-height: 190px;
	max-width: 190px;
}
.homepageblock__headline {
	font-size: 18px;
	text-align: left;
	font-weight: 700;
	text-transform: uppercase;
	border-bottom: 1px solid black;
	margin-top: 20px;
}
.homepageblock__paragraph {
	text-align: left;
	font-weight: 400;
	line-height: 2em;
}
.homepage__content {
	margin-top: 145px;
	padding-bottom: 50px;
	text-align: left;
	min-height: calc(100vh - 520px);
}

@media only screen and (max-width: 768px) {
	.homepage__content{
		margin-top: 0;
	}
}
@media only screen and (min-width: 610px) and (max-width: 819px) {
	.homepage__grid {
		grid-template-columns: 190px 190px 190px;
		width: 610px;
		margin: 0 auto;
	}
	.homepageblock__paragraph ,
	.homepageblock__headline {
		padding-left: calc((100% - 610px)/2 );
		padding-right: calc((100% - 610px)/2 );
	}
}
@media only screen and (min-width: 400px) and (max-width: 609px) {
	.homepage__grid {
		grid-template-columns: 190px 190px;
		width: 400px;
		margin: 0 auto;
	}
	.homepageblock__paragraph ,
	.homepageblock__headline {
		padding-left: calc((100% - 400px)/2 );
		padding-right: calc((100% - 400px)/2 );
	}
}
/* HomePageComponent end */
`
}
