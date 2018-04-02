package staticPresentation

import (
	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
)

// Creates a new EntryPageComponent
func NewEntryPageComponent() *EntryPageComponent {
	return new(EntryPageComponent)
}

type EntryPageComponent struct {
	abstractComponent
	wrapper
}

func (e *EntryPageComponent) VisitPage(p staticIntf.Page) {
	containers := e.renderer.Site().Containers()

	div := htmlDoc.NewNode("div", "", "class", "mainpage__content")
	for _, c := range containers {
		reps := c.Representationals()
		if len(reps) > 0 {
			div.AddChild(htmlDoc.NewNode("p", c.Variant()))
			for _, r := range reps {
				p := htmlDoc.NewNode("p", "")
				a := htmlDoc.NewNode("a", r.Title(),
					"href", r.PathFromDocRootWithName())
				p.AddChild(a)
				div.AddChild(p)
			}
		}
	}
	wn := e.wrap(div)
	p.AddBodyNodes([]*htmlDoc.Node{wn})
}

func (e *EntryPageComponent) GetCss() string {
	return `
.mainpage__content {
	padding-top: 146px;
	padding-bottom: 50px;
	text-align: left;
	min-height: calc(100vh - 520px);
}
`
}
