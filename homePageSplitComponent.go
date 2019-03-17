package staticPresentation

import (
	"fmt"
	"strings"

	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
	"github.com/ingmardrewing/staticUtil"

	log "github.com/sirupsen/logrus"
)

// Creates a new EntryPageComponent
func NewHomePageSplitComponent(r staticIntf.Renderer) *HomePageSplitComponent {
	h := new(HomePageSplitComponent)
	h.abstractComponent.Renderer(r)
	return h
}

type HomePageSplitComponent struct {
	abstractComponent
	wrapper
	mainDiv *htmlDoc.Node
}

func (e *HomePageSplitComponent) VisitPage(p staticIntf.Page) {
	e.mainDiv = htmlDoc.NewNode("div", "", "class", "homePageSplitComponent__content")

	tool := staticUtil.NewPagesContainerCollectionTool(p.Site())
	containers := tool.ContainersOrderedByVariants("blog")
	for _, cb := range e.createBlocksFrom(containers) {
		e.mainDiv.AddChild(cb)
	}

	w := e.wrap(e.mainDiv, "homePageSplitComponent__wrapperouter")
	p.AddBodyNodes([]*htmlDoc.Node{w})
}

func (e *HomePageSplitComponent) createBlocksFrom(
	containers []staticIntf.PagesContainer) []*htmlDoc.Node {

	blocks := []*htmlDoc.Node{}
	for _, c := range containers {
		block := e.createBlockFrom(c)
		if block != nil {
			blocks = append(blocks, block)
		}
	}
	return blocks
}

func (e *HomePageSplitComponent) createBlockFrom(
	c staticIntf.PagesContainer) *htmlDoc.Node {

	pages := c.Representationals()
	log.Debugf("HomePageSplitComponent.createBlockFrom(), found %d representational pages\n", len(pages))
	if len(pages) > 0 {
		block := htmlDoc.NewNode(
			"div", "",
			"class", "homePageSplitComponent")
		block.AddChild(e.createGridWithLinksFrom(pages))
		return block
	}
	return nil
}

func (e *HomePageSplitComponent) createGridWithLinksFrom(
	pages []staticIntf.Page) *htmlDoc.Node {

	grid := htmlDoc.NewNode(
		"div", " ",
		"class", "homePageSplitComponent__grid")
	grid.AddChild(e.getIntroElement(pages[0].Site()))
	for i := len(pages) - 1; i >= 0; i-- {
		grid.AddChild(e.getElementLinkingToPages(pages[i]))
	}
	return grid
}

func (e *HomePageSplitComponent) getIntroElement(site staticIntf.Site) *htmlDoc.Node {
	h2 := htmlDoc.NewNode(
		"h2",
		site.HomeHeadline(),
		"class", "homePageSplitComponent__headline")
	txt := htmlDoc.NewNode(
		"p",
		site.HomeText())
	div := htmlDoc.NewNode(
		"div", "",
		"class", "homePageSplitComponent_intro")
	div.AddChild(h2)
	div.AddChild(txt)
	return div
}

func (e *HomePageSplitComponent) getElementLinkingToPages(
	page staticIntf.Page) *htmlDoc.Node {

	a := htmlDoc.NewNode(
		"a", " ",
		"href", page.Link(),
		"title", page.Title(),
		"class", "homePageSplitComponent__tile")
	a.AddChild(htmlDoc.NewNode(
		"img", "",
		"src", e.findSrc(page.Images()),
		"srcset", e.findSrcSet(page.Images()),
		"width", "80",
		"height", "80",
		"alt", page.Title(),
		"class", "homePageSplitComponent__tileImg"))
	a.AddChild(htmlDoc.NewNode(
		"span", page.PublishedTime()+" ",
		"class", "homePageSplitComponent__tileDate"))
	a.AddChild(htmlDoc.NewNode(
		"span", e.getBlogExcerpt(page),
		"class", "homePageSplitComponent__tileText"))
	return a
}

func (e *HomePageSplitComponent) findSrc(images []staticIntf.Image) string {
	if len(images) > 0 {
		if len(images[0].W80Square()) > 0 {
			return images[0].W80Square()
		}
		return images[0].W185Square()
	}
	return ""
}

func (e *HomePageSplitComponent) findSrcSet(images []staticIntf.Image) string {
	if len(images) > 0 {
		if len(images[0].W185Square()) > 0 {
			return fmt.Sprintf("%s 2x", images[0].W185Square())
		}
	}
	return ""
}

func (e *HomePageSplitComponent) getBlogExcerpt(page staticIntf.Page) string {
	if page.Description() == page.Site().HomeText() ||
		strings.HasPrefix(page.Description(), "A blog containing") {
		return "Just an image ..."
	} else if len(page.Description()) > 0 {
		l := len(page.Description())
		if l > 80 {
			l = 80
		}
		n := strings.Split(
			page.Description()[:l], " ")
		return strings.Join(n[:len(n)-1], " ") + " ..."
	}
	return ""
}

func (b *HomePageSplitComponent) GetCss() string {
	return `
/* HomePageSplitComponent start */
.homePageSplitComponent {
	text-align: center;
}
.homePageSplitComponent__grid {
	display: grid;
	grid-template-columns: 390px 390px;
	grid-gap: 20px;
}
.homePageSplitComponent_intro {
	grid-row: 1 / span 4;
	text-align: left;
	line-height: 2em;
	font-weight: 400;
}
.homePageSplitComponent__tile {
	text-align: left;
	display: block;
	overflow: hidden;
	max-height: 80px;
}
.homePageSplitComponent__tileImg {
	max-height: 80px;
	max-width: 80px;
	margin-right: 20px;
	float: left;
}
.homePageSplitComponent__tileDate,
.homePageSplitComponent__tileText{
	line-height: 1.6em;
}
.homePageSplitComponent__tileText{
	color: #000;
}
.homePageSplitComponent__headline {
	font-size: 18px;
	text-align: left;
	font-weight: 700;
	text-transform: uppercase;
	margin-top: 0;
	margin-bottom: 0;
}
.homePageSplitComponent__paragraph {
	text-align: left;
	font-weight: 400;
	line-height: 2em;
}
.homePageSplitComponent__content {
	text-align: left;
}
@media only screen and (max-width: 768px) {
	.homePageSplitComponent__content{
		margin-top: 0;
	}
}
@media only screen and (min-width: 610px) and (max-width: 819px) {
	.homePageSplitComponent__grid {
		grid-template-columns: 285px 285px;
		width: 610px;
		margin: 0 auto;
	}
	.homePageSplitComponent__paragraph ,
	.homePageSplitComponent__headline {
		margin-left: auto;
		margin-right: auto;
	}
	.homePageSplitComponent__tileText{
		line-height: 1.3em;
	}
	.homePageSplitComponent__tile {
		max-height: 100px;
	}
}
@media only screen and (min-width: 400px) and (max-width: 609px) {
	.homePageSplitComponent__grid {
		grid-template-columns: 400px;
		width: 400px;
		margin: 0 auto;
	}
	.homePageSplitComponent__paragraph,
	.homePageSplitComponent__headline {
		width: 400px;
		max-width: 400px;
		margin-left: auto;
		margin-right: auto;
	}
}
@media only screen and (max-width: 399px) {
	.homePageSplitComponent__grid {
		grid-template-columns: 100%;
		width: 100%;
		margin: 0 auto;
		padding: 0 10px;
		box-sizing: border-box;
	}
	.homePageSplitComponent__paragraph,
	.homePageSplitComponent__headline {
		margin-left: 10px;
		margin-right: 10px;
	}
	.homePageSplitComponent__tile {
		max-height: none;
	}
	.homePageSplitComponent__grid {
		grid-gap: 7px;
	}
	.homePageSplitComponent__paragraph {
		line-height: 1.5em;
	}
	.homePageSplitComponent__tileText {
		line-height: 1em;
	}
}
/* HomePageSplitComponent end */
`
}
