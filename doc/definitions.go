package doc

import (
	"bufio"
	"bytes"
	"errors"
	"go/ast"
	"go/doc"
	"go/printer"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/projectbadger/autodoc/config"
	"github.com/projectbadger/autodoc/templates/md"
)

type Var struct {
	Name       string
	Definition string
	Doc        string
	Filename   string
	Line       int
}

func AddVar(data *Package, node *doc.Value, path string) {
	var buf = bytes.NewBuffer(nil)
	printer.Fprint(buf, filesets[path], node.Decl)
	position := filesets[path].Position(node.Decl.Pos())
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

func AddConst(data *Package, node *doc.Value, path string) {
	var buf = bytes.NewBuffer(nil)
	position := filesets[path].Position(node.Decl.Pos())
	printer.Fprint(buf, filesets[path], node.Decl)
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

func AddExample(data *Package, node *doc.Example, path string) {
	var buf = bytes.NewBuffer(nil)
	printer.Fprint(buf, filesets[path], node.Code)
	data.Examples = append(data.Examples, &Example{
		Name:       node.Name,
		Doc:        parseComment(0, node.Doc),
		Definition: buf.String(),
	})
}
func AddFuncExample(data *Func, node *doc.Example, path string) {
	var buf = bytes.NewBuffer(nil)
	printer.Fprint(buf, filesets[path], node.Code)
	data.Examples = append(data.Examples, &Example{
		Name:       node.Name,
		Doc:        parseComment(0, node.Doc),
		Definition: buf.String(),
	})
}
func AddTypeExample(data *Type, node *doc.Example, path string) {
	var buf = bytes.NewBuffer(nil)
	printer.Fprint(buf, filesets[path], node.Code)
	data.Examples = append(data.Examples, &Example{
		Name:       node.Name,
		Doc:        parseComment(0, node.Doc),
		Definition: buf.String(),
	})
}

type Func struct {
	Name       string
	Definition string
	Recv       FuncVar
	Params     []FuncVar
	Results    []FuncVar
	Doc        string
	Examples   []*Example
	Filename   string
	Line       int
}

type FuncVar struct {
	Name    string
	Type    string
	Pointer bool
}

func (f *Func) FormatParams() string {
	str := ""
	noTypeCount := 0
	for _, param := range f.Params {
		if param.Type != "" {
			str += param.Type + ", "
			if noTypeCount > 0 {
				for ; noTypeCount > 0; noTypeCount-- {
					str += param.Type + ", "
				}
			}
			continue
		}
		noTypeCount++
	}
	return strings.TrimRight(str, ", ")
}

func (f *Func) FormatParamsBrackets() string {
	return "(" + f.FormatResults() + ")"
}

func (f *Func) FormatResults() string {
	str := ""
	if len(f.Results) == 1 {
		return f.Results[0].Type
	}
	noTypeCount := 0
	for _, result := range f.Results {
		if result.Type != "" {
			str += result.Type + ", "
			if noTypeCount > 0 {
				for ; noTypeCount > 0; noTypeCount-- {
					str += result.Type + ", "
				}
			}
			continue
		}
		noTypeCount++
	}
	return strings.TrimRight(str, ", ")
}

func (f *Func) FormatResultsBrackets() string {
	if len(f.Results) < 2 {
		return f.FormatResults()
	}
	return "(" + f.FormatResults() + ")"
}

var matchChars = regexp.MustCompile(`[^a-z0-9-/]+|-+`)
var matchFuncDefinition = regexp.MustCompile(`func\s*(\([\w\s*]+\))?\s*\[(.+)\]\(`)
var matchTypeDefinition = regexp.MustCompile(`type\s*\[(.+)\]\(`)

func getHeadingHREF(str string) string {
	str = strings.ToLower(str)
	return strings.Trim(matchChars.ReplaceAllLiteralString(str, "-"), "-")
}

func (f *Func) GetHeadingHREF() string {
	var b bytes.Buffer
	err := md.TemplateDoc.ExecuteTemplate(&b, "functionHeading", f)
	if err != nil {
		return err.Error()
	}
	def := matchFuncDefinition.FindStringSubmatch(b.String())
	return getHeadingHREF(def[0])
}

func NewFunc(node *doc.Func, path string) *Func {
	var buf = bytes.NewBuffer(nil)
	printer.Fprint(buf, filesets[path], node.Decl)
	position := filesets[path].Position(node.Decl.Pos())
	var params []FuncVar
	for _, val := range node.Decl.Type.Params.List {
		var funcVar FuncVar
		var typeName string
		var isPointer bool
		if se, ok := val.Type.(*ast.SelectorExpr); ok {
			typeName = se.Sel.Name
		}
		if ident, ok := val.Type.(*ast.Ident); ok {
			typeName = ident.Name
		}
		if val.Names != nil {
			for _, name := range val.Names {
				funcVar.Name = name.Name
				params = append(params, FuncVar{
					Type:    typeName,
					Name:    name.Name,
					Pointer: isPointer,
				})
			}
			continue
		}
		params = append(params, FuncVar{
			Type:    typeName,
			Name:    "",
			Pointer: isPointer,
		})
	}
	var results []FuncVar
	if node.Decl.Type.Results != nil && node.Decl.Type.Results.List != nil {
		for _, val := range node.Decl.Type.Results.List {
			var funcVar FuncVar
			var typeName string
			var isPointer bool
			if se, ok := val.Type.(*ast.SelectorExpr); ok {
				if ident, ok := se.X.(*ast.Ident); ok {
					typeName = ident.Name
					isPointer = true
				}
			}
			if se, ok := val.Type.(*ast.StarExpr); ok {
				if ident, ok := se.X.(*ast.Ident); ok {
					typeName = ident.Name
					isPointer = true
				}
			}
			if ident, ok := val.Type.(*ast.Ident); ok {
				typeName = ident.Name
			}
			if typeName == "" {
				continue
			}
			if val.Names != nil {
				for _, name := range val.Names {
					funcVar.Name = name.Name
					results = append(results, FuncVar{
						Type:    typeName,
						Name:    name.Name,
						Pointer: isPointer,
					})
				}
				continue
			}
			results = append(results, FuncVar{
				Type:    typeName,
				Name:    "",
				Pointer: isPointer,
			})
		}
	}
	var recv FuncVar
	if node.Decl.Recv != nil && node.Decl.Recv.List != nil {
		recvTypeName := ""
		if recvType, ok := node.Decl.Recv.List[0].Type.(*ast.Ident); ok {
			recvTypeName = recvType.Name
		}
		// if recvType, ok := node.Decl.Recv.List[0].Type.(*ast.SelectorExpr); ok {
		// 	recvTypeName = recvType.Sel.Name
		// }
		if se, ok := node.Decl.Recv.List[0].Type.(*ast.StarExpr); ok {
			if ident, ok := se.X.(*ast.Ident); ok {
				recvTypeName = ident.Name
				recv.Pointer = true
			}
		}
		if se, ok := node.Decl.Recv.List[0].Type.(*ast.SelectorExpr); ok {
			if ident, ok := se.X.(*ast.Ident); ok {
				recvTypeName = ident.Name
				recv.Pointer = true
			}
		}
		recv.Type = recvTypeName
		if node.Decl.Recv.List[0].Names != nil {
			recv.Name = node.Decl.Recv.List[0].Names[0].Name
		}
	}
	f := &Func{
		Name:       node.Decl.Name.Name,
		Doc:        parseComment(0, node.Doc),
		Params:     params,
		Results:    results,
		Recv:       recv,
		Definition: buf.String(),
		Line:       position.Line,
		Filename:   filepath.Base(position.Filename),
	}
	for _, val := range node.Examples {
		AddFuncExample(f, val, path)
	}
	return f
}

func AddFunc(data *Package, node *doc.Func, path string) {
	f := NewFunc(node, path)
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

func (t *Type) GetHeadingHREF() string {
	var b bytes.Buffer
	err := md.TemplateDoc.ExecuteTemplate(&b, "typeHeading", t)
	if err != nil {
		return err.Error()
	}
	def := matchTypeDefinition.FindStringSubmatch(b.String())
	return getHeadingHREF(def[0])
}

func AddType(data *Package, node *doc.Type, path string) {
	var buf = bytes.NewBuffer(nil)
	position := filesets[path].Position(node.Decl.Pos())
	printer.Fprint(buf, filesets[path], node.Decl)
	t := &Type{
		Name:       node.Name,
		Doc:        parseComment(0, node.Doc),
		Definition: buf.String(),
		Line:       position.Line,
		Filename:   filepath.Base(position.Filename),
	}
	for _, val := range node.Funcs {
		AddTypeFunc(t, val, path)
	}
	for _, val := range node.Methods {
		AddTypeMethod(t, val, path)
	}
	for _, val := range node.Examples {
		AddTypeExample(t, val, path)
	}
	data.Types = append(data.Types, t)
}

func AddTypeFunc(data *Type, node *doc.Func, path string) {
	// var buf = bytes.NewBuffer(nil)
	// position := filesets[path].Position(node.Decl.Pos())
	// printer.Fprint(buf, filesets[path], node.Decl)
	data.Funcs = append(data.Funcs, NewFunc(node, path))
}

func AddTypeMethod(data *Type, node *doc.Func, path string) {
	data.Methods = append(data.Methods, NewFunc(node, path))

}

type Package struct {
	ImportPath      string
	Name            string
	Definition      string
	Doc             string
	Subpackages     map[string]*Package
	Examples        []*Example
	Funcs           []*Func
	Types           []*Type
	Constants       []*Const
	Vars            []*Var
	CustomVars      map[string]string
	Path            string
	PathAbs         string
	ShowName        bool
	ShowDoc         bool
	ShowExamples    bool
	ShowIndex       bool
	ShowFuncs       bool
	ShowTypes       bool
	ShowConsts      bool
	ShowVars        bool
	ShowSubpackages bool
	ShowOverview    bool
	ShowImportPath  bool
}

func getCustomVars() map[string]string {
	vars := make(map[string]string)
	if len(config.Cfg.Templates.CustomVars) > 0 {
		for _, pair := range config.Cfg.Templates.CustomVars {
			split := strings.Split(pair, "=")
			switch len(split) {
			case 1:
				vars[split[0]] = split[0]
				// vars[split[0]] = "true"
			case 2:
				vars[split[0]] = split[1]
			default:
				vars[split[0]] = strings.Join(split[1:], "=")
			}
		}
	}
	r := regexp.MustCompile(`{{\s*[a-zA-Z0-9_-]+\s*}}`)
	rVarName := regexp.MustCompile(`[a-zA-Z0-9_-]+`)
	for key, val := range vars {
		vars[key] = r.ReplaceAllStringFunc(val, func(match string) string {
			names := rVarName.FindAllString(match, -1)
			if len(names) == 0 {
				return match
			}
			value, ok := vars[names[0]]
			if ok {
				return value
			}
			return match
		})
	}
	return vars
}

func ParsePackage(docs *doc.Package, path string) (*Package, error) {
	if docs == nil {
		return nil, errors.New("doc is nil")
	}
	p := &Package{
		Name:       docs.Name,
		Doc:        parseComment(0, docs.Doc),
		CustomVars: getCustomVars(),
	}
	if docs.ImportPath != "" {
		p.ImportPath = docs.ImportPath
	} else {
		SeekGoMod(p, path, 3)
	}
	for _, val := range docs.Examples {
		AddExample(p, val, path)
	}
	for _, val := range docs.Funcs {
		AddFunc(p, val, path)
	}
	for _, val := range docs.Types {
		AddType(p, val, path)
	}
	for _, val := range docs.Vars {
		AddVar(p, val, path)
	}
	for _, val := range docs.Consts {
		AddConst(p, val, path)
	}
	for _, v := range config.Cfg.Templates.IncludedData {
		switch v {
		case "all", "a":
			p.ShowName = true
			p.ShowDoc = true
			p.ShowExamples = true
			p.ShowIndex = true
			p.ShowFuncs = true
			p.ShowTypes = true
			p.ShowConsts = true
			p.ShowVars = true
			p.ShowSubpackages = true
			p.ShowOverview = true
			p.ShowImportPath = true
			// return p, nil
		case "name", "n":
			p.ShowName = true
		case "-name", "-n":
			p.ShowName = false
		case "doc", "d":
			p.ShowDoc = true
		case "-doc", "-d":
			p.ShowDoc = false
		case "examples", "ex":
			p.ShowExamples = true
		case "-examples", "-ex":
			p.ShowExamples = false
		case "index", "i":
			p.ShowIndex = true
		case "-index", "-i":
			p.ShowIndex = false
		case "functions", "f":
			p.ShowFuncs = true
		case "-functions", "-f":
			p.ShowFuncs = false
		case "types", "t":
			p.ShowTypes = true
		case "-types", "-t":
			p.ShowTypes = false
		case "constants", "c":
			p.ShowConsts = true
		case "-constants", "-c":
			p.ShowConsts = false
		case "variables", "v":
			p.ShowVars = true
		case "-variables", "-v":
			p.ShowVars = false
		case "subpackages", "s":
			p.ShowSubpackages = true
		case "-subpackages", "-s":
			p.ShowSubpackages = false
		case "overview", "o":
			p.ShowOverview = true
		case "-overview", "-o":
			p.ShowOverview = false
		case "importpath", "import-path", "import":
			p.ShowImportPath = true
		case "-importpath", "-import-path", "-import":
			p.ShowImportPath = false
		}
	}

	return p, nil
}

func GetPackageDataFromDirRecursive(path string) (*Package, error) {

	p, err := GetPackageDataFromDir(path)
	if err != nil {
		return nil, err
	}
	// SeekGoMod(p, path, 3)
	p.Subpackages, err = GetPackagesDataFromDirRecursive(path, false, p.ImportPath)
	return p, err
}

func GetPackageDataFromDir(path string) (*Package, error) {
	docs, err := GetPackageDocumentation(path, config.Cfg.Templates.ImportPath)
	if err != nil {
		return nil, err
	}
	return ParsePackage(docs, path)
}

func GetPackagesDataFromDirRecursive(dirPath string, includeRoot bool, rootImportPath string) (map[string]*Package, error) {
	dirPath, err := filepath.Abs(dirPath)
	if err != nil {
		return nil, err
	}
	packages := make(map[string]*Package)
	filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(info.Name()) == ".go" {
			relPath := strings.TrimPrefix(filepath.Dir(path), dirPath)
			if relPath == "" && !includeRoot {
				return nil
			}
			if _, ok := packages[relPath]; !ok {
				packages[relPath], err = GetPackageDataFromDir(filepath.Join(dirPath, relPath))
				if rootImportPath != "" {
					packages[relPath].ImportPath = filepath.Join(rootImportPath, relPath)
				}
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
	return packages, nil
}

func ParseGoMod(pkg *Package, path string) error {
	file, err := os.Open(filepath.Join(path, "go.mod"))
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "module ") {
			pkg.ImportPath = strings.Trim(strings.TrimPrefix(line, "module "), " \n\r\t")
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func SeekGoMod(pkg *Package, path string, levels int) error {
	path, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	rel := ""
	for i := 1; i <= levels; i++ {
		err = ParseGoMod(pkg, path)
		if err == nil {
			pkg.ImportPath += "/" + rel
			break
		}
		if rel == "" {
			rel = filepath.Base(path)
		} else {
			rel = filepath.Base(path) + "/" + rel
		}
		path = filepath.Dir(path)
	}
	return nil
}
