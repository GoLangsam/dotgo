// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsa

import (
	"github.com/GoLangsam/container/ccsafe/fs"
	"github.com/GoLangsam/container/ccsafe/fscache"
)

// These variables support the file system analysis.
var (
	cacheFsFile chan *fs.FsFile         // populates FsCache, closed at end of analysis
	doneCache   <-chan *fscache.FsCache // delivers FsCache upon Close of analysis
)

func openCache() {
	cacheFsFile = make(chan *fs.FsFile, 64) // buffered, as FileRead is slow
	doneCache = fscache.Cache(cacheFsFile)  // done Cache listener
}

func closeCache() *fscache.FsCache {
	close(cacheFsFile) // close the cache listener
	return <-doneCache // wait for cache
}

func cacheFiles(files ...*fs.FsFile) {
	for i := range files {
		cacheFsFile <- files[i]
	}
}
