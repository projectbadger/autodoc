package main

import (
	"bytes"
	"fmt"

	"github.com/projectbadger/autodoc/config"
	"github.com/projectbadger/autodoc/doc"
	"github.com/projectbadger/autodoc/templates"
	"github.com/projectbadger/autodoc/templates/md"
)

func main() {
	err := templates.SetupTemplates()
	if err != nil {
		panic(err)
	}
	// if len(os.Args) < 3 {
	// 	fmt.Println("Usage: go run main.go [ name | doc ] package_path")
	// 	os.Exit(1)
	// }

	docs, err := doc.GetPackageDocumentation(config.Cfg.PackageDir, config.Cfg.ImportPath)
	if err != nil {
		panic(err)
	}
	data := doc.ParsePackage(docs)
	buf := bytes.NewBuffer(nil)
	md.TemplateDoc.ExecuteTemplate(buf, "doc.md", data)
	fmt.Println(buf.String())
}
