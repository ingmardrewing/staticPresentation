package staticPresentation

import (
	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
)

// Creates a new CssLinkComponent
func NewCssLinkComponent(r staticIntf.Renderer) *CssLinkComponent {
	clc := new(CssLinkComponent)
	clc.abstractComponent.Renderer(r)
	return clc
}

type CssLinkComponent struct {
	abstractComponent
}

func (clc *CssLinkComponent) VisitPage(p staticIntf.Page) {
	link := htmlDoc.NewNode("link", "", "href", p.Site().Css(), "rel", "stylesheet", "type", "text/css")
	p.AddHeaderNodes([]*htmlDoc.Node{link})
}
