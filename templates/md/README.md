
# md

```go
import github.com/projectbadger/autodoc/templates/md
```

## Index

- func [ExecuteTemplate](#func-executetemplate)
- func [GetTemplatesBytes](#func-gettemplatesbytes)
- func [OutputTemplatesToDir](#func-outputtemplatestodir)
- func [ReplaceTemplates](#func-replacetemplates)
- func [SetupTemplates](#func-setuptemplates)

- [Variables](#variables)

## Variables
```go
var (

	//go:embed vars.md constants.md example.md function.md type.md index.md package.md doc.md
	TemplatesFS	embed.FS
	TemplateDoc	*template.Template
	TemplateNames	= []string{
		"doc.md", "example.md", "function.md", "index.md",
		"type.md", "package.md", "vars.md", "constants.md",
	}
)
```

## func [ExecuteTemplate](<md.go#L213>)

```go
func ExecuteTemplate(name string, data interface{}) (string, error)
```
## func [GetTemplatesBytes](<md.go#L169>)

```go
func GetTemplatesBytes() (t [][]byte)
```
## func [OutputTemplatesToDir](<md.go#L191>)

SaveToFile saves the config to a file in YAML format

```go
func OutputTemplatesToDir(path string) error
```
## func [ReplaceTemplates](<md.go#L80>)

ReplaceTemplates replaces the currently set templates
with the ones set up by the config.

```go
func ReplaceTemplates() (err error)
```
## func [SetupTemplates](<md.go#L50>)

SetupTemplates sets up the TemplateDoc *template.Template
variable from the template strings.

```go
func SetupTemplates() error
```

