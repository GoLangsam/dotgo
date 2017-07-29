// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsa

import (
	"github.com/golangsam/container/ccsafe/fs"
)

// MatchingBaseS returns the base names matching the patterns
func (a *Analysis) MatchingBaseS(files fs.FsFileS, patterns ...*fs.Pattern) (nameS fs.FsBaseS) {
	for _, file := range files {
		if ok, _ := file.BaseMatches(patterns...); ok { // ignore errors
			nameS = append(nameS, file.Base())
		}
	}
	return nameS
}
