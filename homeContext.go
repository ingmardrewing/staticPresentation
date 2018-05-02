package staticPresentation

import "github.com/ingmardrewing/staticIntf"

func NewHomeContextGroup(s staticIntf.Site) staticIntf.Context {

	cg := new(abstractContext)
	cg.site = s
	cg.renderer = NewHomePageRenderer(s)
	cg.renderer.Pages(s.Home()...)

	return cg
}
