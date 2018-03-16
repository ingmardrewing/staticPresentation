package staticPresentation

import (
	"github.com/ingmardrewing/fs"
	"github.com/ingmardrewing/staticIntf"
)

type abstractContextGroup struct {
	pages   []staticIntf.Page
	context staticIntf.SubContext
	site    staticIntf.Site
}

func (a *abstractContextGroup) GetComponents() []staticIntf.Component {
	return a.context.GetComponents()
}

func (a *abstractContextGroup) RenderPages() []fs.FileContainer {
	return a.context.RenderPages()
}

func (a *abstractContextGroup) Domain() string { return a.site.Domain() }

func (a *abstractContextGroup) naviPageDescription() string { return "" }

func (a *abstractContextGroup) naviPageTitle() string { return "" }

func (a *abstractContextGroup) naviPagePathFromDocRoot() string { return "" }

func (a *abstractContextGroup) Init() {}

func (b *blogContextGroup) getLastTenReversedPages() []staticIntf.Page {
	pages := b.context.GetElements()
	if len(pages) > 10 {
		reversed := make([]staticIntf.Page, 10)
		for i := 0; i < 10; i++ {
			reversed[i] = pages[len(pages)-i-1]
		}
		return reversed
	}
	return pages
}
