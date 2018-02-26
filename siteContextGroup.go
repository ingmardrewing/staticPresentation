package staticPresentation

import (
	"github.com/ingmardrewing/fs"
	"github.com/ingmardrewing/staticIntf"
)

func NewSiteContextGroup(
	pages []staticIntf.Page,
	marginalPages []staticIntf.Page,
	dto staticIntf.ContextDto,
	mainNavi []staticIntf.Location,
	footerNavi []staticIntf.Location) staticIntf.ContextGroup {

	cg := new(siteContextGroup)
	cg.pagesContext = NewPagesContext(mainNavi, footerNavi)
	cg.pagesContext.SetElements(pages)
	cg.pagesContext.SetContextDto(dto)

	cg.marginalContext = NewMarginalContext(mainNavi, footerNavi)
	cg.marginalContext.SetElements(marginalPages)
	cg.marginalContext.SetContextDto(dto)

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
