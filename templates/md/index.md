{{ define "index" }}{{ if .ShowFuncs }}{{ if .Funcs }}{{ range .Funcs }}
- [{{ .Definition }}](#{{ headingLink .Definition}})
{{- end }}
{{ end }}{{- end }}{{ if .ShowTypes }}{{ if .Types }}{{ range .Types }}
- [type {{ .Name }}](#{{ headingLink .Name}}){{ if .Funcs }}{{ range .Funcs }}
  - [{{ .Definition }}](#{{ headingLink .Definition}})
{{- end }}{{ end }}{{ if .Methods }}{{ range .Methods }}
  - [{{ .Definition }}](#{{ headingLink .Definition}})
{{- end }}{{ end }}
{{- end }}{{- end }}
{{- end }}{{ if .ShowConsts }}{{ if .Constants }}
- [Constants](#constants)
{{- end }}{{- end }}{{ if .ShowVars }}{{ if .Vars }}
- [Variables](#variables)
{{- end }}
{{- end }}{{- end }}