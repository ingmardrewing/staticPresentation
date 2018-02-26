package staticPresentation

import "github.com/ingmardrewing/staticIntf"

func NewBlogContextGroup(
	posts []staticIntf.Page,
	dto staticIntf.ContextDto,
	mainNavi []staticIntf.Location,
	footerNavi []staticIntf.Location) staticIntf.ContextGroup {

	cg := new(blogContextGroup)

	cg.pagesContext = NewBlogContext(mainNavi, footerNavi)
	cg.pagesContext.SetContextDto(dto)
	cg.pagesContext.SetElements(posts)
	cg.pagesContext.FsSetOff("/blog/")
	cg.pagesContext.AddRss()

	cg.naviContext = NewBlogNaviContext(mainNavi, footerNavi)
	cg.naviContext.SetContextDto(dto)
	cg.naviContext.FsSetOff("/blog/")

	cg.generateNaviPages()

	return cg
}

type blogContextGroup struct {
	navigationalContextGroup
}
