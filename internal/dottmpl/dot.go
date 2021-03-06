// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dottmpl

import (
	"fmt"

	"github.com/GoLangsam/container/ccsafe/dot"
)

// constants borrowed from package `dot`
const (
	// ErrorName is the name of a node-type error
	ErrorName = dot.ErrorName
	// ErrorID is the ID of of a node of type error
	ErrorID = dot.ErrorID
)

// Dot defines what is used from "container/ccsafe/dot" to register errors
type Dot interface {
	String() string
	G(keys ...string) *dot.Dot
	Clone() *dot.Dot
	PrintTree(prefix ...string) *dot.Dot

	Fetch(key string) (interface{}, bool) // for ShowErrors
	SeeError(myName, myThing string, err error) bool
	//SeeNotOk(myName, myThing string, ok bool, complain string) bool
}

// NewData returns a fresh named dot
func NewData(name string) *dot.Dot {
	return dot.New(name)
}

// flagPrintDataTree prints the data tree, iff flag is true
func flagPrintDataTree(flag bool, data Dot, prefix string) {
	if flag {
		fmt.Println(prefix + "\t<- " + "Data: >>")
		data.PrintTree(">>")
	}
}

// flagPrintErrors prints the error(s), iff any
func flagPrintErrors(data *dot.Dot, prefix string) bool {
	e, ok := HaveErrors(data)
	switch {
	case ok:
		flagPrintDataTree(true, e, ErrorName)
		return true
	default:
		return false
	}
}

// HaveErrors returns the subnode with errors nad true, iff any - or nil, false
func HaveErrors(d *dot.Dot) (*dot.Dot, bool) {
	_, ok := d.Fetch(ErrorID)
	switch {
	case ok:
		return d.G(ErrorID), true
	default:
		return nil, false
	}
}

// SeeError returns true iff err is non-nil (after registering it)
func SeeError(data *dot.Dot, err error, prefix string) bool {
	switch {
	case err == nil:
		return false
	default:
		data.SeeError("DotGo", prefix, err)
		fmt.Println(prefix+":\t"+ErrorName+"\t", err.Error())
		return true
	}
}
