package staticPresentation

import (
	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
)

// Creates a new NarrativeArchiveComponent
func NewNarrativeArchiveComponent(r staticIntf.Renderer) *NarrativeArchiveComponent {
	n := new(NarrativeArchiveComponent)
	n.abstractComponent.Renderer(r)
	return n
}

type NarrativeArchiveComponent struct {
	abstractComponent
	wrapper
	categoryNames []string
}

func (na *NarrativeArchiveComponent) addCategory(c string) {
	for _, n := range na.categoryNames {
		if n == c {
			return
		}
	}
	na.categoryNames = append(na.categoryNames, c)
}

func (na *NarrativeArchiveComponent) VisitPage(p staticIntf.Page) {

	div := htmlDoc.NewNode("div", " ", "style", "text-align:left;")

	categories := make(map[string][]staticIntf.Page)
	for _, page := range p.NavigatedPages() {
		c := page.Category()
		if len(c) == 0 {
			c = "-"
		}
		na.addCategory(c)
		categories[c] = append(categories[c], page)
	}

	for _, name := range na.categoryNames {

		pages := categories[name]
		h2 := htmlDoc.NewNode("h2", name)
		ul := htmlDoc.NewNode("ul", "", "class", "narrativearchive__wrapper")
		for _, page := range pages {
			span := htmlDoc.NewNode("span", page.Title(), "class", "narrativearchive__title")
			img := htmlDoc.NewNode("img", "", "src", "data:image/png;base64,"+page.ThumbBase64(), "alt", page.Title())
			a := htmlDoc.NewNode("a", "", "href", page.Link(), "class", "narrativearchive__link")
			a.AddChild(span)
			a.AddChild(img)
			li := htmlDoc.NewNode("li", "", "class", "narrativearchive__tile")
			li.AddChild(a)
			ul.AddChild(li)
		}
		div.AddChild(h2)
		div.AddChild(ul)
	}

	wn := na.wrap(div, "narrativearchive__metawrapper")
	p.AddBodyNodes([]*htmlDoc.Node{wn})
}

func (na *NarrativeArchiveComponent) GetCss() string {
	return `.narrativearchive__tile{
	display: block;
	width: 200px;
	height: 240px;
	float: left;
	text-align: center;
	margin-bottom: 60px;
}
.narrativearchive__title{
	display: block;
}
.narrativearchive__wrapper {
	padding-left: 0;
}
.narrativearchive__wrapper:after {
	content: "";
	display: table;
	clear: both;
}`
}
