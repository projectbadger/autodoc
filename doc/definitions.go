package doc

import (
	"bytes"
	"go/doc"
	"go/printer"
	"path/filepath"
)

type Var struct {
	Name       string
	Definition string
	Doc        string
	Filename   string
	Line       int
}

func AddVar(data *Package, node *doc.Value) {
	var buf = bytes.NewBuffer(nil)
	printer.Fprint(buf, fileset, node.Decl)
	position := fileset.Position(node.Decl.Pos())
	data.Vars = append(data.Vars, &Var{
		Name:       node.Names[0],
		Doc:        parseComment(0, node.Doc),
		Definition: buf.String(),
		Line:       position.Line,
		Filename:   filepath.Base(position.Filename),
	})
}

type Const struct {
	Name       string
	Definition string
	Doc        string
	Filename   string
	Line       int
}

func AddConst(data *Package, node *doc.Value) {
	var buf = bytes.NewBuffer(nil)
	position := fileset.Position(node.Decl.Pos())
	printer.Fprint(buf, fileset, node.Decl)
	data.Vars = append(data.Vars, &Var{
		Name:       node.Names[0],
		Doc:        parseComment(0, node.Doc),
		Definition: buf.String(),
		Line:       position.Line,
		Filename:   filepath.Base(position.Filename),
	})
}

type Example struct {
	Name       string
	Definition string
	Doc        string
}

func AddExample(data *Package, node *doc.Example) {
	var buf = bytes.NewBuffer(nil)
	printer.Fprint(buf, fileset, node.Code)
	data.Examples = append(data.Examples, &Example{
		Name:       node.Name,
		Doc:        parseComment(0, node.Doc),
		Definition: buf.String(),
	})
}
func AddFuncExample(data *Func, node *doc.Example) {
	var buf = bytes.NewBuffer(nil)
	printer.Fprint(buf, fileset, node.Code)
	data.Examples = append(data.Examples, &Example{
		Name:       node.Name,
		Doc:        parseComment(0, node.Doc),
		Definition: buf.String(),
	})
}
func AddTypeExample(data *Type, node *doc.Example) {
	var buf = bytes.NewBuffer(nil)
	printer.Fprint(buf, fileset, node.Code)
	data.Examples = append(data.Examples, &Example{
		Name:       node.Name,
		Doc:        parseComment(0, node.Doc),
		Definition: buf.String(),
	})
}

type Func struct {
	Name       string
	Definition string
	Doc        string
	Examples   []*Example
	Filename   string
	Line       int
}

func AddFunc(data *Package, node *doc.Func) {
	var buf = bytes.NewBuffer(nil)
	printer.Fprint(buf, fileset, node.Decl)
	position := fileset.Position(node.Decl.Pos())
	f := &Func{
		Name:       node.Name,
		Doc:        parseComment(0, node.Doc),
		Definition: buf.String(),
		Line:       position.Line,
		Filename:   filepath.Base(position.Filename),
	}
	for _, val := range node.Examples {
		AddFuncExample(f, val)
	}
	data.Funcs = append(data.Funcs, f)
}

type Type struct {
	Name       string
	Definition string
	Doc        string
	Examples   []*Example
	Methods    []*Func
	Funcs      []*Func
	Filename   string
	Line       int
}

func AddType(data *Package, node *doc.Type) {
	var buf = bytes.NewBuffer(nil)
	position := fileset.Position(node.Decl.Pos())
	printer.Fprint(buf, fileset, node.Decl)
	t := &Type{
		Name:       node.Name,
		Doc:        parseComment(0, node.Doc),
		Definition: buf.String(),
		Line:       position.Line,
		Filename:   filepath.Base(position.Filename),
	}
	for _, val := range node.Funcs {
		AddTypeFunc(t, val)
	}
	for _, val := range node.Methods {
		AddTypeMethod(t, val)
	}
	for _, val := range node.Examples {
		AddTypeExample(t, val)
	}
	data.Types = append(data.Types, t)
}

func AddTypeFunc(data *Type, node *doc.Func) {
	var buf = bytes.NewBuffer(nil)
	position := fileset.Position(node.Decl.Pos())
	printer.Fprint(buf, fileset, node.Decl)
	data.Funcs = append(data.Funcs, &Func{
		Name:       node.Name,
		Doc:        parseComment(0, node.Doc),
		Definition: buf.String(),
		Line:       position.Line,
		Filename:   filepath.Base(position.Filename),
	})
}

func AddTypeMethod(data *Type, node *doc.Func) {
	var buf = bytes.NewBuffer(nil)
	position := fileset.Position(node.Decl.Pos())
	printer.Fprint(buf, fileset, node.Decl)
	data.Methods = append(data.Methods, &Func{
		Name:       node.Name,
		Doc:        parseComment(0, node.Doc),
		Definition: buf.String(),
		Line:       position.Line,
		Filename:   filepath.Base(position.Filename),
	})
}

type Package struct {
	Name       string
	Definition string
	Doc        string
	Examples   []*Example
	Funcs      []*Func
	Types      []*Type
	Constants  []*Const
	Vars       []*Var
}

func ParsePackage(docs *doc.Package) *Package {
	p := &Package{
		Name: docs.Name,
		Doc:  parseComment(0, docs.Doc),
	}
	for _, val := range docs.Examples {
		AddExample(p, val)
	}
	for _, val := range docs.Funcs {
		AddFunc(p, val)
	}
	for _, val := range docs.Types {
		AddType(p, val)
	}
	for _, val := range docs.Consts {
		AddConst(p, val)
	}
	for _, val := range docs.Vars {
		AddVar(p, val)
	}

	return p
}
