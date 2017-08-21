// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsa

import (
	"fmt"
)

// IfPrintAnalysis prints the analyses, iff flag is true
func IfPrintAnalysis(flag bool, suffix string, aS ...*Analysis) {
	if flag {
		for i := range aS {
			fmt.Println(aS[i].String() + "\t" + suffix + "\t")
		}
	}
}
