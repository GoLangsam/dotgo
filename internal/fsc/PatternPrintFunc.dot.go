package fsc

import (
	"fmt"

	"github.com/golangsam/container/ccsafe/fs"
)

func PatternPrintFunc(prefix string) func(fp *fs.Pattern) *fs.Pattern {
	return func(fp *fs.Pattern) *fs.Pattern {
		fmt.Println(prefix, fp.String())
		return fp
	}
}
