package staticPresentation

import "github.com/ingmardrewing/staticIntf"

func NewSiteContextGroup(s staticIntf.Site) staticIntf.Context {

	cg := new(abstractContext)
	cg.site = s
	cg.renderer = NewPagesRenderer(s)
	cg.renderer.Pages(s.Pages()...)

	return cg
}
