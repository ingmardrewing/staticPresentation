package staticPresentation

import (
	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
)

// Creates a new FBComponent
type FBComponent struct {
	abstractComponent
}

func NewFBComponent(r staticIntf.Renderer) *FBComponent {
	fb := new(FBComponent)
	fb.abstractComponent.Renderer(r)
	return fb
}

func (fbc *FBComponent) VisitPage(p staticIntf.Page) {
	description := p.Description()
	if len(description) == 0 {
		description = p.Site().Description()
	}
	m := []*htmlDoc.Node{
		htmlDoc.NewNode("meta", "", "property", "og:title", "content", p.Title()),
		htmlDoc.NewNode("meta", "", "property", "og:url", "content", p.Link()),
		htmlDoc.NewNode("meta", "", "property", "og:image", "content", p.ImageUrl()),
		htmlDoc.NewNode("meta", "", "property", "og:description", "content", description),
		htmlDoc.NewNode("meta", "", "property", "og:site_name", "content", p.Site().Domain()),
		htmlDoc.NewNode("meta", "", "property", "og:type", "content", p.Site().Section()),
		htmlDoc.NewNode("meta", "", "property", "article:published_time", "content", p.PublishedTime()),
		htmlDoc.NewNode("meta", "", "property", "article:modified_time", "content", p.PublishedTime()),
		htmlDoc.NewNode("meta", "", "property", "article:section", "content", p.Site().Section()),
		htmlDoc.NewNode("meta", "", "property", "article:tag", "content", p.Site().Tags())}

	p.AddHeaderNodes(m)
}
