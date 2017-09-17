// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dottmpl

import (
	"context"

	a "github.com/GoLangsam/dotgo/internal/fsa" // adapter to file system analysis

	"github.com/GoLangsam/container/ccsafe/dot"
	"github.com/GoLangsam/do/cli/cancel"
)

type toDo struct {
	data *dot.Dot
	tmpl Template
	dir  *a.Analysis
	ctx  context.Context
	can  context.CancelFunc
}

func doIt(data *dot.Dot, tmpl Template, dir *a.Analysis) *toDo {
	ctx, can := cancel.WithCancel()
	return &toDo{data, tmpl, dir, ctx, can}
}

func (t *toDo) doIt(data *dot.Dot, tmpl Template, dir *a.Analysis) (todo *toDo, ok bool) {
	select {
	case <-t.ctx.Done():
		ok = false
	default:
		ok = true
	}
	return &toDo{data, tmpl, dir, t.ctx, t.can}, ok
}
