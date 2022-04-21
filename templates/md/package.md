{{ define "package" }}{{ if .ShowName }}
{{ if .CustomVars }}{{ if index .CustomVars "name" }}{{ index .CustomVars "name" }}{{ else }}{{ if .Name }}# {{ .Name }}
{{ end }}{{ end }}{{ else }}{{ if .Name }}# {{ .Name }}
{{ end }}{{ end }}{{ end }}{{ if .ShowImportPath }}{{ if .ImportPath }}
```go
import {{ .ImportPath }}
```
{{ end }}{{ end }}{{ if .ShowOverview }}
{{ template "overview" . }}
{{ end }}{{ if .ShowDoc }}{{ if .Doc }}
{{ .Doc }}
{{ end }}{{ end }}{{ if .ShowExamples }}{{ if .Examples }}
{{ range .Examples }}
{{ template "example" . }}
{{ end }}
{{- end }}{{ end }}{{ if .ShowSubpackages }}{{ if .Subpackages }}
## Subpackages
{{ template "subpackages" . }}{{ end }}
{{ end }}{{ if .ShowIndex }}
## Index
{{ template "index" . }}
{{ end }}{{ if .ShowFuncs }}{{ if .Funcs }}
{{ range .Funcs }}{{ template "function" . }}{{ end }}
{{- end }}{{ end }}{{ if .ShowTypes }}{{ if .Types }}
{{ range .Types }}{{ template "type" . }}{{ end }}
{{- end }}{{ end }}{{- end }}