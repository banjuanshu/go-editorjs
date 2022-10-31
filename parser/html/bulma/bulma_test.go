package bulma

import (
	"github.com/banjuanshu/go-editorjs/support"
	"github.com/banjuanshu/go-editorjs/support/domain"
	"github.com/matryer/is"
	"strconv"

	"testing"
)

var obj = Init(true)

func TestHeaderBlock(t *testing.T) {
	is := is.New(t)

	for i := 1; i <= 6; i++ {
		level := strconv.Itoa(i)

		input1 := `{
    "blocks": [
        {
            "type": "header",
            "data": {
                "level": ` + level + `,
                "text": "Level ` + level + ` Header"
            }
        }
    ]
}`

		editorJSON1 := support.ParseEditorJSON(input1)
		obj.Data = support.PrepareData(editorJSON1.Blocks[0]).(*domain.EditorJSDataHeader)

		expected1 := `<div class="content"><h` + level + `  class="title is-` + level + `">Level ` + level + ` Header</h` + level + `></div>`

		obj.Header()

		actual1 := obj.Result[i-1]

		is.Equal(expected1, actual1) // Header 1 is different from expected
	}

	obj.Result = []string{}

	for i := 1; i <= 6; i++ {
		level := strconv.Itoa(i)

		input1 := `{
    "blocks": [
        {
            "type": "header",
            "data": {
                "level": ` + level + `,
                "text": "Level ` + level + ` Header",
				"anchor": "Anchor Text ` + level + `"
            }
        }
    ]
}`

		editorJSON1 := support.ParseEditorJSON(input1)
		obj.Data = support.PrepareData(editorJSON1.Blocks[0])

		expected1 := `<div class="content"><h` + level + ` id="anchor-text-` + level + `" class="title is-` + level + `">Level ` + level + ` Header</h` + level + `></div>`

		obj.Header()

		actual1 := obj.Result[i-1]

		is.Equal(expected1, actual1) // Header 2 is different from expected
	}
}

func TestParagraphBlock(t *testing.T) {
	is := is.New(t)

	obj.Result = []string{}

	input1 := `{
    "blocks": [
        {
			"type": "paragraph",
            "data": {
                "text": "I am a paragraph!"
            }
        }
    ]
}`

	editorJSON1 := support.ParseEditorJSON(input1)
	obj.Data = support.PrepareData(editorJSON1.Blocks[0])

	expected1 := `<div class="content"><p class=" ">I am a paragraph!</p></div>`

	obj.Paragraph()

	actual1 := obj.Result[0]

	is.Equal(expected1, actual1) // Paragraph 1 is different from expected

	obj.Result = []string{}

	input2 := `{
    "blocks": [
        {
			"type": "paragraph",
            "data": {
                "alignment": "center",
                "text": "I am a paragraph!"
            }
        }
    ]
}`

	editorJSON2 := support.ParseEditorJSON(input2)
	obj.Data = support.PrepareData(editorJSON2.Blocks[0])

	expected2 := `<div class="content"><p class=" has-text-centered">I am a paragraph!</p></div>`

	obj.Paragraph()

	actual2 := obj.Result[0]

	is.Equal(expected2, actual2) // Paragraph 2 is different from expected
}

func TestQuoteBlock(t *testing.T) {
	is := is.New(t)

	obj.Result = []string{}

	input := `{
    "blocks": [
        {
			"type": "quote",
            "data": {
                "alignment": "center",
                "caption": "Lao Tzu",
                "text": "The journey of a thousand miles begins with one step."
            }
        }
    ]
}`

	editorJSON := support.ParseEditorJSON(input)
	obj.Data = support.PrepareData(editorJSON.Blocks[0])

	expected := `<div class="content">
<blockquote class=" has-text-centered">
The journey of a thousand miles begins with one step.
<p class="is-italic">
Lao Tzu
</p>
</blockquote>
</div>`

	obj.Quote()

	actual := obj.Result[0]

	is.Equal(expected, actual) // Quote is different from expected
}

func TestWarningBlock(t *testing.T) {
	is := is.New(t)

	obj.Result = []string{}

	input := `{
    "blocks": [
        {
			"type": "warning",
            "data": {
                "message": "Avoid using this method just for lulz. It can be very dangerous opposite your daily fun stuff.",
                "title": "Note:"
            }
        }
    ]
}`

	editorJSON := support.ParseEditorJSON(input)
	obj.Data = support.PrepareData(editorJSON.Blocks[0])

	expected := `<div class="notification is-warning is-light">
<button class="delete"></button>
<span class="has-text-weight-bold">
Note:
</span>
Avoid using this method just for lulz. It can be very dangerous opposite your daily fun stuff.
</div>`

	obj.Warning()

	actual := obj.Result[0]

	is.Equal(expected, actual) // Warning is different from expected
}

func TestDelimiterBlock(t *testing.T) {
	is := is.New(t)

	obj.Result = []string{}

	expected := `<div class="has-text-centered is-size-4">***</div>`

	obj.Delimiter()

	actual := obj.Result[0]

	is.Equal(expected, actual) // Delimiter is different from expected
}

func TestAlertBlock(t *testing.T) {
	is := is.New(t)

	obj.Result = []string{}

	input := `{
    "blocks": [
        {
			"type": "alert",
            "data": {
                "message": "Something happened that you should know about.",
                "type": "primary"
            }
        }
    ]
}`

	editorJSON := support.ParseEditorJSON(input)
	obj.Data = support.PrepareData(editorJSON.Blocks[0])

	expected := `<div class="notification is-primary">
<button class="delete"></button>
Something happened that you should know about.
</div>`

	obj.Alert()

	actual := obj.Result[0]

	is.Equal(expected, actual) // Alert is different from expected
}

func TestListBlock(t *testing.T) {
	is := is.New(t)

	obj.Result = []string{}

	input1 := `{
    "blocks": [
        {
			"type": "list",
            "data": {
                "items": [
                    "This is a block-styled editor",
                    "Clean output data",
                    "Simple and powerful API"
                ],
                "style": "unordered"
            }
        }
    ]
}`

	editorJSON1 := support.ParseEditorJSON(input1)
	obj.Data = support.PrepareData(editorJSON1.Blocks[0])

	expected1 := `<div class="content"><ul class="content">
<li class="">This is a block-styled editor</li>
<li class="">Clean output data</li>
<li class="">Simple and powerful API</li>
</ul></div>`

	obj.List()

	actual1 := obj.Result[0]

	is.Equal(expected1, actual1) // List 1 is different from expected

	obj.Result = []string{}

	input2 := `{
    "blocks": [
        {
			"type": "list",
            "data": {
                "items": [
                    {
                        "content": "Cars",
                        "items": [
                            {
                                "content": "BMW",
                                "items": [
                                    {
                                        "content": "Z3",
                                        "items": []
                                    },
                                    {
                                        "content": "Z4",
                                        "items": []
                                    }
                                ]
                            },
                            {
                                "content": "Audi",
                                "items": [
                                    {
                                        "content": "A3",
                                        "items": []
                                    },
                                    {
                                        "content": "A1",
                                        "items": []
                                    }
                                ]
                            }
                        ]
                    },
                    {
                        "content": "Motorcycle",
                        "items": [
                            {
                                "content": "Ducati",
                                "items": [
                                    {
                                        "content": "916",
                                        "items": []
                                    }
                                ]
                            },
                            {
                                "content": "Yamanha",
                                "items": [
                                    {
                                        "content": "DT 180",
                                        "items": []
                                    }
                                ]
                            },
                            {
                                "content": "Honda",
                                "items": [
                                    {
                                        "content": "VFR 750R",
                                        "items": []
                                    }
                                ]
                            }
                        ]
                    }
                ],
                "style": "ordered"
            }
        }
    ]
}`

	editorJSON2 := support.ParseEditorJSON(input2)
	obj.Data = support.PrepareData(editorJSON2.Blocks[0])

	expected2 := `<div class="content"><ol class="content">
<li class="">Cars</li>
<ol class="">
<li class="">BMW</li>
<ol class="">
<li class="">Z3</li>
</li>
<li class="">Z4</li>
</li>
</ol>
</li>
<li class="">Audi</li>
<ol class="">
<li class="">A3</li>
</li>
<li class="">A1</li>
</li>
</ol>
</li>
</ol>
</li>
<li class="">Motorcycle</li>
<ol class="">
<li class="">Ducati</li>
<ol class="">
<li class="">916</li>
</li>
</ol>
</li>
<li class="">Yamanha</li>
<ol class="">
<li class="">DT 180</li>
</li>
</ol>
</li>
<li class="">Honda</li>
<ol class="">
<li class="">VFR 750R</li>
</li>
</ol>
</li>
</ol>
</li>
</ol></div>`

	obj.List()

	actual2 := obj.Result[0]

	is.Equal(expected2, actual2) // List 2 is different from expected
}

func TestChecklistBlock(t *testing.T) {
	is := is.New(t)

	obj.Result = []string{}

	input := `{
    "blocks": [
        {
            "type": "checklist",
            "data": {
                "items": [
                    {
                        "checked": true,
                        "text": "This is a block-styled editor"
                    },
                    {
                        "checked": false,
                        "text": "Clean output data"
                    },
                    {
                        "checked": true,
                        "text": "Simple and powerful API"
                    }
                ]
            }
        }
    ]
}`

	editorJSON := support.ParseEditorJSON(input)
	obj.Data = support.PrepareData(editorJSON.Blocks[0])

	expected := `<div class="">
<div class="">
<span class="tag is-link is-rounded">&#10004;</span>
<span class="">This is a block-styled editor</span>
</div>
<div class="">
<span class="tag is-rounded">&nbsp;-&nbsp;</span>
<span class="">Clean output data</span>
</div>
<div class="">
<span class="tag is-link is-rounded">&#10004;</span>
<span class="">Simple and powerful API</span>
</div>
</div>`

	obj.Checklist()

	actual := obj.Result[0]

	is.Equal(expected, actual) // Checklist is different from expected
}

func TestTableBlock(t *testing.T) {
	is := is.New(t)

	obj.Result = []string{}

	input1 := `{
    "blocks": [
        {
            "type": "table",
            "data": {
                "content": [
                    [
                        "Kine",
                        "Pigs",
                        "Chicken"
                    ],
                    [
                        "1 pcs",
                        "3 pcs",
                        "12 pcs"
                    ],
                    [
                        "100$",
                        "200$",
                        "150$"
                    ]
                ],
                "withHeadings": true
            }
        }
    ]
}`

	editorJSON1 := support.ParseEditorJSON(input1)
	obj.Data = support.PrepareData(editorJSON1.Blocks[0])

	expected1 := `<table class="table is-striped is-bordered">
<tr class="">
<th class="">Kine</th>
<th class="">Pigs</th>
<th class="">Chicken</th>
</tr>
<tr class="">
<td class="">1 pcs</td>
<td class="">3 pcs</td>
<td class="">12 pcs</td>
</tr>
<tr class="">
<td class="">100$</td>
<td class="">200$</td>
<td class="">150$</td>
</tr>
</table>`

	obj.Table()

	actual1 := obj.Result[0]

	is.Equal(expected1, actual1) // Table 1 is different from expected

	obj.Result = []string{}

	input2 := `{
    "blocks": [
        {
            "type": "table",
            "data": {
                "content": [
                    [
                        "Kine",
                        "1 pcs",
                        "100$"
                    ],
                    [
                        "Pigs",
                        "3 pcs",
                        "200$"
                    ],
                    [
                        "Chickens",
                        "12 pcs",
                        "150$"
                    ]
                ]
            }
        }
    ]
}`

	editorJSON2 := support.ParseEditorJSON(input2)
	obj.Data = support.PrepareData(editorJSON2.Blocks[0])

	expected2 := `<table class="table is-striped is-bordered">
<tr class="">
<td class="">Kine</td>
<td class="">1 pcs</td>
<td class="">100$</td>
</tr>
<tr class="">
<td class="">Pigs</td>
<td class="">3 pcs</td>
<td class="">200$</td>
</tr>
<tr class="">
<td class="">Chickens</td>
<td class="">12 pcs</td>
<td class="">150$</td>
</tr>
</table>`

	obj.Table()

	actual2 := obj.Result[0]

	is.Equal(expected2, actual2) // Table 2 is different from expected
}

func TestAnyButtonBlock(t *testing.T) {
	is := is.New(t)

	obj.Result = []string{}

	input := `{
    "blocks": [
        {
            "type": "AnyButton",
            "data": {
                "link": "https://editorjs.io/",
                "text": "editorjs official"
            }
        }
    ]
}`

	editorJSON := support.ParseEditorJSON(input)
	obj.Data = support.PrepareData(editorJSON.Blocks[0])

	expected := `<a class="button is-link" href="https://editorjs.io/">editorjs official</a>`

	obj.AnyButton()

	actual := obj.Result[0]

	is.Equal(expected, actual) // AnyButton is different from expected
}

func TestCodeBlock(t *testing.T) {
	is := is.New(t)

	obj.Result = []string{}

	input1 := `{
    "blocks": [
        {
            "type": "code",
            "data": {
                "code": "body {\n font-size: 14px;\n line-height: 16px;\n}"
            }
        }
    ]
}`

	editorJSON1 := support.ParseEditorJSON(input1)
	obj.Data = support.PrepareData(editorJSON1.Blocks[0])

	expected1 := `<pre class="">
<code class="">body {
 font-size: 14px;
 line-height: 16px;
}
</code></pre>`

	obj.Code()

	actual1 := obj.Result[0]

	is.Equal(expected1, actual1) // Code 1 is different from expected

	obj.Result = []string{}

	input2 := `{
    "blocks": [
        {
            "type": "code",
            "data": {
                "code": "body {\n font-size: 14px;\n line-height: 16px;\n}",
                "languageCode": "css"
            }
        }
    ]
}`

	editorJSON2 := support.ParseEditorJSON(input2)
	obj.Data = support.PrepareData(editorJSON2.Blocks[0])

	expected2 := `<pre class="">
<code class="">body {
 font-size: 14px;
 line-height: 16px;
}
</code></pre>`

	obj.Code()

	actual2 := obj.Result[0]

	is.Equal(expected2, actual2) // Code 2 is different from expected
}

func TestRawBlock(t *testing.T) {
	is := is.New(t)

	obj.Result = []string{}

	input := `{
    "blocks": [
        {
            "type": "raw",
            "data": {
                "html": "<div style=\"background: #000; color: #fff; font-size: 30px; padding: 50px;\">Any HTML code</div>"
            }
        }
    ]
}`

	editorJSON := support.ParseEditorJSON(input)
	obj.Data = support.PrepareData(editorJSON.Blocks[0])

	expected := `<pre class="">
<code class="">&lt;div style="background: #000; color: #fff; font-size: 30px; padding: 50px;"&gt;Any HTML code&lt;/div&gt;
</code></pre>`

	obj.Raw()

	actual := obj.Result[0]

	is.Equal(expected, actual) // Raw is different from expected
}

func TestImageBlock(t *testing.T) {
	is := is.New(t)

	obj.Result = []string{}

	input1 := `{
    "blocks": [
        {
		  "type" : "image",
		  "data" : {
			"url" : "https://images.freeimages.com/images/large-previews/2d8/mountains-1384887.jpg",
			"caption" : "Mountain",
			"withBorder" : true,
			"withBackground" : true,
			"stretched" : true
		  }
		}
    ]
}`

	editorJSON1 := support.ParseEditorJSON(input1)
	obj.Data = support.PrepareData(editorJSON1.Blocks[0])

	expected1 := `<figure class="image has-background-primary p-5" ><img class="  is-fullwidth" src="https://images.freeimages.com/images/large-previews/2d8/mountains-1384887.jpg" alt="Mountain" title="Mountain" /></figure>`

	obj.Image()

	actual1 := obj.Result[0]

	is.Equal(expected1, actual1) // Image 1 is different from expected

	obj.Result = []string{}

	input2 := `{
    "blocks": [
        {
		  "type" : "image",
		  "data" : {
			"file": {
			  "url" : "https://images.freeimages.com/images/large-previews/2d8/mountains-1384887.jpg"
			},
			"caption" : "Mountain",
			"withBorder" : false,
			"withBackground" : true,
			"stretched" : false
		  }
		}
    ]
}`

	editorJSON2 := support.ParseEditorJSON(input2)
	obj.Data = support.PrepareData(editorJSON2.Blocks[0])

	expected2 := `<figure class="image has-background-primary p-5" ><img class=" " src="https://images.freeimages.com/images/large-previews/2d8/mountains-1384887.jpg" alt="Mountain" title="Mountain" /></figure>`

	obj.Image()

	actual2 := obj.Result[0]

	is.Equal(expected2, actual2) // Image 2 is different from expected
}

func TestLinkToolBlock(t *testing.T) {
	is := is.New(t)

	obj.Result = []string{}

	input := `{
    "blocks": [
        {
		  "type" : "linkTool",
		  "data" : {
			"link" : "https://codex.so",
			"meta" : {
			  "title" : "CodeX Team",
			  "site_name" : "CodeX",
			  "description" : "Club of web-development, design and marketing. We build team learning how to build full-valued projects on the world market.",
			  "image" : {
				"url" : "https://pbs.twimg.com/profile_images/993612654861344768/wMPEM5XW_400x400.jpg"
			  }
			}
		  }
		}
    ]
}`

	editorJSON := support.ParseEditorJSON(input)
	obj.Data = support.PrepareData(editorJSON.Blocks[0])

	expected := `<a href="https://codex.so" target="_Blank" rel="nofollow noindex noreferrer" class="has-text-black-bis">
<div class="columns is-full p-5 has-background-light m-5">
<div class="column is-10">
<div class="column is-12 has-text-weight-bold">
CodeX Team
</div>
<div class="column is-12">
Club of web-development, design and marketing. We build team learning how to build full-valued projects on the world market.
</div>
<div class="column is-12 has-text-grey-light">
codex.so
</div>
</div>
<div class="column is-2">
<img class="image is-96x96" src="https://pbs.twimg.com/profile_images/993612654861344768/wMPEM5XW_400x400.jpg" />
</div>
</div>
</a>`

	obj.LinkTool()

	actual := obj.Result[0]

	is.Equal(expected, actual) // LinkTool is different from expected
}

func TestAttachesBlock(t *testing.T) {
	is := is.New(t)

	obj.Result = []string{}

	input := `{
    "blocks": [
        {
            "type": "attaches",
            "data": {
                "file": {
                    "extension": "jpg",
                    "name": "hero.jpg",
                    "size": 260096,
                    "url": "https://www.tesla.com/tesla_theme/assets/img/_vehicle_redesign/roadster_and_semi/roadster/hero.jpg"
                },
                "title": "Hero"
            }
        }
    ]
}`

	editorJSON := support.ParseEditorJSON(input)
	obj.Data = support.PrepareData(editorJSON.Blocks[0])

	expected := `<a href="https://www.tesla.com/tesla_theme/assets/img/_vehicle_redesign/roadster_and_semi/roadster/hero.jpg" rel="noopener noreferrer" target="_blank" class="has-text-black-bis">
<div class="columns is-full p-5 has-background-light m-5">
<div class="column is-2" >
<img class="image is-96x96" src="https://i.ibb.co/K7Myr2k/file-icon.png" />
</div>
<div class="column is-9">
<div class="column is-12 has-text-weight-bold">
hero.jpg
</div>
<div class="column is-12 has-text-grey-light">
254 KiB
</div>
</div>
<div class="column is-1" >
<img class="image is-48x48" src="https://i.ibb.co/VYyHr6C/download-icon.png" />
</div>
</div>
</a>`

	obj.Attaches()

	actual := obj.Result[0]

	is.Equal(expected, actual) // Attaches is different from expected
}

func TestEmbedBlock(t *testing.T) {
	is := is.New(t)

	obj.Result = []string{}

	input := `{
    "blocks": [
        {
            "type": "embed",
            "data": {
                "caption": "Lamborghini Aventador SVJ",
                "embed": "https://www.youtube.com/embed/viW44cUfxCE",
                "height": 315,
                "service": "Youtube",
                "source": "https://www.youtube.com/watch?v=viW44cUfxCE",
                "width": 560
            }
        }
    ]
}`

	editorJSON := support.ParseEditorJSON(input)
	obj.Data = support.PrepareData(editorJSON.Blocks[0])

	expected := `<div class="" style="max-width: 560px">
<div class="has-text-grey-dark p-2 has-text-weight-bold">Lamborghini Aventador SVJ</div>
<iframe width="560" height="315" src="https://www.youtube.com/embed/viW44cUfxCE" title="Lamborghini Aventador SVJ" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>
<div class="p-2 has-text-right is-italic">
<a class="has-text-danger-dark" href="https://www.youtube.com/watch?v=viW44cUfxCE" target="_Blank">Watch on Youtube</a>
</div>
</div>`

	obj.Embed()

	actual := obj.Result[0]

	is.Equal(expected, actual) // Embed is different from expected
}

func TestImageGalleryBlock(t *testing.T) {
	is := is.New(t)

	obj.Result = []string{}

	input := `{
    "blocks": [
        {
            "type": "imageGallery",
            "data": {
                "bkgMode": false,
                "editImages": true,
                "layoutDefault": true,
                "layoutHorizontal": false,
                "layoutSquare": false,
                "layoutWithFixedSize": false,
                "layoutWithGap": false,
                "urls": [
                    "https://www.nawpic.com/media/2020/ocean-nawpic-15.jpg",
                    "https://www.nawpic.com/media/2020/ocean-nawpic-18.jpg",
                    "https://wallpapercave.com/wp/6L4TVMP.jpg",
                    "https://wallpapercave.com/wp/wp9810772.jpg",
                    "https://wallpapercave.com/wp/wp9121482.jpg",
                    "https://wallpapercave.com/wp/wp9100484.jpg",
                    "https://cdn.wallpapersafari.com/94/22/4H3mFp.jpg"
                ]
            }
        }
    ]
}`

	editorJSON := support.ParseEditorJSON(input)
	obj.Data = support.PrepareData(editorJSON.Blocks[0])

	expected := `<div class="gg-container">
<div class="gg-box" id="">
<img src="https://www.nawpic.com/media/2020/ocean-nawpic-15.jpg" id="gg-image-0" />
<img src="https://www.nawpic.com/media/2020/ocean-nawpic-18.jpg" id="gg-image-1" />
<img src="https://wallpapercave.com/wp/6L4TVMP.jpg" id="gg-image-2" />
<img src="https://wallpapercave.com/wp/wp9810772.jpg" id="gg-image-3" />
<img src="https://wallpapercave.com/wp/wp9121482.jpg" id="gg-image-4" />
<img src="https://wallpapercave.com/wp/wp9100484.jpg" id="gg-image-5" />
<img src="https://cdn.wallpapersafari.com/94/22/4H3mFp.jpg" id="gg-image-6" />
</div>
</div>`

	obj.ImageGallery()

	actual := obj.Result[0]

	is.Equal(expected, actual) // ImageGallery is different from expected
}
