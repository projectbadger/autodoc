{{ define "index" }}{{ if .ShowFuncs }}{{ if .Funcs }}{{ range .Funcs }}
- [{{ template "functionDefinition" . }}](#{{ .GetHeadingHREF}})
{{- end }}
{{ end }}{{- end }}{{ if .ShowTypes }}{{ if .Types }}{{ range .Types }}
- [type {{ .Name }}](#{{ .GetHeadingHREF }}){{ if .Funcs }}{{ range .Funcs }}
  - [{{ template "functionDefinition" . }}](#{{ .GetHeadingHREF}})
{{- end }}{{ end }}{{ if .Methods }}{{ range .Methods }}
  - [{{ template "functionDefinition" . }}](#{{ .GetHeadingHREF}})
{{- end }}{{ end }}{{ end }}{{ end }}
{{- end }}{{ if .ShowVars }}{{ if .Vars }}
- [Variables](#variables)
{{- end }}
{{- end }}{{- end }}