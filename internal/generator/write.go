// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

import (
	f "github.com/golangsam/dotgo/internal/fs" // adapter to file system (via "container/ccsafe/fs"
)

// Write byteS into dirPath with name resolved against data
func Write(data Dot, filename string, byteS []byte) error {
	fi := f.NewFile(filename)
	return fi.WriteFile(byteS)
}

// FileName resolves name as a template, executed against data
func FileName(data Dot, name string) string {
	id := "FileName"
	fileName := name
	template := NewTemplate(id)
	if tmpl, err := template.Parse(fileName); err == nil {
		if byteS, err := Apply(data, Template{tmpl}, id); err == nil {
			fileName = string(byteS)
		} else {
			panic(id + ": Apply: Error: " + err.Error())
		}
	} else {
		panic(id + ": Parse: Error: " + err.Error())
	}
	return fileName
}
