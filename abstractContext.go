package staticPresentation

import (
	"github.com/ingmardrewing/fs"
	"github.com/ingmardrewing/staticIntf"
)

type abstractContext struct {
	pages    []staticIntf.Page
	renderer staticIntf.Renderer
	site     staticIntf.Site
}

func (a *abstractContext) GetComponents() []staticIntf.Component {
	return a.renderer.GetComponents()
}

func (a *abstractContext) RenderPages() []fs.FileContainer {
	return a.renderer.Render()
}

func (a *abstractContext) Domain() string { return a.site.Domain() }

func (b *abstractContext) getLastTenReversedPages() []staticIntf.Page {
	pages := b.renderer.GetPages()
	if len(pages) > 10 {
		reversed := make([]staticIntf.Page, 10)
		for i := 0; i < 10; i++ {
			reversed[i] = pages[len(pages)-i-1]
		}
		return reversed
	}
	return pages
}
