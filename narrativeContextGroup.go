package staticPresentation

import (
	"github.com/ingmardrewing/fs"
	"github.com/ingmardrewing/staticIntf"
)

func NewNarrativeContextGroup(
	pages []staticIntf.Page,
	cd staticIntf.Site) staticIntf.ContextGroup {

	cg := new(narrativeContextGroup)

	cg.pagesContext = NewNarrativeContext(cd)
	cg.pagesContext.FsSetOff("")
	cg.pagesContext.SetElements(pages)

	return cg
}

type narrativeContextGroup struct {
	abstractContextGroup
}

func (a *narrativeContextGroup) RenderPages(dir string) []fs.FileContainer {
	fcs := a.pagesContext.RenderPages(dir)

	if len(fcs) > 1 {
		// copy the content of the last page
		// of the narrative and add a page with
		// this content as index page
		inx := len(fcs) - 2
		lastPageFc := fcs[inx]
		index := fs.NewFileContainer()
		index.SetData(lastPageFc.GetData())
		index.SetPath(a.pagesContext.CommonData().ContextDto().TargetDir())
		index.SetFilename("index.html")
		fcs = append(fcs, index)
	}
	rss := a.rss(dir)
	if rss != nil {
		fcs = append(fcs, rss)
	}

	return fcs
}
