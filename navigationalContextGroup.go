package staticPresentation

import (
	"path"
	"strconv"

	"github.com/ingmardrewing/fs"
	"github.com/ingmardrewing/staticIntf"
	"github.com/ingmardrewing/staticModel"
)

type navigationalContextGroup struct {
	abstractContextGroup
	naviContext staticIntf.SubContext
}

func (n *navigationalContextGroup) GetComponents() []staticIntf.Component {
	components := n.context.GetComponents()
	return append(components, n.naviContext.GetComponents()...)
}

func (n *navigationalContextGroup) RenderPages() []fs.FileContainer {
	pages := n.context.RenderPages()
	return append(pages, n.naviContext.RenderPages()...)
}

func (n *navigationalContextGroup) Init() {
	bundles := n.generateBundles()
	last := len(bundles) - 1
	naviPages := []staticIntf.Page{}
	for i, bundle := range bundles {
		filename := "index" + strconv.Itoa(i) + ".html"
		if i == last {
			filename = "index.html"
		}

		p := path.Join(n.naviContext.FsSetOff(), n.naviPagePathFromDocRoot())

		np := staticModel.NewEmptyNaviPage(n.Domain())
		np.NavigatedPages(bundle...)
		np.Title(n.naviPageTitle())
		np.Description(n.naviPageDescription())

		np.Domain(n.Domain())
		np.PathFromDocRoot(p)
		np.HtmlFilename(filename)

		naviPages = append(naviPages, np)
	}
	n.naviContext.SetElements(naviPages)
}

func (n *navigationalContextGroup) getReversedPages() []staticIntf.Page {
	pages := n.context.GetElements()
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
