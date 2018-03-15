package staticPresentation

import (
	"path"

	"github.com/ingmardrewing/fs"
	"github.com/ingmardrewing/staticIntf"
)

type abstractContextGroup struct {
	pages   []staticIntf.Page
	context staticIntf.SubContext
	site    staticIntf.Site
}

func (a *abstractContextGroup) GetComponents() []staticIntf.Component {
	return a.context.GetComponents()
}

func (a *abstractContextGroup) RenderPages() []fs.FileContainer {
	return a.context.RenderPages()
}

func (a *abstractContextGroup) Domain() string { return a.site.Domain() }

func (a *abstractContextGroup) naviPageDescription() string { return "" }

func (a *abstractContextGroup) naviPageTitle() string { return "" }

func (a *abstractContextGroup) naviPagePathFromDocRoot() string { return "" }

func (a *abstractContextGroup) Init() {}

func (a *abstractContextGroup) rss(targetDir string) fs.FileContainer {
	if len(a.context.GetPages()) > 0 {
		r := NewRssRenderer(a.getLastPages(10))
		rssPath := path.Join(targetDir, a.context.SiteDto().Rss())
		rssFilename := "rss.xml"
		return r.Render(rssPath, rssFilename)
	}
	return nil
}

func (a *abstractContextGroup) getLastPages(nr int) []staticIntf.Page {
	pgs := a.context.GetPages()
	if len(pgs) > nr {
		return pgs[len(pgs)-nr:]
	}
	return pgs
}
