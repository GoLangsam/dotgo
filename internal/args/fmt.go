// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package args

import (
	"fmt"

	"github.com/golangsam/container/ccsafe/dotpath"
)

func IfPrintFlagArgs(flag bool, args ...string) {
	if flag {
		ds := dotpath.FilePathS(args...)
		fmt.Println("===============================================================================")
		for _, dp := range ds {
			dp.Print()
			fmt.Println("-------------------------------------------------------------------------------")
		}
		as := ToFolds(args...)
		for _, a := range as {
			fmt.Println(a.String(), "\t", a.Recurse())
		}
		fmt.Println("===============================================================================")
	}
}
