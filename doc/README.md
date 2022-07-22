
# doc

```go
import github.com/projectbadger/autodoc/doc
```

## Index

- [AddConst(string, string, string)](#func-addconststring-string-string)
- [AddExample(string, string, string)](#func-addexamplestring-string-string)
- [AddFunc(string, string, string)](#func-addfuncstring-string-string)
- [AddFuncExample(string, string, string)](#func-addfuncexamplestring-string-string)
- [AddType(string, string, string)](#func-addtypestring-string-string)
- [AddTypeExample(string, string, string)](#func-addtypeexamplestring-string-string)
- [AddTypeFunc(string, string, string)](#func-addtypefuncstring-string-string)
- [AddTypeMethod(string, string, string)](#func-addtypemethodstring-string-string)
- [AddVar(string, string, string)](#func-addvarstring-string-string)
- [GetDirectories(string)](#func-getdirectoriesstring)
- [GetGoFiles(string)](#func-getgofilesstring)
- [GetGoFilesInDir(string)](#func-getgofilesindirstring)
- [GetPackageDocumentation(string, string) error](#func-getpackagedocumentationstring-string-error)
- [GetPackagesDataFromDirRecursive(string, bool, string) error](#func-getpackagesdatafromdirrecursivestring-bool-string-error)
- [ParseGoMod(string, string) error](#func-parsegomodstring-string-error)
- [ParseGoModFile(string, string) error](#func-parsegomodfilestring-string-error)
- [SeekGoMod(string, string, int) error](#func-seekgomodstring-string-int-error)

- [type Const](#type-const)
- [type Example](#type-example)
- [type Func](#type-func)
  - [NewFunc(string, string) Func](#func-newfuncstring-string-func)
  - [FormatParams() string](#func-func-formatparams-string)
  - [FormatParamsBrackets() string](#func-func-formatparamsbrackets-string)
  - [FormatResults() string](#func-func-formatresults-string)
  - [FormatResultsBrackets() string](#func-func-formatresultsbrackets-string)
  - [GetHeadingHREF() string](#func-func-getheadinghref-string)
- [type FuncVar](#type-funcvar)
- [type Module](#type-module)
  - [ParseModule(string) (Module, error)](#func-parsemodulestring-module-error)
- [type Package](#type-package)
  - [GetPackageDataFromDir(string) (Package, error)](#func-getpackagedatafromdirstring-package-error)
  - [GetPackageDataFromDirRecursive(string) (Package, error)](#func-getpackagedatafromdirrecursivestring-package-error)
  - [ParsePackage(string, string) (Package, error)](#func-parsepackagestring-string-package-error)
  - [PathIndent()](#func-package-pathindent)
  - [PathSplit()](#func-package-pathsplit)
- [type Type](#type-type)
  - [GetHeadingHREF() string](#func-type-getheadinghref-string)
- [type Var](#type-var)

## func [AddConst(string, string, string)](<definitions.go#L54>)

AddConst parses *doc.Value that contains constants data
and appends a new Var struct to the provided data
*Package.


```go
func AddConst(data *Package, node *doc.Value, path string)
```
## func [AddExample(string, string, string)](<definitions.go#L77>)

AddVar parses *doc.Example containing example data and
appends a new *Example to the provided data
*Package.


```go
func AddExample(data *Package, node *doc.Example, path string)
```
## func [AddFunc(string, string, string)](<definitions.go#L330>)

AddFunc parses *doc.Func and appends a new *Func to the
provided data *Package.


```go
func AddFunc(data *Package, node *doc.Func, path string)
```
## func [AddFuncExample(string, string, string)](<definitions.go#L90>)

AddFuncExample parses *doc.Example containing example
data and appends a new *Example to the provided data
*Package.


```go
func AddFuncExample(data *Func, node *doc.Example, path string)
```
## func [AddType(string, string, string)](<definitions.go#L361>)

AddType parses *doc.Type and appends a new *Type to the
provided data *Package.


```go
func AddType(data *Package, node *doc.Type, path string)
```
## func [AddTypeExample(string, string, string)](<definitions.go#L103>)

AddTypeExample parses *doc.Example containing example
data and appends a new *Example to the provided data
*Package.


```go
func AddTypeExample(data *Type, node *doc.Example, path string)
```
## func [AddTypeFunc(string, string, string)](<definitions.go#L386>)

AddTypeFunc parses *doc.Func from a *doc.Type and appends a
new *Func to the provided data *Package.


```go
func AddTypeFunc(data *Type, node *doc.Func, path string)
```
## func [AddTypeMethod(string, string, string)](<definitions.go#L395>)

AddTypeMethod parses *doc.Func from a *doc.Type and appends a
new *Func to the provided data *Package.


```go
func AddTypeMethod(data *Type, node *doc.Func, path string)
```
## func [AddVar(string, string, string)](<definitions.go#L30>)

AddVar parses *doc.Value that contains variables data and
appends a new Var struct to the provided data *Package.


```go
func AddVar(data *Package, node *doc.Value, path string)
```
## func [GetDirectories(string)](<doc.go#L87>)

```go
func GetDirectories(path string) []string
```
## func [GetGoFiles(string)](<doc.go#L21>)

```go
func GetGoFiles(path string) []*ast.File
```
## func [GetGoFilesInDir(string)](<doc.go#L51>)

```go
func GetGoFilesInDir(path string) []*ast.File
```
## func [GetPackageDocumentation(string, string) error](<doc.go#L109>)

```go
func GetPackageDocumentation(packageFilePath, packageImportPath string) (*doc.Package, error)
```
## func [GetPackagesDataFromDirRecursive(string, bool, string) error](<definitions.go#L608>)

GetPackagesDataFromDirRecursive parses all the .go files
in the directories in the path recursively and returns
them as a map[string]*Package with path as key.


```go
func GetPackagesDataFromDirRecursive(dirPath string, includeRoot bool, rootImportPath string) (map[string]*Package, error)
```
## func [ParseGoMod(string, string) error](<definitions.go#L643>)

ParseGoMod parses a go.mod file and fills the data in the
provided *Package.


```go
func ParseGoMod(pkg *Package, path string) error
```
## func [ParseGoModFile(string, string) error](<module.go#L37>)

ParseGoModFile parses a go.mod file.


```go
func ParseGoModFile(module *Module, path string) error
```
## func [SeekGoMod(string, string, int) error](<definitions.go#L667>)

SeekGoMod seeks a go.mod file for a provided number of
levels upwards.


```go
func SeekGoMod(pkg *Package, path string, levels int) error
```


## type [Const](<definitions.go#L43>)
```go
type Const struct {
	Name		string
	Definition	string
	Doc		string
	Filename	string
	Line		int
}
```

## type [Example](<definitions.go#L68>)

Example holds documentation example data.
```go
type Example struct {
	Name		string
	Definition	string
	Doc		string
}
```

## type [Func](<definitions.go#L114>)

Func holds function data.
```go
type Func struct {
	Name		string
	Definition	string
	Recv		FuncVar
	Params		[]FuncVar
	Results		[]FuncVar
	Doc		string
	Examples	[]*Example
	Filename	string
	Line		int
}
```

## func [NewFunc(string, string) Func](<definitions.go#L212>)

NewFunc parses *doc.Func and returns a *Func.


```go
func NewFunc(node *doc.Func, path string) *Func
```

## func (*Func) [FormatParams() string](<definitions.go#L134>)

FormatParams returns function parameters with names
removed.


```go
func (f *Func) FormatParams() string
```
## func (*Func) [FormatParamsBrackets() string](<definitions.go#L154>)

FormatParams returns function parameters in brackets
with names removed.


```go
func (f *Func) FormatParamsBrackets() string
```
## func (*Func) [FormatResults() string](<definitions.go#L160>)

FormatResults returns function results with names
removed.


```go
func (f *Func) FormatResults() string
```
## func (*Func) [FormatResultsBrackets() string](<definitions.go#L183>)

FormatResultsBrackets returns function results with names
removed in brackets.


```go
func (f *Func) FormatResultsBrackets() string
```
## func (*Func) [GetHeadingHREF() string](<definitions.go#L201>)

GetHeadingHREF returns function definition, formated as a
name for a relative link.


```go
func (f *Func) GetHeadingHREF() string
```

## type [FuncVar](<definitions.go#L126>)
```go
type FuncVar struct {
	Name	string
	Type	string
	Pointer	bool
}
```

## type [Module](<module.go#L11>)

Module holds module info.
```go
type Module struct {
	Name		string			`json:"name" yaml:"name"`
	GoVersion	string			`json:"go_version" yaml:"go_version"`
	Heading		string			`json:"heading" yaml:"heading"`
	Overview	string			`json:"overview" yaml:"overview"`
	ImportPath	string			`json:"import_path" yaml:"import_path"`
	Submodules	map[string]*Package	`json:"submodules" yaml:"submodules"`
}
```

## func [ParseModule(string) (Module, error)](<module.go#L22>)

ParseModule parses a go module in the provided path and
returns a *Module.


```go
func ParseModule(path string) (*Module, error)
```

## type [Package](<definitions.go#L401>)

Package holds package data for use in templates.
```go
type Package struct {
	ImportPath	string
	Name		string
	Definition	string
	Doc		string
	Subpackages	map[string]*Package
	Examples	[]*Example
	Funcs		[]*Func
	Types		[]*Type
	Constants	[]*Const
	Vars		[]*Var
	CustomVars	map[string]string
	Path		string
	PathAbs		string
	ShowName	bool
	ShowDoc		bool
	ShowExamples	bool
	ShowIndex	bool
	ShowFuncs	bool
	ShowTypes	bool
	ShowConsts	bool
	ShowVars	bool
	ShowSubpackages	bool
	ShowOverview	bool
	ShowImportPath	bool
}
```

## func [GetPackageDataFromDir(string) (Package, error)](<definitions.go#L597>)

GetPackageDataFromDirRecursive parses all the .go files
in the provided directory.


```go
func GetPackageDataFromDir(path string) (*Package, error)
```
## func [GetPackageDataFromDirRecursive(string) (Package, error)](<definitions.go#L584>)

GetPackageDataFromDirRecursive parses all the .go files
in the path recursively.


```go
func GetPackageDataFromDirRecursive(path string) (*Package, error)
```
## func [ParsePackage(string, string) (Package, error)](<definitions.go#L488>)

ParsePackage parses go files in a directory and returns
a *Package.


```go
func ParsePackage(docs *doc.Package, path string) (*Package, error)
```

## func (Package) [PathIndent()](<definitions.go#L436>)

PathIdent returns 2 spaces for every '/' character in the
path.


```go
func (p Package) PathIndent() func(string) string
```
## func (*Package) [PathSplit()](<definitions.go#L430>)

PathSplit returns the package relative path, split by the
'/' character.


```go
func (p *Package) PathSplit() []string
```

## type [Type](<definitions.go#L336>)

Type holds type data.
```go
type Type struct {
	Name		string
	Definition	string
	Doc		string
	Examples	[]*Example
	Methods		[]*Func
	Funcs		[]*Func
	Filename	string
	Line		int
}
```

## func (*Type) [GetHeadingHREF() string](<definitions.go#L349>)

GetHeadingHREF returns type definition, formated as a
name for a relative link.


```go
func (t *Type) GetHeadingHREF() string
```

## type [Var](<definitions.go#L20>)

Var holds variable data.
```go
type Var struct {
	Name		string
	Definition	string
	Doc		string
	Filename	string
	Line		int
}
```

