package fsc

import (
	"fmt"

	"github.com/golangsam/container/ccsafe/fs"
)

func FsFileSPrintFunc(prefix string) func(fp fs.FsFileS) fs.FsFileS {
	return func(fp fs.FsFileS) fs.FsFileS {
		fmt.Println(prefix, fp.String())
		return fp
	}
}
