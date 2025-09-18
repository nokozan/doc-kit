package service

import (
	"doc-kit/core/parser"
	"doc-kit/db"
	"doc-kit/models"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gorm.io/gorm"
)

func SyncRepoStructs(dbConn *gorm.DB, gitRepo models.GitRepo) error {

	repoPath := fmt.Sprintf("%s/%s", os.Getenv("REPO_BASE_PATH"), gitRepo.Alias)

	var parsed []*parser.StructMeta
	err := filepath.Walk(
		repoPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil || info.IsDir() || !strings.HasSuffix(info.Name(), ".go") {
				return nil
			}
			structs, err := parser.ExtractStructsFromFile(path)
			if err != nil {
				return fmt.Errorf("failed to extract structs from file %s: %w", path, err)
			}
			parsed = append(parsed, structs...)
			return nil
		})

	if err != nil {
		return fmt.Errorf("error walking the path %s: %w", repoPath, err)
	}

	var allStructs []models.Struct
	for _, p := range parsed {
		allStructs = append(allStructs, MapStructMetaToModel(p))
	}

	return db.SaveStructs(dbConn, allStructs)
}
