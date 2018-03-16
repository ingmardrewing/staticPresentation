package staticPresentation

import "github.com/ingmardrewing/staticIntf"

func NewSiteContextGroup(s staticIntf.Site) staticIntf.ContextGroup {

	cg := new(siteContextGroup)
	cg.site = s
	cg.context = NewPagesContext(s)
	cg.context.SetElements(s.Pages())

	return cg
}

type siteContextGroup struct {
	abstractContextGroup
	marginalContext staticIntf.SubContext
}
