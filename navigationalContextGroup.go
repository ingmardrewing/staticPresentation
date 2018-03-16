package staticPresentation

import (
	"github.com/ingmardrewing/fs"
	"github.com/ingmardrewing/staticIntf"
)

type navigationalContextGroup struct {
	abstractContextGroup
	naviContext staticIntf.SubContext
}

func (n *navigationalContextGroup) GetComponents() []staticIntf.Component {
	components := n.context.GetComponents()
	return append(components, n.naviContext.GetComponents()...)
}

func (n *navigationalContextGroup) RenderPages() []fs.FileContainer {
	pages := n.context.RenderPages()
	return append(pages, n.naviContext.RenderPages()...)
}
