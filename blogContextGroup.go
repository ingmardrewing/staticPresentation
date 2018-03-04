package staticPresentation

import "github.com/ingmardrewing/staticIntf"

func NewBlogContextGroup(
	posts []staticIntf.Page,
	cd staticIntf.CommonData) staticIntf.ContextGroup {

	cg := new(blogContextGroup)

	cg.pagesContext = NewBlogContext(cd)
	cg.pagesContext.SetContextDto(cd.ContextDto())
	cg.pagesContext.FsSetOff("/blog/")
	cg.pagesContext.SetElements(posts)
	cg.pagesContext.AddRss()

	cg.naviContext = NewBlogNaviContext(cd)
	cg.naviContext.SetContextDto(cd.ContextDto())
	cg.naviContext.FsSetOff("/blog/")

	cg.Init()

	return cg
}

type blogContextGroup struct {
	navigationalContextGroup
}
