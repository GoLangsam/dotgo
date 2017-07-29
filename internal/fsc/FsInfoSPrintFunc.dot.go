package fsc

import (
	"fmt"

	"github.com/golangsam/container/ccsafe/fs"
)

func FsInfoSPrintFunc(prefix string) func(fp fs.FsInfoS) fs.FsInfoS {
	return func(fp fs.FsInfoS) fs.FsInfoS {
		fmt.Println(prefix, fp.String())
		return fp
	}
}
