package staticPresentation

import (
	"fmt"
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

func (a *abstractComponent) previousFromDocRoot(
	p staticIntf.Page,
	label, class string) *htmlDoc.Node {

	if p.Container() != nil {
		pageBefore := p.Container().GetPageBefore(p)
		return a.abs(pageBefore, label, class, "prev")
	}
	return a.abs(nil, label, class, "prev")
}

func (a *abstractComponent) nextFromDocRoot(
	p staticIntf.Page,
	label, class string) *htmlDoc.Node {

	if p.Container() != nil {
		pageAfter := p.Container().GetPageAfter(p)
		return a.abs(pageAfter, label, class, "next")
	}
	return a.abs(nil, label, class, "next")
}

func (a *abstractComponent) abs(relativePage staticIntf.Page, label, class, rel string) *htmlDoc.Node {
	if relativePage == nil {
		return htmlDoc.NewNode(
			"span", label,
			"class", class)
	}
	href := relativePage.Link()
	return htmlDoc.NewNode(
		"a", label,
		"href", href,
		"rel", rel,
		"class", class)
}

func (a *abstractComponent) previous(
	p staticIntf.Page,
	label, class string) *htmlDoc.Node {

	if p.Container() != nil {
		pageBefore := p.Container().GetPageBefore(p)
		return a.rel(pageBefore, label, class, "prev")
	} else {
		fmt.Println("-- container == nil")
	}
	return a.rel(nil, label, class, "prev")
}

func (a *abstractComponent) next(
	p staticIntf.Page,
	label, class string) *htmlDoc.Node {

	if p.Container() != nil {
		pageBefore := p.Container().GetPageAfter(p)
		return a.rel(pageBefore, label, class, "next")
	}
	return a.rel(nil, label, class, "next")
}

func (a *abstractComponent) rel(relativePage staticIntf.Page, label, class, rel string) *htmlDoc.Node {
	if relativePage == nil {
		return htmlDoc.NewNode(
			"span", label,
			"class", class)
	}
	href := relativePage.Link()
	return htmlDoc.NewNode(
		"a", label,
		"href", href,
		"rel", rel,
		"class", class)
}

// wrapper struct used to generate extra html nodes
type wrapper struct{}

func (w *wrapper) wrap(n *htmlDoc.Node, addedclasses ...string) *htmlDoc.Node {
	inner := htmlDoc.NewNode("div", "", "class", "wrapperInner")
	inner.AddChild(n)
	classes := "wrapperOuter " + strings.Join(addedclasses, " ")
	wrapperNode := htmlDoc.NewNode("div", "", "class", classes)
	wrapperNode.AddChild(inner)
	return wrapperNode
}
