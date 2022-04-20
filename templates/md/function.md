{{ define "function" }}## func [{{ .Name }}]({{ if .Filename }}<{{ .Filename }}{{ if .Line }}#L{{ .Line }}{{ end }}>)
{{ end }}{{ if .Doc }}
{{ .Doc }}
{{ end }}
```go
{{ .Definition }}
```
{{ if .Examples }}
{{ range .Examples }}{{ template "example" . }}{{ end }}
{{- end }}{{- end }}