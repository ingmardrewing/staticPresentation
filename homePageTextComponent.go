package staticPresentation

import (
	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
)

// Creates a new EntryPageComponent
func NewHomePageTextComponent(r staticIntf.Renderer) *HomePageTextComponent {
	h := new(HomePageTextComponent)
	h.abstractComponent.Renderer(r)
	return h
}

type HomePageTextComponent struct {
	abstractComponent
	wrapper
	mainDiv *htmlDoc.Node
}

func (e *HomePageTextComponent) VisitPage(p staticIntf.Page) {
	e.mainDiv = htmlDoc.NewNode("div", "", "class", "homePageTextComponent__content")

	textBlock := e.getHomeTextBlock(p.Site())
	e.mainDiv.AddChild(textBlock)

	w := e.wrap(e.mainDiv, "homePageTextComponent__wrapperouter")
	p.AddBodyNodes([]*htmlDoc.Node{w})
}

func (e *HomePageTextComponent) getHomeTextBlock(site staticIntf.Site) *htmlDoc.Node {
	hl := site.HomeHeadline()
	txt := site.HomeText()
	return e.createBlockFromTexts(hl, txt)
}

func (e *HomePageTextComponent) createBlockFromTexts(headlineTxt, bodyCopy string) *htmlDoc.Node {
	block := htmlDoc.NewNode(
		"div", "",
		"class", "homePageTextComponent")
	block.AddChild(htmlDoc.NewNode(
		"h2", headlineTxt,
		"class", "homePageTextComponent__headline"))
	block.AddChild(htmlDoc.NewNode(
		"p", bodyCopy,
		"class", "homePageTextComponent__paragraph"))
	return block
}

func (b *HomePageTextComponent) GetCss() string {
	return `
/* HomePageTextComponent start */
.homePageTextComponent {
	text-align: center;
}
.homePageTextComponent__grid {
	display: grid;
	grid-template-columns: 190px 190px 190px 190px;
	grid-gap: 20px;
}
.homePageTextComponent__tile {
	display: block;
	overflow: hidden;
	max-height: 190px;
}
.homePageTextComponent__tileImg {
	max-height: 190px;
	max-width: 190px;
}
.homePageTextComponent__headline {
	font-size: 18px;
	text-align: left;
	font-weight: 700;
	text-transform: uppercase;
	border-bottom: 1px solid black;
	margin-top: 20px;
}
.homePageTextComponent__paragraph {
	text-align: left;
	font-weight: 400;
	line-height: 2em;
}
.homePageTextComponent__content {
	margin-top: 145px;
	text-align: left;
}

@media only screen and (max-width: 768px) {
	.homePageTextComponent__content{
		margin-top: 0;
	}
}
@media only screen and (min-width: 630px) and (max-width: 819px) {
	.homePageTextComponent__paragraph {
		line-height: 1.8em;
	}
	.homePageTextComponent__grid {
		grid-template-columns: 190px 190px 190px;
		width: 610px;
		margin: 0 auto;
	}
	.homePageTextComponent__paragraph ,
	.homePageTextComponent__headline {
		width: 610px;
		max-width: 610px;
		margin-left: auto;
		margin-right: auto;
	}
}
@media only screen and (min-width: 420px) and (max-width: 629px) {
	.homePageTextComponent__grid {
		grid-template-columns: 190px 190px;
		width: 400px;
		margin: 0 auto;
	}
	.homePageTextComponent__paragraph,
	.homePageTextComponent__headline {
		width: 400px;
		max-width: 400px;
		margin-left: auto;
		margin-right: auto;
	}
	.homePageTextComponent__paragraph {
		line-height: 1.5em;
	}
}
@media only screen and (max-width: 419px) {
	.homePageTextComponent__grid {
		grid-template-columns: 100%;
		width: 100%;
		margin: 0 auto;
	}
	.homePageTextComponent__paragraph,
	.homePageTextComponent__headline {
		margin-left: 10px;
		margin-right: 10px;
	}
	.homePageTextComponent__tileImg {
		width: calc(100% - 20px);
		max-width: calc(100% - 20px);
		height: auto;
		max-height: none;
	}
	.homePageTextComponent__tile {
		max-height: none;
	}
	.homePageTextComponent__grid {
		grid-gap: 7px;
	}
	.homePageTextComponent__paragraph {
		line-height: 1.5em;
	}
}
/* HomePageTextComponent end */
`
}
