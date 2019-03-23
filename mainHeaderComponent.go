package staticPresentation

import (
	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
)

// Creates a new MainHeaderComponent
func NewMainHeaderComponent(r staticIntf.Renderer) *MainHeaderComponent {
	mhc := new(MainHeaderComponent)
	mhc.abstractComponent.Renderer(r)
	return mhc
}

type MainHeaderComponent struct {
	abstractComponent
	wrapper
}

func (mhc *MainHeaderComponent) VisitPage(p staticIntf.Page) {
	logo := htmlDoc.NewNode(
		"img", "",
		"src", p.Site().SvgLogo(),
		"class", "headerbar__logo")
	logocontainer := htmlDoc.NewNode(
		"a", "",
		"href", "https://"+p.Site().Domain(),
		"class", "headerbar__logocontainer")
	logocontainer.AddChild(logo)

	header := htmlDoc.NewNode(
		"header", "",
		"class", "headerbar")
	header.AddChild(logocontainer)

	wn := mhc.wrap(header, "headerbar__wrapper")
	p.AddBodyNodes([]*htmlDoc.Node{wn})
}

func (mhc *MainHeaderComponent) GetCss() string {
	return `
.headerbar__logocontainer {
	display: block;
	height: 80px;
	box-sizing: border-box;
}
.headerbar__logo {
	height: 100%;
}
.headerbar__wrapper {
	margin-top: 0;
}
@media only screen and (max-width: 768px) {
	.headerbar__wrapper {
		width: 100%;
		background-color: white;
	}
	.headerbar__navelement {
		display: block;
		font-weight: 700;
		font-size: 1.2em;
		line-height: 2em;
		text-transform: uppercase;
		text-decoration: none;
		color: black;
		padding: 10px 20px;
	}
}
@media only screen and (min-width: 769px) {
	.headerbar__wrapper {
		z-index: 100;
		position: fixed;
		width: 100%;
		top: 0;
		background-color: white;
	}
	.headerbar__navelement {
		display: inline-block;
		font-weight: 700;
		font-size: 18px;
		line-height: 20px;
		text-transform: uppercase;
		text-decoration: none;
		color: black;
		padding: 10px 20px;
	}
}
`
}
