{{ define "subpackages" }}{{ $importPath:=.ImportPath }}{{ if .Subpackages }}{{ range $path, $package := .Subpackages }}
- [{{ $package.Name }}]({{ headingGitLink $importPath $path}})
{{- end}}
{{- end}}

{{ $importPath:=.ImportPath }}{{ if .Subpackages }}{{ range $path, $package := .Subpackages }}
- [{{ $package.Name }}](blob/main/{{ $path }})
{{- end}}
{{- end}}
{{- end }}