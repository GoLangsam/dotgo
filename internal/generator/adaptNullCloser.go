// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

type nullItem struct{}

func null() nullItem {
	return nullItem{}
}

func (n nullItem) S() []string {
	return []string{}
}

func (n nullItem) Close() error {
	return nil
}

func (n nullItem) Walker(t *toDo, out ...maker) func() {
	return func() { return }
}

func (n nullItem) Add(item string) {
	// nop
}
