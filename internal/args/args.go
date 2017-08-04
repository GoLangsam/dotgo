// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package args is used to convert the command line arguments
package args

import (
	"github.com/golangsam/container/ccsafe/dotpath"
	"github.com/golangsam/container/ccsafe/fs"
)

// ToFolds uses dotpath to convert args to fs.FsFoldS - with or without recurse flag
func ToFolds(args ...string) fs.FsFoldS {
	dirS := make([]*fs.FsFold, 0, len(args))
	dirS = append(dirS, toFoldS(dotpath.FilePathS(args...)...)...)
	return dirS
}

// toFolds converts pathS to fs.FsFoldS - with or without recurse flag
func toFoldS(pathS ...*dotpath.DotPath) (dirS fs.FsFoldS) {
	for i := range pathS {
		dirS = append(dirS, dotFoldS(pathS[i])...)
	}
	return dirS
}

// dotFoldS returns Recurse / NotDown FoldS from given DotPath
func dotFoldS(dotPath *dotpath.DotPath) (dirS fs.FsFoldS) {

	var waydown = make(map[string]bool)

	dpS := dotPath.PathS()
	for i := range dpS {
		waydown[dpS[i]] = false
	}

	rpS := dotPath.RecursePathS()
	for i := range rpS {
		if _, ok := waydown[rpS[i]]; ok {
			waydown[rpS[i]] = true
		} else {
			dirS = append(dirS, fs.Recurse(rpS[i]))
		}
	}

	for i := range dpS {
		if waydown[dpS[i]] {
			dirS = append(dirS, fs.Recurse(dpS[i]))
		} else {
			dirS = append(dirS, fs.NotDown(dpS[i]))
		}
	}

	return dirS
}
