package db

import (
	"doc-kit/models"

	"gorm.io/gorm"
)

// replace if exists
func SaveStructs(db *gorm.DB, metas []models.Struct) error {
	for _, meta := range metas {
		structEntity := models.Struct{
			Name:    meta.Name,
			Comment: meta.Comment,
		}
		err := db.Where(
			"name = ?", meta.Name).
			Assign(structEntity).
			FirstOrCreate(&structEntity).Error
		if err != nil {
			return err
		}

		// Clear existing fields
		err = db.Where("struct_id = ?", structEntity.ID).Delete(&models.Field{}).Error
		if err != nil {
			return err
		}

		// Add new fields
		for _, fieldMeta := range meta.Fields {
			fieldEntity := models.Field{
				StructID: structEntity.ID,
				Name:     fieldMeta.Name,
				Type:     fieldMeta.Type,
				Tag:      fieldMeta.Tag,
				Comment:  fieldMeta.Comment,
			}
			if err := db.Create(&fieldEntity).Error; err != nil {
				return err
			}
		}
	}
	return nil
}
