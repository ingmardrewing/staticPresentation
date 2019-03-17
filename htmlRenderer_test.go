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
	dto := staticPersistence.NewFilledDto(
		"Archive",
		"An archive overview of pages",
		"",
		"narrative archive",
		"",
		"",
		"archive.html",
		[]string{},
		[]staticIntf.Image{})
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
	return staticModel.NewPage(dto, "testdomain", site)
}
