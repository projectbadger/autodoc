{{ define "package" }}{{ if .Name }}# {{ .Name }}
{{ end }}{{ if .Doc }}
{{ .Doc }}
{{ end }}{{ if .Examples }}
{{ range .Examples }}{{ template "example" . }}{{ end }}
{{- end }}
{{ template "index" . }}
{{ if .Funcs }}
{{ range .Funcs }}{{ template "function" . }}{{ end }}
{{- end }}{{ if .Types }}
{{ range .Types }}{{ template "type" . }}{{ end }}
{{- end }}{{- end }}