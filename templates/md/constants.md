{{ define "constants" }}## Constants
```go
{{ range .Vars }}{{ .Definition }}{{- end }}
```
{{ end }}