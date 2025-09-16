package parser

import (
	"go/ast"
	"regexp"
	"strings"
)

var docRegex = regexp.MustCompile(`desc\s*=\s*"(.*?)"`)

func extractDescriptionFromDoc(text string) string {
	matches := docRegex.FindStringSubmatch(text)
	if len(matches) == 2 {
		return matches[1]
	}
	return ""
}

// parses @Doc(...) annotation from struct comment
// Example: // @Doc(desc="Payment struct for user payments", author="John Doe")
func extractStructDoc(genDecl *ast.GenDecl) string {
	if genDecl.Doc == nil {
		return ""
	}

	for _, comment := range genDecl.Doc.List {
		text := strings.TrimSpace(comment.Text)
		if strings.HasPrefix(text, "// @Doc") {
			return extractDescriptionFromDoc(text)
		}
	}

	return ""
}

// extractStructComment reads comment above the struct
func extractStructComment(decl *ast.GenDecl) string {
	if decl.Doc == nil {
		return ""
	}
	return decl.Doc.Text()
}
