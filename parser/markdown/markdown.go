package markdown

import (
	"encoding/json"
	"fmt"
	"log"
	"rbsite/internal/editorjs/support"
	"rbsite/internal/editorjs/support/domain"
	"strconv"
	"strings"
)

func Header(el *domain.EditorJSDataHeader) string {
	headerLevel := ""

	for i := 0; i < el.Level; i++ {
		headerLevel += "#"
	}

	return fmt.Sprintf("%s %s", headerLevel, el.Text)
}

func Paragraph(el *domain.EditorJSDataParagraph) string {
	return fmt.Sprintf("%s", el.Text)
}

func Quote(el *domain.EditorJSDataQuote) string {
	quoteMD := `> ` + el.Text + `
>
> --- ` + el.Caption

	return fmt.Sprintf("%s", quoteMD)
}

func Warning(el *domain.EditorJSDataWarning) string {
	warningMD := `> ` + el.Title + `
>
> --- ` + el.Message

	return fmt.Sprintf("%s", warningMD)
}

func Delimiter() string {
	delimiterMD := `
***
`
	return fmt.Sprintf("%s", delimiterMD)
}

func Alert(el *domain.EditorJSDataAlert) string {
	alertMD := `
| ` + el.Message + ` |
| --- |`

	return fmt.Sprintf("%s", alertMD)
}

func List(el *domain.EditorJSDataList) string {
	var result []string

	items, err := json.Marshal(el.Items)
	if err != nil {
		log.Fatal(err)
	}

	var itemsList []domain.NestedListItem

	err = json.Unmarshal(items, &itemsList)
	if err == nil {
		result = append(result, support.CreateMarkDownNestedList(itemsList, el.Style, ""))

	} else {
		if el.Style == "unordered" {
			for _, item := range el.Items {
				result = append(result, "- "+fmt.Sprintf("%v", item))
			}

		} else {
			for i, item := range el.Items {
				n := strconv.Itoa(i+1) + "."
				result = append(result, fmt.Sprintf("%s %s", n, item))
			}
		}
	}

	return strings.Join(result[:], "\n")
}

func Checklist(el *domain.EditorJSDataChecklist) string {
	var result []string

	items, err := json.Marshal(el.Items)
	if err != nil {
		log.Fatal(err)
	}

	var itemsList []domain.ChecklistItem

	err = json.Unmarshal(items, &itemsList)
	if err == nil {

		for _, item := range itemsList {
			if item.Checked {
				result = append(result, `- [x] `+item.Text)
			} else {
				result = append(result, `- [ ] `+item.Text)
			}
		}
	}

	return strings.Join(result[:], "\n")
}

func Table(el *domain.EditorJSDataTable) string {
	var result []string

	for index, line := range el.Content {

		if el.WithHeadings && index == 0 {
			lineTitle := `|`
			lineSeparator := `|`

			for _, info := range line {
				lineTitle += ` ` + info + ` |`
				lineSeparator += `---|`
			}

			result = append(result, lineTitle)
			result = append(result, lineSeparator)

		} else if index == 0 {
			lineTitle := `|`
			lineSeparator := `|`
			lineContent := `|`

			for _, info := range line {
				lineTitle += ` |`
				lineSeparator += `---|`
				lineContent += ` ` + info + ` |`
			}

			result = append(result, lineTitle)
			result = append(result, lineSeparator)
			result = append(result, lineContent)

		} else {
			lineData := `|`

			for _, info := range line {
				lineData += ` ` + info + ` |`
			}

			result = append(result, lineData)
		}
	}

	return strings.Join(result, "\n")
}

func AnyButton(el *domain.EditorJSDataAnyButton) string {
	return fmt.Sprintf(`[%s](%s)`, el.Text, el.Link)
}

func Code(el *domain.EditorJSDataCode) string {
	var result []string

	result = append(result, "```"+el.LanguageCode)
	result = append(result, el.Code)
	result = append(result, "```")

	return strings.Join(result, "\n")
}

func Raw(el *domain.EditorJSDataRaw) string {
	var result []string

	result = append(result, "```")
	result = append(result, el.Html)
	result = append(result, "```")

	return strings.Join(result, "\n")
}

func Image(el *domain.EditorJSDataImage) string {
	url := ""

	if el.File.URL != "" {
		url = el.File.URL
	} else {
		url = el.URL
	}

	return fmt.Sprintf(`![%s](%s)`, el.Caption, url)
}

func LinkTool(el *domain.EditorJSDataLinkTool) string {
	var result []string

	result = append(result, "---")
	result = append(result, "")
	result = append(result, `# `+el.Meta.Title)
	result = append(result, "")
	result = append(result, `![`+el.Meta.Title+`](`+el.Meta.Image.URL+`)`)
	result = append(result, "")
	result = append(result, `*`+el.Meta.Description+`*`)
	result = append(result, "")
	result = append(result, `[`+strings.ReplaceAll(strings.ReplaceAll(el.Link, "https://", ""), "http://", "")+`](`+el.Link+`)`)
	result = append(result, "")
	result = append(result, "---")

	return strings.Join(result, "\n")
}

func Attaches(el *domain.EditorJSDataAttaches) string {
	var result []string

	result = append(result, "---")
	result = append(result, "")
	result = append(result, `### `+el.File.Name)
	result = append(result, "")
	result = append(result, `###### `+support.HumanFileSize(el.File.Size))
	result = append(result, "")
	result = append(result, `[Download](`+el.File.URL+`)`)
	result = append(result, "")
	result = append(result, "---")

	return strings.Join(result, "\n")
}

func Embed(el *domain.EditorJSDataEmbed) string {
	var result []string

	result = append(result, "---")
	result = append(result, "")
	result = append(result, `### `+el.Caption)
	result = append(result, "")
	result = append(result, `[Watch on `+el.Service+`](`+el.Source+`)`)
	result = append(result, "")
	result = append(result, "---")

	return strings.Join(result, "\n")
}

func ImageGallery(el *domain.EditorJSDataImageGallery) string {
	var result []string

	result = append(result, "---")

	for index, img := range el.URLs {
		result = append(result, `![Image `+strconv.Itoa(index)+`](`+img+`)`)
	}

	result = append(result, "---")

	return strings.Join(result, "\n")
}
