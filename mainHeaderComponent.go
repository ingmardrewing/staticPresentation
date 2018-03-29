package staticPresentation

import (
	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
)

// Creates a new MainHeaderComponent
func NewMainHeaderComponent() *MainHeaderComponent {
	mhc := new(MainHeaderComponent)
	return mhc
}

type MainHeaderComponent struct {
	abstractComponent
	wrapper
}

func (mhc *MainHeaderComponent) VisitPage(p staticIntf.Page) {
	logo := htmlDoc.NewNode("a", "<!-- logo -->",
		"href", "https://"+mhc.abstractComponent.renderer.SiteName(),
		"class", "headerbar__logo")
	logocontainer := htmlDoc.NewNode("div", "",
		"class", "headerbar__logocontainer")
	logocontainer.AddChild(logo)

	header := htmlDoc.NewNode("header", "", "class", "headerbar")
	header.AddChild(logocontainer)

	wn := mhc.wrap(header, "headerbar__wrapper")
	p.AddBodyNodes([]*htmlDoc.Node{wn})
}

func (mhc *MainHeaderComponent) GetCss() string {
	return `
.headerbar__wrapper {
	position: fixed;
	width: 100%;
	top: 0;
	background-color: white;
}
.headerbar__logo {
	background-image: url(https://s3.amazonaws.com/drewingdeblog/drewing_de_logo.png);
	background-repeat: no-repeat;
	background-position: center center;
	display: block;
	width: 100%;
	height: 80px;
}
.headerbar__navelement {
	display: inline-block;
	font-family: Arial Black, Arial, Helvetica, sans-serif;
	font-weight: 900;
	font-size: 18px;
	line-height: 20px;
	text-transform: uppercase;
	text-decoration: none;
	color: black;
	padding: 10px 20px;
}
`
}
