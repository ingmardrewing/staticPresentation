package staticPresentation

import "github.com/ingmardrewing/staticIntf"

func NewNarrativeContextGroup(
	pages []staticIntf.Page,
	dto staticIntf.ContextDto,
	mainNavi []staticIntf.Location,
	footerNavi []staticIntf.Location) staticIntf.ContextGroup {

	narrativeCtx := NewNarrativeContext(mainNavi, footerNavi)
	narrativeCtx.SetContextDto(dto)
	narrativeCtx.SetElements(pages)
	narrativeCtx.AddRss()

	// TODO: Genrate archive pages, etc.

	cg := new(narrativeContextGroup)
	cg.pagesContext = narrativeCtx
	return cg
}

type narrativeContextGroup struct {
	abstractContextGroup
}
