package staticPresentation

import (
	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
)

// Creates a new GeneralMetaComponent
func NewGeneralMetaComponent(r staticIntf.Renderer) *GeneralMetaComponent {
	g := new(GeneralMetaComponent)
	g.abstractComponent.Renderer(r)
	return g
}

type GeneralMetaComponent struct {
	abstractComponent
}

func (g *GeneralMetaComponent) VisitPage(p staticIntf.Page) {
	description := p.Description()
	author := p.Site().Author()
	keyWords := p.Site().KeyWords()
	subject := p.Site().Subject()
	topic := p.Site().Topic()

	if len(p.Description()) == 0 {
		description = p.Site().Description()
	}

	viewport := "initial-scale=1, maximum-scale=1, minimum-scale=1, user-scalable=no"
	m := []*htmlDoc.Node{
		htmlDoc.NewNode("meta", "",
			"name", "viewport",
			"content", viewport),
		htmlDoc.NewNode("meta", "",
			"name", "robots",
			"content", "index,follow"),
		htmlDoc.NewNode("meta", "",
			"name", "author",
			"content", author),
		htmlDoc.NewNode("meta", "",
			"name", "publisher",
			"content", author),
		htmlDoc.NewNode("meta", "",
			"name", "keywords",
			"content", keyWords),
		htmlDoc.NewNode("meta", "",
			"name", "description",
			"content", description),
		htmlDoc.NewNode("meta", "",
			"name", "DC.subject",
			"content", subject),
		htmlDoc.NewNode("meta", "",
			"name", "page-topic",
			"content", topic),
		htmlDoc.NewNode("meta", "",
			"charset", "UTF-8")}
	p.AddHeaderNodes(m)
}
