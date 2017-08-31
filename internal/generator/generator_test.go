// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

import (
	_ "fmt"
)

// Do all adapted collections implement Item?
var _ Some = Actor{}

// var _ Item = NewData(aDot) // missing: Add Walker
var _ Some = NewDict()
var _ Some = NewNull()
var _ Some = NewNext(0, 0)
var _ Some = NewPrev(0, 0)

// var _ Item = NewTemplate(aDot) // missing: Add Walker
