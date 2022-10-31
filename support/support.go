package support

import (
	"embed"
	"encoding/json"
	"fmt"
	"github.com/banjuanshu/go-editorjs/support/config"
	"github.com/banjuanshu/go-editorjs/support/domain"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/js"
	"io/ioutil"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var SM domain.StyleMap

//go:embed libs/*
var libsFiles embed.FS

//go:embed assets/*
var assetsFiles embed.FS

func ReadJsonFile(jsonFilePath string) (jsonData string, err error) {
	data, err := ioutil.ReadFile(jsonFilePath)
	if err != nil {
		log.Println("Error reading the input json file\n", err)
	}

	jsonData = string(data)

	return
}

func MinifyLib(libPath, libType string) (contentMinified []byte) {
	contentFile, err := libsFiles.ReadFile("libs/" + libPath)
	if err == nil {
		contentMinified, err = MinifyContent(contentFile, libType)
		if err != nil {
			log.Println("Error minifying lib file\n", err)
		}
	}

	return
}

func LoadAsset(assetFile, assetType string) (contentFile []byte) {
	contentFile, err := assetsFiles.ReadFile("assets/" + assetType + "/" + assetFile)
	if err != nil {
		log.Println("Error loading asset file\n", err)
	}

	return
}

func MinifyAsset(assetPath, assetType string) (contentMinified []byte) {
	contentFile, err := assetsFiles.ReadFile(assetPath)
	if err == nil {
		contentMinified, err = MinifyContent(contentFile, assetType)
		if err != nil {
			log.Println("Error minifying asset file\n", err)
		}
	}

	return
}

func MinifyExternalStyle(libPath string) (contentMinified []byte, err error) {
	contentFile, err := ioutil.ReadFile(libPath)
	if err == nil {
		contentMinified, err = MinifyContent(contentFile, "css")
		if err != nil {
			log.Println("Error minifying external lib file\n", err)
		}
	}

	return
}

func MinifyContent(content []byte, format string) (contentMinified []byte, err error) {

	m := minify.New()

	switch format {

	case "css":
		m.AddFunc("text/css", css.Minify)
		contentMinified, err = m.Bytes("text/css", content)
		if err != nil {
			log.Println("Error minifying CSS file\n", err)
		}

	case "js":
		m.AddFuncRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), js.Minify)
		contentMinified, err = m.Bytes("application/javascript", content)
		if err != nil {
			log.Println("Error minifying JS file\n", err)
		}
	}

	return
}

func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func CreateHTMLNestedList(items []domain.NestedListItem, listStyle string, first bool) string {
	var result []string

	if first {
		result = append(result, `<`+listStyle+` class="`+SM.Blocks.List.Group+`">`)
	} else {
		result = append(result, `<`+listStyle+` class="`+SM.Blocks.List.NestedGroup+`">`)
	}

	for _, item := range items {
		result = append(result, `<li class="`+SM.Blocks.List.Item+`">`+item.Content+`</li>`)

		if len(item.Items) > 0 {
			result = append(result, CreateHTMLNestedList(item.Items, listStyle, false))
		}

		result = append(result, `</li>`)
	}

	result = append(result, `</`+listStyle+`>`)

	return strings.Join(result[:], "\n")
}

func CreateMarkDownNestedList(items []domain.NestedListItem, listStyle, spaceLeft string) string {
	var result []string

	for i, item := range items {
		if listStyle == "unordered" {
			result = append(result, spaceLeft+"- "+fmt.Sprintf("%v", item.Content))
		} else {
			n := spaceLeft + strconv.Itoa(i+1) + "."
			result = append(result, fmt.Sprintf("%s %s", n, item.Content))
		}

		if len(item.Items) > 0 {
			result = append(result, CreateMarkDownNestedList(item.Items, listStyle, spaceLeft+"    "))
		}

	}

	return strings.Join(result[:], "\n")
}

func PrepareData(el domain.EditorJSBlock) (data interface{}) {
	jsonData, err := json.Marshal(el.Data)
	if err != nil {
		log.Println("Error when trying to marshall EditorJS block data\n", err)
	}

	switch el.Type {
	case "header":
		data = new(domain.EditorJSDataHeader)
	case "paragraph":
		data = new(domain.EditorJSDataParagraph)
	case "quote":
		data = new(domain.EditorJSDataQuote)
	case "warning":
		data = new(domain.EditorJSDataWarning)
	case "alert":
		data = new(domain.EditorJSDataAlert)
	case "list":
		data = new(domain.EditorJSDataList)
	case "checklist":
		data = new(domain.EditorJSDataChecklist)
	case "table":
		data = new(domain.EditorJSDataTable)
	case "AnyButton":
		data = new(domain.EditorJSDataAnyButton)
	case "code":
		data = new(domain.EditorJSDataCode)
	case "raw":
		data = new(domain.EditorJSDataRaw)
	case "image":
		data = new(domain.EditorJSDataImage)
	case "linkTool":
		data = new(domain.EditorJSDataLinkTool)
	case "attaches":
		data = new(domain.EditorJSDataAttaches)
	case "embed":
		data = new(domain.EditorJSDataEmbed)
	case "imageGallery":
		data = new(domain.EditorJSDataImageGallery)

	}

	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Println("Error when trying to unmarshall EditorJS block data\n", err)
	}

	return
}

func Separator(class string) string {
	return `<div class="` + class + `">&nbsp;</div>`
}

func HumanFileSize(fileSize float64) string {
	var sizePrefix string
	var formattedSize float64

	if math.Log10(+fileSize) >= 6 {
		sizePrefix = "MiB"
		formattedSize = fileSize / math.Pow(2, 20)
	} else {
		sizePrefix = "KiB"
		formattedSize = fileSize / math.Pow(2, 10)
	}

	return fmt.Sprintf("%v %v", toFixed(formattedSize, 1), sizePrefix)
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func WriteOutputFile(outputFilePath, outputContent, outputType string) (err error) {
	outputPath := outputFilePath

	if outputType == "html" && !strings.HasSuffix(outputFilePath, ".html") {
		outputPath = "output.html"
	} else if outputType == "markdown" && !strings.HasSuffix(outputFilePath, ".md") {
		outputPath = "output.md"
	}

	err = os.WriteFile(outputPath, []byte(outputContent), 0644)
	if err != nil {
		log.Println("Error writing the output file\n", err)
	}

	return
}

func ParseEditorJSON(editorJS string) domain.EditorJS {
	var result domain.EditorJS

	err := json.Unmarshal([]byte(editorJS), &result)
	if err != nil {
		log.Fatal("Error unmarshalling the input json file\n", err)
	}

	return result
}

func LoadStyleMap(path string) {
	content := LoadAsset(path, "json")

	err := json.Unmarshal(content, &SM)
	if err != nil {
		log.Fatal("Error unmarshalling the style config json file\n", err)
	}

	return
}

func LoadExternalStyleMap(path string) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("Error reading the external style file\n", err)
	}

	err = json.Unmarshal(content, &SM)
	if err != nil {
		log.Fatal("Error unmarshalling the style config json file\n", err)
	}

	return
}

func AppendBlockScript(blockScript string) (blockScriptOut string) {
	blockScriptMinified, _ := MinifyContent([]byte(blockScript), "js")
	blockScriptOut = string(blockScriptMinified)
	return
}

func IsValidStyle(style string) bool {
	for _, s := range config.AvailableStyles() {
		if s == style {
			return true
		}
	}
	return false
}
