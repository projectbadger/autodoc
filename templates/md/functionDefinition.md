{{ define "functionDefinition" }}{{ .Name }}({{ if .Params }}{{ .FormatParams }}{{ end }}){{ if len .Results | eq 0 | not }} {{ .FormatResultsBrackets }}{{ end }}{{ end }}
