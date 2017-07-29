// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dottmpl

import (
	"fmt"

	"github.com/golangsam/container/ccsafe/dot"
)

const (
	ErrorName = dot.ErrorName
	ErrorID   = dot.ErrorID
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

func NewData(name string) *dot.Dot {
	return dot.New(name)
}

func flagPrintDataTree(flag bool, data Dot, prefix string) {
	if flag {
		fmt.Println(prefix + "\t<- " + "Data: >>")
		data.PrintTree(">>")
	}
}

func flagPrintErrors(data *dot.Dot, prefix string) bool {
	if e, ok := HaveErrors(data); ok {
		flagPrintDataTree(true, e, ErrorName)
		return true
	} else {
		return false
	}
}

func HaveErrors(d *dot.Dot) (*dot.Dot, bool) {
	if _, ok := d.Fetch(ErrorID); ok {
		return d.G(ErrorID), true
	} else {
		return nil, false
	}
}

func SeeError(data *dot.Dot, err error, prefix string) bool {
	if err == nil {
		return false
	} else {
		data.SeeError("DotGo", prefix, err)
		fmt.Println(prefix+":\t"+ErrorName+"\t", err.Error())
		return true
	}
}
