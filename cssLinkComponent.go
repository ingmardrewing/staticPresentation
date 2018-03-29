package staticPresentation

import (
	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
)

// Creates a new CssLinkComponent
func NewCssLinkComponent() *CssLinkComponent {
	clc := new(CssLinkComponent)
	return clc
}

type CssLinkComponent struct {
	abstractComponent
}

func (clc *CssLinkComponent) VisitPage(p staticIntf.Page) {
	link := htmlDoc.NewNode("link", "", "href", clc.abstractComponent.renderer.CssUrl(), "rel", "stylesheet", "type", "text/css")
	p.AddHeaderNodes([]*htmlDoc.Node{link})
}
