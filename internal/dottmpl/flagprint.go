// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dottmpl

import (
	"fmt"
)

func flagOpen(flag bool, prefix string) {
	if flag {
		fmt.Print(prefix + "\t<- ")
	}
}

func flagDot(flag bool) {
	if flag {
		fmt.Print(". ")
	}
}
func flagClose(flag bool) {
	if flag {
		fmt.Println("\t<- " + "Done!" + "\t")
	}
}

func flagPrintByteS(flag bool, byteS []byte, prefix string) {
	if flag {
		fmt.Println(prefix + "\t")
		fmt.Println(string(byteS) + "\t")
	}
}

func flagPrintString(flag bool, prefix string, suffix string) {
	if flag {
		fmt.Println(prefix + "\t<- " + suffix + "\t")
	}
}
