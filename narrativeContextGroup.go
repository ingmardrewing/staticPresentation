package staticPresentation

import (
	"github.com/ingmardrewing/fs"
	"github.com/ingmardrewing/staticIntf"
)

func NewNarrativeContextGroup(
	pages []staticIntf.Page,
	dto staticIntf.ContextDto,
	mainNavi []staticIntf.Location,
	footerNavi []staticIntf.Location) staticIntf.ContextGroup {

	narrativeCtx := NewNarrativeContext(mainNavi, footerNavi)
	narrativeCtx.SetContextDto(dto)
	narrativeCtx.SetElements(pages)
	narrativeCtx.AddRss()

	// TODO: Genrate archive pages, separate rss, etc.

	cg := new(narrativeContextGroup)
	cg.pagesContext = narrativeCtx
	cg.dto = dto
	return cg
}

type narrativeContextGroup struct {
	abstractContextGroup
	dto staticIntf.ContextDto
}

func (a *narrativeContextGroup) RenderPages(dir string) []fs.FileContainer {
	fcs := a.pagesContext.RenderPages(dir)
	if len(fcs) > 1 {
		lastPageFc := fcs[len(fcs)-2]
		index := fs.NewFileContainer()
		index.SetData(lastPageFc.GetData())
		index.SetPath(a.dto.TargetDir())
		index.SetFilename("index.html")
		return append(fcs, index)
	}
	return fcs
}
