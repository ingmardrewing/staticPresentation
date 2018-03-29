package staticPresentation

import (
	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
)

// Creates a new GoogleComponent
func NewGoogleComponent() *GoogleComponent {
	gc := new(GoogleComponent)
	return gc
}

type GoogleComponent struct {
	abstractComponent
}

func (goo *GoogleComponent) VisitPage(p staticIntf.Page) {
	m := []*htmlDoc.Node{
		htmlDoc.NewNode("meta", "", "itemprop", "name", "content", p.Title()),
		htmlDoc.NewNode("meta", "", "itemprop", "description", "content", p.Description()),
		htmlDoc.NewNode("meta", "", "itemprop", "image", "content", p.ImageUrl())}
	p.AddHeaderNodes(m)
}
