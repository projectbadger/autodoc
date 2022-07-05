
# config

```go
import github.com/projectbadger/autodoc/config
```

Package config handles all the configuration data.

Before use, the package must be initialized by the function `SetupConfig() error`. It is normally already initialized by the `init()` function.
Example:

```go
err := config.SetupConfig()
```
The configuration data is accessible from the `Cfg` variable.

Example:

```go
configFilePath := config.Cfg.File
```


<details>
<summary>Example</summary>
<p>

```go
{

	fmt.Println(config.Cfg.PackageDir)

}
```
</p></details>

## Index

- func [SetupConfig](#func-setupconfig)

- [type Config](#config)
- [type ConfigApp](#configapp)
  - func [GetParsedConfig](#func-getparsedconfig)
  - func [NewDefaultConfigApp](#func-newdefaultconfigapp)
  - func [(c *ConfigApp) CheckIsVersion](#func--configapp-checkisversion)
  - func [(c *ConfigApp) IsVersion](#func--configapp-isversion)
  - func [(cfg *ConfigApp) SaveToFile](#func--configapp-savetofile)
  - func [(c *ConfigApp) SetupDefault](#func--configapp-setupdefault)
- [type ConfigTemplates](#configtemplates)
  - func [(c ConfigTemplates) GetLinkPrefix](#func--configtemplates-getlinkprefix)
  - func [(c *ConfigTemplates) SetupDefault](#func--configtemplates-setupdefault)
- [Variables](#variables)

## Variables
```go
var (
	// Main application configuration data.
	Cfg	*ConfigApp
	// supplied through ldflags
	PackageName	= "github.com/projectbadger/autodoc"
	Version		= "development"
	BuildTime	= ""
)
```

## func [SetupConfig](<config.go#L70>)

SetupConfig sets Cfg variable to the parsed *ConfigApp

```go
func SetupConfig() error
```


## type [Config](<config.go#L19>)
```go
type Config interface {
	SetupDefault()
}
```

## type [ConfigApp](<configApp.go#L11>)
ConfigApp holds all the application configuration data.
```go
type ConfigApp struct {
	File		string		`default:"" json:"-" yaml:"-" cli:"config Config file path\n     "`
	CreateConfig	string		`default:"" yaml:"-" json:"-" cli:"create-config Create a named default config file with cli parameters and environment variables.\n   "`
	Output		string		`json:"output" yaml:"output" cli:"output Output the generated documentation to file\n    "`
	PackageDir	string		`default:"." json:"package_dir" yaml:"package_dir" cli:"package Package directory filepath.\nThe contents of this directory will be parsed as a Go package\n      "`
	ModuleDir	string		`default:"." json:"module_dir" yaml:"module_dir" cli:"module Module directory filepath.\nThe contents of this directory will be parsed as a Go module\n      "`
	ExcludeDirs	[]string	`default:"" json:"exclude_dirs" yaml:"exclude_dirs" cli:"exclude Exclude directories from the package search.\nDefault:\n  node_nodules;.git\n      "`
	Version		bool		`default:"false" json:"-" yaml:"-" cli:"version Print version"`
	VersionShort	bool		`default:"false" json:"-" yaml:"-" cli:"v Print version"`
	Templates	ConfigTemplates	`yaml:"templates" json:"templates"`
}
```

## func [GetParsedConfig](<config.go#L46>)

GetParsedConfig returns a config, filled with
environment variables, config file and CLI arguments
values.

Variable source parsing order:
```go
1. config file (if defined)	2. environment variables
	3. CLI arguments

```
CLI arguments always take precedence.

If the CLI flag "-create-config" was provided along with
a filepath, a config file will be created with default
values and any parsed environment variables and CLI
arguments.

```go
func GetParsedConfig() (cfg *ConfigApp, err error)
```
## func [NewDefaultConfigApp](<config.go#L25>)

NewDefaultConfigApp returns a *ConfigApp with all the
default values filled.

```go
func NewDefaultConfigApp() *ConfigApp
```

## func [(c *ConfigApp) CheckIsVersion](<configApp.go#L29>)

CheckIfVersion checks whether the version flag has been
set and prints the version and exits if it has.

```go
func (c *ConfigApp) CheckIsVersion()
```
## func [(c *ConfigApp) IsVersion](<configApp.go#L23>)

```go
func (c *ConfigApp) IsVersion() bool
```
## func [(cfg *ConfigApp) SaveToFile](<configApp.go#L39>)

SaveToFile saves the config to a file in YAML format

```go
func (cfg *ConfigApp) SaveToFile(path string) error
```
## func [(c *ConfigApp) SetupDefault](<configApp.go#L51>)

SetupDefault sets up default config data.

```go
func (c *ConfigApp) SetupDefault()
```

## type [ConfigTemplates](<configTemplates.go#L4>)
ConfigTemplates holds data pertaining the templates
```go
type ConfigTemplates struct {
	TemplatesDir		string		`default:"" json:"templates_dir" yaml:"templates_dir" cli:"templates Templates directory filepath.\nThe templates within must have same names as the original ones:\n  doc.md, package.md, index.md, example.md, function.md, type.md\nDefault templates will be used instead the missing ones.\n      "`
	OutputTemplates		string		`default:"" yaml:"-" json:"-" cli:"output-templates Output template files to the provided directory path.\nIf empty, current working directory will be used.\n   "`
	LinkConstruction	string		`default:"" json:"link_construction" yaml:"link_construction" cli:"link-construction Links construction\n  Options: [ direct | github | gitlab | gitea ]\n      "`
	IncludedData		[]string	`default:"" json:"included_data" yaml:"included_data" cli:"included-data Data to be included in the rendered doc.\n  Options:\n    name,doc,examples,variables,constants,functions,function_examples,types,type_examples,type_functions,type_methods,index\n      " separator:","`
	CustomVars		[]string	`default:"" json:"custom_vars" yaml:"custom_vars" cli:"custom-vars Custom data to be included in the template overrides\n  Example:\n    -custom-vars var1=value1,var2=value2\n      " separator:","`
	ImportPath		string		`default:"" json:"import_path" yaml:"import_path" cli:"import-path Package import path. Will be parsed as a git server repository URL for links in the documentation.\n      "`
}
```

## func [(c ConfigTemplates) GetLinkPrefix](<configTemplates.go#L33>)

```go
func (c ConfigTemplates) GetLinkPrefix() string
```
## func [(c *ConfigTemplates) SetupDefault](<configTemplates.go#L14>)

SetupDefault sets the default data

```go
func (c *ConfigTemplates) SetupDefault()
```

