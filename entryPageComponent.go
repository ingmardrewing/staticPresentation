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

func (cc *EntryPageComponent) VisitPage(p staticIntf.Page) {

	//	n.Renderer().
	n := htmlDoc.NewNode("main", p.Content(),
		"class", "narrativemarginal")
	wn := cc.wrap(n)
	p.AddBodyNodes([]*htmlDoc.Node{wn})
}
