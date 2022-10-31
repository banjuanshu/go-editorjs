package html

import (
	"log"
	"os"
	"rbsite/internal/editorjs/parser/html/bootstrap"
	"rbsite/internal/editorjs/parser/html/bulma"
	"rbsite/internal/editorjs/parser/html/sample"
	"rbsite/internal/editorjs/support"
	"rbsite/internal/editorjs/support/config"
	"rbsite/internal/editorjs/support/domain"
	"reflect"
	"strings"
)

func Parser(jsonstr, styleName string) string {

	useDefault := true
	if styleName == "custom" {
		useDefault = false
		styleName = support.SM.StyleName
	}

	var f domain.EditorJSMethods
	switch styleName {
	case sample.StyleName:
		samplePkg := sample.Init(useDefault)
		f = &samplePkg
	case bootstrap.StyleName:
		bootstrapPkg := bootstrap.Init(useDefault)
		f = &bootstrapPkg
	case bulma.StyleName:
		bulmaPkg := bulma.Init(useDefault)
		f = &bulmaPkg
	}

	if reflect.DeepEqual(support.SM, domain.StyleMap{}) {
		log.Fatal("Style map is empty\n", nil)
	}

	if !support.IsValidStyle(support.SM.StyleName) {
		log.Fatal("Invalid style name: "+support.SM.StyleName+"\n", nil)
	}

	f.LoadLibrary()

	//input, err := support.ReadJsonFile(jsonFilePath)
	//if err != nil {
	//	log.Println("It was not possible to read the input json file\n", err)
	//}

	editorJSON := support.ParseEditorJSON(jsonstr)

	for _, el := range editorJSON.Blocks {

		styles, scripts := appendLibs(el)
		f.SetStyles(styles)
		f.SetScripts(scripts)
		//f.Separator()
		f.SetData(support.PrepareData(el))

		switch el.Type {

		case "header":
			f.Header()
		case "paragraph":
			f.Paragraph()
		case "quote":
			f.Quote()
		case "warning":
			f.Warning()
		case "delimiter":
			f.Delimiter()
		case "alert":
			f.Alert()
		case "list":
			f.List()
		case "checklist":
			f.Checklist()
		case "table":
			f.Table()
		case "AnyButton":
			f.AnyButton()
		case "code":
			f.Code()
		case "raw":
			f.Raw()
		case "image":
			f.Image()
		case "linkTool":
			f.LinkTool()
		case "attaches":
			f.Attaches()
		case "embed":
			f.Embed()
		case "imageGallery":
			f.ImageGallery()
		}

	}

	f.Separator()
	htmlStr := f.GetHtml()

	//err = support.WriteOutputFile(outputFilePath, f.CreatePage(), "html")
	//if err != nil {
	//	log.Println("It was not possible to write the output html file\n", err)
	//}

	return htmlStr
}

func appendLibs(block domain.EditorJSBlock) (styles []string, scripts []string) {
	libName := strings.ToLower(block.Type)
	libPath := config.LibsPath + libName + "/"
	if _, err := os.Stat(libPath); !os.IsNotExist(err) {
		styleMinified := string(support.MinifyLib(libName+"/"+libName+".css", "css"))
		if styleMinified != "" {
			styles = append(styles, `<style>`+styleMinified+`</style>`)
		}

		scriptMinified := string(support.MinifyLib(libName+"/"+libName+".js", "js"))
		if scriptMinified != "" {
			scripts = append(scripts, scriptMinified)
		}
	}

	return
}
