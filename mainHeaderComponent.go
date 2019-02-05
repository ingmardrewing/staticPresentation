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
		"a", "<!-- logo -->",
		"href", "https://"+p.Site().Domain(),
		"class", "headerbar__logo")
	logocontainer := htmlDoc.NewNode(
		"div", "",
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
@media only screen and (max-width: 768px) {
	.headerbar__wrapper {
		width: 100%;
		background-color: white;
	}
	.headerbar__logo {
		background-image: url(https://s3.amazonaws.com/drewingdeblog/drewing_de_logo.png);
		background-repeat: no-repeat;
		background-position: center center;
		background-size: 60%;
		display: block;
		width: 100%;
		height: 40px;
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
