package staticPresentation

import (
	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
)

// Creates a new NarrativeComponent
func NewNarrativeComponent(r staticIntf.Renderer) *NarrativeComponent {
	n := new(NarrativeComponent)
	n.abstractComponent.Renderer(r)

	return n
}

type NarrativeComponent struct {
	abstractComponent
	wrapper
}

func (cc *NarrativeComponent) VisitPage(p staticIntf.Page) {

	n := htmlDoc.NewNode("main", "", "class", "mainnarrativecontent")
	img := htmlDoc.NewNode("img", "",
		"src", p.ImageUrl(), "width", "800")

	if p.Container() == nil {
		n.AddChild(img)
	} else {
		np := p.Container().GetPageAfter(p)
		if np == nil {
			n.AddChild(img)
		} else {
			a := htmlDoc.NewNode("a", "",
				"href", np.Link())
			a.AddChild(img)
			n.AddChild(a)
		}
	}

	wn := cc.wrap(n)
	p.AddBodyNodes([]*htmlDoc.Node{wn})
}
