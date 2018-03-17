package staticPresentation

import "github.com/ingmardrewing/staticIntf"

func NewSiteContextGroup(s staticIntf.Site) staticIntf.Context {

	cg := new(abstractContext)
	cg.site = s
	cg.renderer = NewPagesContext(s)
	cg.renderer.SetPages(s.Pages())

	return cg
}
