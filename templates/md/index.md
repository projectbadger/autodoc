{{ define "index" }}{{ if .Funcs }}{{ range .Funcs }}
- [{{ .Definition }}](#{{ headingLink .Definition}})
{{- end}}
{{- end }}{{ if .Types }}{{ range .Types }}
- [type {{ .Name }}](#{{ headingLink .Name}}){{ if .Funcs }}{{ range .Funcs }}
  - [{{ .Definition }}](#{{ headingLink .Definition}})
{{- end }}{{ end }}{{ if .Methods }}{{ range .Methods }}
  - [{{ .Definition }}](#{{ headingLink .Definition}})
{{- end }}{{ end }}
{{- end }}
{{- end }}
{{- end }}