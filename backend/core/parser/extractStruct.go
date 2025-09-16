package parser

import (
	"go/ast"
	"go/token"
)

type StructMeta struct {
	Name    string      `json:"name"`
	Comment string      `json:"comment"`
	Fields  []FieldMeta `json:"fields"`
}

func extractStruct(decl ast.Decl) *StructMeta {
	genDecl, ok := decl.(*ast.GenDecl)
	if !ok || genDecl.Tok != token.TYPE || len(genDecl.Specs) == 0 {
		return nil
	}

	for _, spec := range genDecl.Specs {

		typeSpec, ok := spec.(*ast.TypeSpec)
		if !ok {
			continue
		}

		structType, ok := typeSpec.Type.(*ast.StructType)
		if !ok {
			continue
		}

		name := extractStructName(decl)
		comment := extractStructComment(genDecl)
		fields := extractStructFields(structType)

		return &StructMeta{
			Name:    name,
			Comment: comment,
			Fields:  fields,
		}
	}

	return nil
}
