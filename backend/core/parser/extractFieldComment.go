package parser

import (
	"go/ast"
	"strings"
)

// extracts the first comment found beside the field or above the field
// returns "" if no comment found
func extractFieldComment(field *ast.Field) string {
	var commentGroup *ast.CommentGroup

	if field.Doc != nil {
		commentGroup = field.Doc
	} else if field.Comment != nil {
		commentGroup = field.Comment
	}

	if commentGroup != nil {
		for _, comment := range commentGroup.List {
			text := strings.TrimSpace(comment.Text)
			// remove leading // or /* */
			if strings.HasPrefix(text, "//") {
				return strings.TrimPrefix(text, "//")
			}
		}
	}
	return ""
}
