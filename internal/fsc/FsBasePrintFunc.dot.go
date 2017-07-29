package fsc

import (
	"fmt"

	"github.com/golangsam/container/ccsafe/fs"
)

func FsBasePrintFunc(prefix string) func(fp *fs.FsBase) *fs.FsBase {
	return func(fp *fs.FsBase) *fs.FsBase {
		fmt.Println(prefix, fp.String())
		return fp
	}
}
