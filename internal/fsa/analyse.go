// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fsa provides File System Analysis based on the fs File System abstraction
package fsa

import (
	"github.com/golangsam/container/ccsafe/fs"
	"github.com/golangsam/container/ccsafe/fscache"
)

const (
	aDot = fs.Dot // just a "."
)

// Open prepares the patterns and a fresh cache and returns the root of new analysis.
//  Note: Do not Open analyses concurrently.
func Open(patterns fs.PatternS) *Analysis {
	a := newAnalysis(Root, fs.ForceFold(aDot)) // Root analysis
	setPatterns(patterns)                      // register patterns
	openCache()                                // open the cache listener
	return a
}

// Close returns the populated cache.
func Close() *fscache.FsCache {
	return closeCache()
}
