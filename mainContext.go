package staticPresentation

import "github.com/ingmardrewing/staticIntf"

func NewEntryContextGroup(s staticIntf.Site) staticIntf.Context {
	cg := new(abstractContext)
	cg.site = s
	cg.renderer = NewEntryPageRenderer(s)
	cg.renderer.Pages(s.Home()...)

	return cg
}
