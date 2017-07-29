// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dottmpl

import (
	a "github.com/golangsam/dotgo/internal/fsa" // adapter to file system analysis

	"github.com/golangsam/container/ccsafe/dot"
)

type toDo struct {
	data *dot.Dot
	tmpl Template
	dir  *a.Analysis
}

func doIt(data *dot.Dot, tmpl Template, dir *a.Analysis) *toDo {
	return &toDo{data, tmpl, dir}
}
