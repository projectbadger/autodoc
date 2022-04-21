// Autodoc main package
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/projectbadger/autodoc/config"
	"github.com/projectbadger/autodoc/doc"
	"github.com/projectbadger/autodoc/templates"
	"github.com/projectbadger/autodoc/templates/md"
)

func main() {
	if len(os.Args) > 2 {
		// Extract module path
		log.Println("args:", os.Args)

		os.Args = append(os.Args[:1], os.Args[2:]...)
		log.Println("args:", os.Args)
	}
	err := templates.SetupTemplates()
	if err != nil {
		log.Fatalln((err))
	}
	if config.Cfg.ModuleDir != "" {
		module, err := doc.ParseModule(config.Cfg.PackageDir)
		if err != nil {
			log.Fatalln((err))
		}
		str, err := md.ExecuteTemplate("mod.md", module)
		if err != nil {
			log.Fatalln((err))
		}
		Output(str)
	} else if config.Cfg.PackageDir != "" {
		data, err := doc.GetPackageDataFromDirRecursive(config.Cfg.PackageDir)
		if err != nil {
			log.Fatalln((err))
		}
		str, err := md.ExecuteTemplate("doc.md", data)
		if err != nil {
			log.Fatalln((err))
		}
		Output(str)
	}
}

func Output(str string) {
	if config.Cfg.Output != "" {
		err := os.WriteFile(config.Cfg.Output, []byte(config.Cfg.Output), 0664)
		if err != nil {
			log.Fatalln(err)
		}
	}
	fmt.Printf("%s", str)
}
