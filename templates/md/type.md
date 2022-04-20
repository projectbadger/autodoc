{{ define "type" }}{{ if .Name }}
## type [{{ .Name }}]({{ if .Filename }}<{{ .Filename }}{{ if .Line }}#L{{ .Line }}{{ end }}>{{ end }})
{{- end }}{{ if .Doc }}
{{ .Doc }}
{{- end }}
```go
{{ .Definition }}
```
{{ if .Examples }}
{{ range .Examples }}{{ template "example" . }}{{ end }}
{{- end }}{{ if .Funcs }}
{{ range .Funcs }}{{ template "function" . }}{{ end }}
{{- end }}{{ if .Methods }}
{{ range .Methods }}{{ template "function" . }}{{ end }}
{{- end }}{{- end }}