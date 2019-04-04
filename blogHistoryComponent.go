package staticPresentation

import (
	"strings"

	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
	"github.com/ingmardrewing/staticUtil"
)

func NewBlogHistoryComponent(r staticIntf.Renderer) *BlogHistoryComponent {
	h := new(BlogHistoryComponent)
	h.abstractComponent.Renderer(r)
	return h
}

type BlogHistoryComponent struct {
	abstractComponent
	wrapper
	mainDiv *htmlDoc.Node
}

func (e *BlogHistoryComponent) VisitPage(p staticIntf.Page) {
	e.mainDiv = htmlDoc.NewNode("div", "", "class", "blogHistoryComponent__content")
	grid := htmlDoc.NewNode(
		"div", " ",
		"class", "blogHistoryComponent__grid")

	limit := 100
	for _, year := range e.getAllYears(p) {
		grid.AddChild(e.getYearHeadline(year))
		for i, p := range e.getBlogPostsReversedByYear(year, p) {
			if i == limit {
				break
			}
			grid.AddChild(e.getElementLinkingToPages(p))
		}
	}

	e.mainDiv.AddChild(grid)

	w := e.wrap(e.mainDiv, "blogHistoryComponent__wrapperouter")
	p.AddBodyNodes([]*htmlDoc.Node{w})
}

func (e *BlogHistoryComponent) getYearHeadline(year string) *htmlDoc.Node {
	return htmlDoc.NewNode(
		"h2", year,
		"class", "blogHistoryComponent__year")
}

func (e *BlogHistoryComponent) getAllYears(p staticIntf.Page) []string {
	years := []string{}
	for _, p := range e.getAllPosts(p) {
		years = append(years, e.getYear(p))
	}
	return years
}

func (e *BlogHistoryComponent) getAllPosts(p staticIntf.Page) []staticIntf.Page {
	tool := staticUtil.NewPagesContainerCollectionTool(p.Site())
	containers := tool.ContainersOrderedByVariants("blog")

	if len(containers) > 0 {
		return containers[0].Pages()
	}
	return []staticIntf.Page{}
}

func (e *BlogHistoryComponent) getBlogPostsReversedByYear(year string, p staticIntf.Page) []staticIntf.Page {
	posts := e.getPostsByYear(year, p)
	for i := len(posts)/2 - 1; i >= 0; i-- {
		opp := len(posts) - 1 - i
		posts[i], posts[opp] = posts[opp], posts[i]
	}
	return posts
}

func (e *BlogHistoryComponent) getPostsByYear(year string, p staticIntf.Page) []staticIntf.Page {
	posts := []staticIntf.Page{}
	for _, p := range e.getAllPosts(p) {
		if e.getYear(p) == year {
			posts = append(posts, p)
		}
	}
	return posts
}

func (e *BlogHistoryComponent) getYear(p staticIntf.Page) string {
	parts := strings.Split(p.PublishedTime(), "-")
	if len(parts) > 0 {
		return parts[0]
	}
	return ""
}

func (e *BlogHistoryComponent) getElementLinkingToPages(
	page staticIntf.Page) *htmlDoc.Node {

	return htmlDoc.NewNode(
		"a", page.PublishedTime(),
		"href", page.Link(),
		"title", page.Title(),
		"class", "blogHistoryComponent__tile")
}

func (e *BlogHistoryComponent) GetCss() string {
	return `.blogHistoryComponent__tile {
	display: block;
}
`
}
