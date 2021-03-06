package staticPresentation

import (
	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/rx"
	"github.com/ingmardrewing/staticIntf"
	log "github.com/sirupsen/logrus"
)

// Creates a new ContentComponent
func NewContentComponent(r staticIntf.Renderer) *ContentComponent {
	c := new(ContentComponent)
	c.abstractComponent.Renderer(r)
	return c
}

type ContentComponent struct {
	abstractComponent
	wrapper
}

func (cc *ContentComponent) VisitPage(p staticIntf.Page) {
	log.Debug("rendering: " + p.Link())
	h1 := htmlDoc.NewNode("h1", p.Title(),
		"class", "maincontent__h1")
	date := cc.getDateFromPublishedTime(p.PublishedTime())
	h2 := htmlDoc.NewNode("h2", date,
		"class", "maincontent__h2")
	n := htmlDoc.NewNode("main", p.Content(),
		"class", "maincontent")
	n.AddChild(h1)
	n.AddChild(h2)
	wn := cc.wrap(n)
	p.AddBodyNodes([]*htmlDoc.Node{wn})
}

func (cc *ContentComponent) getDateFromPublishedTime(publishedTime string) string {
	rgx, err := rx.NewRx(" [0-9]{2}:[0-9]{2}:[0-9]{2}.*$")
	if err != nil {
		log.Error(err)
	}
	return rgx.SubstituteAllOccurences(publishedTime, "")
}

func (cc *ContentComponent) GetCss() string {
	return `
.maincontent__h1 ,
.maincontent__h2 {
	margin-top: 0;
}
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
		padding-bottom: 50px;
		text-align: left;
		line-height: 20px;
	}
	.maincontent p {
		line-height: 30px;
	}
	.maincontent h2,
	.maincontent__h1,
	.maincontent__h2 {
		text-transform: uppercase;
		display: inline-block;
		text-transform: uppercase;
		font-size: 18px;
		line-height: initial;
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
