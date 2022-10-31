package bootstrap

import (
	"github.com/banjuanshu/go-editorjs/parser/html/common"
	"github.com/banjuanshu/go-editorjs/support"
	"github.com/banjuanshu/go-editorjs/support/config"
	"github.com/banjuanshu/go-editorjs/support/domain"
)

type Object struct {
	Data    interface{}
	Result  []string
	Styles  []string
	Scripts []string
}

const (
	StyleName  = "bootstrap"
	MapFile    = "bootstrap.json"
	ScriptFile = "bootstrap.js"
	ScriptType = "js"
)

func Init(useDefaultMap bool) (framework Object) {
	if useDefaultMap {
		support.LoadStyleMap(MapFile)
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
	for _, l := range support.SM.LibraryPaths {
		o.Styles = append(o.Styles, `<link rel="stylesheet" href="`+l+`">`)
	}

	o.Scripts = append(o.Scripts, string(support.MinifyAsset(config.AssetsScriptPath+ScriptFile, ScriptType)))
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
	o.Result = append(o.Result, common.Header(obj))
}

func (o *Object) Paragraph() {
	obj := o.Data.(*domain.EditorJSDataParagraph)
	o.Result = append(o.Result, common.Paragraph(obj))
}
func (o *Object) Quote() {
	obj := o.Data.(*domain.EditorJSDataQuote)
	o.Result = append(o.Result, common.Quote(obj))
}

func (o *Object) Warning() {
	obj := o.Data.(*domain.EditorJSDataWarning)
	o.Result = append(o.Result, common.Warning(obj))
}

func (o *Object) Delimiter() {
	o.Result = append(o.Result, common.Delimiter())
}

func (o *Object) Alert() {
	obj := o.Data.(*domain.EditorJSDataAlert)
	o.Result = append(o.Result, common.Alert(obj))
}

func (o *Object) List() {
	obj := o.Data.(*domain.EditorJSDataList)
	o.Result = append(o.Result, common.List(obj))
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
	o.Result = append(o.Result, common.Image(obj))
}

func (o *Object) LinkTool() {
	obj := o.Data.(*domain.EditorJSDataLinkTool)
	o.Result = append(o.Result, common.LinkTool(obj))
}

func (o *Object) Attaches() {
	obj := o.Data.(*domain.EditorJSDataAttaches)
	o.Result = append(o.Result, common.Attaches(obj))
}

func (o *Object) Embed() {
	obj := o.Data.(*domain.EditorJSDataEmbed)
	o.Result = append(o.Result, common.Embed(obj))
}

func (o *Object) ImageGallery() {
	obj := o.Data.(*domain.EditorJSDataImageGallery)
	r, s := common.ImageGallery(obj)

	o.Result = append(o.Result, r)
	o.Scripts = append(o.Scripts, support.AppendBlockScript(s))
}
