{{ define "mod.md" }}# Module {{ .Name }}

go {{ .GoVersion }}

## Submodules
{{ template "subpackages" . }}
{{ end }}