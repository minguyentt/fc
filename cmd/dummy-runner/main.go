package main

import (
	"fmt"
	"path/filepath"

	"github.com/minguyentt/fc/internal/filecopy"
)

func main() {
	filecopy.CreateDummyFiles()

	dummyPath := "./test_dir"

	err := filepath.WalkDir(dummyPath, filecopy.Visited)
	if err != nil {
		fmt.Printf("error walking directory tree: %v\n", err)
	}

	filecopy.CleanupDummyFiles()
}
