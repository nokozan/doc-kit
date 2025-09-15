package parser

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestExtractStructFields(t *testing.T) {

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "testStruct.go", nil, parser.AllErrors)
	if err != nil {
		t.Fatalf("Failed to parse source: %v", err)
	}

	var structType *ast.StructType
	for _, decl := range node.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.TYPE {
			continue
		}

		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			if st, ok := typeSpec.Type.(*ast.StructType); ok {
				structType = st

			}
		}
	}

	fields := extractStructFields(structType)
	if len(fields) != 4 {
		t.Fatalf("Expected 4 fields, got %d", len(fields))
	}
	if fields[0].Name != "ID" || fields[0].Type != "int" || fields[0].Tag != "`json:\"id\"`" {
		t.Errorf("Field 0 mismatch: %+v", fields[0])
	}
	if fields[1].Name != "Name" || fields[1].Type != "string" || fields[1].Tag != "`json:\"name\"`" {
		t.Errorf("Field 1 mismatch: %+v", fields[1])
	}
	if fields[2].Name != "Other" || fields[2].Type != "*string" || fields[2].Tag != "" {
		t.Errorf("Field 2 mismatch: %+v", fields[2])
	}
	if fields[3].Name != "Items" || fields[3].Type != "[]int" || fields[3].Tag != "" {
		t.Errorf("Field 3 mismatch: %+v", fields[3])
	}
}
