package staticPresentation

import (
	"path"
	"strings"

	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
)

// abstractComponent implementing default functions
// for implementing components
type abstractComponent struct {
	renderer staticIntf.Renderer
}

func (ac *abstractComponent) Renderer(r staticIntf.Renderer) {
	ac.renderer = r
}

func (a *abstractComponent) GetCss() string { return "" }

func (a *abstractComponent) GetJs() string { return "" }

func (a *abstractComponent) VisitPage(p staticIntf.Page) {}

func (a *abstractComponent) getIndexOfPage(p staticIntf.Page) int {
	if a.renderer == nil {
		return -1
	}
	for i, l := range a.renderer.Pages() {
		lurl := l.PathFromDocRoot() + l.HtmlFilename()
		purl := p.PathFromDocRoot() + p.HtmlFilename()
		if lurl == purl {
			return i
		}
	}
	return -1
}

func (a *abstractComponent) getFirstPage() staticIntf.Page {
	pages := a.renderer.Pages()
	if len(pages) > 0 {
		return pages[0]
	}
	return nil
}

func (a *abstractComponent) getLastPage() staticIntf.Page {
	pages := a.renderer.Pages()
	if len(pages) > 0 {
		return pages[len(pages)-1]
	}
	return nil
}

func (a *abstractComponent) getPageBefore(p staticIntf.Page) staticIntf.Page {
	index := a.getIndexOfPage(p)
	pages := a.renderer.Pages()
	if index > 0 {
		return pages[index-1]
	}
	return nil
}

func (a *abstractComponent) previousFromDocRoot(
	p staticIntf.Page,
	label, class string) *htmlDoc.Node {

	pageBefore := a.getPageBefore(p)
	return a.abs(pageBefore, label, class, "prev")
}

func (a *abstractComponent) nextFromDocRoot(
	p staticIntf.Page,
	label, class string) *htmlDoc.Node {

	pageBefore := a.getPageAfter(p)
	return a.abs(pageBefore, label, class, "next")
}

func (a *abstractComponent) abs(relativePage staticIntf.Page, label, class, rel string) *htmlDoc.Node {
	if relativePage == nil {
		return htmlDoc.NewNode(
			"span", label,
			"class", class)
	}
	href := "/" + path.Join(relativePage.PathFromDocRoot(),
		relativePage.HtmlFilename())
	return htmlDoc.NewNode(
		"a", label,
		"href", href,
		"rel", rel,
		"class", class)
}

func (a *abstractComponent) previous(
	p staticIntf.Page,
	label, class string) *htmlDoc.Node {

	pageBefore := a.getPageBefore(p)
	return a.rel(pageBefore, label, class, "prev")
}

func (a *abstractComponent) next(
	p staticIntf.Page,
	label, class string) *htmlDoc.Node {

	pageBefore := a.getPageAfter(p)
	return a.rel(pageBefore, label, class, "next")
}

func (a *abstractComponent) rel(relativePage staticIntf.Page, label, class, rel string) *htmlDoc.Node {
	if relativePage == nil {
		return htmlDoc.NewNode(
			"span", label,
			"class", class)
	}
	href := path.Join(relativePage.PathFromDocRoot(),
		relativePage.HtmlFilename())
	return htmlDoc.NewNode(
		"a", label,
		"href", href,
		"rel", rel,
		"class", class)
}

func (a *abstractComponent) getPageAfter(p staticIntf.Page) staticIntf.Page {
	index := a.getIndexOfPage(p)
	if a.renderer == nil {
		return nil
	}
	pages := a.renderer.Pages()
	if index+1 < len(pages) {
		return pages[index+1]
	}
	return nil
}

// wrapper
type wrapper struct{}

func (w *wrapper) wrap(n *htmlDoc.Node, addedclasses ...string) *htmlDoc.Node {
	inner := htmlDoc.NewNode("div", "", "class", "wrapperInner")
	inner.AddChild(n)
	classes := "wrapperOuter " + strings.Join(addedclasses, " ")
	wrapperNode := htmlDoc.NewNode("div", "", "class", classes)
	wrapperNode.AddChild(inner)
	return wrapperNode
}
