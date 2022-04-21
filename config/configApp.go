package config

import (
	"fmt"

	// "github.com/projectbadger/autodoc/filters"
	"os"

	"gopkg.in/yaml.v2"
)

// ConfigApp holds all the application configuration data.
type ConfigApp struct {
	File         string `default:"" json:"-" yaml:"-" cli:"config Config file path\n     "`
	CreateConfig string `default:"" yaml:"-" json:"-" cli:"create-config Create a named default config file with cli parameters and environment variables.\n   "`
	Output       string `json:"output" yaml:"output" cli:"output Output the generated documentation to file\n    "`
	// DebugMode    bool   `default:"false" yaml:"-" json:"-" cli:"debug Run the app in debug mode\n   "`
	// LoginUrl     string                `default:"" json:"login_url" yaml:"login_url" cli:"login-url Login URL to appear at the top of the page\n      "`
	// LogoutUrl    string                `default:"" json:"logout_url" yaml:"logout_url" cli:"logout-url Logout URL to appear at the top of the page\n      "`
	// Filters     filters.RequestFilter `json:"filters" yaml:"filters"`
	PackageDir   string          `default:"." json:"package_dir" yaml:"package_dir" cli:"package Package directory filepath.\nThe contents of this directory will be parsed as a Go package\n      "`
	ModuleDir    string          `default:"." json:"module_dir" yaml:"module_dir" cli:"module Module directory filepath.\nThe contents of this directory will be parsed as a Go module\n      "`
	ExcludeDirs  []string        `default:"" json:"exclude_dirs" yaml:"exclude_dirs" cli:"exclude Exclude directories from the package search.\nDefault:\n  node_nodules;.git\n      "`
	Version      bool            `default:"false" json:"-" yaml:"-" cli:"version Print version"`
	VersionShort bool            `default:"false" json:"-" yaml:"-" cli:"v Print version"`
	Templates    ConfigTemplates `yaml:"templates" json:"templates"`
}

func (c *ConfigApp) IsVersion() bool {
	return c.Version || c.VersionShort
}

func (c *ConfigApp) CheckIsVersion() {
	if c.Version || c.VersionShort {
		fmt.Printf("Package:    %s\n", PackageName)
		fmt.Printf("Version:    %s\n", Version)
		fmt.Printf("Build Time: %s\n", BuildTime)
		os.Exit(0)
	}
}

// SaveToFile saves the config to a file in YAML format
func (cfg *ConfigApp) SaveToFile(path string) error {
	if path == "" {
		return nil
	}
	configYaml, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	return os.WriteFile(path, configYaml, 0644)
}

// SetupDefault sets up default config data.
func (c *ConfigApp) SetupDefault() {
	c.PackageDir = "."
	// c.ImportPath = "git.example.com/project/repository"
	c.Templates.SetupDefault()
}
