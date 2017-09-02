// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

import (
	"fmt"
)

type Null struct{}

func NewNull() Null {
	return Null{}
}

func (n Null) S() []string {
	return []string{}
}

func (n Null) Len() int {
	return 0
}

func (n Null) Close() error {
	return nil
}

func (n Null) Walker(quit func() bool, out ...*Actor) func() {
	return func() { return }
}

// flagPrint prints nothing but header, iff flag is true
func (n Null) flagPrint(flag, verbose bool, header string) {
	if flag {
		fmt.Println(header, tab, cnt, n.Len(), tab, tab)
		if verbose {
			fmt.Println(tab, tab, tab)
		}
	}
}
