
# doc

```go
import github.com/projectbadger/autodoc/doc
```

## Index

- func [AddConst](#func-addconst)
- func [AddExample](#func-addexample)
- func [AddFunc](#func-addfunc)
- func [AddFuncExample](#func-addfuncexample)
- func [AddType](#func-addtype)
- func [AddTypeExample](#func-addtypeexample)
- func [AddTypeFunc](#func-addtypefunc)
- func [AddTypeMethod](#func-addtypemethod)
- func [AddVar](#func-addvar)
- func [GetDirectories](#func-getdirectories)
- func [GetGoFiles](#func-getgofiles)
- func [GetGoFilesInDir](#func-getgofilesindir)
- func [GetPackageDocumentation](#func-getpackagedocumentation)
- func [GetPackagesDataFromDirRecursive](#func-getpackagesdatafromdirrecursive)
- func [ParseGoMod](#func-parsegomod)
- func [ParseGoModFile](#func-parsegomodfile)
- func [SeekGoMod](#func-seekgomod)

- [type Const](#const)
- [type Example](#example)
- [type Func](#func)
- [type Module](#module)
  - func [ParseModule](#func-parsemodule)
- [type Package](#package)
  - func [GetPackageDataFromDir](#func-getpackagedatafromdir)
  - func [GetPackageDataFromDirRecursive](#func-getpackagedatafromdirrecursive)
  - func [ParsePackage](#func-parsepackage)
- [type Type](#type)
- [type Var](#var)

## func [AddConst](<definitions.go#L46>)

```go
func AddConst(data *Package, node *doc.Value, path string)
```
## func [AddExample](<definitions.go#L65>)

```go
func AddExample(data *Package, node *doc.Example, path string)
```
## func [AddFunc](<definitions.go#L102>)

```go
func AddFunc(data *Package, node *doc.Func, path string)
```
## func [AddFuncExample](<definitions.go#L74>)

```go
func AddFuncExample(data *Func, node *doc.Example, path string)
```
## func [AddType](<definitions.go#L130>)

```go
func AddType(data *Package, node *doc.Type, path string)
```
## func [AddTypeExample](<definitions.go#L83>)

```go
func AddTypeExample(data *Type, node *doc.Example, path string)
```
## func [AddTypeFunc](<definitions.go#L153>)

```go
func AddTypeFunc(data *Type, node *doc.Func, path string)
```
## func [AddTypeMethod](<definitions.go#L166>)

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
## func [GetPackagesDataFromDirRecursive](<definitions.go#L358>)

```go
func GetPackagesDataFromDirRecursive(dirPath string, includeRoot bool, rootImportPath string) (map[string]*Package, error)
```
## func [ParseGoMod](<definitions.go#L394>)

```go
func ParseGoMod(pkg *Package, path string) error
```
## func [ParseGoModFile](<module.go#L33>)

```go
func ParseGoModFile(module *Module, path string) error
```
## func [SeekGoMod](<definitions.go#L416>)

```go
func SeekGoMod(pkg *Package, path string, levels int) error
```


## type [Const](<definitions.go#L38>)
```go
type Const struct {
	Name		string
	Definition	string
	Doc		string
	Filename	string
	Line		int
}
```

## type [Example](<definitions.go#L59>)
```go
type Example struct {
	Name		string
	Definition	string
	Doc		string
}
```

## type [Func](<definitions.go#L93>)
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

## type [Package](<definitions.go#L179>)
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

## func [GetPackageDataFromDir](<definitions.go#L349>)

```go
func GetPackageDataFromDir(path string) (*Package, error)
```
## func [GetPackageDataFromDirRecursive](<definitions.go#L338>)

```go
func GetPackageDataFromDirRecursive(path string) (*Package, error)
```
## func [ParsePackage](<definitions.go#L240>)

```go
func ParsePackage(docs *doc.Package, path string) (*Package, error)
```

## type [Type](<definitions.go#L119>)
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

