// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

type Null struct{}

func NewNull() Null {
	return Null{}
}

func (n Null) S() []string {
	return []string{}
}

func (n Null) Close() error {
	return nil
}

func (n Null) Walker(quit func() bool, out ...Actor) func() {
	return func() { return }
}
