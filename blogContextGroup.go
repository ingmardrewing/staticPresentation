package staticPresentation

import (
	"fmt"

	"github.com/ingmardrewing/fs"
	"github.com/ingmardrewing/staticIntf"
)

func NewBlogContextGroup(
	posts []staticIntf.Page,
	cd staticIntf.CommonData) staticIntf.ContextGroup {

	cg := new(blogContextGroup)

	cg.pagesContext = NewBlogContext(cd)
	cg.pagesContext.FsSetOff("/blog/")
	cg.pagesContext.SetElements(posts)

	cg.naviContext = NewBlogNaviContext(cd)
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
	fmt.Println("group size fcs: ", len(fcs))

	return fcs
}
