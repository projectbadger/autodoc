package doc

import (
	"bufio"
	"bytes"
	"fmt"
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

var (
	filesets = make(map[string]*token.FileSet)
)

func GetGoFiles(path string) []*ast.File {
	var astFiles []*ast.File
	_, ok := filesets[path]
	if !ok {
		filesets[path] = token.NewFileSet()
	}

	filepath.WalkDir(path, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			fmt.Println("error (start):", e)
			return e
		}
		if d.IsDir() {
			return nil
		}
		if filepath.Ext(d.Name()) == ".go" {
			newFile, err := parser.ParseFile(filesets[path], s, nil, parser.ParseComments|parser.AllErrors)
			if err != nil {
				fmt.Println("error:", err)
				return err
			}
			astFiles = append(astFiles, newFile)
			// fmt.Printf("appended: '%#v'\n", d.Name())

		}
		return nil
	})
	return astFiles
}

func GetGoFilesInDir(path string) []*ast.File {
	var astFiles []*ast.File
	_, ok := filesets[path]
	if !ok {
		filesets[path] = token.NewFileSet()
	}
	dirFiles, err := os.ReadDir(path)
	if err != nil {
		return nil
	}
	for _, d := range dirFiles {
		if d.IsDir() {
			continue
		}
		if filepath.Ext(d.Name()) == ".go" {
			newFile, err := parser.ParseFile(filesets[path], filepath.Join(path, d.Name()), nil, parser.ParseComments|parser.AllErrors)
			if err != nil {
				fmt.Println("error:", err)
				return nil
			}
			astFiles = append(astFiles, newFile)
			// fmt.Printf("appended: '%#v'\n", d.Name())

		}
	}
	return astFiles
}

var (
	blacklistedDirNames = []string{
		"node_modules",
		".git",
		"vendor",
	}
)

func GetDirectories(path string) []string {
	var dirs []string

	filepath.WalkDir(path, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			fmt.Println("error (start):", e)
			return e
		}
		if d.IsDir() && d.Name()[0] != '.' {
			for _, b := range blacklistedDirNames {
				if b == d.Name() {
					return nil
				}
			}
			dirs = append(dirs, filepath.Join(path, d.Name()))
			return nil
		}
		return nil
	})
	return dirs
}

func GetPackageDocumentation(packageFilePath, packageImportPath string) (*doc.Package, error) {
	_, ok := filesets[packageFilePath]
	if !ok {
		filesets[packageFilePath] = token.NewFileSet()
	}
	files := GetGoFilesInDir(packageFilePath)
	// fmt.Printf("Got files: '%#v', files, from %s", files, packageFilePath)
	return doc.NewFromFiles(filesets[packageFilePath], files, packageImportPath)
}

func parseComment(format uint, comment string) string {
	buf := bytes.NewBuffer(nil)
	textBuf := bytes.NewReader([]byte(comment))
	scanner := bufio.NewScanner(textBuf)
	isCode := false
	isCodePrev := false
	codeIndentPrefix := "\t"
	for scanner.Scan() {
		isCodePrev = isCode
		line := scanner.Text()
		if len(line) > 0 && (line[0] == '\t' || line[0] == ' ') {
			isCode = true
		} else {
			isCode = false
		}
		if isCode && !isCodePrev {
			// Code start
			codeIndentPrefix = line[:1]
			lineLen := len(line)
			for i := 0; i < lineLen; i++ {
				if line[i] != codeIndentPrefix[0] {
					codeIndentPrefix = line[:i]
					break
				}
			}
			buf.WriteString("```go\n")
			buf.WriteString(strings.TrimPrefix(line, codeIndentPrefix))
		} else if !isCode && isCodePrev {
			// Code end
			buf.WriteString(strings.TrimPrefix(line, codeIndentPrefix))
			buf.WriteString("\n```\n")
		} else {
			buf.WriteString(line)
			buf.WriteString("\n")
		}
	}
	if isCode {
		buf.WriteString("\n```")
	}
	return strings.Trim(buf.String(), "\n")
}
