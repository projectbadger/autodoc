package md

import (
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"text/template"

	"github.com/projectbadger/autodoc/config"
	"github.com/projectbadger/autodoc/templates"
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

func init() {
	var err error
	// TemplateDoc, _ = template.New("doc.md").Funcs(templates.GetTemplateFuncMap()).
	TemplateDoc, err = template.New("doc.md").Funcs(templates.GetTemplateFuncMap()).ParseFS(TemplatesFS, TemplateNames...)
	if err != nil {
		fmt.Println("error:", err)
	}
}

func ReplaceTemplates() (err error) {
	dirStat, err := os.Stat(config.Cfg.TemplatesDir)
	if err != nil {
		return err
	}
	if !dirStat.IsDir() {
		return errors.New("not dir")
	}
	var filepaths []string

	filepath.WalkDir(dirStat.Name(), func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			fmt.Println("error (start):", e)
			return e
		}
		for _, templateFile := range TemplateNames {
			if templateFile == filepath.Base(d.Name()) {
				filepaths = append(filepaths, d.Name())
			}
		}
		return nil
	})
	TemplateDoc, err = TemplateDoc.ParseFiles(filepaths...)
	return
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
