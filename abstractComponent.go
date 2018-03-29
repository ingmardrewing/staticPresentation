package staticPresentation

import (
	"strings"

	"github.com/ingmardrewing/htmlDoc"
	"github.com/ingmardrewing/staticIntf"
)

// abstractComponent implementing default functions
// for implementing components
type abstractComponent struct {
	renderer staticIntf.Renderer
}

func (ac *abstractComponent) Renderer(r ...staticIntf.Renderer) staticIntf.Renderer {
	if len(r) == 1 {
		ac.renderer = r[0]
	}
	return ac.renderer
}

func (a *abstractComponent) GetCss() string { return "" }

func (a *abstractComponent) GetJs() string { return "" }

func (a *abstractComponent) VisitPage(p staticIntf.Page) {}

func (a *abstractComponent) getIndexOfPage(p staticIntf.Page) int {
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

func (a *abstractComponent) getPageAfter(p staticIntf.Page) staticIntf.Page {
	index := a.getIndexOfPage(p)
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
