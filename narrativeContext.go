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
	cg.renderer.Pages(s.Narratives()...)

	cg.narrativeArchiveContext = NewNarrativeArchiveRename(s)

	cg.narrativeMarginalContext = NewNarrativeMarginalRenderer(s)
	cg.narrativeMarginalContext.Pages(s.NarrativeMarginals()...)

	cg.GenerateArchivePage()

	return cg
}

type narrativeContext struct {
	abstractContext
	narrativeArchiveContext  staticIntf.Renderer
	narrativeMarginalContext staticIntf.Renderer
}

func (a *narrativeContext) GetComponents() []staticIntf.Component {
	cmps := a.renderer.Components()
	cmps = append(cmps, a.narrativeArchiveContext.Components()...)
	return append(cmps, a.narrativeMarginalContext.Components()...)
}

func (a *narrativeContext) GenerateArchivePage() {
	dto := staticPersistence.NewFilledDto(0,
		"Archive", "Archive", "",
		"", "", "",
		"", "", "", "",
		"", "", "archive.html", "")
	np := staticModel.NewPage(dto, a.site.Domain())

	np.NavigatedPages(a.renderer.Pages()...)

	a.narrativeArchiveContext.Pages(np)
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
		a.site.Narratives(),
		a.site.TargetDir(),
		a.site.RssPath(),
		a.site.RssFilename())
	rssFc := rr.Render()

	if rssFc != nil {
		fcs = append(fcs, rssFc)
	}

	return fcs
}
