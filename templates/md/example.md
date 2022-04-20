{{ define "example" }}<details><summary>Example</summary>
<p>{{ if .Name }}
[{{ .Name }}]({{ if .Filename }}<{{ .Filename }}{{ if .Line }}#L{{ .Line }}{{ end }}>{{ end }})
{{- end }}{{ if .Doc }}
{{ .Doc }}
{{- end }}

```go
{{ .Definition }}
```
</p></details>{{ end }}