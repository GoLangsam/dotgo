package fsc

import (
	"fmt"

	"github.com/golangsam/container/ccsafe/fs"
)

func FsPathSPrintFunc(prefix string) func(fp fs.FsPathS) fs.FsPathS {
	return func(fp fs.FsPathS) fs.FsPathS {
		fmt.Println(prefix, fp.String())
		return fp
	}
}
