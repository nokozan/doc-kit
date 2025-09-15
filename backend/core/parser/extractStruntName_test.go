package parser

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestExtractStructName(t *testing.T) {

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "testStruct.go", nil, parser.AllErrors)
	if err != nil {
		t.Fatalf("Failed to parse source: %v", err)
	}

	for _, decl := range node.Decls {
		structName := extractStructName(decl)
		if structName == "MyTestStruct" {
			t.Logf("Extracted struct name: %s", structName)
			return
		}
	}

	t.Errorf("Struct name not found or incorrect")
}
