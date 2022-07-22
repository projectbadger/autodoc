
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

## func [AddConst(string, string, string)](<definitions.go#L49>)

```go
func AddConst(data *Package, node *doc.Value, path string)
```
## func [AddExample(string, string, string)](<definitions.go#L68>)

```go
func AddExample(data *Package, node *doc.Example, path string)
```
## func [AddFunc(string, string, string)](<definitions.go#L299>)

```go
func AddFunc(data *Package, node *doc.Func, path string)
```
## func [AddFuncExample(string, string, string)](<definitions.go#L77>)

```go
func AddFuncExample(data *Func, node *doc.Example, path string)
```
## func [AddType(string, string, string)](<definitions.go#L325>)

```go
func AddType(data *Package, node *doc.Type, path string)
```
## func [AddTypeExample(string, string, string)](<definitions.go#L86>)

```go
func AddTypeExample(data *Type, node *doc.Example, path string)
```
## func [AddTypeFunc(string, string, string)](<definitions.go#L348>)

```go
func AddTypeFunc(data *Type, node *doc.Func, path string)
```
## func [AddTypeMethod(string, string, string)](<definitions.go#L355>)

```go
func AddTypeMethod(data *Type, node *doc.Func, path string)
```
## func [AddVar(string, string, string)](<definitions.go#L28>)

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
## func [GetPackagesDataFromDirRecursive(string, bool, string) error](<definitions.go#L554>)

```go
func GetPackagesDataFromDirRecursive(dirPath string, includeRoot bool, rootImportPath string) (map[string]*Package, error)
```
## func [ParseGoMod(string, string) error](<definitions.go#L587>)

```go
func ParseGoMod(pkg *Package, path string) error
```
## func [ParseGoModFile(string, string) error](<module.go#L33>)

```go
func ParseGoModFile(module *Module, path string) error
```
## func [SeekGoMod(string, string, int) error](<definitions.go#L609>)

```go
func SeekGoMod(pkg *Package, path string, levels int) error
```


## type [Const](<definitions.go#L41>)
```go
type Const struct {
	Name		string
	Definition	string
	Doc		string
	Filename	string
	Line		int
}
```

## type [Example](<definitions.go#L62>)
```go
type Example struct {
	Name		string
	Definition	string
	Doc		string
}
```

## type [Func](<definitions.go#L96>)
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

## func [NewFunc(string, string) Func](<definitions.go#L183>)

```go
func NewFunc(node *doc.Func, path string) *Func
```

## func (*Func) [FormatParams() string](<definitions.go#L114>)

```go
func (f *Func) FormatParams() string
```
## func (*Func) [FormatParamsBrackets() string](<definitions.go#L132>)

```go
func (f *Func) FormatParamsBrackets() string
```
## func (*Func) [FormatResults() string](<definitions.go#L136>)

```go
func (f *Func) FormatResults() string
```
## func (*Func) [FormatResultsBrackets() string](<definitions.go#L157>)

```go
func (f *Func) FormatResultsBrackets() string
```
## func (*Func) [GetHeadingHREF() string](<definitions.go#L173>)

```go
func (f *Func) GetHeadingHREF() string
```

## type [FuncVar](<definitions.go#L108>)
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

## type [Package](<definitions.go#L360>)
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

## func [GetPackageDataFromDir(string) (Package, error)](<definitions.go#L546>)

```go
func GetPackageDataFromDir(path string) (*Package, error)
```
## func [GetPackageDataFromDirRecursive(string) (Package, error)](<definitions.go#L535>)

```go
func GetPackageDataFromDirRecursive(path string) (*Package, error)
```
## func [ParsePackage(string, string) (Package, error)](<definitions.go#L441>)

```go
func ParsePackage(docs *doc.Package, path string) (*Package, error)
```

## func (Package) [PathIndent()](<definitions.go#L391>)

```go
func (p Package) PathIndent() func(string) string
```
## func (*Package) [PathSplit()](<definitions.go#L387>)

```go
func (p *Package) PathSplit() []string
```

## type [Type](<definitions.go#L304>)
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

## func (*Type) [GetHeadingHREF() string](<definitions.go#L315>)

```go
func (t *Type) GetHeadingHREF() string
```

## type [Var](<definitions.go#L20>)
```go
type Var struct {
	Name		string
	Definition	string
	Doc		string
	Filename	string
	Line		int
}
```

