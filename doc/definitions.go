package doc

import (
	"bufio"
	"bytes"
	"errors"
	"go/doc"
	"go/printer"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/projectbadger/autodoc/config"
)

type Var struct {
	Name       string
	Definition string
	Doc        string
	Filename   string
	Line       int
}

func AddVar(data *Package, node *doc.Value, path string) {
	// _, ok := filesets[path]
	// if !ok {
	// 	filesets[path] = token.NewFileSet()
	// }
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
	Doc        string
	Examples   []*Example
	Filename   string
	Line       int
}

func AddFunc(data *Package, node *doc.Func, path string) {
	var buf = bytes.NewBuffer(nil)
	printer.Fprint(buf, filesets[path], node.Decl)
	position := filesets[path].Position(node.Decl.Pos())
	f := &Func{
		Name:       node.Name,
		Doc:        parseComment(0, node.Doc),
		Definition: buf.String(),
		Line:       position.Line,
		Filename:   filepath.Base(position.Filename),
	}
	for _, val := range node.Examples {
		AddFuncExample(f, val, path)
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
	var buf = bytes.NewBuffer(nil)
	position := filesets[path].Position(node.Decl.Pos())
	printer.Fprint(buf, filesets[path], node.Decl)
	data.Funcs = append(data.Funcs, &Func{
		Name:       node.Name,
		Doc:        parseComment(0, node.Doc),
		Definition: buf.String(),
		Line:       position.Line,
		Filename:   filepath.Base(position.Filename),
	})
}

func AddTypeMethod(data *Type, node *doc.Func, path string) {
	var buf = bytes.NewBuffer(nil)
	position := filesets[path].Position(node.Decl.Pos())
	printer.Fprint(buf, filesets[path], node.Decl)
	data.Methods = append(data.Methods, &Func{
		Name:       node.Name,
		Doc:        parseComment(0, node.Doc),
		Definition: buf.String(),
		Line:       position.Line,
		Filename:   filepath.Base(position.Filename),
	})
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
		// fmt.Println("Import:", docs.ImportPath)
	} else {
		SeekGoMod(p, path, 3)
		// fmt.Println("parsed import:", p.ImportPath)
	}
	// fmt.Println("p.Doc:", p.Doc)
	// fmt.Println("docs.ImportPath:", docs.ImportPath)
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
	// fmt.Println("got package", docs.Name, "from path", path)
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
	// for p, val := range packages {
	// 	fmt.Println("Package", val.Name, "from", p)
	// }
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
	// absPathSPlit := strings.Split(absPath, "/")
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
