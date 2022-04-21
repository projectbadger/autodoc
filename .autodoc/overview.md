{{ define "overview" }}# {{ index .CustomVars "name" }}

Autodoc is a golang documentation tool. It  generates package documentation from a provided path.

This README.md has been generated using `.autodoc/overview.md` template override and the command:
```{{ if index .CustomVars "created-command" }}
{{ index .CustomVars "created-command" }}{{ else }}
go run main.go -config .autodoc/config.root.yml > README.md
{{ end }}
```
{{ if index .CustomVars "is-root" }}
For the subpackages, the README.md files have been generated using the command
```
PACKAGE_DIR=config go run main.go -package $PACKAGE_DIR > $PACKAGE_DIR/README.md
```
{{ end }}
Autodoc works with templates in the `templates/md` directory, combining them to create documentation. Any and all templates can be overriden by providing the flag `-templates /path/to/templates/directory`, where the templates contained in the directory have the same name as the original template.

To output all the templates into a directory for inspection, use the flag `-output-templates /path/to/new/templates/directory`

By default, all data is included in the generated documentation. To render a specific subset of data, use the flag `-included-data` and list the data to be included.

All the included data:
```
autodoc -package . -included-data name,overview,doc,examples,subpackages,index,variables,constants,functions,types
```
or
```
autodoc -package . -included-data all
```
You may also remove data from all included data:
```
autodoc -package . -included-data all,-types,-vars,-constants,
```
{{ end }}