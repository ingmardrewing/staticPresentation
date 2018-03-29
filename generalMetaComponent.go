package staticPresentation

import (
	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
)

// Creates a new GeneralMetaComponent
func NewGeneralMetaComponent() *GeneralMetaComponent {
	return new(GeneralMetaComponent)
}

type GeneralMetaComponent struct {
	abstractComponent
}

func (g *GeneralMetaComponent) VisitPage(p staticIntf.Page) {
	m := []*htmlDoc.Node{
		htmlDoc.NewNode("meta", "", "name", "viewport", "content", "width=device-width, initial-scale=1.0"),
		htmlDoc.NewNode("meta", "", "name", "robots", "content", "index,follow"),
		htmlDoc.NewNode("meta", "", "name", "author", "content", "Ingmar Drewing"),
		htmlDoc.NewNode("meta", "", "name", "publisher", "content", "Ingmar Drewing"),
		htmlDoc.NewNode("meta", "", "name", "keywords", "content", "storytelling, illustration, drawing, web comic, comic, cartoon, caricatures"),
		htmlDoc.NewNode("meta", "", "name", "DC.subject", "content", "storytelling, illustration, drawing, web comic, comic, cartoon, caricatures"),
		htmlDoc.NewNode("meta", "", "name", "page-topic", "content", "art"),
		htmlDoc.NewNode("meta", "", "charset", "UTF-8")}
	p.AddHeaderNodes(m)
}
