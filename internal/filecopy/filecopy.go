package filecopy

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)


func join(elem ...string) string {
	return filepath.Join(elem...)
}

func createDir(dir string) error {
	err := os.Mkdir(dir, 0750)
	if err != nil && !os.IsExist(err) {
		return err
	}

	return nil
}

// copies the contents from src path file to the destination path
func CopyFiles(src string, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open file from src: %w\n", err)
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w\n", err)
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return fmt.Errorf("failed to copy file contents: %w\n", err)
	}

	fmt.Printf("Copying %s to %s\n", src, dst)
	return out.Close()
}

func VisitWithMatch(src string, dst string, fName string, copyCount *int) error {
	return filepath.WalkDir(src, func(currPath string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("error encountered at path %s: %v\n", currPath, err)
			return err
		}

		// skip every directory except parent directory
		if d.IsDir() && src != currPath {
			fmt.Printf("skipping sub-directory: %s\n", currPath)
			return filepath.SkipDir
		}

		if !d.IsDir() && strings.Contains(d.Name(), fName) {
			*copyCount++
			fmt.Printf("match found, path: %s\n", d.Name())

			// join the current file name to the destination path
			fullDstPath := join(dst, d.Name())
			return CopyFiles(currPath, fullDstPath)
		}

		return nil
	})
}

//note: use these funcs for testing scenarios
func CreateDummyFiles() {
	os.MkdirAll("./test_dir/subdir1", 0755)
	os.MkdirAll("./test_dir/subdir_to_skip", 0755)
	os.WriteFile("./test_dir/file1.txt", []byte("content1"), 0644)
	os.WriteFile("./test_dir/subdir1/file2.txt", []byte("content2"), 0644)
	os.WriteFile("./test_dir/subdir_to_skip/skipped_file.txt", []byte("content_skipped"), 0644)
}

func CleanupDummyFiles() {
	os.RemoveAll("./test_dir")
}

// callback function for each file or dir encountered
func Visited(path string, d fs.DirEntry, err error) error {
	if err != nil {
		fmt.Printf("error encountered at path %s: %v\n", path, err)
		return err
	}

	// skip every directory in the src path
	if d.IsDir() {
		fmt.Printf(" (Directory)\n")
		// Example: Skip a specific subdirectory
		if d.Name() == "subdir_to_skip" {
			fmt.Printf("Skipping directory: %s\n", path)
			return filepath.SkipDir
		}
	}
	fmt.Printf("Visited: %s (Is Directory: %t)\n", path, d.IsDir())

	return nil
}
