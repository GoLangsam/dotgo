package fsc

import (
	"fmt"

	"github.com/golangsam/container/ccsafe/fs"
)

func FsBaseSPrintFunc(prefix string) func(fp fs.FsBaseS) fs.FsBaseS {
	return func(fp fs.FsBaseS) fs.FsBaseS {
		fmt.Println(prefix, fp.String())
		return fp
	}
}
