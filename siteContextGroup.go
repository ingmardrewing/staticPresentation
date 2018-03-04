package staticPresentation

import (
	"github.com/ingmardrewing/fs"
	"github.com/ingmardrewing/staticIntf"
)

func NewSiteContextGroup(
	pages []staticIntf.Page,
	marginalPages []staticIntf.Page,
	cd staticIntf.CommonData) staticIntf.ContextGroup {

	cg := new(siteContextGroup)

	cg.pagesContext = NewPagesContext(cd)
	cg.pagesContext.SetElements(pages)

	cg.marginalContext = NewMarginalContext(cd)
	cg.marginalContext.SetElements(marginalPages)

	return cg
}

type siteContextGroup struct {
	abstractContextGroup
	marginalContext staticIntf.Context
}

func (s *siteContextGroup) GetComponents() []staticIntf.Component {
	components := s.pagesContext.GetComponents()
	return append(components, s.marginalContext.GetComponents()...)
}

func (s *siteContextGroup) RenderPages(dir string) []fs.FileContainer {
	fcs := s.pagesContext.RenderPages(dir)
	return append(fcs, s.marginalContext.RenderPages(dir)...)
}
