{{ define "subpackages" }}{{ $importPath:=.ImportPath }}{{ if .Subpackages }}{{ range $path, $package := .Subpackages }}
- [{{ $package.Name }}]({{ trimSlashes $path }})
{{- end}}
{{- end}}{{- end }}