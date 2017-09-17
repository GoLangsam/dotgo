// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dottmpl

import (
	f "github.com/GoLangsam/dotgo/internal/fs" // adapter to file system (via "container/ccsafe/fs"
)

// Write byteS into dirPath with name resolved against data
func Write(data Dot, filename string, byteS []byte) error {
	fi := f.NewFile(filename)
	return fi.WriteFile(byteS)
}

// FileName resolves name as a template, executed against data
func FileName(data Dot, name string) string {
	fileName := name
	if tmpl, err := NewTemplate("FileName").Parse(fileName); err == nil {
		if byteS, err := Apply(data, tmpl, "FileName"); err == nil {
			fileName = string(byteS)
		} else {
			panic("Filename: Apply: Error: " + err.Error())
		}
	} else {
		panic("Filename: Parse: Error: " + err.Error())
	}
	return fileName
}
