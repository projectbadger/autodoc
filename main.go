package main

import (
	"bytes"
	"fmt"

	"os"

	"github.com/projectbadger/autodoc/doc"
	"github.com/projectbadger/autodoc/templates/md"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go [ name | doc ] package_path")
		os.Exit(1)
	}

	docs, err := doc.GetPackageDocumentation(os.Args[2], "github.com/owner/package")
	if err != nil {
		panic(err)
	}

	switch os.Args[1] {
	case "name":
		fmt.Println(docs.Name)
	case "doc":
		data := doc.ParsePackage(docs)
		buf := bytes.NewBuffer(nil)
		// fmt.Println("data:", data, "templates:", md.TemplateDoc.DefinedTemplates())
		md.TemplateDoc.ExecuteTemplate(buf, "package", data)
		fmt.Println(buf.String())
	default:
		os.Exit(1)
	}
}
