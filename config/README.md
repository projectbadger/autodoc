
# config

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

- [func SetupConfig() error](#func-setupconfig-error)
- [type Config](#config)
- [type ConfigApp](#configapp)
  - [func GetParsedConfig() (cfg *ConfigApp, err error)](#func-getparsedconfig--configapp-err-error)
  - [func NewDefaultConfigApp() *ConfigApp](#func-newdefaultconfigapp-configapp)
  - [func (c *ConfigApp) CheckIsVersion()](#func--configapp-checkisversion)
  - [func (c *ConfigApp) IsVersion() bool](#func--configapp-isversion-bool)
  - [func (cfg *ConfigApp) SaveToFile(path string) error](#func--configapp-savetofile-string-error)
  - [func (c *ConfigApp) SetupDefault()](#func--configapp-setupdefault)
- [type ConfigTemplates](#configtemplates)
  - [func (c ConfigTemplates) GetLinkPrefix() string](#func--configtemplates-getlinkprefix-string)
  - [func (c *ConfigTemplates) SetupDefault()](#func--configtemplates-setupdefault)

## func [SetupConfig](<config.go#L86>)

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

## type [ConfigApp](<configApp.go#L13>)
ConfigApp holds all the application configuration data.
```go
type ConfigApp struct {
	File		string	`default:"" json:"-" yaml:"-" cli:"config Config file path\n     "`
	CreateConfig	string	`default:"" yaml:"-" json:"-" cli:"create-config Create a named default config file with cli parameters and environment variables.\n   "`
	Output		string	`json:"output" yaml:"output" cli:"output Output the generated documentation to file\n    "`
	DebugMode	bool	`default:"false" yaml:"-" json:"-" cli:"debug Run the app in debug mode\n   "`
	// LoginUrl     string                `default:"" json:"login_url" yaml:"login_url" cli:"login-url Login URL to appear at the top of the page\n      "`
	// LogoutUrl    string                `default:"" json:"logout_url" yaml:"logout_url" cli:"logout-url Logout URL to appear at the top of the page\n      "`
	// Filters     filters.RequestFilter `json:"filters" yaml:"filters"`
	PackageDir	string		`default:"" json:"package_dir" yaml:"package_dir" cli:"package Package directory filepath.\nThe contents of this directory will be parsed as a Go package\n      "`
	ModuleDir	string		`default:"" json:"module_dir" yaml:"module_dir" cli:"module Module directory filepath.\nThe contents of this directory will be parsed as a Go module\n      "`
	ExcludeDirs	string		`default:"" json:"exclude_dirs" yaml:"exclude_dirs" cli:"exclude Exclude directories from the package search.\nDefault:\n  node_nodules;.git\n      "`
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

## func [CheckIsVersion](<configApp.go#L33>)

```go
func (c *ConfigApp) CheckIsVersion()
```
## func [IsVersion](<configApp.go#L29>)

```go
func (c *ConfigApp) IsVersion() bool
```
## func [SaveToFile](<configApp.go#L43>)

SaveToFile saves the config to a file in YAML format

```go
func (cfg *ConfigApp) SaveToFile(path string) error
```
## func [SetupDefault](<configApp.go#L55>)

SetupDefault sets up default config data.

```go
func (c *ConfigApp) SetupDefault()
```

## type [ConfigTemplates](<configTemplates.go#L3>)
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

## func [GetLinkPrefix](<configApp.go#L72>)

```go
func (c ConfigTemplates) GetLinkPrefix() string
```
## func [SetupDefault](<configTemplates.go#L12>)

```go
func (c *ConfigTemplates) SetupDefault()
```

