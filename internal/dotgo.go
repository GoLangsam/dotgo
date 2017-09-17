// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	doit "github.com/GoLangsam/dotgo/internal/generator"
)

func main() {
	if err := doit.DoIt(); err != nil {
		fmt.Println("Error:", err)
	}
}
