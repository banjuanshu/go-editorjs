package common

import (
	"encoding/json"
	"fmt"
	sup "github.com/banjuanshu/go-editorjs/support"
	"github.com/banjuanshu/go-editorjs/support/domain"
	"log"
	"strconv"
	"strings"
)

func CreatePage(scripts, styles, result []string) string {
	script := "\n\n<script>\n" + strings.Join(scripts[:], "\n") + "\n</script>\n\n"

	page := `<!DOCTYPE html>
<html>
  <head>
`
	for _, h := range sup.SM.PageHead {
		page += h + `
`
	}

	page += strings.Join(styles[:], "\n") + ` 
  </head>
  <body>
` + strings.Join(result[:], "\n\n") + script + ` 
</body>
</html>
`

	return page
}

func GetHtml(result []string) string {

	return strings.Join(result[:], "\n\n")
}

func Separator() (separator string) {
	if sup.SM.SpaceBetweenBlocks != "" {
		separator = sup.Separator(sup.SM.SpaceBetweenBlocks)
	}
	return
}

func Header(el *domain.EditorJSDataHeader) string {
	anchor := ""
	if el.Anchor != "" {
		anchor = `id="` + strings.ToLower(strings.ReplaceAll(el.Anchor, " ", "-")) + `"`
	}

	tag := `h` + strconv.Itoa(el.Level)

	class := `class="` + sup.SM.Blocks.Header[tag] + `"`

	level := strconv.Itoa(el.Level)

	return fmt.Sprintf("<h%s %s %s>%s</h%s>", level, anchor, class, el.Text, level)
}

func Paragraph(el *domain.EditorJSDataParagraph) string {
	return fmt.Sprintf("<p%s>%s</p>", ` class="`+sup.SM.Blocks.Paragraph+` `+sup.SM.Alignment[el.Alignment]+`"`, el.Text)
}

func Quote(el *domain.EditorJSDataQuote) string {
	var output []string

	output = append(output, `<figure class="`+sup.SM.Blocks.Quote.Figure+` `+sup.SM.Alignment[el.Alignment]+`">`,
		`<blockquote class="`+sup.SM.Blocks.Quote.Blockquote+`">`,
		el.Text,
		`</blockquote>`,
		`<figcaption class="`+sup.SM.Blocks.Quote.Figcaption+`">`,
		el.Caption,
		`</figcaption>`,
		`</figure>`)

	return strings.Join(output[:], "\n")
}

func Warning(el *domain.EditorJSDataWarning) string {
	var output []string

	output = append(output, `<div class="`+sup.SM.Blocks.Warning.Block+`">`,
		`<b>`,
		el.Title,
		`</b>`,
		el.Message,
		`</div>`)

	return strings.Join(output[:], "\n")
}

func Delimiter() string {
	var output []string

	output = append(output, `<div class="`+sup.SM.Blocks.Delimiter+`">***</div>`)

	return strings.Join(output[:], "\n")
}

func Alert(el *domain.EditorJSDataAlert) string {
	var output []string

	output = append(output, `<div class="`+sup.SM.Blocks.Alert.Block+` `+sup.SM.Blocks.Alert.Types[el.Type]+`">`,
		el.Message,
		`</div>`)

	return strings.Join(output[:], "\n")
}

func List(el *domain.EditorJSDataList) string {
	var output []string

	listStyle := "ol"
	if el.Style == "unordered" {
		listStyle = "ul"
	}

	items, err := json.Marshal(el.Items)
	if err != nil {
		log.Fatal(err)
	}

	var itemsList []domain.NestedListItem

	err = json.Unmarshal(items, &itemsList)
	if err == nil {
		output = append(output, sup.CreateHTMLNestedList(itemsList, listStyle, true))
	} else {
		output = append(output, `<`+listStyle+` class="`+sup.SM.Blocks.List.Group+`">`)

		for _, item := range el.Items {
			output = append(output, `<li class="`+sup.SM.Blocks.List.Item+`">`+fmt.Sprintf("%v", item)+`</li>`)
		}

		output = append(output, `</`+listStyle+`>`)
	}

	return strings.Join(output[:], "\n")
}

func Checklist(el *domain.EditorJSDataChecklist) string {
	var output []string

	items, err := json.Marshal(el.Items)
	if err != nil {
		log.Fatal(err)
	}

	var itemsList []domain.ChecklistItem

	err = json.Unmarshal(items, &itemsList)
	if err == nil {
		output = append(output, `<div class="`+sup.SM.Blocks.Checklist.Block+`">`)

		for _, item := range itemsList {
			output = append(output, `<div class="`+sup.SM.Blocks.Checklist.Item+`">`)

			if item.Checked {
				output = append(output, `<span class="`+sup.SM.Blocks.Checklist.CheckboxChecked+`">&#10004;</span>`)
			} else {
				output = append(output, `<span class="`+sup.SM.Blocks.Checklist.CheckboxUnchecked+`">&nbsp;-&nbsp;</span>`)
			}

			output = append(output, `<span class="`+sup.SM.Blocks.Checklist.Text+`">`+item.Text+`</span>`,
				`</div>`)
		}

		output = append(output, `</div>`)
	}

	return strings.Join(output[:], "\n")
}

func Table(el *domain.EditorJSDataTable) string {
	var output []string

	output = append(output, `<table class="`+sup.SM.Blocks.Table.Table+`">`)

	for index, line := range el.Content {
		output = append(output, `<tr class="`+sup.SM.Blocks.Table.Row+`">`)

		tag := `td`
		tagClass := `class="` + sup.SM.Blocks.Table.CellTD + `"`
		if el.WithHeadings && index == 0 {
			tag = `th`
			tagClass = `class="` + sup.SM.Blocks.Table.CellTH + `"`
		}

		for _, info := range line {
			output = append(output, `<`+tag+` `+tagClass+`>`+info+`</`+tag+`>`)
		}

		output = append(output, `</tr>`)
	}

	output = append(output, `</table>`)

	return strings.Join(output[:], "\n")
}

func AnyButton(el *domain.EditorJSDataAnyButton) string {
	var output []string

	output = append(output, `<a class="`+sup.SM.Blocks.AnyButton+`" href="`+el.Link+`">`+el.Text+`</a>`)

	return strings.Join(output[:], "\n")
}

func Code(el *domain.EditorJSDataCode) string {
	var output []string

	output = append(output, `<pre class="`+sup.SM.Blocks.Code.Pre+`">`,
		`<code class="`+sup.SM.Blocks.Code.Code+`">`+el.Code,
		`</code></pre>`)

	return strings.Join(output[:], "\n")
}

func Raw(el *domain.EditorJSDataRaw) string {
	var output []string

	content := strings.ReplaceAll(el.Html, "<", "&lt;")
	content = strings.ReplaceAll(content, ">", "&gt;")

	output = append(output, `<pre class="`+sup.SM.Blocks.Raw.Pre+`">`,
		`<code class="`+sup.SM.Blocks.Raw.Code+`">`+content,
		`</code></pre>`)

	return strings.Join(output[:], "\n")
}

func Image(el *domain.EditorJSDataImage) string {
	classes := ""
	classDiv := ""
	url := ""

	if el.File.URL != "" {
		url = el.File.URL
	} else {
		url = el.URL
	}

	if el.WithBorder {
		classes += sup.SM.Blocks.Image.Border + " "
	}

	if el.Stretched {
		classes += sup.SM.Blocks.Image.Stretched
	}

	if el.WithBackground {
		classDiv = sup.SM.Blocks.Image.Background
	}

	return fmt.Sprintf(`<div class="%s" ><img class="%s" src="%s" alt="%s" title="%s" /></div>`, classDiv, classes, url, el.Caption, el.Caption)
}

func LinkTool(el *domain.EditorJSDataLinkTool) string {
	var output []string

	output = append(output, `<a href="`+el.Link+`" target="_Blank" rel="nofollow noindex noreferrer" class="`+sup.SM.Blocks.LinkTool.Link+`">`,
		`<div class="`+sup.SM.Blocks.LinkTool.Container+`">`,
		`<div class="`+sup.SM.Blocks.LinkTool.Row+`">`,
		`<div class="`+sup.SM.Blocks.LinkTool.LeftColumn+`">`,
		`<div class="`+sup.SM.Blocks.LinkTool.Title+`">`,
		el.Meta.Title,
		`</div>`,
		`<div class="`+sup.SM.Blocks.LinkTool.Description+`">`,
		el.Meta.Description,
		`</div>`,
		`<div class="`+sup.SM.Blocks.LinkTool.LinkDescription+`">`,
		strings.ReplaceAll(strings.ReplaceAll(el.Link, "https://", ""), "http://", ""),
		`</div>`,
		`</div>`,
		`<div class="`+sup.SM.Blocks.LinkTool.RightColumn+`">`,
		`<img class="`+sup.SM.Blocks.LinkTool.Image+`" src="`+el.Meta.Image.URL+`" />`,
		`</div>`,
		`</div>`,
		`</div>`,
		`</a>`)

	return strings.Join(output[:], "\n")
}

func Attaches(el *domain.EditorJSDataAttaches) string {
	var output []string

	output = append(output, `<a href="`+el.File.URL+`" rel="noopener noreferrer" target="_blank" class="`+sup.SM.Blocks.Attaches.Link+`">`,
		`<div class="`+sup.SM.Blocks.Attaches.Container+`">`,
		`<div class="`+sup.SM.Blocks.Attaches.Row+`" >`,
		`<div class="`+sup.SM.Blocks.Attaches.LeftColumn+`" >`,
		`<img class="`+sup.SM.Blocks.Attaches.LeftImage+`" src="https://i.ibb.co/K7Myr2k/file-icon.png" />`,
		`</div>`,
		`<div class="`+sup.SM.Blocks.Attaches.CenterColumn+`">`,
		`<div class="`+sup.SM.Blocks.Attaches.Filename+`">`,
		el.File.Name,
		`</div>`,
		`<div class="`+sup.SM.Blocks.Attaches.Size+`">`,
		sup.HumanFileSize(el.File.Size),
		`</div>`,
		`</div>`,
		`<div class="`+sup.SM.Blocks.Attaches.RightColumn+`" >`,
		`<img class="`+sup.SM.Blocks.Attaches.RightImage+`" src="https://i.ibb.co/VYyHr6C/download-icon.png" />`,
		`</div>`,
		`</div>`,
		`</div>`,
		`</a>`)

	return strings.Join(output[:], "\n")
}

func Embed(el *domain.EditorJSDataEmbed) string {
	var output []string

	output = append(output, `<div class="`+sup.SM.Blocks.Embed.Block+`" style="max-width: `+strconv.Itoa(el.Width)+`px">`,
		`<div class="`+sup.SM.Blocks.Embed.Title+`">`+el.Caption+`</div>`,
		`<iframe width="`+strconv.Itoa(el.Width)+`" height="`+strconv.Itoa(el.Height)+`" src="`+el.Embed+`" title="`+el.Caption+`" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>`,
		`<div class="`+sup.SM.Blocks.Embed.Bottom+`">`,
		`<a class="`+sup.SM.Blocks.Embed.Link+`" href="`+el.Source+`" target="_Blank">Watch on `+el.Service+`</a>`,
		`</div>`,
		`</div>`)
	return strings.Join(output[:], "\n")
}

func ImageGallery(el *domain.EditorJSDataImageGallery) (string, string) {
	galleryScript := `
    const gg = new GalleryGrid();
    gg.loadGallery();
`
	galleryId := ""
	galleryClass := "gg-box"
	galleryHTML := `<div class="gg-container">`

	if el.BkgMode {
		galleryClass += " dark"
		galleryScript += `
	gg.galleryOptions({
	    selector: ".dark",
	    darkMode: true
	});`
	}

	if el.LayoutDefault {
		galleryId = ""
	} else if el.LayoutHorizontal {
		galleryId = "horizontal"
		galleryScript += `
	gg.galleryOptions({
		selector: "#horizontal",
		layout: "horizontal"
	});`
	} else if el.LayoutSquare {
		galleryId = "square"
		galleryScript += `
	gg.galleryOptions({
		selector: "#square",
		layout: "square"
	});`
	} else if el.LayoutWithGap {
		galleryId = "gap"
		galleryScript += `
	gg.galleryOptions({
		selector: "#gap",
		gapLength: 10
	});`
	} else if el.LayoutWithFixedSize {
		galleryId = "heightWidth"
		galleryScript += `
	gg.galleryOptions({
		selector: "#heightWidth",
		rowHeight: 180,
		columnWidth: 280
	});`
	}

	galleryHTML += fmt.Sprintf(`
<div class="%s" id="%s">`, galleryClass, galleryId)

	for index, url := range el.URLs {
		galleryHTML += fmt.Sprintf(`
<img src="%s" id="gg-image-%s" />`, url, strconv.Itoa(index))
	}

	galleryHTML += `
</div>
</div>`

	return galleryHTML, galleryScript
}
