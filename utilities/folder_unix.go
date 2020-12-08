// +build !windows

package utilities

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func HideFolder(path string) {
	if !strings.HasPrefix(filepath.Base(path), ".") {
		err := os.Rename(path, "."+filepath.Base(path))
		if err != nil {
			log.Fatal(err)
		}
	}
}
