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
func NewHomePageBlogComponent(r staticIntf.Renderer) *HomePageBlogComponent {
	h := new(HomePageBlogComponent)
	h.abstractComponent.Renderer(r)
	return h
}

type HomePageBlogComponent struct {
	abstractComponent
	wrapper
	mainDiv *htmlDoc.Node
}

func (e *HomePageBlogComponent) VisitPage(p staticIntf.Page) {
	e.mainDiv = htmlDoc.NewNode("div", "", "class", "homePageBlogComponent__content")

	tool := staticUtil.NewPagesContainerCollectionTool(p.Site())
	containers := tool.ContainersOrderedByVariants("blog")
	for _, cb := range e.createBlocksFrom(containers) {
		e.mainDiv.AddChild(cb)
	}

	w := e.wrap(e.mainDiv, "homePageBlogComponent__wrapperouter")
	p.AddBodyNodes([]*htmlDoc.Node{w})
}

func (e *HomePageBlogComponent) createBlocksFrom(
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

func (e *HomePageBlogComponent) createBlockFrom(
	c staticIntf.PagesContainer) *htmlDoc.Node {

	pages := c.Representationals()
	log.Debugf("HomePageBlogComponent.createBlockFrom(), found %d representational pages\n", len(pages))
	if len(pages) > 0 {
		block := htmlDoc.NewNode(
			"div", "",
			"class", "homePageBlogComponent")
		block.AddChild(htmlDoc.NewNode(
			"h2", c.Headline(),
			"class", "homePageBlogComponent__headline"))
		block.AddChild(e.createGridWithLinksFrom(pages))
		return block
	}
	return nil
}

func (e *HomePageBlogComponent) createGridWithLinksFrom(
	pages []staticIntf.Page) *htmlDoc.Node {

	grid := htmlDoc.NewNode(
		"div", " ",
		"class", "homePageBlogComponent__grid")
	for i := len(pages) - 1; i >= 0; i-- {
		grid.AddChild(e.getElementLinkingToPages(pages[i]))
	}
	return grid
}

func (e *HomePageBlogComponent) getElementLinkingToPages(
	page staticIntf.Page) *htmlDoc.Node {

	a := htmlDoc.NewNode(
		"a", " ",
		"href", page.Link(),
		"title", page.Title(),
		"class", "homePageBlogComponent__tile")
	a.AddChild(htmlDoc.NewNode(
		"img", "",
		"src", e.findSrc(page.Images()),
		"srcset", e.findSrcSet(page.Images()),
		"alt", page.Title(),
		"class", "homePageBlogComponent__tileImg"))
	a.AddChild(htmlDoc.NewNode(
		"span", page.PublishedTime()+" ",
		"class", "homePageBlogComponent__tileDate"))
	a.AddChild(htmlDoc.NewNode(
		"span", e.getBlogExcerpt(page),
		"class", "homePageBlogComponent__tileText"))
	return a
}

func (e *HomePageBlogComponent) findSrc(images []staticIntf.Image) string {
	if len(images) > 0 {
		if len(images[0].W80Square()) > 0 {
			return images[0].W80Square()
		}
		return images[0].W185Square()
	}
	return ""
}

func (e *HomePageBlogComponent) findSrcSet(images []staticIntf.Image) string {
	if len(images) > 0 {
		if len(images[0].W185Square()) > 0 {
			return fmt.Sprintf("%s 2x", images[0].W185Square())
		}
	}
	return ""
}

func (e *HomePageBlogComponent) getBlogExcerpt(page staticIntf.Page) string {
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

func (b *HomePageBlogComponent) GetCss() string {
	return `
/* HomePageComponent start */
.homePageBlogComponent {
	text-align: center;
}
.homePageBlogComponent__grid {
	display: grid;
	grid-template-columns: 390px 390px;
	grid-gap: 20px;
}
.homePageBlogComponent__tile {
	text-align: left;
	display: block;
	overflow: hidden;
	max-height: 80px;
}
.homePageBlogComponent__tileImg {
	max-height: 80px;
	max-width: 80px;
	margin-right: 20px;
	float: left;
}
.homePageBlogComponent__tileDate,
.homePageBlogComponent__tileText{
	line-height: 1.6em;
}
.homePageBlogComponent__tileText{
	color: #000;
}
.homePageBlogComponent__headline {
	font-size: 18px;
	text-align: left;
	font-weight: 700;
	text-transform: uppercase;
	border-bottom: 1px solid black;
	margin-top: 30px;
}
.homePageBlogComponent__paragraph {
	text-align: left;
	font-weight: 400;
	line-height: 2em;
}
.homePageBlogComponent__content {
	text-align: left;
}
@media only screen and (max-width: 768px) {
	.homePageBlogComponent__content{
		margin-top: 0;
	}
}
@media only screen and (min-width: 610px) and (max-width: 819px) {
	.homePageBlogComponent__grid {
		grid-template-columns: 285px 285px;
		width: 610px;
		margin: 0 auto;
	}
	.homePageBlogComponent__paragraph ,
	.homePageBlogComponent__headline {
		width: 610px;
		max-width: 610px;
		margin-left: auto;
		margin-right: auto;
	}
	.homePageBlogComponent__tileText{
		line-height: 1.3em;
	}
	.homePageBlogComponent__tile,
	.homePageBlogComponent__grid {
		max-height: 100px;
	}
}
@media only screen and (min-width: 400px) and (max-width: 609px) {
	.homePageBlogComponent__grid {
		grid-template-columns: 400px;
		width: 400px;
		margin: 0 auto;
	}
	.homePageBlogComponent__paragraph,
	.homePageBlogComponent__headline {
		width: 400px;
		max-width: 400px;
		margin-left: auto;
		margin-right: auto;
	}
}
@media only screen and (max-width: 399px) {
	.homePageBlogComponent__grid {
		grid-template-columns: 100%;
		width: 100%;
		margin: 0 auto;
		padding: 0 10px;
		box-sizing: border-box;
	}
	.homePageBlogComponent__paragraph,
	.homePageBlogComponent__headline {
		margin-left: 10px;
		margin-right: 10px;
	}
	.homePageBlogComponent__tile {
		max-height: none;
	}
	.homePageBlogComponent__grid {
		grid-gap: 7px;
	}
	.homePageBlogComponent__paragraph {
		line-height: 1.5em;
	}
	.homePageBlogComponent__tileText {
		line-height: 1em;
	}
}
/* HomePageComponent end */
`
}
