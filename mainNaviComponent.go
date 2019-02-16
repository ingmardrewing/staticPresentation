package staticPresentation

import (
	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
)

// Creates a new MainNaviComponent
func NewMainNaviComponent(r staticIntf.Renderer) *MainNaviComponent {
	nc := new(MainNaviComponent)
	nc.abstractComponent.Renderer(r)
	return nc
}

type MainNaviComponent struct {
	abstractComponent
	wrapper
	cssClass string
}

func (nv *MainNaviComponent) VisitPage(p staticIntf.Page) {
	nav := htmlDoc.NewNode(
		"nav", "",
		"class", "mainnavi")
	for _, l := range p.Site().Main() {
		if len(l.ExternalLink()) > 0 {
			a := htmlDoc.NewNode(
				"a", l.Title(),
				"href", l.ExternalLink(),
				"class", "mainnavi__navelement")
			nav.AddChild(a)
		} else if p.Url() == l.Url() {
			span := htmlDoc.NewNode(
				"span", l.Title(),
				"class", "mainnavi__navelement--current")
			nav.AddChild(span)
		} else {
			a := htmlDoc.NewNode(
				"a", l.Title(),
				"href", l.PathFromDocRootWithName(),
				"class", "mainnavi__navelement")
			nav.AddChild(a)
		}
	}
	node := htmlDoc.NewNode(
		"div", "",
		"class", nv.cssClass)
	node.AddChild(nav)
	wn := nv.wrap(node, "mainnavi__wrapper")
	p.AddBodyNodes([]*htmlDoc.Node{wn})
}

func (mhc *MainNaviComponent) GetCss() string {
	return `
@media only screen and (max-width: 768px) {
	.mainnavi {
		border-bottom: 1px solid black;
		border-top: 1px solid black;
	}
	.mainnavi__wrapper {
		width: 100%;
		background-color: white;
	}
	.mainnavi__navelement--current,
	a.mainnavi__navelement {
		display: block;
		font-weight: 700;
		font-size: 18px;
		line-height: 20px;
		text-transform: uppercase;
		color: black;
		padding: 10px 20px;
	}
	.mainnavi__navelement--current,
	a.mainnavi__navelement:hover {
		text-decoration: none;
		color: gray;
	}
	.mainnavi__navelement + .mainnavi__navelement--current,
	.mainnavi__navelement--current + .mainnavi__navelement,
	.mainnavi__navelement + .mainnavi__navelement {
		border-top: 1px solid black;
	}
}
@media only screen and (min-width: 769px) {
	.mainnavi {
		border-top: 1px solid black;
		border-bottom: 1px solid black;
	}
	.mainnavi__wrapper {
		z-index: 100;
		position: fixed;
		width: 100%;
		top: 80px;
		background-color: white;
	}
	.mainnavi__navelement--current,
	a.mainnavi__navelement {
		display: inline-block;
		font-weight: 700;
		font-size: 18px;
		line-height: 20px;
		text-transform: uppercase;
		color: black;
		padding: 10px 20px;
	}
	.mainnavi__navelement--current,
	a.mainnavi__navelement:hover {
		text-decoration: none;
		color: gray;
	}
	.mainnavi__nav {
		border-bottom: 1px solid black;
	}
}
`
}
