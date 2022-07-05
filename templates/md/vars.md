{{ define "vars" }}## Variables
```go
{{ range .Vars }}{{ .Definition }}{{ end }}
```
{{ end }}