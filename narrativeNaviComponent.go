package staticPresentation

import (
	"path"

	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
)

// Creates anew NarrativeNaviComponent
func NewNarrativeNaviComponent(r staticIntf.Renderer) *NarrativeNaviComponent {
	nc := new(NarrativeNaviComponent)
	nc.abstractComponent.Renderer(r)
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
	return nv.absRel(p, nv.getFirstPage(),
		"&lt;&lt; first page",
		"narrativenavigation__last narrativenavigation__item narrativenavigation__placeholder", "fist")
}

func (nv *NarrativeNaviComponent) last(p staticIntf.Page) *htmlDoc.Node {
	return nv.absRel(p, nv.getLastPage(),
		"last page &gt;&gt;",
		"narrativenavigation__last narrativenavigation__item narrativenavigation__placeholder", "last")
}

func (nv *NarrativeNaviComponent) absRel(curPage, relPage staticIntf.Page, label, class, rel string) *htmlDoc.Node {
	if relPage == nil || relPage.Id() == curPage.Id() {
		return htmlDoc.NewNode("span", label, "class", class)
	}
	href := path.Join(relPage.PathFromDocRoot(), relPage.HtmlFilename())
	return htmlDoc.NewNode("a", label, "href", href, "rel", rel, "class", class)
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
span.narrativenavigation__item {
	color: lightgrey;
}
.narrativenavigation__item + .narrativenavigation__item {
	margin-left: 10px;
}
`
}
