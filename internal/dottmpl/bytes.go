// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dottmpl

import (
	"bytes"
)

// Apply template name to data; evaluate name (less Ext) against data
// as filenameWriteFile into dirPath
func Apply(data Dot, tmpl Template, name string) (byteS []byte, err error) {
	var buf bytes.Buffer
	err = tmpl.ExecuteTemplate(&buf, name, data)
	return buf.Bytes(), err
}
