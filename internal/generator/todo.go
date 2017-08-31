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
	ctx  context.Context
	can  context.CancelFunc
	wg   sync.WaitGroup
}

func doIt(data *dot.Dot) *toDo {
	ctx, can := cancel.WithCancel()
	wg := new(sync.WaitGroup)
	return &toDo{data, ctx, can, *wg}
}

func (t *toDo) doIt(data *dot.Dot) *toDo {
	return &toDo{data, t.ctx, t.can, t.wg}
}

func (t *toDo) ok() bool {
	return t.ctx.Err() == nil
}

func (t *toDo) do(do func()) {
	t.wg.Add(1)

	go func(t *toDo, do func()) {
		defer t.wg.Done()
		do()
	}(t, do)
}
