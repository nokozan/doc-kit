package parser

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestExtractStruct(t *testing.T) {

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "testStruct.go", nil, parser.ParseComments)
	if err != nil {
		t.Fatalf("Failed to parse source: %v", err)
	}

	var meta *StructMeta
	for _, decl := range node.Decls {
		meta = extractStruct(decl)
		if meta != nil {
			break
		}

	}

	if meta == nil {
		t.Fatalf("Struct 'MyTestStruct' not found")
		return
	}

	if meta.Name != "MyTestStruct" {
		t.Errorf("Expected struct name 'MyTestStruct', got '%s'", meta.Name)
	}
	if meta.Comment != `@Doc(desc="MyTestStruct is a sample struct for testing", author="Jane Doe")
MyTestStruct is a sample struct for testing
` {
		t.Errorf("Expected doc 'MyTestStruct is a sample struct for testing', got '%s'", meta.Comment)
	}
	if len(meta.Fields) != 4 {
		t.Fatalf("Expected 4 fields, got %d", len(meta.Fields))
	}

	expectedFields := map[string]string{
		"ID":    "ID is the unique identifier",
		"Name":  "Name is the name of the entity",
		"Other": "Other is an optional field",
		"Items": "Items is a list of integers(only can above field)",
	}

	for _, field := range meta.Fields {
		if desc, ok := expectedFields[field.Name]; !ok {
			t.Errorf("Unexpected field: %+v", field)
		} else if field.Comment != desc {
			t.Errorf("Field %s: expected comment '%s', got '%s'", field.Name, desc, field.Comment)
		}
	}
}
