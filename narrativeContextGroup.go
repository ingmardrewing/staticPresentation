package staticPresentation

import (
	"github.com/ingmardrewing/fs"
	"github.com/ingmardrewing/staticIntf"
	"github.com/ingmardrewing/staticModel"
)

func NewNarrativeContextGroup(
	pages []staticIntf.Page,
	site staticIntf.Site) staticIntf.ContextGroup {

	cg := new(narrativeContextGroup)
	cg.site = site

	cg.pagesContext = NewNarrativeContext(site)
	cg.pagesContext.FsSetOff("")
	cg.pagesContext.SetElements(pages)

	cg.narrativeArchiveContext = NewNarrativeArchiveContext(site)
	cg.narrativeArchiveContext.FsSetOff("")

	cg.Init()

	return cg
}

type narrativeContextGroup struct {
	abstractContextGroup
	narrativeArchiveContext staticIntf.Context
	site                    staticIntf.Site
}

func (a *narrativeContextGroup) Init() {
	np := staticModel.NewNaviPage()
	np.NavigatedPages(a.pagesContext.GetElements()...)
	np.Title("Archive")
	np.Domain(a.site.ContextDto().Domain())
	filename := "archive.html"
	np.Url(filename)
	np.HtmlFilename(filename)
	np.PathFromDocRoot("")
	a.narrativeArchiveContext.SetElements([]staticIntf.Page{np})
	a.site.AddMarginal(np)
}

func (a *narrativeContextGroup) RenderPages(dir string) []fs.FileContainer {
	fcs := a.pagesContext.RenderPages(dir)

	fcs = append(fcs, a.narrativeArchiveContext.RenderPages(dir)...)

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
