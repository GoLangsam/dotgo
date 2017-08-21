// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsa

import (
	"github.com/golangsam/container/ccsafe/fs"
)

// AnalyseDown returns the analysis recursing into any subdirs rooted at *Fold without
func (a *Analysis) AnalyseDown(dir *fs.FsFold) *Analysis {
	ca := a.Descend(Down, dir, dir.Recurse())

	dS, fS, _ := fs.MatchDisk(dir.String())
	ca.takeFileS(fS...)

	if len(dS) > 0 {
		for _, d := range dS {
			if dir.Recurse() {
				d = d.AsRecurse()
			}
			ca.collectDown(d)
		}
	} else {
		ca.collectDown(dir)
	}

	return ca
}

func (a *Analysis) collectDown(dir *fs.FsFold) *Analysis {
	a = a.takeFileS(a.FileS(extPatternS...)...)
	if dir.Recurse() {
		a = a.collectDownSubDirS()
	}
	return a
}

func (a *Analysis) collectDownSubDirS() *Analysis {
	dirInfoS, _ := a.ReadDir()
	for i := range dirInfoS { // ReadDir returned one level
		if dirInfoS[i].IsDir() {
			downDir := fs.Recurse(a.JoinWith(dirInfoS[i].Name()))
			ca := a.Descend(Down, downDir, true)
			ca = ca.takeFileS(downDir.FileS(extPatternS...)...)
			ca.collectDownSubDirS() // recurse down
		}
	}
	return a
}
