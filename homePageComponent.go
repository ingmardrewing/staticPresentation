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
		block := e.createBlockNode(c.Headline())
		ctr := 1
		for _, l := range e.createLinksFrom(pages) {
			block.AddChild(l)
			if ctr%4 == 0 {
				block.AddChild(htmlDoc.NewNode("br", ""))
			}
			ctr++
		}
		block.AddChild(htmlDoc.NewNode("span", " ", "class", "clearfix"))
		return block
	}
	return nil
}

func (e *HomePageComponent) createBlockFromTexts(headlineTxt, bodyCopy string) *htmlDoc.Node {
	n := e.createBlockNode(headlineTxt)
	paragraph := htmlDoc.NewNode("p", bodyCopy, "class", "homepageblock__paragraph")
	n.AddChild(paragraph)
	return n
}

func (e *HomePageComponent) createBlockNode(headlineTxt string) *htmlDoc.Node {
	block := htmlDoc.NewNode("div", "", "class", "homepageblock")
	blockHeadline := htmlDoc.NewNode("h2", headlineTxt, "class", "homepageblock__headline")
	block.AddChild(blockHeadline)
	return block
}

func (e *HomePageComponent) createLinksFrom(pages []staticIntf.Page) []*htmlDoc.Node {
	links := []*htmlDoc.Node{}
	for i := len(pages) - 1; i >= 0; i-- {
		link := e.getElementLinkingToPages(pages[i])
		links = append(links, link)
	}
	return links
}

func (e *HomePageComponent) getElementLinkingToPages(page staticIntf.Page) *htmlDoc.Node {
	a := htmlDoc.NewNode("a", " ",
		"href", page.Link(),
		"title", page.Title(),
		"class", "homepage__thumb")
	a.AddChild(htmlDoc.NewNode("img", " ",
		"src", page.MicroThumbnailUrl(),
		"srcset", staticUtil.MakeMicroSrcSet(page),
		"class", "homepage__thumbimg"))
	return a
}

func (e *HomePageComponent) GetCss() string {
	return `
.homepage__thumbimg {
	max-width: 100%;
	height: auto;
	clip: rect(0px,0px,190px, 190px);
}
@media only screen and (max-width: 768px) {
	.homepageblock__headline {
		font-family: Arial Black, Arial, Helvetica, sans-serif;
		text-transform: uppercase;
	}
	.homepageblock__headline ,
	.homepageblock__paragraph {
		padding-left: 10px;
		padding-right: 10px;
	}
	.homepageblock + .homepageblock h2 {
		border-top: 1px solid black;
	}
	.homepageblock__paragraph {
		font-family: Arial, Helvetica, sans-serif;
		line-height: 1.4em;
	}
	.homepage__content {
		padding-bottom: 50px;
		text-align: left;
		min-height: calc(100vh - 520px);
	}
	.homepage__thumb {
		display: block;
		float: left;
		width: 50%;
		height: auto;
		background-size: cover;
		margin-left: 0px;
		margin-right: 0px;
		margin-top: 10px;
		margin-bottom: 10px;
	}
	.clearfix {
		display:block;
		visibility:hidden;
		clear:both;
		height:0;
		font-size:0;
		content:" ";
	}
}
@media only screen and (min-width: 769px) {
	.homepage__wrapperouter {
		margin-top: 165px;
	}
	.homepageblock__headline {
		font-family: Arial Black, Arial, Helvetica, sans-serif;
		text-transform: uppercase;
		border-bottom: 1px solid black;
	}
	.homepageblock__paragraph {
		font-family: Arial, Helvetica, sans-serif;
		line-height: 2em;
	}
	.homepage__content {
		padding-bottom: 50px;
		text-align: left;
		min-height: calc(100vh - 520px);
	}
	.homepage__thumb {
		display: block;
		float: left;
		background-size: cover;
		width: 190px;
		height: 190px;
		margin-left: 0px;
		margin-top: 10px;
		margin-bottom: 10px;
	}
	.homepage__thumb + .homepage__thumb {
		margin-left: 13px;
	}
	.clearfix {
		display:block;
		visibility:hidden;
		clear:both;
		height:0;
		font-size:0;
		content:" ";
	}
}
`
}
