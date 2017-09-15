// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gen

import "fmt"

type Error struct {
	name string
	what string
	fail error
}

type ErrorS []Error

var all ErrorS = []Error{} // global errors

// Ok tells, iff all is fine
// and collects any error
func (e ErrorS) Ok(myName, myThing string, err error) bool {
	if err != nil {
		e = append(e, Error{myName, myThing, err})
	}
	return err == nil
}

// NotOk prints
// the error(s), iff any
// and tells, if there were some
func (e ErrorS) NotOk(header string) bool {

	for i := range e {
		fmt.Println(header, e[i].fail.Error())
	}

	return len(e) > 0
}
