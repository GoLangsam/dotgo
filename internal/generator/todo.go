// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

import (
	"context"

	"github.com/GoLangsam/do/cli/cancel"
)

type toDo struct {
	ctx context.Context
	can context.CancelFunc
}

func doIt() *toDo {
	ctx, can := cancel.WithCancel()
	return &toDo{ctx, can}
}

func (s *step) done() bool {
	return s.todo.ctx.Err() != nil
}

var noquit = func() bool { return false }
