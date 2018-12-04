package staticPresentation

import (
	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
)

// Creates a new TwitterComponent
func NewTwitterComponent(r staticIntf.Renderer) *TwitterComponent {
	t := new(TwitterComponent)
	t.abstractComponent.Renderer(r)
	return t
}

type TwitterComponent struct {
	abstractComponent
}

func (tw *TwitterComponent) VisitPage(p staticIntf.Page) {
	m := []*htmlDoc.Node{
		htmlDoc.NewNode("meta", "",
			"name", "t:card",
			"content", tw.abstractComponent.renderer.TwitterCardType()),
		htmlDoc.NewNode("meta", "",
			"name", "t:site",
			"content", tw.abstractComponent.renderer.TwitterHandle()),
		htmlDoc.NewNode("meta", "",
			"name", "t:title",
			"content", p.Title()),
		htmlDoc.NewNode("meta", "",
			"name", "t:text:description",
			"content", p.Description()),
		htmlDoc.NewNode("meta", "",
			"name", "t:creator",
			"content", tw.abstractComponent.renderer.TwitterHandle()),
		htmlDoc.NewNode("meta", "",
			"name", "t:image",
			"content", p.ImageUrl())}
	p.AddHeaderNodes(m)
}
