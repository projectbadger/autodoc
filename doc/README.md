
# doc

```go
import github.com/projectbadger/autodoc/doc
```

## Index

- [func AddConst(data *Package, node *doc.Value, path string)](#func-addconst-package-node-docvalue-path-string)
- [func AddExample(data *Package, node *doc.Example, path string)](#func-addexample-package-node-docexample-path-string)
- [func AddFunc(data *Package, node *doc.Func, path string)](#func-addfunc-package-node-docfunc-path-string)
- [func AddFuncExample(data *Func, node *doc.Example, path string)](#func-addfuncexample-func-node-docexample-path-string)
- [func AddType(data *Package, node *doc.Type, path string)](#func-addtype-package-node-doctype-path-string)
- [func AddTypeExample(data *Type, node *doc.Example, path string)](#func-addtypeexample-type-node-docexample-path-string)
- [func AddTypeFunc(data *Type, node *doc.Func, path string)](#func-addtypefunc-type-node-docfunc-path-string)
- [func AddTypeMethod(data *Type, node *doc.Func, path string)](#func-addtypemethod-type-node-docfunc-path-string)
- [func AddVar(data *Package, node *doc.Value, path string)](#func-addvar-package-node-docvalue-path-string)
- [func GetDirectories(path string) []string](#func-getdirectories-string-string)
- [func GetGoFiles(path string) []*ast.File](#func-getgofiles-string-astfile)
- [func GetGoFilesInDir(path string) []*ast.File](#func-getgofilesindir-string-astfile)
- [func GetPackageDocumentation(packageFilePath, packageImportPath string) (*doc.Package, error)](#func-getpackagedocumentation-packageimportpath-string-docpackage-error)
- [func GetPackagesDataFromDirRecursive(dirPath string, includeRoot bool, rootImportPath string) (map[string]*Package, error)](#func-getpackagesdatafromdirrecursive-string-includeroot-bool-rootimportpath-string-stringpackage-error)
- [func ParseGoMod(pkg *Package, path string) error](#func-parsegomod-package-path-string-error)
- [func ParseGoModFile(module *Module, path string) error](#func-parsegomodfile-module-path-string-error)
- [func SeekGoMod(pkg *Package, path string, levels int) error](#func-seekgomod-package-path-string-levels-int-error)
- [type Const](#const)
- [type Example](#example)
- [type Func](#func)
- [type Module](#module)
  - [func ParseModule(path string) (*Module, error)](#func-parsemodule-string-module-error)
- [type Package](#package)
  - [func GetPackageDataFromDir(path string) (*Package, error)](#func-getpackagedatafromdir-string-package-error)
  - [func GetPackageDataFromDirRecursive(path string) (*Package, error)](#func-getpackagedatafromdirrecursive-string-package-error)
  - [func ParsePackage(docs *doc.Package, path string) (*Package, error)](#func-parsepackage-docpackage-path-string-package-error)
- [type Type](#type)
- [type Var](#var)

## func [AddConst](<definitions.go#L50>)

```go
func AddConst(data *Package, node *doc.Value, path string)
```
## func [AddExample](<definitions.go#L69>)

```go
func AddExample(data *Package, node *doc.Example, path string)
```
## func [AddFunc](<definitions.go#L106>)

```go
func AddFunc(data *Package, node *doc.Func, path string)
```
## func [AddFuncExample](<definitions.go#L78>)

```go
func AddFuncExample(data *Func, node *doc.Example, path string)
```
## func [AddType](<definitions.go#L134>)

```go
func AddType(data *Package, node *doc.Type, path string)
```
## func [AddTypeExample](<definitions.go#L87>)

```go
func AddTypeExample(data *Type, node *doc.Example, path string)
```
## func [AddTypeFunc](<definitions.go#L157>)

```go
func AddTypeFunc(data *Type, node *doc.Func, path string)
```
## func [AddTypeMethod](<definitions.go#L170>)

```go
func AddTypeMethod(data *Type, node *doc.Func, path string)
```
## func [AddVar](<definitions.go#L25>)

```go
func AddVar(data *Package, node *doc.Value, path string)
```
## func [GetDirectories](<doc.go#L87>)

```go
func GetDirectories(path string) []string
```
## func [GetGoFiles](<doc.go#L21>)

```go
func GetGoFiles(path string) []*ast.File
```
## func [GetGoFilesInDir](<doc.go#L51>)

```go
func GetGoFilesInDir(path string) []*ast.File
```
## func [GetPackageDocumentation](<doc.go#L109>)

```go
func GetPackageDocumentation(packageFilePath, packageImportPath string) (*doc.Package, error)
```
## func [GetPackagesDataFromDirRecursive](<definitions.go#L362>)

```go
func GetPackagesDataFromDirRecursive(dirPath string, includeRoot bool, rootImportPath string) (map[string]*Package, error)
```
## func [ParseGoMod](<definitions.go#L398>)

```go
func ParseGoMod(pkg *Package, path string) error
```
## func [ParseGoModFile](<module.go#L33>)

```go
func ParseGoModFile(module *Module, path string) error
```
## func [SeekGoMod](<definitions.go#L420>)

```go
func SeekGoMod(pkg *Package, path string, levels int) error
```


## type [Const](<definitions.go#L42>)
```go
type Const struct {
	Name		string
	Definition	string
	Doc		string
	Filename	string
	Line		int
}
```

## type [Example](<definitions.go#L63>)
```go
type Example struct {
	Name		string
	Definition	string
	Doc		string
}
```

## type [Func](<definitions.go#L97>)
```go
type Func struct {
	Name		string
	Definition	string
	Doc		string
	Examples	[]*Example
	Filename	string
	Line		int
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

## func [ParseModule](<module.go#L19>)

```go
func ParseModule(path string) (*Module, error)
```

## type [Package](<definitions.go#L183>)
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

## func [GetPackageDataFromDir](<definitions.go#L353>)

```go
func GetPackageDataFromDir(path string) (*Package, error)
```
## func [GetPackageDataFromDirRecursive](<definitions.go#L342>)

```go
func GetPackageDataFromDirRecursive(path string) (*Package, error)
```
## func [ParsePackage](<definitions.go#L244>)

```go
func ParsePackage(docs *doc.Package, path string) (*Package, error)
```

## type [Type](<definitions.go#L123>)
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

## type [Var](<definitions.go#L17>)
```go
type Var struct {
	Name		string
	Definition	string
	Doc		string
	Filename	string
	Line		int
}
```

