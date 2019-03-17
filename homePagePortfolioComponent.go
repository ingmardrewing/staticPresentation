package staticPresentation

import (
	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
	"github.com/ingmardrewing/staticUtil"

	log "github.com/sirupsen/logrus"
)

// Creates a new EntryPageComponent
func NewHomePagePortfolioComponent(r staticIntf.Renderer) *HomePagePortfolioComponent {
	h := new(HomePagePortfolioComponent)
	h.abstractComponent.Renderer(r)
	return h
}

type HomePagePortfolioComponent struct {
	abstractComponent
	wrapper
	mainDiv *htmlDoc.Node
}

func (e *HomePagePortfolioComponent) VisitPage(p staticIntf.Page) {
	e.mainDiv = htmlDoc.NewNode("div", "", "class", "homePagePortfolioComponent__content")

	tool := staticUtil.NewPagesContainerCollectionTool(p.Site())
	containers := tool.ContainersOrderedByVariants("portfolio")
	for _, cb := range e.createBlocksFrom(containers) {
		e.mainDiv.AddChild(cb)
	}

	w := e.wrap(e.mainDiv, "homePagePortfolioComponent__wrapperouter")
	p.AddBodyNodes([]*htmlDoc.Node{w})
}

func (e *HomePagePortfolioComponent) createBlocksFrom(containers []staticIntf.PagesContainer) []*htmlDoc.Node {
	blocks := []*htmlDoc.Node{}
	for _, c := range containers {
		block := e.createBlockFrom(c)
		if block != nil {
			blocks = append(blocks, block)
		}
	}
	return blocks
}

func (e *HomePagePortfolioComponent) createBlockFrom(c staticIntf.PagesContainer) *htmlDoc.Node {
	pages := c.Representationals()
	log.Debugf("HomePagePortfolioComponent.createBlockFrom(), found %d representational pages\n", len(pages))
	if len(pages) > 0 {
		block := htmlDoc.NewNode(
			"div", "",
			"class", "homePagePortfolioComponent")
		block.AddChild(htmlDoc.NewNode(
			"h2", c.Headline(),
			"class", "homePagePortfolioComponent__headline"))
		block.AddChild(e.createGridWithLinksFrom(pages))
		return block
	}
	return nil
}

func (e *HomePagePortfolioComponent) createGridWithLinksFrom(pages []staticIntf.Page) *htmlDoc.Node {
	grid := htmlDoc.NewNode(
		"div", " ",
		"class", "homePagePortfolioComponent__grid")
	for i := len(pages) - 1; i >= 0; i-- {
		grid.AddChild(e.getElementLinkingToPages(pages[i]))
	}
	return grid
}

func (e *HomePagePortfolioComponent) getElementLinkingToPages(page staticIntf.Page) *htmlDoc.Node {
	a := htmlDoc.NewNode(
		"a", " ",
		"href", page.Link(),
		"title", page.Title(),
		"class", "homePagePortfolioComponent__tile")
	a.AddChild(htmlDoc.NewNode(
		"img", " ",
		"src", page.MicroThumbnailUrl(),
		"srcset", staticUtil.MakeMicroSrcSet(page),
		"alt", page.Title(),
		"class", "homePagePortfolioComponent__tileImg"))
	return a
}

func (b *HomePagePortfolioComponent) GetCss() string {
	return `
/* HomePagePortfolioComponent start */
.homePagePortfolioComponent {
	text-align: center;
}
.homePagePortfolioComponent__grid {
	display: grid;
	grid-template-columns: 190px 190px 190px 190px;
	grid-gap: 20px;
}
.homePagePortfolioComponent__tile {
	display: block;
	overflow: hidden;
	max-height: 190px;
}
.homePagePortfolioComponent__tileImg {
	max-height: 190px;
	max-width: 190px;
}
.homePagePortfolioComponent__headline {
	font-size: 18px;
	text-align: left;
	font-weight: 700;
	text-transform: uppercase;
	border-bottom: 1px solid black;
	margin-top: 30px;
}
.homePagePortfolioComponent__paragraph {
	text-align: left;
	font-weight: 400;
	line-height: 2em;
}
.homePagePortfolioComponent__content {
	padding-bottom: 50px;
	text-align: left;
}

@media only screen and (max-width: 768px) {
	.homePagePortfolioComponent__content{
		margin-top: 0;
	}
}
@media only screen and (min-width: 610px) and (max-width: 819px) {
	.homePagePortfolioComponent__grid {
		grid-template-columns: 190px 190px 190px;
		width: 610px;
		margin: 0 auto;
	}
	.homePagePortfolioComponent__paragraph ,
	.homePagePortfolioComponent__headline {
		width: 610px;
		max-width: 610px;
		margin-left: auto;
		margin-right: auto;
	}
}
@media only screen and (min-width: 400px) and (max-width: 609px) {
	.homePagePortfolioComponent__grid {
		grid-template-columns: 190px 190px;
		width: 400px;
		margin: 0 auto;
	}
	.homePagePortfolioComponent__paragraph,
	.homePagePortfolioComponent__headline {
		width: 400px;
		max-width: 400px;
		margin-left: auto;
		margin-right: auto;
	}
}
@media only screen and (max-width: 399px) {
	.homePagePortfolioComponent__grid {
		grid-template-columns: 100%;
		width: 100%;
		margin: 0 auto;
	}
	.homePagePortfolioComponent__grid *:nth-child(n+5){
		display: none;
	}
	.homePagePortfolioComponent__paragraph,
	.homePagePortfolioComponent__headline {
		margin-left: 10px;
		margin-right: 10px;
	}
	.homePagePortfolioComponent__tileImg {
		width: calc(100% - 20px);
		max-width: calc(100% - 20px);
		height: auto;
		max-height: none;
	}
	.homePagePortfolioComponent__tile {
		max-height: none;
	}
	.homePagePortfolioComponent__grid {
		grid-gap: 7px;
	}
	.homePagePortfolioComponent__paragraph {
		line-height: 1.5em;
	}
}
/* HomePagePortfolioComponent end */
`
}
