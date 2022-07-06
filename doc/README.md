
# doc

```go
import github.com/projectbadger/autodoc/doc
```

## Index

- [AddConst(string, string, string)](#func-addconst-string-string-string)
- [AddExample(string, string, string)](#func-addexample-string-string-string)
- [AddFunc(string, string, string)](#func-addfunc-string-string-string)
- [AddFuncExample(string, string, string)](#func-addfuncexample-string-string-string)
- [AddType(string, string, string)](#func-addtype-string-string-string)
- [AddTypeExample(string, string, string)](#func-addtypeexample-string-string-string)
- [AddTypeFunc(string, string, string)](#func-addtypefunc-string-string-string)
- [AddTypeMethod(string, string, string)](#func-addtypemethod-string-string-string)
- [AddVar(string, string, string)](#func-addvar-string-string-string)
- [GetDirectories(string) ](#func-getdirectories-string)
- [GetGoFiles(string) ](#func-getgofiles-string)
- [GetGoFilesInDir(string) ](#func-getgofilesindir-string)
- [GetPackageDocumentation(string, string) (error, error)](#func-getpackagedocumentation-string-string-error-error)
- [GetPackagesDataFromDirRecursive(string, bool, string) (error, error)](#func-getpackagesdatafromdirrecursive-string-bool-string-error-error)
- [ParseGoMod(string, string) error](#func-parsegomod-string-string-error)
- [ParseGoModFile(string, string) error](#func-parsegomodfile-string-string-error)
- [SeekGoMod(string, string, int) error](#func-seekgomod-string-string-int-error)

- [type Const](#type-const)
- [type Example](#type-example)
- [type Func](#type-func)
  - [FormatParams()](#func-formatparams)
  - [FormatParamsBrackets()](#func-formatparamsbrackets)
  - [FormatResults()](#func-formatresults)
  - [FormatResultsBrackets()](#func-formatresultsbrackets)
  - [GetHeadingHREF()](#func-getheadinghref)
- [type FuncVar](#type-funcvar)
- [type Module](#type-module)
  - [ParseModule()](#func-parsemodule)
- [type Package](#type-package)
  - [GetPackageDataFromDir()](#func-getpackagedatafromdir)
  - [GetPackageDataFromDirRecursive()](#func-getpackagedatafromdirrecursive)
  - [ParsePackage()](#func-parsepackage)
- [type Type](#type-type)
  - [GetHeadingHREF()](#func-getheadinghref)
- [type Var](#type-var)

## func [AddConst(string, string, string)](<definitions.go#L48>)
```go
func AddConst(data *Package, node *doc.Value, path string)
```
## func [AddExample(string, string, string)](<definitions.go#L67>)
```go
func AddExample(data *Package, node *doc.Example, path string)
```
## func [AddFunc(string, string, string)](<definitions.go#L180>)
```go
func AddFunc(data *Package, node *doc.Func, path string)
```
## func [AddFuncExample(string, string, string)](<definitions.go#L76>)
```go
func AddFuncExample(data *Func, node *doc.Example, path string)
```
## func [AddType(string, string, string)](<definitions.go#L274>)
```go
func AddType(data *Package, node *doc.Type, path string)
```
## func [AddTypeExample(string, string, string)](<definitions.go#L85>)
```go
func AddTypeExample(data *Type, node *doc.Example, path string)
```
## func [AddTypeFunc(string, string, string)](<definitions.go#L297>)
```go
func AddTypeFunc(data *Type, node *doc.Func, path string)
```
## func [AddTypeMethod(string, string, string)](<definitions.go#L310>)
```go
func AddTypeMethod(data *Type, node *doc.Func, path string)
```
## func [AddVar(string, string, string)](<definitions.go#L27>)
```go
func AddVar(data *Package, node *doc.Value, path string)
```
## func [GetDirectories(string) ](<doc.go#L87>)
```go
func GetDirectories(path string) []string
```
## func [GetGoFiles(string) ](<doc.go#L21>)
```go
func GetGoFiles(path string) []*ast.File
```
## func [GetGoFilesInDir(string) ](<doc.go#L51>)
```go
func GetGoFilesInDir(path string) []*ast.File
```
## func [GetPackageDocumentation(string, string) (error, error)](<doc.go#L109>)
```go
func GetPackageDocumentation(packageFilePath, packageImportPath string) (*doc.Package, error)
```
## func [GetPackagesDataFromDirRecursive(string, bool, string) (error, error)](<definitions.go#L497>)
```go
func GetPackagesDataFromDirRecursive(dirPath string, includeRoot bool, rootImportPath string) (map[string]*Package, error)
```
## func [ParseGoMod(string, string) error](<definitions.go#L530>)
```go
func ParseGoMod(pkg *Package, path string) error
```
## func [ParseGoModFile(string, string) error](<module.go#L33>)
```go
func ParseGoModFile(module *Module, path string) error
```
## func [SeekGoMod(string, string, int) error](<definitions.go#L552>)
```go
func SeekGoMod(pkg *Package, path string, levels int) error
```


## type [Const](<definitions.go#L40>)
```go
type Const struct {
	Name		string
	Definition	string
	Doc		string
	Filename	string
	Line		int
}
```

## type [Example](<definitions.go#L61>)
```go
type Example struct {
	Name		string
	Definition	string
	Doc		string
}
```

## type [Func](<definitions.go#L95>)
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

## func [FormatParams()](<definitions.go#L112>)
```go
func (f *Func) FormatParams() string
```
## func [FormatParamsBrackets()](<definitions.go#L130>)
```go
func (f *Func) FormatParamsBrackets() string
```
## func [FormatResults()](<definitions.go#L134>)
```go
func (f *Func) FormatResults() string
```
## func [FormatResultsBrackets()](<definitions.go#L155>)
```go
func (f *Func) FormatResultsBrackets() string
```
## func [GetHeadingHREF()](<definitions.go#L170>)
```go
func (f *Func) GetHeadingHREF() string
```

## type [FuncVar](<definitions.go#L107>)
```go
type FuncVar struct {
	Name	string
	Type	string
}
```

## type [Module](<module.go#L10>)
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

## func [ParseModule()](<module.go#L19>)
```go
func ParseModule(path string) (*Module, error)
```

## type [Package](<definitions.go#L323>)
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

## func [GetPackageDataFromDir()](<definitions.go#L489>)
```go
func GetPackageDataFromDir(path string) (*Package, error)
```
## func [GetPackageDataFromDirRecursive()](<definitions.go#L478>)
```go
func GetPackageDataFromDirRecursive(path string) (*Package, error)
```
## func [ParsePackage()](<definitions.go#L384>)
```go
func ParsePackage(docs *doc.Package, path string) (*Package, error)
```

## type [Type](<definitions.go#L259>)
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

## func [GetHeadingHREF()](<definitions.go#L270>)
```go
func (t *Type) GetHeadingHREF() string
```

## type [Var](<definitions.go#L19>)
```go
type Var struct {
	Name		string
	Definition	string
	Doc		string
	Filename	string
	Line		int
}
```

