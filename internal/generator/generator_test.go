// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

import (
	_ "fmt"
)

// Do all adapted collections implement Some?
var _ Some = new(Actor)
var _ Some = Actor{}

var _ Some = DirS{}
var _ Some = NewData(aDot)
var _ Some = NewTemplate(aDot)

var _ SomeWithAction = NewDict()
var _ SomeWithAction = NewNull()
var _ SomeWithAction = NewNext(0, 0)
var _ SomeWithAction = NewPrev(0, 0)
