package staticPresentation

import (
	"fmt"

	"github.com/ingmardrewing/fs"
	"github.com/ingmardrewing/staticIntf"
	"github.com/ingmardrewing/staticModel"
	"github.com/ingmardrewing/staticPersistence"
	"github.com/ingmardrewing/staticUtil"
	log "github.com/sirupsen/logrus"
)

func NewNarrativeContextGroup(s staticIntf.Site) staticIntf.Context {

	tool := staticUtil.NewPagesContainerCollectionTool(s)

	cg := new(narrativeContext)
	cg.site = s

	cg.renderer = NewNarrativeRenderer(s)
	pages := tool.GetPagesByVariant(staticIntf.NARRATIVES)
	msg := fmt.Sprintf("narrativeContext, nr of pages: %d", len(pages))
	log.Debug(msg)
	cg.renderer.Pages(pages...)

	cg.narrativeArchiveRenderer = NewNarrativeArchiveRenderer(s)

	cg.narrativeMarginalRenderer = NewNarrativeMarginalRenderer(s)
	cg.narrativeMarginalRenderer.Pages(tool.GetPagesByVariant(staticIntf.NARRATIVEMARGINALS)...)

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
	// TODO Reimplement archive generation
	dto := staticPersistence.NewFilledDto(
		"Archive",
		"An archive overview of pages within "+a.site.Domain(),
		"",
		"narrative archive",
		"",
		"",
		"archive.html",
		[]string{},
		[]staticIntf.Image{})
	np := staticModel.NewPage(dto, a.site.Domain(), a.site)
	np.NavigatedPages(a.renderer.Pages()...)

	a.narrativeArchiveRenderer.Pages(np)
	a.site.AddMarginal(np)

	tool := staticUtil.NewPagesContainerCollectionTool(a.site)
	for _, n := range tool.GetPagesByVariant(
		staticIntf.NARRATIVEMARGINALS) {
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

	tool := staticUtil.NewPagesContainerCollectionTool(a.site)
	rr := NewRssRenderer(
		tool.GetPagesByVariant(staticIntf.NARRATIVES),
		a.site.TargetDir(),
		a.site.RssPath(),
		a.site.RssFilename())
	rssFc := rr.Render()

	if rssFc != nil {
		fcs = append(fcs, rssFc)
	}

	nr := fmt.Sprintf("narrativeContext, files: %d", len(fcs))
	log.Debug(nr)

	return fcs
}
