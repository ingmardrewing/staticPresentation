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
	naviContext  staticIntf.Context
}

func (a *abstractContextGroup) generateNaviPages() {
	bundles := a.generateBundles()
	last := len(bundles) - 1
	naviPages := []staticIntf.Page{}
	for i, bundle := range bundles {
		filename := "index" + strconv.Itoa(i) + ".html"
		if i == last {
			filename = "index.html"
		}

		np := staticModel.NewNaviPage()
		np.NavigatedPages(bundle...)
		np.Domain(a.Domain())
		np.Title(a.naviPageTitle())
		np.Description(a.naviPageDescription())
		np.HtmlFilename(filename)
		np.PathFromDocRoot(a.naviPagePathFromDocRoot())

		naviPages = append(naviPages, np)
	}
	a.naviContext.SetElements(naviPages)
}

func (a *abstractContextGroup) getReversedPages() []staticIntf.Page {
	pages := a.pagesContext.GetElements()
	length := len(pages)
	reversed := []staticIntf.Page{}
	for i := length - 1; i >= 0; i-- {
		reversed = append(reversed, pages[i])
	}
	return reversed
}

func (a *abstractContextGroup) generateReversedBundles() []*elementBundle {
	reversed := a.getReversedPages()
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

func (a *abstractContextGroup) generateBundles() [][]staticIntf.Page {
	reversedBundles := a.generateReversedBundles()
	length := len(reversedBundles)
	pageBundles := [][]staticIntf.Page{}
	for i := length - 1; i >= 0; i-- {
		pageBundles = append(pageBundles, reversedBundles[i].getElements())
	}
	return pageBundles
}

func (a *abstractContextGroup) GetComponents() []staticIntf.Component {
	components := a.pagesContext.GetComponents()
	return append(components, a.naviContext.GetComponents()...)
}

func (a *abstractContextGroup) RenderPages(dir string) []fs.FileContainer {
	pages := a.pagesContext.RenderPages(dir)
	return append(pages, a.naviContext.RenderPages(dir)...)
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
