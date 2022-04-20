// Documentation templates
package templates

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/projectbadger/autodoc/config"
	"github.com/projectbadger/autodoc/templates/md"
)

// SaveToFile saves the config to a file in YAML format
func OutputTemplatesToDir(path string) error {
	if path == "" {
		path = "."
	}
	pathStat, err := os.Stat(path)
	if err != nil {
		parent := filepath.Dir(path)
		parentStat, err := os.Stat(parent)
		if err != nil {
			return err
		}
		if !parentStat.IsDir() {
			return errors.New("invalid path")
		}
		err = os.Mkdir(path, 0755)
		if err != nil {
			return err
		}
		pathStat, _ = os.Stat(path)
	}
	if !pathStat.IsDir() {
		return errors.New("path is not dir")
	}
	files := md.GetTemplatesBytes()
	for i, templateName := range md.TemplateNames {
		err := os.WriteFile(filepath.Join(path, templateName), files[i], 0664)
		if err != nil {
			return err
		}
	}
	return nil
	// return os.WriteFile(path, configYaml, 0644)
}

func ReplaceTemplates() error {
	return md.ReplaceTemplates()
}

func SetupTemplates() error {
	if config.Cfg.OutputTemplates != "" {
		fmt.Printf("Outputting templates to '%s'\n", config.Cfg.OutputTemplates)
		err := OutputTemplatesToDir(config.Cfg.OutputTemplates)
		if err != nil {
			fmt.Printf("Error saving to file: '%q'", err)
			os.Exit(1)
		}
		os.Exit(0)
	}
	if config.Cfg.TemplatesDir != "" {
		// fmt.Printf("Parsing templates from '%s'\n", config.Cfg.TemplatesDir)
		err := ReplaceTemplates()
		if err != nil {
			fmt.Printf("Error saving to file: '%q'", err)
			os.Exit(1)
		}
	}
	// fmt.Printf("Set up templates\n")

	return nil
}
