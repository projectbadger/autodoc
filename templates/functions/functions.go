package functions

import (
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/projectbadger/autodoc/config"
)

var headingMDLinkRegex = regexp.MustCompile(`\(\w*|[^a-z0-9-/]+|-{2,}|^/`)
var methodHeadingRegex = regexp.MustCompile(`func\s?(\(\w+\s\*?\w+\))?\s?\w+`)

func toMDHeadingHREF(heading string) string {
	if strings.HasPrefix(heading, "func") {
		heading = "func " + toMethodHeading(heading)
	}
	heading = strings.ToLower(strings.ReplaceAll(heading, " ", "-"))
	return headingMDLinkRegex.ReplaceAllLiteralString(heading, "")
}

func toMDHeadingHREFGit(importPath, heading string) string {
	if heading == "" {
		return ""
	}
	heading = strings.ToLower(strings.ReplaceAll(heading, " ", "-"))
	return filepath.Join(importPath, config.Cfg.Templates.GetLinkPrefix(), headingMDLinkRegex.ReplaceAllLiteralString(heading, ""))
}

func toMethodHeading(definition string) string {
	return strings.TrimSpace(strings.TrimPrefix(methodHeadingRegex.FindString(definition), "func"))
}

func trimSlashes(str string) string {
	return strings.Trim(str, "/")
}

// GetTemplateFuncMap returns the template function map
func GetTemplateFuncMap() template.FuncMap {
	return template.FuncMap{
		"headingLink":    toMDHeadingHREF,
		"headingGitLink": toMDHeadingHREFGit,
		"methodHeading":  toMethodHeading,
		"trimSlashes":    trimSlashes,
	}
}
