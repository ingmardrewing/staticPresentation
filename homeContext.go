package staticPresentation

import "github.com/ingmardrewing/staticIntf"
import log "github.com/sirupsen/logrus"

func NewHomeContextGroup(s staticIntf.Site) staticIntf.Context {

	log.Info("XXX - XXX")

	cg := new(abstractContext)
	cg.site = s
	cg.renderer = NewHomePageRenderer(s)
	cg.renderer.Pages(s.Home()...)

	return cg
}
