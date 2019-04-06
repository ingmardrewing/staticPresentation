package staticPresentation

import (
	"testing"

	"github.com/ingmardrewing/staticIntf"
	"github.com/ingmardrewing/staticModel"
	"github.com/ingmardrewing/staticPersistence"
)

func TestTemplatedRenderer(t *testing.T) {
	templatedRenderer := NewTemplatedRenderer("test")
	templatedRenderer.Pages(GetPage())

	expected := `<!doctype html><html itemscope lang="en"><head></head><body></body></html>`
	actual := templatedRenderer.Render()[0].GetDataAsString()

	if actual != expected {
		t.Error("Expeted", expected, "but got", actual)
	}
}

func GetPage() staticIntf.Page {
	site := staticModel.NewSiteDto(
		"twitterHandle",
		"topic",
		"tags",
		"domain",
		"basePath",
		"cardType",
		"section",
		"fbPage",
		"twitterPage",
		"rssPath",
		"rssFilename",
		"css",
		"disqusId",
		"targetDir",
		"description",
		"keyWords",
		"subject",
		"author",
		"homeText",
		"homeHeadline",
		"svgLogo")
	pm .= staticModel.NewPageMaker()
	pm.Title("Archive")
	pm.Description("An archive overview of pages")
	pm.Category("narrative archive")
	pm.FileName("archive.html")
	pm.Site(site)

	return pm.Make()
}
