package doc

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

// Module holds module info.
type Module struct {
	Name       string              `json:"name" yaml:"name"`
	GoVersion  string              `json:"go_version" yaml:"go_version"`
	Heading    string              `json:"heading" yaml:"heading"`
	Overview   string              `json:"overview" yaml:"overview"`
	ImportPath string              `json:"import_path" yaml:"import_path"`
	Submodules map[string]*Package `json:"submodules" yaml:"submodules"`
}

// ParseModule parses a go module in the provided path and
// returns a *Module.
func ParseModule(path string) (*Module, error) {
	packages, err := GetPackagesDataFromDirRecursive(path, true, "")
	if err != nil {
		return nil, err
	}
	module := &Module{}
	err = ParseGoModFile(module, path)
	module.Submodules = packages
	for _, v := range module.Submodules {
		v.ImportPath = module.ImportPath
	}
	return module, err
}

// ParseGoModFile parses a go.mod file.
func ParseGoModFile(module *Module, path string) error {
	file, err := os.Open(filepath.Join(path, "go.mod"))
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "module ") {
			module.Name = strings.Trim(strings.TrimPrefix(line, "module "), " \n\r\t")
		}
		if strings.HasPrefix(line, "go ") {
			module.GoVersion = strings.Trim(strings.TrimPrefix(line, "go "), " \n\r\t")
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
