package staticPresentation

import (
	"github.com/ingmardrewing/fs"
	"github.com/ingmardrewing/staticIntf"
)

func NewBlogContextGroup(s staticIntf.Site) staticIntf.ContextGroup {

	cg := new(blogContextGroup)

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
}

func (b *blogContextGroup) RenderPages(targetDir string) []fs.FileContainer {
	fcs := b.pagesContext.RenderPages(targetDir)

	fcs = append(fcs, b.naviContext.RenderPages(targetDir)...)
	rss := b.rss(targetDir)
	if rss != nil {
		fcs = append(fcs, rss)
	}

	return fcs
}
