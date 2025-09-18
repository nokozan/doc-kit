package service

import (
	"doc-kit/core/parser"
	"doc-kit/models"
)

func MapStructMetaToModel(p *parser.StructMeta) models.Struct {
	fields := make([]models.Field, len(p.Fields))
	for _, f := range p.Fields {
		fields = append(fields, models.Field{
			Name:    f.Name,
			Type:    f.Type,
			Tag:     f.Tag,
			Comment: f.Comment,
		})
	}
	return models.Struct{
		Name:    p.Name,
		Comment: p.Comment,
		Fields:  fields,
	}
}
