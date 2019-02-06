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
	e.mainDiv = htmlDoc.NewNode("div", "", "class", "homePageComponent__content")

	textBlock := e.getHomeTextBlock(p.Site())
	e.mainDiv.AddChild(textBlock)

	containerBlocks := e.getBlocksFromContainers(p.Site())
	for _, cb := range containerBlocks {
		e.mainDiv.AddChild(cb)
	}

	w := e.wrap(e.mainDiv, "homePageComponent__wrapperouter")
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
			"class", "homePageComponent")
		block.AddChild(htmlDoc.NewNode(
			"h2", c.Headline(),
			"class", "homePageComponent__headline"))
		block.AddChild(e.createGridWithLinksFrom(pages))
		return block
	}
	return nil
}

func (e *HomePageComponent) createBlockFromTexts(headlineTxt, bodyCopy string) *htmlDoc.Node {
	block := htmlDoc.NewNode(
		"div", "",
		"class", "homePageComponent")
	block.AddChild(htmlDoc.NewNode(
		"h2", headlineTxt,
		"class", "homePageComponent__headline"))
	block.AddChild(htmlDoc.NewNode(
		"p", bodyCopy,
		"class", "homePageComponent__paragraph"))
	return block
}

func (e *HomePageComponent) createGridWithLinksFrom(pages []staticIntf.Page) *htmlDoc.Node {
	grid := htmlDoc.NewNode(
		"div", " ",
		"class", "homePageComponent__grid")
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
		"class", "homePageComponent__tile")
	a.AddChild(htmlDoc.NewNode(
		"img", " ",
		"src", page.MicroThumbnailUrl(),
		"srcset", staticUtil.MakeMicroSrcSet(page),
		"alt", page.Title(),
		"class", "homePageComponent__tileImg"))
	return a
}

func (b *HomePageComponent) GetCss() string {
	return `
/* HomePageComponent start */
.homePageComponent {
	text-align: center;
}
.homePageComponent__grid {
	display: grid;
	grid-template-columns: 190px 190px 190px 190px;
	grid-gap: 20px;
}
.homePageComponent__tile {
	display: block;
	overflow: hidden;
	max-height: 190px;
}
.homePageComponent__tileImg {
	max-height: 190px;
	max-width: 190px;
}
.homePageComponent__headline {
	font-size: 18px;
	text-align: left;
	font-weight: 700;
	text-transform: uppercase;
	border-bottom: 1px solid black;
	margin-top: 20px;
}
.homePageComponent__paragraph {
	text-align: left;
	font-weight: 400;
	line-height: 2em;
}
.homePageComponent__content {
	margin-top: 145px;
	padding-bottom: 50px;
	text-align: left;
	min-height: calc(100vh - 520px);
}

@media only screen and (max-width: 768px) {
	.homePageComponent__content{
		margin-top: 0;
	}
}
@media only screen and (min-width: 610px) and (max-width: 819px) {
	.homePageComponent__grid {
		grid-template-columns: 190px 190px 190px;
		width: 610px;
		margin: 0 auto;
	}
	.homePageComponent__paragraph ,
	.homePageComponent__headline {
		padding-left: calc((100% - 610px)/2 );
		padding-right: calc((100% - 610px)/2 );
	}
}
@media only screen and (min-width: 400px) and (max-width: 609px) {
	.homePageComponent__grid {
		grid-template-columns: 190px 190px;
		width: 400px;
		margin: 0 auto;
	}
	.homePageComponent__paragraph,
	.homePageComponent__headline {
		padding-left: calc((100% - 400px)/2 );
		padding-right: calc((100% - 400px)/2 );
	}
}
@media only screen and (max-width: 399px) {
	.homePageComponent__grid {
		grid-template-columns: 100%;
		width: 100%;
		margin: 0 auto;
	}
	.homePageComponent__paragraph,
	.homePageComponent__headline {
		padding-left: 10px;
		padding-right: 10px;
	}
	.homePageComponent__tileImg {
		width: calc(100% - 20px);
		max-width: calc(100% - 20px);
		height: auto;
		max-height: none;
	}
	.homePageComponent__tile {
		max-height: none;
	}
	.homePageComponent__grid {
		grid-gap: 7px;
	}
	.homePageComponent__paragraph {
		line-height: 1.5em;
	}
}
/* HomePageComponent end */
`
}
