package staticPresentation

import (
	"github.com/ingmardrewing/fs"
	"github.com/ingmardrewing/staticIntf"
)

func NewBlogContextGroup(s staticIntf.Site) staticIntf.ContextGroup {

	cg := new(blogContextGroup)
	cg.site = s

	cg.context = NewBlogContext(s)
	cg.context.SetElements(s.Posts())

	cg.naviContext = NewBlogNaviContext(s)
	cg.naviContext.SetElements(s.PostNaviPages())

	cg.Init()

	return cg
}

type blogContextGroup struct {
	navigationalContextGroup
}

func (b *blogContextGroup) RenderPages() []fs.FileContainer {

	fcs := b.context.RenderPages()
	fcs = append(fcs, b.naviContext.RenderPages()...)

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
