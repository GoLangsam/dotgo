// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/golangsam/do/cli/cancel"
	doit "github.com/golangsam/dotgo/internal/dottmpl"
)

func main() {
	_ = cancel.Canceler() // TODO pass ctx on
	doit.DoIt()
}
