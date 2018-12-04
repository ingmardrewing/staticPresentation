package staticPresentation

import (
	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
)

// Creates a new Title component
func NewTitleComponent(r staticIntf.Renderer) *TitleComponent {
	t := new(TitleComponent)
	t.abstractComponent.Renderer(r)
	return t
}

type TitleComponent struct {
	abstractComponent
}

func (tc *TitleComponent) VisitPage(p staticIntf.Page) {
	title := htmlDoc.NewNode("title", p.Title())
	p.AddHeaderNodes([]*htmlDoc.Node{title})
}
