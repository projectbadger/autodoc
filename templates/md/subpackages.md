{{ define "subpackages" }}{{ $importPath:=.ImportPath }}{{ if .Subpackages }}{{ range $path, $package := .Subpackages }}
- [{{ $package.Name }}]({{ headingGitLink $importPath $path}})
{{- end}}
{{- end}}
{{- end }}