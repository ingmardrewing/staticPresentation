package staticPresentation

import (
	"github.com/ingmardrewing/fs"
	"github.com/ingmardrewing/staticIntf"
	"github.com/ingmardrewing/staticModel"
)

func NewNarrativeContextGroup(s staticIntf.Site) staticIntf.Context {

	cg := new(narrativeContext)
	cg.site = s

	cg.renderer = NewNarrativeContext(s)
	cg.renderer.SetPages(s.Narratives())

	cg.narrativeArchiveContext = NewNarrativeArchiveContext(s)

	cg.narrativeMarginalContext = NewNarrativeMarginalContext(s)
	cg.narrativeMarginalContext.SetPages(s.NarrativeMarginals())

	cg.GenerateArchivePage()

	return cg
}

type narrativeContext struct {
	abstractContext
	narrativeArchiveContext  staticIntf.Renderer
	narrativeMarginalContext staticIntf.Renderer
}

func (a *narrativeContext) GetComponents() []staticIntf.Component {
	cmps := a.renderer.GetComponents()
	cmps = append(cmps, a.narrativeArchiveContext.GetComponents()...)
	return append(cmps, a.narrativeMarginalContext.GetComponents()...)
}

func (a *narrativeContext) GenerateArchivePage() {
	np := staticModel.NewEmptyNaviPage(a.site.Domain())
	np.NavigatedPages(a.renderer.GetPages()...)
	np.Title("Archive")
	np.HtmlFilename("archive.html")
	np.PathFromDocRoot("")
	a.narrativeArchiveContext.SetPages([]staticIntf.Page{np})
	a.site.AddMarginal(np)

	for _, n := range a.site.NarrativeMarginals() {
		a.site.AddMarginal(n)
	}
}

func (a *narrativeContext) RenderPages() []fs.FileContainer {
	fcs := a.renderer.Render()
	fcs = append(fcs, a.narrativeArchiveContext.Render()...)
	fcs = append(fcs, a.narrativeMarginalContext.Render()...)

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

	rr := NewRssRenderer(
		a.site.Posts(),
		a.site.TargetDir(),
		a.site.RssPath(),
		a.site.RssFilename())
	rssFc := rr.Render()

	if rssFc != nil {
		fcs = append(fcs, rssFc)
	}

	return fcs
}
