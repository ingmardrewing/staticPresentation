package staticPresentation

import (
	"fmt"

	"github.com/ingmardrewing/fs"
	"github.com/ingmardrewing/staticIntf"
)

func NewNarrativeContextGroup(
	pages []staticIntf.Page,
	cd staticIntf.CommonData) staticIntf.ContextGroup {

	narrativeCtx := NewNarrativeContext(cd)
	narrativeCtx.SetElements(pages)

	// TODO: Genrate archive pages, separate rss, etc.

	cg := new(narrativeContextGroup)
	cg.pagesContext = narrativeCtx
	return cg
}

type narrativeContextGroup struct {
	navigationalContextGroup
}

func (a *narrativeContextGroup) RenderPages(dir string) []fs.FileContainer {
	fcs := a.pagesContext.RenderPages(dir)
	if len(fcs) > 1 {
		// copy the content of the last page
		// of the narrative and add a page with
		// this content as index page
		lastPageFc := fcs[len(fcs)-2]
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
	fmt.Println("group size fcs: ", len(fcs))
	return fcs
}
