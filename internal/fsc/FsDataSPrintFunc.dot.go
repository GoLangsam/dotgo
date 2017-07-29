package fsc

import (
	"fmt"

	"github.com/golangsam/container/ccsafe/fs"
)

func FsDataSPrintFunc(prefix string) func(fp fs.FsDataS) fs.FsDataS {
	return func(fp fs.FsDataS) fs.FsDataS {
		fmt.Println(prefix, fp.String())
		return fp
	}
}
