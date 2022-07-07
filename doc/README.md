
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
- [GetDirectories(string)](#func-getdirectories-string)
- [GetGoFiles(string)](#func-getgofiles-string)
- [GetGoFilesInDir(string)](#func-getgofilesindir-string)
- [GetPackageDocumentation(string, string) error](#func-getpackagedocumentation-string-string-error)
- [GetPackagesDataFromDirRecursive(string, bool, string) error](#func-getpackagesdatafromdirrecursive-string-bool-string-error)
- [ParseGoMod(string, string) error](#func-parsegomod-string-string-error)
- [ParseGoModFile(string, string) error](#func-parsegomodfile-string-string-error)
- [SeekGoMod(string, string, int) error](#func-seekgomod-string-string-int-error)

- [type Const](#type-const)
- [type Example](#type-example)
- [type Func](#type-func)
  - [NewFunc(string, string) Func](#func-newfunc-string-string-func)
  - [FormatParams() string](#func-func-formatparams-string)
  - [FormatParamsBrackets() string](#func-func-formatparamsbrackets-string)
  - [FormatResults() string](#func-func-formatresults-string)
  - [FormatResultsBrackets() string](#func-func-formatresultsbrackets-string)
  - [GetHeadingHREF() string](#func-func-getheadinghref-string)
- [type FuncVar](#type-funcvar)
- [type Module](#type-module)
  - [ParseModule(string) (Module, error)](#func-parsemodule-string-module-error)
- [type Package](#type-package)
  - [GetPackageDataFromDir(string) (Package, error)](#func-getpackagedatafromdir-string-package-error)
  - [GetPackageDataFromDirRecursive(string) (Package, error)](#func-getpackagedatafromdirrecursive-string-package-error)
  - [ParsePackage(string, string) (Package, error)](#func-parsepackage-string-string-package-error)
- [type Type](#type-type)
  - [GetHeadingHREF() string](#func-type-getheadinghref-string)
- [type Var](#type-var)

## func [AddConst(string, string, string)](<definitions.go#L48>)

```go
func AddConst(data *Package, node *doc.Value, path string)
```
## func [AddExample(string, string, string)](<definitions.go#L67>)

```go
func AddExample(data *Package, node *doc.Example, path string)
```
## func [AddFunc(string, string, string)](<definitions.go#L298>)

```go
func AddFunc(data *Package, node *doc.Func, path string)
```
## func [AddFuncExample(string, string, string)](<definitions.go#L76>)

```go
func AddFuncExample(data *Func, node *doc.Example, path string)
```
## func [AddType(string, string, string)](<definitions.go#L324>)

```go
func AddType(data *Package, node *doc.Type, path string)
```
## func [AddTypeExample(string, string, string)](<definitions.go#L85>)

```go
func AddTypeExample(data *Type, node *doc.Example, path string)
```
## func [AddTypeFunc(string, string, string)](<definitions.go#L347>)

```go
func AddTypeFunc(data *Type, node *doc.Func, path string)
```
## func [AddTypeMethod(string, string, string)](<definitions.go#L354>)

```go
func AddTypeMethod(data *Type, node *doc.Func, path string)
```
## func [AddVar(string, string, string)](<definitions.go#L27>)

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
## func [GetPackagesDataFromDirRecursive(string, bool, string) error](<definitions.go#L533>)

```go
func GetPackagesDataFromDirRecursive(dirPath string, includeRoot bool, rootImportPath string) (map[string]*Package, error)
```
## func [ParseGoMod(string, string) error](<definitions.go#L566>)

```go
func ParseGoMod(pkg *Package, path string) error
```
## func [ParseGoModFile(string, string) error](<module.go#L33>)

```go
func ParseGoModFile(module *Module, path string) error
```
## func [SeekGoMod(string, string, int) error](<definitions.go#L588>)

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

## func [NewFunc(string, string) Func](<definitions.go#L182>)

```go
func NewFunc(node *doc.Func, path string) *Func
```

## func (*Func) [FormatParams() string](<definitions.go#L113>)

```go
func (f *Func) FormatParams() string
```
## func (*Func) [FormatParamsBrackets() string](<definitions.go#L131>)

```go
func (f *Func) FormatParamsBrackets() string
```
## func (*Func) [FormatResults() string](<definitions.go#L135>)

```go
func (f *Func) FormatResults() string
```
## func (*Func) [FormatResultsBrackets() string](<definitions.go#L156>)

```go
func (f *Func) FormatResultsBrackets() string
```
## func (*Func) [GetHeadingHREF() string](<definitions.go#L172>)

```go
func (f *Func) GetHeadingHREF() string
```

## type [FuncVar](<definitions.go#L107>)
```go
type FuncVar struct {
	Name	string
	Type	string
	Pointer	bool
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

## func [ParseModule(string) (Module, error)](<module.go#L19>)

```go
func ParseModule(path string) (*Module, error)
```

## type [Package](<definitions.go#L359>)
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

## func [GetPackageDataFromDir(string) (Package, error)](<definitions.go#L525>)

```go
func GetPackageDataFromDir(path string) (*Package, error)
```
## func [GetPackageDataFromDirRecursive(string) (Package, error)](<definitions.go#L514>)

```go
func GetPackageDataFromDirRecursive(path string) (*Package, error)
```
## func [ParsePackage(string, string) (Package, error)](<definitions.go#L420>)

```go
func ParsePackage(docs *doc.Package, path string) (*Package, error)
```

## type [Type](<definitions.go#L303>)
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

## func (*Type) [GetHeadingHREF() string](<definitions.go#L314>)

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

