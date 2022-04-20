package md

import (
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/projectbadger/autodoc/config"
	"github.com/projectbadger/autodoc/templates/functions"
)

var (
	//go:embed doc.md
	templateDoc []byte
	//go:embed package.md
	templatePackage []byte
	//go:embed index.md
	templateIndex []byte
	//go:embed type.md
	templateType []byte
	//go:embed function.md
	templateFunction []byte
	//go:embed example.md
	templateExample []byte
	//go:embed vars.md
	templateVars []byte
	//go:embed constants.md
	templateConstants []byte
	//go:embed vars.md constants.md example.md function.md type.md index.md package.md doc.md
	TemplatesFS   embed.FS
	TemplateDoc   *template.Template
	TemplateNames = []string{
		"doc.md", "example.md", "function.md", "index.md",
		"type.md", "package.md", "vars.md", "constants.md",
	}
)

func SetupTemplates() error {
	var err error
	// TemplateDoc, _ = template.New("doc.md").Funcs(templates.GetTemplateFuncMap()).
	// TemplateDoc, err = template.New("doc.md").Funcs(templates.GetTemplateFuncMap()).ParseFS(TemplatesFS, TemplateNames...)
	TemplateDoc = template.New("doc.md").
		Funcs(functions.GetTemplateFuncMap())
	TemplateDoc = template.Must(TemplateDoc.Parse(string(templateConstants)))
	TemplateDoc = template.Must(TemplateDoc.Parse(string(templateVars)))
	TemplateDoc = template.Must(TemplateDoc.Parse(string(templateExample)))
	TemplateDoc = template.Must(TemplateDoc.Parse(string(templateFunction)))
	TemplateDoc = template.Must(TemplateDoc.Parse(string(templateType)))
	TemplateDoc = template.Must(TemplateDoc.Parse(string(templateIndex)))
	TemplateDoc = template.Must(TemplateDoc.Parse(string(templatePackage)))
	TemplateDoc = template.Must(TemplateDoc.Parse(string(templateDoc)))

	return err
}

func init() {
	err := SetupTemplates()
	if err != nil {
		fmt.Println(err)
	}
}

func ReplaceTemplates() (err error) {
	if config.Cfg.TemplatesDir == "" {
		return nil
	}
	path, err := filepath.Abs(config.Cfg.TemplatesDir)
	if err != nil {
		return err
	}
	dirStat, err := os.Stat(path)
	if err != nil {
		return err
	}
	if !dirStat.IsDir() {
		return errors.New("not dir")
	}

	filepath.WalkDir(dirStat.Name(), func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			fmt.Println("error (start):", e)
			return e
		}
		fileName := filepath.Base(d.Name())
		if d.IsDir() || !strings.HasSuffix(fileName, ".md") {
			return nil
		}
		// fmt.Println("Processing", d.Name())
		switch fileName {
		case "vars.md":
			templateVars, err = os.ReadFile(filepath.Join(path, filepath.Base(d.Name())))
			if err != nil {
				fmt.Println("Error:", err)
				return err
			}
		case "constants.md":
			templateConstants, err = os.ReadFile(filepath.Join(path, filepath.Base(d.Name())))
			if err != nil {
				fmt.Println("Error:", err)
				return err
			}
		case "example.md":
			templateExample, err = os.ReadFile(filepath.Join(path, filepath.Base(d.Name())))
			if err != nil {
				fmt.Println("Error:", err)
				return err
			}
		case "function.md":
			templateFunction, err = os.ReadFile(filepath.Join(path, filepath.Base(d.Name())))
			if err != nil {
				fmt.Println("Error:", err)
				return err
			}
		case "type.md":
			templateType, err = os.ReadFile(filepath.Join(path, filepath.Base(d.Name())))
			if err != nil {
				fmt.Println("Error:", err)
				return err
			}
		case "packindexage.md":
			templateIndex, err = os.ReadFile(filepath.Join(path, filepath.Base(d.Name())))
			if err != nil {
				fmt.Println("Error:", err)
				return err
			}
		case "package.md":
			templatePackage, err = os.ReadFile(filepath.Join(path, filepath.Base(d.Name())))
			if err != nil {
				fmt.Println("Error:", err)
				return err
			}
		case "doc.md":
			templateDoc, err = os.ReadFile(filepath.Join(path, filepath.Base(d.Name())))
			if err != nil {
				fmt.Println("Error:", err)
				return err
			}
		}
		// fmt.Println("Replacing template with", filepath.Abs(fileName))
		return nil
	})
	return SetupTemplates()
}

func GetTemplatesBytes() (t [][]byte) {
	dirs, err := TemplatesFS.ReadDir(".")
	if err != nil {
		return nil
	}
	for _, name := range TemplateNames {
		f, err := TemplatesFS.ReadFile(name)
		if err != nil {
			return nil
		}
		t = append(t, f)
	}
	for _, file := range dirs {
		if file.IsDir() {
			continue
		}
		t = append(t, []byte(filepath.Base(file.Name())))
	}
	return
}

// SaveToFile saves the config to a file in YAML format
func OutputTemplatesToDir(path string) error {
	if path == "" {
		path = "."
	}
	pathStat, err := os.Stat(path)
	if err != nil {
		return err
	}
	if !pathStat.IsDir() {
		return errors.New("path is not dir")
	}
	files := GetTemplatesBytes()
	for i, templateName := range TemplateNames {
		err := os.WriteFile(filepath.Join(path, templateName), files[i], 0664)
		if err != nil {
			return err
		}
	}
	return nil
	// return os.WriteFile(path, configYaml, 0644)
}
