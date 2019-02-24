package staticPresentation

import (
	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
)

// Creates a new GoogleComponent
func NewGoogleComponent(r staticIntf.Renderer) *GoogleComponent {
	gc := new(GoogleComponent)
	gc.abstractComponent.Renderer(r)
	return gc
}

type GoogleComponent struct {
	abstractComponent
}

func (goo *GoogleComponent) VisitPage(p staticIntf.Page) {
	description := p.Description()
	if len(description) == 0 {
		description = p.Site().Description()
	}
	m := []*htmlDoc.Node{
		htmlDoc.NewNode("meta", "", "itemprop", "name", "content", p.Title()),
		htmlDoc.NewNode("meta", "", "itemprop", "description", "content", description),
		htmlDoc.NewNode("meta", "", "itemprop", "image", "content", p.ImageUrl())}
	p.AddHeaderNodes(m)
}
