package parser

import (
	"go/parser"
	"go/token"
	"os"
)

func ExtractStructsFromFile(path string) ([]*StructMeta, error) {
	fset := token.NewFileSet()

	src, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	node, err := parser.ParseFile(fset, path, src, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	var structs []*StructMeta
	for _, decl := range node.Decls {
		if meta := extractStruct(decl); meta != nil {
			structs = append(structs, meta)
		}
	}

	return structs, nil
}
