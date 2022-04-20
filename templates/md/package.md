{{ define "package" }}{{ if .ShowName }}{{ if .Name }}# {{ .Name }}
{{ end }}{{ end }}{{ if .ShowDoc }}{{ if .Doc }}
{{ .Doc }}
{{ end }}{{ end }}{{ if .ShowExamples }}{{ if .Examples }}
{{ range .Examples }}
{{ template "example" . }}
{{ end }}
{{- end }}{{ end }}{{ if .ShowIndex }}
{{ template "index" . }}
{{ end }}{{ if .ShowFuncs }}{{ if .Funcs }}
{{ range .Funcs }}{{ template "function" . }}{{ end }}
{{- end }}{{ end }}{{ if .ShowTypes }}{{ if .Types }}
{{ range .Types }}{{ template "type" . }}{{ end }}
{{- end }}{{ end }}{{- end }}