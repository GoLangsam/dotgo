package fsc

import (
	"fmt"

	"github.com/golangsam/container/ccsafe/fs"
)

func PatternSPrintFunc(prefix string) func(fp fs.PatternS) fs.PatternS {
	return func(fp fs.PatternS) fs.PatternS {
		fmt.Println(prefix, fp.String())
		return fp
	}
}
