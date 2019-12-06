package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func walKDir(dir string, fileSize chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subDir := filepath.Join(dir, entry.Name())
			walKDir(subDir, fileSize)
		} else {
			fileSize <- entry.Size()
		}
	}

}

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

func printDiskUage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	fileSize := make(chan int64)
	go func() {
		for _, root := range roots {
			walKDir(root, fileSize)
		}
		close(fileSize)
	}()

	var nfiles, nbytes int64
	for size := range fileSize {
		nfiles++
		nbytes += size
	}
	printDiskUage(nfiles, nbytes)
}
