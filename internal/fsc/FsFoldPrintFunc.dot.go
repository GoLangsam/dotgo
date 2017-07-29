package fsc

import (
	"fmt"

	"github.com/golangsam/container/ccsafe/fs"
)

func FsFoldPrintFunc(prefix string) func(fp *fs.FsFold) *fs.FsFold {
	return func(fp *fs.FsFold) *fs.FsFold {
		fmt.Println(prefix, fp.String())
		return fp
	}
}
