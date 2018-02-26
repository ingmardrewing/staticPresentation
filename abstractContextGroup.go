package staticPresentation

import (
	"strconv"

	"github.com/ingmardrewing/fs"
	"github.com/ingmardrewing/staticIntf"
	"github.com/ingmardrewing/staticModel"
)

type abstractContextGroup struct {
	pages        []staticIntf.Page
	pagesContext staticIntf.Context
}

func (a *abstractContextGroup) GetComponents() []staticIntf.Component {
	return a.pagesContext.GetComponents()
}

func (a *abstractContextGroup) RenderPages(dir string) []fs.FileContainer {
	return a.pagesContext.RenderPages(dir)
}

func (a *abstractContextGroup) Domain() string {
	return ""
}

func (a *abstractContextGroup) naviPageDescription() string {
	return ""
}

func (a *abstractContextGroup) naviPageTitle() string {
	return ""
}

func (a *abstractContextGroup) naviPagePathFromDocRoot() string {
	return ""
}

//
type navigationalContextGroup struct {
	abstractContextGroup
	naviContext staticIntf.Context
}

func (n *navigationalContextGroup) GetComponents() []staticIntf.Component {
	components := n.pagesContext.GetComponents()
	return append(components, n.naviContext.GetComponents()...)
}

func (n *navigationalContextGroup) RenderPages(dir string) []fs.FileContainer {
	pages := n.pagesContext.RenderPages(dir)
	return append(pages, n.naviContext.RenderPages(dir)...)
}

func (n *navigationalContextGroup) generateNaviPages() {
	bundles := n.generateBundles()
	last := len(bundles) - 1
	naviPages := []staticIntf.Page{}
	for i, bundle := range bundles {
		filename := "index" + strconv.Itoa(i) + ".html"
		if i == last {
			filename = "index.html"
		}

		np := staticModel.NewNaviPage()
		np.NavigatedPages(bundle...)
		np.Domain(n.Domain())
		np.Title(n.naviPageTitle())
		np.Description(n.naviPageDescription())
		np.HtmlFilename(filename)
		np.PathFromDocRoot(n.naviPagePathFromDocRoot())

		naviPages = append(naviPages, np)
	}
	n.naviContext.SetElements(naviPages)
}

func (n *navigationalContextGroup) getReversedPages() []staticIntf.Page {
	pages := n.pagesContext.GetElements()
	length := len(pages)
	reversed := []staticIntf.Page{}
	for i := length - 1; i >= 0; i-- {
		reversed = append(reversed, pages[i])
	}
	return reversed
}

func (n *navigationalContextGroup) generateReversedBundles() []*elementBundle {
	reversed := n.getReversedPages()
	b := newElementBundle()
	bundles := []*elementBundle{}
	for _, p := range reversed {
		b.addElement(p)
		if b.full() {
			bundles = append(bundles, b)
			b = newElementBundle()
		}
	}
	if !b.full() {
		bundles = append(bundles, b)
	}
	return bundles
}

func (n *navigationalContextGroup) generateBundles() [][]staticIntf.Page {
	reversedBundles := n.generateReversedBundles()
	length := len(reversedBundles)
	pageBundles := [][]staticIntf.Page{}
	for i := length - 1; i >= 0; i-- {
		pageBundles = append(pageBundles, reversedBundles[i].getElements())
	}
	return pageBundles
}