package fsc

import (
	"fmt"

	"github.com/golangsam/container/ccsafe/fs"
)

func FsFilePrintFunc(prefix string) func(fp *fs.FsFile) *fs.FsFile {
	return func(fp *fs.FsFile) *fs.FsFile {
		fmt.Println(prefix, fp.String())
		return fp
	}
}
