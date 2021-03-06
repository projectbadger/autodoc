
# md

```go
import github.com/projectbadger/autodoc/templates/md
```

## Index

- [ExecuteTemplate(string) (string, error)](#func-executetemplatestring-string-error)
- [GetTemplatesBytes()](#func-gettemplatesbytes)
- [OutputTemplatesToDir(string) error](#func-outputtemplatestodirstring-error)
- [ReplaceTemplates() error](#func-replacetemplates-error)
- [SetupTemplates() error](#func-setuptemplates-error)

- [type EmbeddedTemplates](#type-embeddedtemplates)
- [Variables](#variables)

## Variables
```go
var (
	Templates	= []embeddedTemplate{
		{
			Name:	"mod.md",

			Content:	[]byte{},
		},
	}

	//go:embed vars.md constants.md example.md function.md type.md index.md package.md doc.md
	TemplatesFS	embed.FS
	TemplateDoc	*template.Template
	TemplateNames	= []templateName{
		"vars.md",
		"constants.md",
		"example.md",
		"functionDefinition.md",
		"functionHeading.md",
		"function.md",
		"typeHeading",
		"type.md",
		"index.md",
		"subpackages.md",
		"overview.md",
		"package.md",
		"doc.md",
		"mod.md",
	}
)

```

## func [ExecuteTemplate(string) (string, error)](<md.go#L198>)

```go
func ExecuteTemplate(name string, data interface{}) (string, error)
```
## func [GetTemplatesBytes()](<md.go#L154>)

```go
func GetTemplatesBytes() (t [][]byte)
```
## func [OutputTemplatesToDir(string) error](<md.go#L176>)

SaveToFile saves the config to a file in YAML format


```go
func OutputTemplatesToDir(path string) error
```
## func [ReplaceTemplates() error](<md.go#L123>)

ReplaceTemplates replaces the currently set templates
with the ones set up by the config.


```go
func ReplaceTemplates() (err error)
```
## func [SetupTemplates() error](<md.go#L93>)

SetupTemplates sets up the TemplateDoc *template.Template
variable from the template strings.


```go
func SetupTemplates() error
```


## type [EmbeddedTemplates](<md.go#L29>)
```go
type EmbeddedTemplates struct {
	Mod		embeddedTemplate
	Subpackages	embeddedTemplate
}
```

