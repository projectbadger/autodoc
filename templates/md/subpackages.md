{{ define "subpackages" }}{{ $importPath:=.ImportPath }}{{ if .Subpackages }}{{ range $path, $package := .Subpackages }}
{{ call .PathIndent $path }}- [{{ $package.Name }}]({{ $path }}){{ end }}
{{- end}}{{- end }}