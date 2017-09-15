// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

import (
	"fmt"

	"github.com/golangsam/container/ccsafe/dot"
)

type Data struct {
	*dot.Dot
}

// NewData returns a fresh named dot
func NewData(name string) Data {
	return Data{dot.New(name)}
}

// Beg implement Some

// S -
// inherited

// Len -
// inherited

// Close -
// pretend to be a Closer (<=> an io.Closer)
func (d Data) Close() error {
	return nil
}

// Walker -
// traverse the (sorted) child names the data tree node
func (d Data) Walker(quit func() bool, out ...*Actor) func() {

	return func() {

		defer ActorsClose(out...)
		for _, item := range d.S() {
			if quit() {
				return // bail out
			}
			ActorsDo(item, out...)
		}
	}
}

// flagPrint prints
// the data tree,
// iff flag is true
func (d Data) flagPrint(flag, verbose bool, header string) {
	if flag {
		fmt.Println(header, tab, cnt, d.Len(), tab, tab)

		if verbose {
			d.PrintTree(">>")
			fmt.Println(tab, tab, tab)
		}
	}
}

// End implement Some

// FileName resolves name as a template, executed against data
func (d Data) FileName(name string) string {
	id := "FileName"
	fileName := name
	template := NewTemplate(id)
	tmpl, err := template.Parse(fileName)
	if all.Ok("Parse", id, err) {
		byteS, err := Apply(d, Template{tmpl}, id)
		if all.Ok("Apply", id, err) {
			fileName = string(byteS)
		}
	}
	return fileName
}
