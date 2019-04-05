package staticPresentation

import (
	"fmt"
	"path"

	"github.com/ingmardrewing/fs"
	"github.com/ingmardrewing/staticIntf"
	"github.com/ingmardrewing/staticModel"
	"github.com/ingmardrewing/staticPersistence"
	"github.com/ingmardrewing/staticUtil"
)

func NewBlogContext(s staticIntf.Site) staticIntf.Context {
	tool := staticUtil.NewPagesContainerCollectionTool(s)

	cg := new(blogContext)
	cg.site = s

	cg.renderer = NewBlogRenderer(s)
	cg.renderer.Pages(tool.GetPagesByVariant(staticIntf.BLOG)...)

	cg.naviRenderer = NewBlogNaviRenderer(s)
	cg.naviRenderer.Pages(tool.GetNaviPagesByVariant(staticIntf.BLOG)...)

	cg.archiveRenderer = NewBlogArchiveRenderer(s)
	cg.addArchiveLocation()
	return cg
}

type blogContext struct {
	abstractContext
	naviRenderer    staticIntf.Renderer
	archiveRenderer staticIntf.Renderer
}

func (b *blogContext) addArchiveLocation() {
	file := "archive.html"
	path := "/blog/"
	fullPath := path + file
	archiveLoc := staticModel.NewLocation(
		"",
		b.site.Domain(),
		"Archive",
		"",
		path,
		file,
		"",
		fullPath,
		fmt.Sprintf("https://%s%s", b.site.Domain(), fullPath))
	b.site.AddMain(archiveLoc)
}

func (b *blogContext) GetComponents() []staticIntf.Component {
	components := b.renderer.Components()
	components = append(components, b.archiveRenderer.Components()...)
	return append(components, b.naviRenderer.Components()...)
}

func (b *blogContext) RenderPages() []fs.FileContainer {
	fcs := b.renderer.Render()
	fcs = append(fcs, b.naviRenderer.Render()...)

	rssFc := b.RenderRssPages()
	if rssFc != nil {
		fcs = append(fcs, rssFc)
	}
	fcs = append(fcs, b.RenderBlogArchive()...)

	return fcs
}

func (b *blogContext) RenderRssPages() fs.FileContainer {
	rr := NewRssRenderer(
		b.getLastTenReversedPages(),
		path.Join(b.site.TargetDir(), "/blog/"),
		"/blog/",
		b.site.RssFilename())
	return rr.Render()
}

func (b *blogContext) RenderBlogArchive() []fs.FileContainer {
	dto := staticPersistence.NewFilledDto(
		b.site.Domain()+" Archive",
		"",
		"",
		"blog archive",
		"",
		"/blog/",
		"archive.html",
		[]string{},
		[]staticIntf.Image{})
	p := staticModel.NewPage(dto, b.site.Domain(), b.site)
	b.archiveRenderer.Pages(p)
	return b.archiveRenderer.Render()
}
