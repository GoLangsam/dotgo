// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

import (
	"context"

	"github.com/golangsam/do/cli/cancel"
)

type toDo struct {
	data Dot
	ctx  context.Context
	can  context.CancelFunc
}

func doIt(data Dot) *toDo {
	ctx, can := cancel.WithCancel()
	return &toDo{data, ctx, can}
}

func (t *toDo) quit() func() bool {
	return func() bool { return t.ctx.Err() != nil }
}

// ifPrintErrors prints the error(s), iff any
func (t *toDo) ifPrintErrors(header string) bool {
	return flagPrintErrors(t.data, header)
}
