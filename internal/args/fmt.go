// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package args

import (
	"fmt"

	"github.com/GoLangsam/container/ccsafe/dotpath"
)

// IfPrintFlagArgs prints the arguments, iff flag is true
func IfPrintFlagArgs(flag bool, args ...string) {
	if flag {
		ds := dotpath.FilePathS(args...)
		fmt.Println("===============================================================================")
		for i := range ds {
			ds[i].Print()
			fmt.Println("-------------------------------------------------------------------------------")
		}
		as := ToFolds(args...)
		for i := range as {
			fmt.Println(as[i].String(), "\t", as[i].Recurse())
		}
		fmt.Println("===============================================================================")
	}
}
