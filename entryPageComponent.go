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

	mainDiv := htmlDoc.NewNode("div", "", "class", "mainpage__content")
	for _, c := range containers {
		reps := c.Representationals()
		if len(reps) > 0 {
			block := htmlDoc.NewNode("div", "",
				"class", "mainpage_block")
			blockHeadline := htmlDoc.NewNode("h2", c.Variant())
			block.AddChild(blockHeadline)
			for i := len(reps) - 1; i >= 0; i-- {
				a := htmlDoc.NewNode("a", "",
					"href", reps[i].PathFromDocRootWithName(),
					"title", reps[i].Title(),
					"class", "mainpage__thumb",
					"style", "background-image: url("+reps[i].ThumbnailUrl()+")")
				block.AddChild(a)
			}
			mainDiv.AddChild(block)
		}
	}
	wn := e.wrap(mainDiv)
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
.mainpage__thumb {
	display: block;
	float: left;
	width: 200px;
	height: 200px;
	background-size: cover;
	margin: 10px;
}
`
}
