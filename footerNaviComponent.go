package staticPresentation

import (
	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
)

// Creates a new FooterNaviComponent
func NewFooterNaviComponent() *FooterNaviComponent {
	nc := new(FooterNaviComponent)
	return nc
}

type FooterNaviComponent struct {
	abstractComponent
	wrapper
	cssClass string
}

func (f *FooterNaviComponent) VisitPage(p staticIntf.Page) {
	nav := htmlDoc.NewNode("nav", "",
		"class", "footernavi")
	for _, l := range f.abstractComponent.renderer.FooterNavigationLocations() {

		if len(l.ExternalLink()) > 0 {
			a := htmlDoc.NewNode("a", l.Title(),
				"href", l.ExternalLink(),
				"class", "footernavi__navelement")
			nav.AddChild(a)
		} else if p.Url() == l.Url() {
			span := htmlDoc.NewNode("span", l.Title(),
				"class", "footernavi__navelement--current")
			nav.AddChild(span)
		} else {
			a := htmlDoc.NewNode("a", l.Title(),
				"href", l.PathFromDocRootWithName(),
				"class", "footernavi__navelement")
			nav.AddChild(a)
		}
	}
	node := htmlDoc.NewNode("div", "", "class", f.cssClass)
	node.AddChild(nav)
	wn := f.wrap(node, "footernavi__wrapper")
	p.AddBodyNodes([]*htmlDoc.Node{wn})
}

func (f *FooterNaviComponent) GetCss() string {
	return `
@media only screen and (max-width: 768px) {
	.footernavi {
		border-top: 1px solid black;
	}
	.footernavi__wrapper {
		width: 100%;
		background-color: white;
		border-top: 1px solid black;
		border-bottom: 1px solid black;
	}
	.footernavi__navelement--current ,
	a.footernavi__navelement {
		display: inline-block;
		font-family: Arial Black, Arial, Helvetica, sans-serif;
		font-weight: 900;
		font-size: 16px;
		line-height: 20px;
		text-transform: uppercase;
		text-decoration: none;
		color: black;
		padding: 10px 15px;
	}
	a.footernavi__navelement:hover,
	.footernavi__navelement--current {
		color: gray;
	}
}
@media only screen and (min-width: 769px) {
	.footernavi {
		border-top: 1px solid black;
	}
	.footernavi__wrapper {
		position: fixed;
		width: 100%;
		bottom: 0;
		background-color: white;
	}
	.footernavi__navelement--current ,
	a.footernavi__navelement {
		display: inline-block;
		font-family: Arial Black, Arial, Helvetica, sans-serif;
		font-weight: 900;
		font-size: 16px;
		line-height: 20px;
		text-transform: uppercase;
		text-decoration: none;
		color: black;
		padding: 10px 15px;
	}
	a.footernavi__navelement:hover,
	.footernavi__navelement--current {
		color: gray;
	}
}
`
}
