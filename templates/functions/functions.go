package functions

import (
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/projectbadger/autodoc/config"
)

var headingMDLinkRegex = regexp.MustCompile(`\(\w*|[^a-z0-9-/]+|-{2,}|^/`)

func toMDHeadingHREF(heading string) string {
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

func GetTemplateFuncMap() template.FuncMap {
	return template.FuncMap{
		"headingLink":    toMDHeadingHREF,
		"headingGitLink": toMDHeadingHREFGit,
	}
}