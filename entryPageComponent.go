package staticPresentation

import (
	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
)

// Creates a new EntryPageComponent
func NewEntryPageComponent() *EntryPageComponent {
	return new(EntryPageComponent)
}

type EntryPageComponent struct {
	abstractComponent
	wrapper
}

func (e *EntryPageComponent) VisitPage(p staticIntf.Page) {
	mainDiv := htmlDoc.NewNode("div", "", "class", "mainpage__content")
	containers := e.renderer.Site().Containers()
	for _, block := range e.createBlocksFrom(containers) {
		mainDiv.AddChild(block)
	}
	wn := e.wrap(mainDiv)
	p.AddBodyNodes([]*htmlDoc.Node{wn})
}

func (e *EntryPageComponent) createBlocksFrom(containers []staticIntf.PagesContainer) []*htmlDoc.Node {
	blocks := []*htmlDoc.Node{}
	for _, c := range containers {
		block := e.createBlockFrom(c)
		if block != nil {
			blocks = append(blocks, block)
		}
	}
	return blocks
}

func (e *EntryPageComponent) createBlockFrom(c staticIntf.PagesContainer) *htmlDoc.Node {
	pages := c.Representationals()
	if len(pages) > 0 {
		block := e.createBlockNode(c)
		for _, l := range e.createLinksFrom(pages) {
			block.AddChild(l)
		}
		block.AddChild(htmlDoc.NewNode("span", " ", "class", "clearfix"))
		return block
	}
	return nil
}

func (e *EntryPageComponent) createBlockNode(c staticIntf.PagesContainer) *htmlDoc.Node {
	block := htmlDoc.NewNode("div", "", "class", "mainpageblock")
	blockHeadline := htmlDoc.NewNode("h2", c.Variant(), "class", "mainpageblock__headline")
	block.AddChild(blockHeadline)
	return block
}

func (e *EntryPageComponent) createLinksFrom(pages []staticIntf.Page) []*htmlDoc.Node {
	links := []*htmlDoc.Node{}
	for i := len(pages) - 1; i >= 0; i-- {
		link := e.getElementLinkingToPages(pages[i])
		links = append(links, link)
	}
	return links
}

func (e *EntryPageComponent) getElementLinkingToPages(page staticIntf.Page) *htmlDoc.Node {
	return htmlDoc.NewNode("a", " ",
		"href", page.PathFromDocRootWithName(),
		"title", page.Title(),
		"class", "mainpage__thumb",
		"style", "background-image: url("+page.ThumbnailUrl()+")")
}

func (e *EntryPageComponent) GetCss() string {
	return `
.mainpageblock__headline {
	text-transform: uppercase;
	border-bottom: 1px solid black;
}
.mainpage__content {
	padding-top: 146px;
	padding-bottom: 50px;
	text-align: left;
	min-height: calc(100vh - 520px);
}
.mainpage__thumb {
	display: block;
	float: left;
	width: 190px;
	height: 190px;
	background-size: cover;
	margin-left: 0px;
	margin-top: 10px;
	margin-bottom: 10px;
}
.mainpage__thumb + .mainpage__thumb {
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
`
}
