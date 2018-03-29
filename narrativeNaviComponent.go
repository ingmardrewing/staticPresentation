package staticPresentation

import (
	"path"

	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
)

// Creates anew NarrativeNaviComponent
func NewNarrativeNaviComponent() *NarrativeNaviComponent {
	nc := new(NarrativeNaviComponent)
	return nc
}

type NarrativeNaviComponent struct {
	abstractComponent
	wrapper
	cssClass string
}

func (nv *NarrativeNaviComponent) VisitPage(p staticIntf.Page) {
	firstNode := nv.first(p)
	prevNode := nv.previous(p,
		"&lt; previous page",
		"narrativenavigation__previous narrativenavigation__item narrativenavigation__placeholder")
	nextNode := nv.next(p,
		"next page &gt;",
		"narrativenavigation__next narrativenavigation__item narrativenavigation__placeholder")
	lastNode := nv.last(p)

	nav := htmlDoc.NewNode("nav", "", "class", "narrativenavigation")
	nav.AddChild(firstNode)
	nav.AddChild(prevNode)
	nav.AddChild(nextNode)
	nav.AddChild(lastNode)

	wn := nv.wrap(nav, "narrativenavi__wrapper")
	p.AddBodyNodes([]*htmlDoc.Node{wn})
}

func (nv *NarrativeNaviComponent) first(p staticIntf.Page) *htmlDoc.Node {
	fPage := nv.getFirstPage()
	if fPage == nil || fPage.Id() == p.Id() {
		return htmlDoc.NewNode("span", "&lt;&lt; first page", "class", "narrativenavigation__first narrativenavigation__item narrativenavigation__placeholder")
	}
	href := path.Join(fPage.PathFromDocRoot(), fPage.HtmlFilename())
	return htmlDoc.NewNode("a", "&lt;&lt; first page", "href", href, "rel", "first", "class", "narrativenavigation__first narrativenavigation__item")
}

func (nv *NarrativeNaviComponent) last(p staticIntf.Page) *htmlDoc.Node {
	lPage := nv.getLastPage()
	if lPage == nil || lPage.Id() == p.Id() {
		return htmlDoc.NewNode("span", "last page &gt;&gt;", "class", "narrativenavigation__last narrativenavigation__item narrativenavigation__placeholder")
	}
	href := path.Join(lPage.PathFromDocRoot(), lPage.HtmlFilename())
	return htmlDoc.NewNode("a", "last page &gt;&gt;", "href", href, "rel", "last", "class", "narrativenavigation__last narrativenavigation__item")
}

func (mhc *NarrativeNaviComponent) GetCss() string {
	return `
.narrativenavigation{
	text-align: right;
	margin-bottom: 50px;
}
.narrativenavigation__item {
	font-family: Arial Black, Arial, Helvetica, sans-serif;
	color: grey;
	text-transform: uppercase;
	font-weight: 900;
	font-size: 16px;
}
.narrativenavigation__item.narrativenavigation__placeholder {
	color: lightgrey;
}
.narrativenavigation__item + .narrativenavigation__item {
	margin-left: 10px;
}
`
}
