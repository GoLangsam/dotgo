// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fs is an adapter to the file system as represented by "container/ccsafe/fs"
package fs

import (
	"github.com/golangsam/container/ccsafe/fs"
)

// NewFile returns new a *File with given name/path
func NewFile(name string) *fs.FsFile {
	return fs.ForceFile(name)
}

// NewFold returns new a *Fold with given name/path
func NewFold(name string) *fs.FsFold {
	return fs.ForceFold(name)
}

// AsFile returns the info as type *File, if possible,
// and panics otherwise
func AsFile(e FsInfo) *fs.FsFile {
	switch e := e.(type) {
	case *fs.FsFile:
		return e
	default:
		panic("AsFile: *FsFile expected")
		// return fs.ForceFile(e.String())

	}
}

// AsFold returns the info as type *Fold, if possible,
// and panics otherwise
func AsFold(e FsInfo) *fs.FsFold {
	switch e := e.(type) {
	case *fs.FsFold:
		return e
	default:
		panic("AsFold: *FsFold expected")
		// return fs.ForceFold(e.String())
	}
}

// FsName defines what is used from fs.FsBase
type FsName interface {
	String() string
	Base() *fs.FsBase
	BaseLessExt() *fs.FsBase
}

// FsInfo defines what is used from fs.FsInfo
type FsInfo interface {
	FsName
	IsFold() bool
	JoinWith(elem ...string) string
	// WriteFile(data []byte) error
}
