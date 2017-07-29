// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsa

import (
	"github.com/golangsam/container/ccsafe/fs"
)

// These variables support the file system analysis.
var (
	extPatternS fs.PatternS // patterns for extensions
	namePattern *fs.Pattern // pattern for basename := first pattern
)

func setPatterns(patterns fs.PatternS) {
	extPatternS = patterns
	namePattern = extPatternS[0] // let it panic, if empty.
}
