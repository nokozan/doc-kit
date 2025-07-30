package git_utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

const RepoBaseDir = "./data/repos"

func CloneRepo(repoURL, alias, branch string) error {

	//Full local path: ./data/repos/alias
	destPath := filepath.Join(RepoBaseDir, alias)

	// If exists, delete first(for now - later can add "update" mode)
	if _, err := os.Stat(destPath); err == nil {
		if err := os.RemoveAll(destPath); err != nil {
			return fmt.Errorf("failed to remove existing repo: %w", err)
		}
	}

	//Run: git clone <url> <dest>
	cmd := exec.Command("git", "clone", repoURL, destPath)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("git clone failed: %s, output: %s", err.Error(), out)
	}

	return nil
}
