// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package template

import (
	"fmt"
)

func IfPrintTemplate(flag bool, tmpl Template, prefix string) {
	if flag {
		fmt.Println(prefix + " Template Names:\t")
		IfPrintNames(flag, tmpl, "\t")
	}
}

func IfPrintNames(flag bool, tmpl Template, prefix string) {
	if flag {
		for _, name := range Names(tmpl) {
			fmt.Println(prefix + name + "\t")
		}
	}
}
