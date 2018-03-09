package staticPresentation

import (
	"fmt"
	"path"

	"github.com/ingmardrewing/fs"
	"github.com/ingmardrewing/staticIntf"
)

func NewBlogContextGroup(s staticIntf.Site) staticIntf.ContextGroup {

	posts := s.Posts()
	for _, p := range posts {
		newPath := path.Join("/blog/", p.PathFromDocRoot())
		p.PathFromDocRoot(newPath)
	}

	cg := new(blogContextGroup)
	cg.site = s
	cg.context = NewBlogContext(s)
	cg.context.SetElements(posts)

	cg.naviContext = NewBlogNaviContext(s)
	cg.naviContext.FsSetOff("/blog/")

	fmt.Println("prmf", s.Domain())
	cg.Init()

	return cg
}

type blogContextGroup struct {
	navigationalContextGroup
}

func (b *blogContextGroup) RenderPages() []fs.FileContainer {

	fcs := b.context.RenderPages()
	fcs = append(fcs, b.naviContext.RenderPages()...)

	rss := b.rss(b.site.TargetDir())
	if rss != nil {
		fcs = append(fcs, rss)
	}

	return fcs
}
