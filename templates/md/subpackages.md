{{ define "subpackages" }}{{ $importPath:=.ImportPath }}{{ if .Subpackages }}{{ range $path, $package := .Subpackages }}
- [{{ $package.Name }}]({{ $path }})
{{- end}}
{{- end}}{{- end }}