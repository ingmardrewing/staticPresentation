package staticPresentation

import (
	"github.com/ingmardrewing/fs"
	"github.com/ingmardrewing/staticIntf"
	"github.com/ingmardrewing/staticModel"
)

func NewNarrativeContextGroup(s staticIntf.Site) staticIntf.ContextGroup {

	cg := new(narrativeContextGroup)
	cg.site = s

	cg.context = NewNarrativeContext(s)
	cg.context.SetElements(s.Narratives())

	cg.narrativeArchiveContext = NewNarrativeArchiveContext(s)

	cg.narrativeMarginalContext = NewNarrativeMarginalContext(s)
	cg.narrativeMarginalContext.SetElements(s.NarrativeMarginals())

	cg.Init()

	return cg
}

type narrativeContextGroup struct {
	abstractContextGroup
	narrativeArchiveContext  staticIntf.Context
	narrativeMarginalContext staticIntf.Context
}

func (a *narrativeContextGroup) GetComponents() []staticIntf.Component {
	cmps := a.context.GetComponents()
	cmps = append(cmps, a.narrativeArchiveContext.GetComponents()...)
	return append(cmps, a.narrativeMarginalContext.GetComponents()...)
}

func (a *narrativeContextGroup) Init() {
	np := staticModel.NewEmptyNaviPage(a.site.Domain())
	np.NavigatedPages(a.context.GetElements()...)
	np.Title("Archive")
	np.HtmlFilename("archive.html")
	np.PathFromDocRoot("")
	a.narrativeArchiveContext.SetElements([]staticIntf.Page{np})
	a.site.AddMarginal(np)

	for _, n := range a.site.NarrativeMarginals() {
		a.site.AddMarginal(n)
	}
}

func (a *narrativeContextGroup) RenderPages() []fs.FileContainer {
	fcs := a.context.RenderPages()

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
	fcs = append(fcs, a.narrativeMarginalContext.RenderPages()...)

	return fcs
}
