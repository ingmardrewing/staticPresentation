package staticPresentation

import (
	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
)

// Creates a new ContentComponent
func NewContentComponent() *ContentComponent {
	return new(ContentComponent)
}

type ContentComponent struct {
	abstractComponent
	wrapper
}

func (cc *ContentComponent) VisitPage(p staticIntf.Page) {
	h1 := htmlDoc.NewNode("h1", p.Title(),
		"class", "maincontent__h1")
	h2 := htmlDoc.NewNode("h2", p.PublishedTime(),
		"class", "maincontent__h2")
	n := htmlDoc.NewNode("main", p.Content(),
		"class", "maincontent")
	n.AddChild(h1)
	n.AddChild(h2)
	wn := cc.wrap(n)
	p.AddBodyNodes([]*htmlDoc.Node{wn})
}

func (cc *ContentComponent) GetCss() string {
	return `
@media only screen and (max-width: 768px) {
	.maincontent{
		padding-bottom: 50px;
		text-align: left;
		line-height: 1.2em;
	}
	.maincontent li,
	.maincontent p {
		line-height: 1.4em;
	}
	.maincontent p {
		padding-left: 10px;
		padding-right: 10px;
	}
	.maincontent__h1,
	.maincontent__h2 {
		text-transform: uppercase;
		display: inline-block;
		font-family: Arial Black, Arial, Helvetica, sans-serif;
		text-transform: uppercase;
		font-size: 18px;
		line-height: 20px;
		text-transform: uppercase;
	}
	.maincontent__h1 {
		margin-left: 10px;
		margin-bottom: 0;
	}
	.maincontent__h1 + .maincontent__h2 {
		color: grey;
		margin-top: 0;
		margin-left: 10px;
	}
	.maincontent__h2 + p ,
	.maincontent__h2 + dl {
		margin-top: 0;
	}
	dd + dt {
		margin-top: 10px;
	}
	.maincontent img {
		max-width: 100%;
		width: auto;
	}
}
@media only screen and (min-width: 769px) {
	.maincontent{
		padding-top: 126px;
		padding-bottom: 50px;
		text-align: left;
		line-height: 20px;
	}
	.maincontent li,
	.maincontent p {
		line-height: 30px;
	}
	.maincontent__h1,
	.maincontent__h2 {
		text-transform: uppercase;
		display: inline-block;
		font-family: Arial Black, Arial, Helvetica, sans-serif;
		text-transform: uppercase;
		font-size: 18px;
		line-height: 20px;
		text-transform: uppercase;
	}
	.maincontent__h1 + .maincontent__h2 {
		color: grey;
		margin-left: 10px;
	}
	.maincontent__h2 + p ,
	.maincontent__h2 + dl {
		margin-top: 0;
	}
	dd + dt {
		margin-top: 10px;
	}
}
`
}
