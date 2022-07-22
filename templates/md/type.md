{{ define "type" }}{{ if .Name }}
{{ template "typeHeading" . }}
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