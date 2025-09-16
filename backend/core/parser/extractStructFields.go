package parser

import (
	"go/ast"
	"go/token"
)

type FieldMeta struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Tag     string `json:"tag"`
	Comment string `json:"comment"`
}

// convert ast.Expr to string type name (e.g., int, *string, []int)
func exprToType(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		return "*" + exprToType(t.X)
	case *ast.ArrayType:
		return "[]" + exprToType(t.Elt)
	case *ast.SelectorExpr:
		return exprToType(t.X) + "." + t.Sel.Name
	case *ast.MapType:
		return "map[" + exprToType(t.Key) + "]" + exprToType(t.Value)
	case *ast.InterfaceType:
		return "interface{}"
	default:
		return "unknown"
	}
}

// extracts the string value from *ast.BasicLit tag (`json:"id"` -> `json:"id"`)
func extractTag(tag *ast.BasicLit) string {
	if tag == nil || tag.Kind != token.STRING {
		return ""
	}
	return tag.Value
}

// extractStructFields extracts fields from a given struct type.
func extractStructFields(structType *ast.StructType) []FieldMeta {
	var fields []FieldMeta

	for _, field := range structType.Fields.List {
		if len(field.Names) == 0 {
			// Embedded field (anonymous)
			continue
		}

		for _, name := range field.Names {
			fields = append(fields, FieldMeta{
				Name:    name.Name,
				Type:    exprToType(field.Type),
				Tag:     extractTag(field.Tag),
				Comment: extractFieldComment(field),
			})
		}
	}

	return fields
}
