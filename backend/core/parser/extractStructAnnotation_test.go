package parser

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestExtractStructDoc(t *testing.T) {

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "testStruct.go", nil, parser.ParseComments)
	if err != nil {
		t.Fatalf("Failed to parse source: %v", err)
	}

	var doc string
	for _, decl := range node.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}

		doc = extractStructDoc(genDecl)
	}
	expected := "MyTestStruct is a sample struct for testing"
	if doc != expected {
		t.Errorf("Expected doc '%s', got '%s'", expected, doc)
	} else {
		// t.Logf("Extracted doc: %s", doc)
	}
}
