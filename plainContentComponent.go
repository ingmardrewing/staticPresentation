package staticPresentation

import (
	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
)

// Creates a new PlainContentComponent
func NewPlainContentComponent(r staticIntf.Renderer) *PlainContentComponent {
	p := new(PlainContentComponent)
	p.abstractComponent.Renderer(r)
	return p
}

type PlainContentComponent struct {
	abstractComponent
	wrapper
}

func (cc *PlainContentComponent) VisitPage(p staticIntf.Page) {
	n := htmlDoc.NewNode("main", p.Content(),
		"class", "narrativemarginal")
	wn := cc.wrap(n)
	p.AddBodyNodes([]*htmlDoc.Node{wn})
}

func (cc *PlainContentComponent) GetCss() string {
	return `
@media only screen and (max-width: 768px) {
	h1 ,
	.narrativemarginal h1 ,
	.narrativemarginal h2 ,
	.narrativemarginal h3 {
		text-transform: uppercase;
		font-family: Arial Black, Arial, Helvetica, sans-serif;
	}
	.narrativemarginal{
		padding-top: 0;
		padding-bottom: 50px;
		text-align: left;
		line-height: 20px;
	}
	.narrativemarginal li,
	.narrativemarginal p {
		line-height: 30px;
	}
	.narrativemarginal h1,
	.narrativemarginal h2 {
		text-transform: uppercase;
	}
	.narrativemarginal__h1,
	.narrativemarginal__h2 {
		display: inline-block;
		font-family: Arial Black, Arial, Helvetica, sans-serif;
		text-transform: uppercase;
	}
	.narrativemarginal__h1 ,
	.narrativemarginal__h2 {
		font-size: 18px;
		line-height: 20px;
		text-transform: uppercase;
	}
	.narrativemarginal__h2 {
		color: grey;
		margin-left: 10px;
	}
}
@media only screen and (min-width: 769px) {
	h1 ,
	.narrativemarginal h1 ,
	.narrativemarginal h2 ,
	.narrativemarginal h3 {
		text-transform: uppercase;
		font-family: Arial Black, Arial, Helvetica, sans-serif;
	}
	.narrativemarginal{
		padding-top: 0;
		padding-bottom: 50px;
		text-align: left;
		line-height: 20px;
	}
	.narrativemarginal li,
	.narrativemarginal p {
		line-height: 30px;
	}
	.narrativemarginal h1,
	.narrativemarginal h2 {
		text-transform: uppercase;
	}
	.narrativemarginal__h1,
	.narrativemarginal__h2 {
		display: inline-block;
		font-family: Arial Black, Arial, Helvetica, sans-serif;
		text-transform: uppercase;
	}
	.narrativemarginal__h1 ,
	.narrativemarginal__h2 {
		font-size: 18px;
		line-height: 20px;
		text-transform: uppercase;
	}
	.narrativemarginal__h2 {
		color: grey;
		margin-left: 10px;
	}
}
`
}
