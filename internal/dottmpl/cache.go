// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dottmpl

import (
	"github.com/golangsam/container/ccsafe/fscache"
	f "github.com/golangsam/dotgo/internal/fs" // adapter to file system (via "container/ccsafe/fs"
)

var (
	fsCache *fscache.FsCache // file data cache
)

// lookupData retrieves the data related to FsInfo from FsCache
func lookupData(file f.FsInfo) string {
	return fsCache.LookupData(f.AsFile(file))
}
