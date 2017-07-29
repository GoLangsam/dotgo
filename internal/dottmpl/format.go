// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dottmpl

import (
	"go/format"
)

// Source is a wrapper for format.Source
func Source(byteS []byte) ([]byte, error) {
	return format.Source(byteS)
}
