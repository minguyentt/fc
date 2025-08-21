package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/minguyentt/fc/internal/filecopy"
)

func main() {
	root, err := os.Getwd()
	if err != nil {
		fmt.Printf("failed to get current working directory: %v\n", err)
	}

	src := flag.String("src", "", "source path")
	dst := flag.String("dst", "", "destination path")
	pat := flag.String("p", "", "pattern to match (optional)")
	flag.Parse()

	if *src == "" || *dst == "" {
		log.Fatalf("usage: %s -src=<source> -dst=<dest> [-p=<pattern>]", flag.CommandLine.Name())
	}

	// build the file path to include the home directory for user
	parts := strings.Split(filepath.Clean(root), string(filepath.Separator))
	home := filepath.Join(string(filepath.Separator), parts[1], parts[2])

	srcPath := filepath.Join(home, *src)
	dstPath := filepath.Join(home, *dst)

	copied := 0
	err = filecopy.VisitWithMatch(srcPath, dstPath, *pat, &copied)
	if err != nil {
		fmt.Println("error encountered for file copying: %w\n", err)
	}

	fmt.Printf("Recursive directory walk count: %d", copied)

}
