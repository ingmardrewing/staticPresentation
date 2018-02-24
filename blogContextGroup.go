package staticPresentation

import "github.com/ingmardrewing/staticIntf"

func NewBlogContextGroup(pagesContext, naviContext staticIntf.Context) staticIntf.ContextGroup {
	cg := new(blogContextGroup)
	cg.pagesContext = pagesContext
	cg.naviContext = naviContext
	cg.generateNaviPages()
	return cg
}

type blogContextGroup struct {
	abstractContextGroup
}
