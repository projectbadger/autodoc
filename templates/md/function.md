{{ define "function" }}{{ template "functionHeading" . }}{{ if .Doc }}
{{ .Doc }}
{{ end }}
```go
{{ .Definition }}
```
{{ if .Examples }}
{{ range .Examples }}{{ template "example" . }}{{ end }}
{{- end }}
{{- end }}