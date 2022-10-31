package bulma

import (
	"fmt"
	"github.com/banjuanshu/go-editorjs/parser/html/common"
	sup "github.com/banjuanshu/go-editorjs/support"
	"github.com/banjuanshu/go-editorjs/support/config"
	"github.com/banjuanshu/go-editorjs/support/domain"
	"strings"
)

type Object struct {
	Data    interface{}
	Result  []string
	Styles  []string
	Scripts []string
}

const (
	StyleName  = "bulma"
	MapFile    = "bulma.json"
	ScriptFile = "bulma.js"
	ScriptType = "js"
)

func Init(useDefaultMap bool) (framework Object) {
	if useDefaultMap {
		sup.LoadStyleMap(MapFile)
	}
	return framework
}

func (o *Object) SetData(data interface{}) {
	o.Data = data
}

func (o *Object) SetStyles(styles []string) {
	for _, style := range styles {
		o.Styles = append(o.Styles, style)
	}
}

func (o *Object) SetResult(result string) {
	o.Result = append(o.Result, result)
}

func (o *Object) SetScripts(scripts []string) {
	for _, script := range scripts {
		o.Scripts = append(o.Scripts, script)
	}
}

func (o *Object) LoadLibrary() {
	for _, l := range sup.SM.LibraryPaths {
		o.Styles = append(o.Styles, `<link rel="stylesheet" href="`+l+`">`)
	}

	o.Scripts = append(o.Scripts, string(sup.MinifyAsset(config.AssetsScriptPath+ScriptFile, ScriptType)))
}

func (o *Object) CreatePage() string {
	return common.CreatePage(o.Scripts, o.Styles, o.Result)
}

func (o *Object) GetHtml() string {
	return common.GetHtml(o.Result)
}

func (o *Object) Separator() {
	o.SetResult(common.Separator())
}

func (o *Object) Header() {
	obj := o.Data.(*domain.EditorJSDataHeader)
	o.Result = append(o.Result, fmt.Sprintf(`<div class="content">%s</div>`, common.Header(obj)))
}

func (o *Object) Paragraph() {
	obj := o.Data.(*domain.EditorJSDataParagraph)
	o.Result = append(o.Result, fmt.Sprintf(`<div class="content">%s</div>`, common.Paragraph(obj)))
}

func (o *Object) Quote() {
	obj := o.Data.(*domain.EditorJSDataQuote)
	var output []string

	output = append(output, `<div class="content">`,
		`<blockquote class="`+sup.SM.Blocks.Quote.Blockquote+` `+sup.SM.Alignment[obj.Alignment]+`">`,
		obj.Text,
		`<p class="`+sup.SM.Blocks.Quote.Author+`">`,
		obj.Caption,
		`</p>`,
		`</blockquote>`,
		`</div>`)

	o.Result = append(o.Result, strings.Join(output[:], "\n"))
}

func (o *Object) Warning() {
	obj := o.Data.(*domain.EditorJSDataWarning)
	var output []string

	output = append(output, `<div class="`+sup.SM.Blocks.Warning.Block+`">`)

	if sup.SM.Blocks.Warning.CloseButton {
		output = append(output, `<button class="delete"></button>`)
	}

	output = append(output, `<span class="`+sup.SM.Blocks.Warning.Title+`">`,
		obj.Title,
		`</span>`,
		obj.Message,
		`</div>`)

	o.Result = append(o.Result, strings.Join(output[:], "\n"))
}

func (o *Object) Delimiter() {
	o.Result = append(o.Result, common.Delimiter())
}

func (o *Object) Alert() {
	obj := o.Data.(*domain.EditorJSDataAlert)
	var output []string

	output = append(output, `<div class="`+sup.SM.Blocks.Alert.Block+` `+sup.SM.Blocks.Alert.Types[obj.Type]+`">`)

	if sup.SM.Blocks.Alert.CloseButton {
		output = append(output, `<button class="delete"></button>`)
	}

	output = append(output, obj.Message,
		`</div>`)

	o.Result = append(o.Result, strings.Join(output[:], "\n"))
}

func (o *Object) List() {
	obj := o.Data.(*domain.EditorJSDataList)
	o.Result = append(o.Result, fmt.Sprintf(`<div class="content">%s</div>`, common.List(obj)))
}

func (o *Object) Checklist() {
	obj := o.Data.(*domain.EditorJSDataChecklist)
	o.Result = append(o.Result, common.Checklist(obj))
}

func (o *Object) Table() {
	obj := o.Data.(*domain.EditorJSDataTable)
	o.Result = append(o.Result, common.Table(obj))
}

func (o *Object) AnyButton() {
	obj := o.Data.(*domain.EditorJSDataAnyButton)
	o.Result = append(o.Result, common.AnyButton(obj))
}

func (o *Object) Code() {
	obj := o.Data.(*domain.EditorJSDataCode)
	o.Result = append(o.Result, common.Code(obj))
}

func (o *Object) Raw() {
	obj := o.Data.(*domain.EditorJSDataRaw)
	o.Result = append(o.Result, common.Raw(obj))
}

func (o *Object) Image() {
	obj := o.Data.(*domain.EditorJSDataImage)
	classes := ""
	classDiv := ""
	url := ""

	if obj.File.URL != "" {
		url = obj.File.URL
	} else {
		url = obj.URL
	}

	if obj.WithBorder {
		classes += sup.SM.Blocks.Image.Border + " "
	}

	if obj.Stretched {
		classes += sup.SM.Blocks.Image.Stretched
	}

	if obj.WithBackground {
		classDiv = sup.SM.Blocks.Image.Background
	}

	o.Result = append(o.Result, fmt.Sprintf(`<figure class="%s %s" ><img class="%s %s" src="%s" alt="%s" title="%s" /></figure>`, sup.SM.Blocks.Image.Block, classDiv, sup.SM.Blocks.Image.Image, classes, url, obj.Caption, obj.Caption))
}

func (o *Object) LinkTool() {
	obj := o.Data.(*domain.EditorJSDataLinkTool)
	var output []string

	output = append(output, `<a href="`+obj.Link+`" target="_Blank" rel="nofollow noindex noreferrer" class="`+sup.SM.Blocks.LinkTool.Link+`">`,
		`<div class="`+sup.SM.Blocks.LinkTool.Container+`">`,
		`<div class="`+sup.SM.Blocks.LinkTool.LeftColumn+`">`,
		`<div class="`+sup.SM.Blocks.LinkTool.Title+`">`,
		obj.Meta.Title,
		`</div>`,
		`<div class="`+sup.SM.Blocks.LinkTool.Description+`">`,
		obj.Meta.Description,
		`</div>`,
		`<div class="`+sup.SM.Blocks.LinkTool.LinkDescription+`">`,
		strings.ReplaceAll(strings.ReplaceAll(obj.Link, "https://", ""), "http://", ""),
		`</div>`,
		`</div>`,
		`<div class="`+sup.SM.Blocks.LinkTool.RightColumn+`">`,
		`<img class="`+sup.SM.Blocks.LinkTool.Image+`" src="`+obj.Meta.Image.URL+`" />`,
		`</div>`,
		`</div>`,
		`</a>`)

	o.Result = append(o.Result, strings.Join(output[:], "\n"))
}

func (o *Object) Attaches() {
	obj := o.Data.(*domain.EditorJSDataAttaches)
	var output []string

	output = append(output, `<a href="`+obj.File.URL+`" rel="noopener noreferrer" target="_blank" class="`+sup.SM.Blocks.Attaches.Link+`">`,
		`<div class="`+sup.SM.Blocks.Attaches.Container+`">`,
		`<div class="`+sup.SM.Blocks.Attaches.LeftColumn+`" >`,
		`<img class="`+sup.SM.Blocks.Attaches.LeftImage+`" src="https://i.ibb.co/K7Myr2k/file-icon.png" />`,
		`</div>`,
		`<div class="`+sup.SM.Blocks.Attaches.CenterColumn+`">`,
		`<div class="`+sup.SM.Blocks.Attaches.Filename+`">`,
		obj.File.Name,
		`</div>`,
		`<div class="`+sup.SM.Blocks.Attaches.Size+`">`,
		sup.HumanFileSize(obj.File.Size),
		`</div>`,
		`</div>`,
		`<div class="`+sup.SM.Blocks.Attaches.RightColumn+`" >`,
		`<img class="`+sup.SM.Blocks.Attaches.RightImage+`" src="https://i.ibb.co/VYyHr6C/download-icon.png" />`,
		`</div>`,
		`</div>`,
		`</a>`)

	o.Result = append(o.Result, strings.Join(output[:], "\n"))
}

func (o *Object) Embed() {
	obj := o.Data.(*domain.EditorJSDataEmbed)
	o.Result = append(o.Result, common.Embed(obj))
}

func (o *Object) ImageGallery() {
	obj := o.Data.(*domain.EditorJSDataImageGallery)
	r, s := common.ImageGallery(obj)

	o.Result = append(o.Result, r)
	o.Scripts = append(o.Scripts, sup.AppendBlockScript(s))
}
