package parser

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestExtractFieldComment(t *testing.T) {

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "testStruct.go", nil, parser.ParseComments)
	if err != nil {
		t.Fatalf("Failed to parse source: %v", err)
	}

	var structType *ast.StructType
	for _, decl := range node.Decls {
		gen, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		for _, spec := range gen.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if ok {
				if st, ok := typeSpec.Type.(*ast.StructType); ok {
					structType = st

				}
			}
		}
	}
	if structType == nil {
		t.Fatalf("Struct 'Sample' not found")
	}
	fields := structType.Fields.List

	tests := []struct {
		index           int
		expectedComment string
	}{
		{0, "ID is the unique identifier"},
		{1, "Name is the name of the entity"},
		{2, "Other is an optional field"},
		{3, "Items is a list of integers(only can above field)"},
	}

	for _, test := range tests {
		comment := extractFieldComment(fields[test.index])
		if comment != test.expectedComment {
			t.Errorf("Field %d: expected comment '%s', got '%s'", test.index, test.expectedComment, comment)
		} else {
			// t.Logf("Field %d: extracted comment: %s", test.index, comment)
		}
	}
}
