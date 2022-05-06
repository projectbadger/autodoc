
# md

```go
import github.com/projectbadger/autodoc/templates/md
```

## Index

- [func ExecuteTemplate(name string, data interface{}) (string, error)](#func-executetemplate-string-data-interface--error)
- [func GetTemplatesBytes() (t [][]byte)](#func-gettemplatesbytes--byte)
- [func OutputTemplatesToDir(path string) error](#func-outputtemplatestodir-string-error)
- [func ReplaceTemplates() (err error)](#func-replacetemplates--error)
- [func SetupTemplates() error](#func-setuptemplates-error)

## func [ExecuteTemplate](<md.go#L222>)

```go
func ExecuteTemplate(name string, data interface{}) (string, error)
```
## func [GetTemplatesBytes](<md.go#L178>)

```go
func GetTemplatesBytes() (t [][]byte)
```
## func [OutputTemplatesToDir](<md.go#L200>)

SaveToFile saves the config to a file in YAML format

```go
func OutputTemplatesToDir(path string) error
```
## func [ReplaceTemplates](<md.go#L78>)

```go
func ReplaceTemplates() (err error)
```
## func [SetupTemplates](<md.go#L50>)

```go
func SetupTemplates() error
```

