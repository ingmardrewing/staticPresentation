package staticPresentation

import (
	"github.com/ingmardrewing/fs"
	"github.com/ingmardrewing/staticIntf"
	"github.com/ingmardrewing/staticModel"
	"github.com/ingmardrewing/staticPersistence"
)

func NewNarrativeContextGroup(s staticIntf.Site) staticIntf.Context {

	cg := new(narrativeContext)
	cg.site = s

	cg.renderer = NewNarrativeRenderer(s)
	cg.renderer.Pages(s.GetPagesByVariant(staticIntf.NARRATIVES)...)

	cg.narrativeArchiveRenderer = NewNarrativeArchiveRenderer(s)

	cg.narrativeMarginalRenderer = NewNarrativeMarginalRenderer(s)
	cg.narrativeMarginalRenderer.Pages(s.GetPagesByVariant(staticIntf.NARRATIVEMARGINALS)...)

	cg.GenerateArchivePage()

	return cg
}

type narrativeContext struct {
	abstractContext
	narrativeArchiveRenderer  staticIntf.Renderer
	narrativeMarginalRenderer staticIntf.Renderer
}

func (a *narrativeContext) GetComponents() []staticIntf.Component {
	cmps := a.renderer.Components()
	cmps = append(cmps, a.narrativeArchiveRenderer.Components()...)
	cmps = append(cmps, a.narrativeMarginalRenderer.Components()...)
	return cmps
}

func (a *narrativeContext) GenerateArchivePage() {
	dto := staticPersistence.NewFilledDto(0,
		"Archive", "Archive", "",
		"", "", "",
		"", "", "", "",
		"", "", "archive.html", "", "narrative archive", "")
	np := staticModel.NewPage(dto, a.site.Domain(), a.site)

	np.NavigatedPages(a.renderer.Pages()...)

	a.narrativeArchiveRenderer.Pages(np)
	a.site.AddMarginal(np)

	for _, n := range a.site.GetPagesByVariant(staticIntf.NARRATIVEMARGINALS) {
		a.site.AddMarginal(n)
	}
}

func (a *narrativeContext) RenderPages() []fs.FileContainer {
	fcs := a.narrativeArchiveRenderer.Render()
	fcs = append(fcs, a.narrativeMarginalRenderer.Render()...)
	fcs = append(fcs, a.renderer.Render()...)

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
		a.site.GetPagesByVariant(staticIntf.NARRATIVES),
		a.site.TargetDir(),
		a.site.RssPath(),
		a.site.RssFilename())
	rssFc := rr.Render()

	if rssFc != nil {
		fcs = append(fcs, rssFc)
	}

	return fcs
}
