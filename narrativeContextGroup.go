package staticPresentation

import (
	"github.com/ingmardrewing/fs"
	"github.com/ingmardrewing/staticIntf"
	"github.com/ingmardrewing/staticModel"
)

func NewNarrativeContextGroup(s staticIntf.Site) staticIntf.ContextGroup {

	cg := new(narrativeContextGroup)
	cg.site = s

	cg.pagesContext = NewNarrativeContext(s)
	cg.pagesContext.FsSetOff("")
	cg.pagesContext.SetElements(s.Narratives())

	cg.narrativeArchiveContext = NewNarrativeArchiveContext(s)
	cg.narrativeArchiveContext.FsSetOff("")

	cg.Init()

	return cg
}

type narrativeContextGroup struct {
	abstractContextGroup
	narrativeArchiveContext staticIntf.Context
}

func (a *narrativeContextGroup) Init() {
	np := staticModel.NewEmptyNaviPage()
	np.NavigatedPages(a.pagesContext.GetElements()...)
	np.Title("Archive")
	np.Domain(a.site.Domain())
	filename := "archive.html"
	np.Url(filename)
	np.HtmlFilename(filename)
	np.PathFromDocRoot("")
	a.narrativeArchiveContext.SetElements([]staticIntf.Page{np})
	a.site.AddMarginal(np)
}

func (a *narrativeContextGroup) RenderPages() []fs.FileContainer {
	fcs := a.pagesContext.RenderPages()

	fcs = append(fcs, a.narrativeArchiveContext.RenderPages()...)

	if len(fcs) > 1 {
		// copy the content of the last page
		// of the narrative and add a page with
		// this content as index page
		inx := len(fcs) - 2
		lastPageFc := fcs[inx]
		index := fs.NewFileContainer()
		index.SetData(lastPageFc.GetData())
		index.SetPath(a.site.TargetDir())
		index.SetFilename("index.html")
		fcs = append(fcs, index)
	}
	rss := a.rss(a.site.TargetDir())
	if rss != nil {
		fcs = append(fcs, rss)
	}

	return fcs
}
