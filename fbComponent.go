package staticPresentation

import (
	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
)

// Creates a new FBComponent
type FBComponent struct {
	abstractComponent
}

func NewFBComponent() *FBComponent {
	fb := new(FBComponent)
	return fb
}

func (fbc *FBComponent) VisitPage(p staticIntf.Page) {
	m := []*htmlDoc.Node{
		htmlDoc.NewNode("meta", "", "property", "og:title", "content", p.Title()),
		htmlDoc.NewNode("meta", "", "property", "og:url", "content", p.PathFromDocRoot()+p.HtmlFilename()),
		htmlDoc.NewNode("meta", "", "property", "og:image", "content", p.ImageUrl()),
		htmlDoc.NewNode("meta", "", "property", "og:description", "content", p.Description()),
		htmlDoc.NewNode("meta", "", "property", "og:site_name", "content", fbc.abstractComponent.renderer.SiteName()),
		htmlDoc.NewNode("meta", "", "property", "og:type", "content", fbc.abstractComponent.renderer.OGType()),
		htmlDoc.NewNode("meta", "", "property", "article:published_time", "content", p.PublishedTime()),
		htmlDoc.NewNode("meta", "", "property", "article:modified_time", "content", p.PublishedTime()),
		htmlDoc.NewNode("meta", "", "property", "article:section", "content", fbc.abstractComponent.renderer.ContentSection()),
		htmlDoc.NewNode("meta", "", "property", "article:tag", "content", fbc.abstractComponent.renderer.ContentTags())}

	p.AddHeaderNodes(m)
}
