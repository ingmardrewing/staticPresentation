package staticPresentation

import (
	"github.com/ingmardrewing/fs"
	"github.com/ingmardrewing/staticIntf"
)

func NewBlogContextGroup(s staticIntf.Site) staticIntf.ContextGroup {

	cg := new(blogContextGroup)
	cg.site = s

	cg.pagesContext = NewBlogContext(s)
	cg.pagesContext.FsSetOff("/blog/")
	cg.pagesContext.SetElements(s.Posts())

	cg.naviContext = NewBlogNaviContext(s)
	cg.naviContext.FsSetOff("/blog/")

	cg.Init()

	return cg
}

type blogContextGroup struct {
	navigationalContextGroup
	site staticIntf.Site
}

func (b *blogContextGroup) RenderPages() []fs.FileContainer {

	fcs := b.pagesContext.RenderPages()
	fcs = append(fcs, b.naviContext.RenderPages()...)

	rss := b.rss(b.site.TargetDir())
	if rss != nil {
		fcs = append(fcs, rss)
	}

	return fcs
}
