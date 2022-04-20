# templates

Documentation templates

## Index


- [func AddConst(data *Package, node *doc.Value)](#func-addconst-package-node-docvalue)
- [func AddExample(data *Package, node *doc.Example)](#func-addexample-package-node-docexample)
- [func AddFunc(data *Package, node *doc.Func)](#func-addfunc-package-node-docfunc)
- [func AddFuncExample(data *Func, node *doc.Example)](#func-addfuncexample-func-node-docexample)
- [func AddType(data *Package, node *doc.Type)](#func-addtype-package-node-doctype)
- [func AddTypeExample(data *Type, node *doc.Example)](#func-addtypeexample-type-node-docexample)
- [func AddTypeFunc(data *Type, node *doc.Func)](#func-addtypefunc-type-node-docfunc)
- [func AddTypeMethod(data *Type, node *doc.Func)](#func-addtypemethod-type-node-docfunc)
- [func AddVar(data *Package, node *doc.Value)](#func-addvar-package-node-docvalue)
- [func GetDirectories(path string) []string](#func-getdirectories-string-string)
- [func GetPackageDocumentation(packageFilePath, packageImportPath string) (*doc.Package, error)](#func-getpackagedocumentation-packageimportpath-string-docpackage-error)
- [func GetTemplateFuncMap() template.FuncMap](#func-gettemplatefuncmap-templatefuncmap)

- [type Const](#const)
- [type Example](#example)
- [type Func](#func)
- [type Package](#package)
  - [func ParsePackage(docs *doc.Package) *Package](#func-parsepackage-docpackage-package)
- [type Type](#type)
- [type Var](#var)

## func [AddConst](&lt;doc.go#L157>)

```go
func AddConst(data *Package, node *doc.Value)
```
## func [AddExample](&lt;doc.go#L176>)

```go
func AddExample(data *Package, node *doc.Example)
```
## func [AddFunc](&lt;doc.go#L213>)

```go
func AddFunc(data *Package, node *doc.Func)
```
## func [AddFuncExample](&lt;doc.go#L185>)

```go
func AddFuncExample(data *Func, node *doc.Example)
```
## func [AddType](&lt;doc.go#L241>)

```go
func AddType(data *Package, node *doc.Type)
```
## func [AddTypeExample](&lt;doc.go#L194>)

```go
func AddTypeExample(data *Type, node *doc.Example)
```
## func [AddTypeFunc](&lt;doc.go#L264>)

```go
func AddTypeFunc(data *Type, node *doc.Func)
```
## func [AddTypeMethod](&lt;doc.go#L277>)

```go
func AddTypeMethod(data *Type, node *doc.Func)
```
## func [AddVar](&lt;doc.go#L136>)

```go
func AddVar(data *Package, node *doc.Value)
```
## func [GetDirectories](&lt;doc.go#L47>)

```go
func GetDirectories(path string) []string
```
## func [GetPackageDocumentation](&lt;doc.go#L64>)

```go
func GetPackageDocumentation(packageFilePath, packageImportPath string) (*doc.Package, error)
```
## func [GetTemplateFuncMap](&lt;templates.go#L17>)

```go
func GetTemplateFuncMap() template.FuncMap
```


## type [Const](&lt;doc.go#L149>)
```go
type Const struct {
	Name		string
	Definition	string
	Doc		string
	Filename	string
	Line		int
}
```

## type [Example](&lt;doc.go#L170>)
```go
type Example struct {
	Name		string
	Definition	string
	Doc		string
}
```

## type [Func](&lt;doc.go#L204>)
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

## type [Package](&lt;doc.go#L290>)
```go
type Package struct {
	Name		string
	Definition	string
	Doc		string
	Examples	[]*Example
	Funcs		[]*Func
	Types		[]*Type
	Constants	[]*Const
	Vars		[]*Var
}
```

## func [ParsePackage](&lt;doc.go#L301>)

```go
func ParsePackage(docs *doc.Package) *Package
```

## type [Type](&lt;doc.go#L230>)
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

## type [Var](&lt;doc.go#L128>)
```go
type Var struct {
	Name		string
	Definition	string
	Doc		string
	Filename	string
	Line		int
}
```

