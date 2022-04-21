package md

import (
	"bytes"
	"embed"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/projectbadger/autodoc/config"
	"github.com/projectbadger/autodoc/templates/functions"
)

var (
	//go:embed mod.md
	templateMod []byte
	//go:embed subpackages.md
	templateSubpackages []byte
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
	//go:embed overview.md
	templateOverview []byte
	//go:embed vars.md constants.md example.md function.md type.md index.md package.md doc.md
	TemplatesFS   embed.FS
	TemplateDoc   *template.Template
	TemplateNames = []string{
		"doc.md", "example.md", "function.md", "index.md",
		"type.md", "package.md", "vars.md", "constants.md",
	}
)

// SetupTemplates sets up the TemplateDoc *template.Template
// variable from the template strings.
func SetupTemplates() error {
	var err error
	// TemplateDoc, _ = template.New("doc.md").Funcs(templates.GetTemplateFuncMap()).
	// TemplateDoc, err = template.New("doc.md").Funcs(templates.GetTemplateFuncMap()).ParseFS(TemplatesFS, TemplateNames...)
	TemplateDoc = template.New("doc.md").
		Funcs(functions.GetTemplateFuncMap())
	TemplateDoc = template.Must(TemplateDoc.Parse(string(templateOverview)))
	TemplateDoc = template.Must(TemplateDoc.Parse(string(templateConstants)))
	TemplateDoc = template.Must(TemplateDoc.Parse(string(templateVars)))
	TemplateDoc = template.Must(TemplateDoc.Parse(string(templateExample)))
	TemplateDoc = template.Must(TemplateDoc.Parse(string(templateFunction)))
	TemplateDoc = template.Must(TemplateDoc.Parse(string(templateType)))
	TemplateDoc = template.Must(TemplateDoc.Parse(string(templateIndex)))
	TemplateDoc = template.Must(TemplateDoc.Parse(string(templatePackage)))
	TemplateDoc = template.Must(TemplateDoc.Parse(string(templateSubpackages)))
	TemplateDoc = template.Must(TemplateDoc.Parse(string(templateDoc)))
	TemplateDoc = template.Must(TemplateDoc.Parse(string(templateMod)))

	return err
}

func init() {
	err := SetupTemplates()
	if err != nil {
		fmt.Println(err)
	}
}

// ReplaceTemplates replaces the currently set templates
// with the ones set up by the config.
func ReplaceTemplates() (err error) {
	if config.Cfg.Templates.TemplatesDir == "" {
		return nil
	}
	path, err := filepath.Abs(config.Cfg.Templates.TemplatesDir)
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
	files, _ := os.ReadDir(dirStat.Name())
	for _, fName := range files {
		switch fName.Name() {
		case "vars.md":
			templateVars, err = os.ReadFile(filepath.Join(path, fName.Name()))
			if err != nil {
				fmt.Println("Error:", err)
				return err
			}
		case "constants.md":
			templateConstants, err = os.ReadFile(filepath.Join(path, fName.Name()))
			if err != nil {
				fmt.Println("Error:", err)
				return err
			}
		case "example.md":
			templateExample, err = os.ReadFile(filepath.Join(path, fName.Name()))
			if err != nil {
				fmt.Println("Error:", err)
				return err
			}
		case "function.md":
			templateFunction, err = os.ReadFile(filepath.Join(path, fName.Name()))
			if err != nil {
				fmt.Println("Error:", err)
				return err
			}
		case "type.md":
			templateType, err = os.ReadFile(filepath.Join(path, fName.Name()))
			if err != nil {
				fmt.Println("Error:", err)
				return err
			}
		case "index.md":
			templateIndex, err = os.ReadFile(filepath.Join(path, fName.Name()))
			if err != nil {
				fmt.Println("Error:", err)
				return err
			}
		case "subpackages.md":
			templateSubpackages, err = os.ReadFile(filepath.Join(path, fName.Name()))
			if err != nil {
				fmt.Println("Error:", err)
				return err
			}
		case "overview.md":
			templateOverview, err = os.ReadFile(filepath.Join(path, fName.Name()))
			if err != nil {
				fmt.Println("Error:", err)
				return err
			}
		case "mod.md":
			templateMod, err = os.ReadFile(filepath.Join(path, fName.Name()))
			if err != nil {
				fmt.Println("Error:", err)
				return err
			}
		case "package.md":
			templatePackage, err = os.ReadFile(filepath.Join(path, fName.Name()))
			if err != nil {
				fmt.Println("Error:", err)
				return err
			}
		case "doc.md":
			templateDoc, err = os.ReadFile(filepath.Join(path, fName.Name()))
			if err != nil {
				fmt.Println("Error:", err)
				return err
			}
		}
	}
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

func ExecuteTemplate(name string, data interface{}) (string, error) {
	if name == "" {
		name = "doc.md"
	}
	buf := bytes.NewBuffer(nil)
	err := TemplateDoc.ExecuteTemplate(buf, name, data)
	return buf.String(), err
}
