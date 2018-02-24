package staticPresentation

import "github.com/ingmardrewing/staticIntf"

func NewBlogContextGroup(
	posts []staticIntf.Page,
	dto staticIntf.ContextDto,
	mainNavi []staticIntf.Location,
	footerNavi []staticIntf.Location) staticIntf.ContextGroup {

	bc := NewBlogContext(mainNavi, footerNavi)
	bc.SetContextDto(dto)
	bc.SetElements(posts)

	bnc := NewBlogNaviContext(mainNavi, footerNavi)
	bnc.SetContextDto(dto)
	bnc.FsSetOff("/blog/")

	cg := new(blogContextGroup)
	cg.pagesContext = bc
	cg.naviContext = bnc
	cg.generateNaviPages()
	return cg
}

type blogContextGroup struct {
	abstractContextGroup
}
