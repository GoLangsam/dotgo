package fsc

import (
	"fmt"

	"github.com/golangsam/container/ccsafe/fs"
)

func FsDataPrintFunc(prefix string) func(fp *fs.FsData) *fs.FsData {
	return func(fp *fs.FsData) *fs.FsData {
		fmt.Println(prefix, fp.String())
		return fp
	}
}
