package staticPresentation

import (
	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
)

// Creates a new StartPageComponent
func NewStartPageComponent() *StartPageComponent {
	return new(StartPageComponent)
}

// start page component
type StartPageComponent struct {
	abstractComponent
	wrapper
}

func (cc *StartPageComponent) VisitPage(p staticIntf.Page) {
	c1 := htmlDoc.NewNode("div", "portfoliocontent", "class", "home__portfolio")
	c2 := htmlDoc.NewNode("div", "devabode", "class", "home__devabode")
	c3 := htmlDoc.NewNode("div", "blog", "class", "home__blog")

	n := htmlDoc.NewNode("main", "", "class", "maincontent")
	n.AddChild(c1)
	n.AddChild(c2)
	n.AddChild(c3)

	wn := cc.wrap(n)
	p.AddBodyNodes([]*htmlDoc.Node{wn})
}
