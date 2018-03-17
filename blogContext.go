package staticPresentation

import (
	"github.com/ingmardrewing/fs"
	"github.com/ingmardrewing/staticIntf"
)

func NewBlogContextGroup(s staticIntf.Site) staticIntf.ContextGroup {

	cg := new(blogContext)
	cg.site = s

	cg.renderer = NewBlogRenderer(s)
	cg.renderer.SetPages(s.Posts())

	cg.naviRenderer = NewBlogNaviRenderer(s)
	cg.naviRenderer.SetPages(s.PostNaviPages())

	return cg
}

type blogContext struct {
	abstractContext
	naviRenderer staticIntf.Renderer
}

func (b *blogContext) GetComponents() []staticIntf.Component {
	components := b.renderer.GetComponents()
	return append(components, b.naviRenderer.GetComponents()...)
}

func (b *blogContext) RenderPages() []fs.FileContainer {
	fcs := b.renderer.Render()
	fcs = append(fcs, b.naviRenderer.Render()...)

	rr := NewRssRenderer(
		b.getLastTenReversedPages(),
		b.site.TargetDir(),
		b.site.RssPath(),
		b.site.RssFilename())
	rssFc := rr.Render()

	if rssFc != nil {
		fcs = append(fcs, rssFc)
	}

	return fcs
}
