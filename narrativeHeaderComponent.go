package staticPresentation

import (
	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
)

// Creates a new NarrativeHeaderComponent
func NewNarrativeHeaderComponent(r staticIntf.Renderer) *NarrativeHeaderComponent {
	nc := new(NarrativeHeaderComponent)
	nc.abstractComponent.Renderer(r)
	return nc
}

type NarrativeHeaderComponent struct {
	abstractComponent
	wrapper
	cssClass string
}

func (nv *NarrativeHeaderComponent) VisitPage(p staticIntf.Page) {
	a1 := htmlDoc.NewNode("a", "<!-- Devabo.de-->", "href", "https://devabo.de", "class", "home")
	a2 := htmlDoc.NewNode("a", "New Reader? Start here!", "href", "https://devabo.de/2013/08/01/a-step-in-the-dark/", "class", "orange")
	h1 := htmlDoc.NewNode("h1", p.Title(), "class", "maincontent__h1")

	n := htmlDoc.NewNode("header", "")
	n.AddChild(a1)
	n.AddChild(a2)
	n.AddChild(h1)

	wn := nv.wrap(n, "header__wrapper")
	p.AddBodyNodes([]*htmlDoc.Node{wn})
}

func (mhc *NarrativeHeaderComponent) GetCss() string {
	return `body {
	padding-top: 0;
}

.header__wrapper {
	margin-top: 0;
}

header .home {
    display: block;
    line-height: 80px;
    height: 30px;
    width: 800px;
    text-align: left;
    color: rgb(0, 0, 0);
    margin-bottom: 0px;
    margin-top: 0px;
    background: url(https://devabo.de/imgs/header_devabo_de.png) 0px 0px no-repeat transparent;
}

header .orange {
    display: block;
    height: 2.2em;
    background-color: rgb(255, 136, 0);
    color: rgb(255, 255, 255);
    line-height: 1em;
    box-sizing: border-box;
    width: 100%;
    font-size: 24px;
	font-weight: 700;
    text-transform: uppercase;
    margin-bottom: 1rem;
    padding: 0.5em;
    text-decoration: underline;
}

header {
	text-align: left;
}
`
}
