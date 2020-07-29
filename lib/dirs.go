package lib

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var imgExts = []string{
	".png",
	".jpg",
	".jpeg",
}

// Dirwalk : Extracts only the image file path from the specified directory.
func Dirwalk(dir string) ([]string, error) {
	var stmts []string
	var count int = 1

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("Fail to read %v: %v", path, err)
		}

		if info.IsDir() {
			nest := strings.Count(path, "/") + 1
			stmt := fmt.Sprintf("%v %v\n", strings.Repeat("#", nest), info.Name())
			stmts = append(stmts, stmt)
		} else {
			if ext := filepath.Ext(path); isImg(ext) {
				stmt := fmt.Sprintf("![figure%v](%v)\n", count, path)
				stmts = append(stmts, stmt)
				count++
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return stmts, nil
}

// isImg : It is determined whether the extension is for an image file.
func isImg(ext string) bool {
	for _, e := range imgExts {
		if e == ext {
			return true
		}
	}
	return false
}
