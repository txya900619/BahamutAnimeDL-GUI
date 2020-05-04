//+build windows

package utilities

import (
	"golang.org/x/sys/windows"
	"log"
	"runtime"
)

func HideFolder(path string) {
	pathW, err := windows.UTF16PtrFromString(path)
	if err != nil {
		log.Fatal(err)
	}
	err = syscall.SetFileAttributes(pathW, windows.FILE_ATTRIBUTE_HIDDEN)
	if err != nil {
		log.Fatal(err)
	}
}
