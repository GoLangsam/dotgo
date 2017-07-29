package fsc

import (
	"fmt"

	"github.com/golangsam/container/ccsafe/fs"
)

func FsFoldSPrintFunc(prefix string) func(fp fs.FsFoldS) fs.FsFoldS {
	return func(fp fs.FsFoldS) fs.FsFoldS {
		fmt.Println(prefix, fp.String())
		return fp
	}
}
