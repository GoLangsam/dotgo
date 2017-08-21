// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsa

import (
	"github.com/golangsam/container/ccsafe/fs"
)

// MatchingFileS returns the file names matching the patterns
func (a *Analysis) MatchingFileS(files fs.FsFileS, patterns ...*fs.Pattern) (nameS fs.FsFileS) {
	for i := range files {
		if ok, _ := files[i].BaseMatches(patterns...); ok { // ignore errors
			nameS = append(nameS, files[i])
		}
	}
	return nameS
}
