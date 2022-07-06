package config

// ConfigTemplates holds data pertaining the templates
type ConfigTemplates struct {
	TemplatesDir     string   `default:"" json:"templates_dir" yaml:"templates_dir" cli:"templates Templates directory filepath.\nThe templates within must have same names as the original ones:\n vars.md, constants.md, example.md, functionDefinition.md, functionHeading.md, function.md, type.md, index.md, subpackages.md overview.md, package.md, doc.md, mod.md\nDefault templates will be used instead the missing ones.\n      "`
	OutputTemplates  string   `default:"" yaml:"-" json:"-" cli:"output-templates Output template files to the provided directory path.\nIf empty, current working directory will be used.\n   "`
	LinkConstruction string   `default:"" json:"link_construction" yaml:"link_construction" cli:"link-construction Links construction\n  Options: [ direct | github | gitlab | gitea ]\n      "`
	IncludedData     []string `default:"" json:"included_data" yaml:"included_data" cli:"included-data Data to be included in the rendered doc.\n  Options:\n    name,doc,examples,variables,constants,functions,function_examples,types,type_examples,type_functions,type_methods,index\n      " separator:","`
	CustomVars       []string `default:"" json:"custom_vars" yaml:"custom_vars" cli:"custom-vars Custom data to be included in the template overrides\n  Example:\n    -custom-vars var1=value1,var2=value2\n      " separator:","`
	ImportPath       string   `default:"" json:"import_path" yaml:"import_path" cli:"import-path Package import path. Will be parsed as a git server repository URL for links in the documentation.\n      "`
}

// SetupDefault sets the default data
func (c *ConfigTemplates) SetupDefault() {
	// c.ImportPath = "git.example.com/project/repository"
	c.IncludedData = []string{
		"name", "doc", "examples", "variables", "constants",
		"functions", "types", "index", "import",
	}
}

func (c ConfigTemplates) parseLinkConstruction() {
	if c.LinkConstruction == "" {
		c.LinkConstruction = "direct"
	}
	switch c.LinkConstruction {
	case "direct", "github", "gitlab", "gitea":
	default:

	}
}

func (c ConfigTemplates) GetLinkPrefix() string {
	switch c.LinkConstruction {
	case "direct":
		return ""
	case "github", "gitlab":
		return "tree/main/"
	case "gitea", "gogs":
		return "src/branch/main"
	}
	return ""
}
