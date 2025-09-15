package parser

import (
	"go/ast"
	"go/token"
)

// Only handles type Xxx Struct {...}
// returns "" if not found
// "user" for type user struct {...}
func extractStructName(decl ast.Decl) string {
	genDecl, ok := decl.(*ast.GenDecl)
	if !ok || genDecl.Tok != token.TYPE || len(genDecl.Specs) == 0 {
		return ""
	}

	for _, spec := range genDecl.Specs {
		typeSpec, ok := spec.(*ast.TypeSpec)
		if !ok {
			continue
		}

		if _, ok := typeSpec.Type.(*ast.StructType); ok {
			return typeSpec.Name.Name
		}
	}

	return ""

}
