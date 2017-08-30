// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

import (
	"context"
	"sync"

	"github.com/golangsam/container/ccsafe/dot"
	"github.com/golangsam/do/cli/cancel"
)

type toDo struct {
	data *dot.Dot
	tmpl Template
	ctx  context.Context
	can  context.CancelFunc
	wg   sync.WaitGroup
}

func doIt(data *dot.Dot, tmpl Template) *toDo {
	ctx, can := cancel.WithCancel()
	wg := new(sync.WaitGroup)
	return &toDo{data, tmpl, ctx, can, *wg}
}

func (t *toDo) doIt(data *dot.Dot, tmpl Template) (todo *toDo, ok bool) {
	select {
	case <-t.ctx.Done():
		ok = false
	default:
		ok = true
	}
	wg := new(sync.WaitGroup)
	return &toDo{data, tmpl, t.ctx, t.can, *wg}, ok
}

func (t *toDo) ok() bool {
	return t.ctx.Err() == nil
}

func (t *toDo) do(do func()) {
	t.wg.Add(1)

	f := func(t *toDo, do func()) {
		defer t.wg.Done()
		do()
	}

	if seq {
		f(t, do)
	} else {
		go f(t, do)
	}
}
