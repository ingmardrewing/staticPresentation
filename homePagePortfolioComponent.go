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
	nodes := []*htmlDoc.Node{}

	tool := staticUtil.NewPagesContainerCollectionTool(p.Site())
	containers := tool.ContainersOrderedByVariants("portfolio")
	for _, current := range e.createBlocksFrom(containers) {
		nodes = append(nodes, e.wrap(current, "homePagePortfolioComponent__wrapperouter"))
	}

	p.AddBodyNodes(nodes)
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
		mainDiv := htmlDoc.NewNode("div", "", "class", "homePagePortfolioComponent__content")
		mainDiv.AddChild(block)
		return mainDiv
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
		//"title", page.Title(),
		"class", "homePagePortfolioComponent__tile")

	// TODO: It might make sense to replace
	// the following linesl the picture tag
	// to avoid loading huge piles of images on mobile phones
	a.AddChild(htmlDoc.NewNode(
		"img", " ",
		"width", "185",
		"height", "185",
		"src", page.MicroThumbnailUrl(),
		"srcset", staticUtil.MakeMicroSrcSet(page),
		"alt", page.Title(),
		"class", "homePagePortfolioComponent__tileImg"))
	titleContainer := htmlDoc.NewNode(
		"div", " ",
		"class", "homePagePortfolioComponent__titleContainer")
	titleContainer.AddChild(htmlDoc.NewNode(
		"div", page.Title(),
		"class", "homePagePortfolioComponent__titleText"))
	a.AddChild(titleContainer)
	return a
}

func (b *HomePagePortfolioComponent) GetCss() string {
	return `
/* HomePagePortfolioComponent start */
.homePagePortfolioComponent__titleContainer {
	position: absolute;
	bottom: 0;
	left: 0;
	right: 0;
	overflow: hidden;
	width: 100%;
	height: 0;

	background-color: #FFFFFFCC;

	-webkit-transition: 0.4s ease;
    -moz-transition: 0.4s ease;
    -o-transition: 0.4s ease;
    transition: 0.4s ease;
}
@media (hover) {
	.homePagePortfolioComponent__tile:hover .homePagePortfolioComponent__titleContainer {
		height: 35%;
	}
}
.homePagePortfolioComponent__titleText {
	box-sizing: border-box;
	padding: 10px;
	width: 100%;
	height: 100%;
	text-align: left;
	color: #000;
}

.homePagePortfolioComponent {
	text-align: center;
}
.homePagePortfolioComponent__grid {
	display: grid;
	grid-template-columns: 185px 185px 185px 185px;
	grid-gap: 20px;
}
.homePagePortfolioComponent__tile {
	position: relative;
	display: block;
	overflow: hidden;
	max-height: 185px;
}
.homePagePortfolioComponent__tileImg {
	max-height: 185px;
	max-width: 185px;
}
.homePagePortfolioComponent__headline {
	font-size: 18px;
	text-align: left;
	font-weight: 700;
	text-transform: uppercase;
	border-bottom: 1px solid black;
}
.homePagePortfolioComponent__paragraph {
	text-align: left;
	font-weight: 400;
	line-height: 2em;
}
.homePagePortfolioComponent__content {
	text-align: left;
}

@media only screen and (max-width: 768px) {
	.homePagePortfolioComponent__content{
		margin-top: 0;
	}
}
@media only screen and (min-width: 630px) and (max-width: 819px) {
	.homePagePortfolioComponent__grid {
		grid-template-columns: 185px 185px 185px;
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
@media only screen and (min-width: 420px) and (max-width: 629px) {
	.homePagePortfolioComponent__grid {
		grid-template-columns: 185px 185px;
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
@media only screen and (max-width: 419px) {
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
