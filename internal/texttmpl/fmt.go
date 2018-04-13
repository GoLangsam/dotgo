// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package template

import (
	"fmt"
)

// IfPrintTemplate prints the template names, iff flag is true
func IfPrintTemplate(flag bool, tmpl Template, prefix string) {
	if flag {
		fmt.Println(prefix + " Template Names:\t")
		IfPrintNames(flag, tmpl, "\t")
	}
}

// IfPrintNames prints the names of the templates, iff flag is true
func IfPrintNames(flag bool, tmpl Template, prefix string) {
	if flag {
		nameS := Names(tmpl)
		for i := range nameS {
			fmt.Println(prefix + nameS[i] + "\t")
		}
	}
}
