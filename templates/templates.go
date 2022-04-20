// Documentation templates
package templates

import (
	"regexp"
	"strings"
	"text/template"
)

var headingMDLinkRegex = regexp.MustCompile(`\(\w*|[^a-z0-9-]+|-{2,}`)

func toMDHeadingHREF(heading string) string {
	heading = strings.ToLower(strings.ReplaceAll(heading, " ", "-"))
	return headingMDLinkRegex.ReplaceAllLiteralString(heading, "")
}

func GetTemplateFuncMap() template.FuncMap {
	return template.FuncMap{
		"headingLink": toMDHeadingHREF,
	}
}
