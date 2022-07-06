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

// OutputTemplatesToDir creates template files in the
// specified directory
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
		err := os.WriteFile(filepath.Join(path, templateName.Name()), files[i], 0664)
		if err != nil {
			return err
		}
	}
	return nil
}

// SetupTemplates sets the templates and template overrides
// according to the config.
func SetupTemplates() error {
	if config.Cfg.Templates.OutputTemplates != "" {
		fmt.Printf("Outputting templates to '%s'\n", config.Cfg.Templates.OutputTemplates)
		err := OutputTemplatesToDir(config.Cfg.Templates.OutputTemplates)
		if err != nil {
			fmt.Printf("Error saving to file: '%q'", err)
			os.Exit(1)
		}
		os.Exit(0)
	}
	if config.Cfg.Templates.TemplatesDir != "" {
		err := md.ReplaceTemplates()
		if err != nil {
			fmt.Printf("Error saving to file: '%q'", err)
			os.Exit(1)
		}
	}

	return nil
}
