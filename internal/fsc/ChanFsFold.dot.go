// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
)

// MakeFsFoldChan returns a new open channel
// (simply a 'chan *fs.FsFold' that is).
//
// Note: No 'FsFold-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myFsFoldPipelineStartsHere := MakeFsFoldChan()
//	// ... lot's of code to design and build Your favourite "myFsFoldWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myFsFoldPipelineStartsHere <- drop
//	}
//	close(myFsFoldPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeFsFoldBuffer) the channel is unbuffered.
//
func MakeFsFoldChan() (out chan *fs.FsFold) {
	return make(chan *fs.FsFold)
}

func sendFsFold(out chan<- *fs.FsFold, inp ...*fs.FsFold) {
	defer close(out)
	for _, i := range inp {
		out <- i
	}
}

// ChanFsFold returns a channel to receive all inputs before close.
func ChanFsFold(inp ...*fs.FsFold) (out <-chan *fs.FsFold) {
	cha := make(chan *fs.FsFold)
	go sendFsFold(cha, inp...)
	return cha
}

func sendFsFoldSlice(out chan<- *fs.FsFold, inp ...[]*fs.FsFold) {
	defer close(out)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
}

// ChanFsFoldSlice returns a channel to receive all inputs before close.
func ChanFsFoldSlice(inp ...[]*fs.FsFold) (out <-chan *fs.FsFold) {
	cha := make(chan *fs.FsFold)
	go sendFsFoldSlice(cha, inp...)
	return cha
}

func joinFsFold(done chan<- struct{}, out chan<- *fs.FsFold, inp ...*fs.FsFold) {
	defer close(done)
	for _, i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinFsFold
func JoinFsFold(out chan<- *fs.FsFold, inp ...*fs.FsFold) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinFsFold(cha, out, inp...)
	return cha
}

func joinFsFoldSlice(done chan<- struct{}, out chan<- *fs.FsFold, inp ...[]*fs.FsFold) {
	defer close(done)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
	done <- struct{}{}
}

// JoinFsFoldSlice
func JoinFsFoldSlice(out chan<- *fs.FsFold, inp ...[]*fs.FsFold) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinFsFoldSlice(cha, out, inp...)
	return cha
}

func joinFsFoldChan(done chan<- struct{}, out chan<- *fs.FsFold, inp <-chan *fs.FsFold) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinFsFoldChan
func JoinFsFoldChan(out chan<- *fs.FsFold, inp <-chan *fs.FsFold) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinFsFoldChan(cha, out, inp)
	return cha
}

func doitFsFold(done chan<- struct{}, inp <-chan *fs.FsFold) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneFsFold returns a channel to receive one signal before close after inp has been drained.
func DoneFsFold(inp <-chan *fs.FsFold) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitFsFold(cha, inp)
	return cha
}

func doitFsFoldSlice(done chan<- ([]*fs.FsFold), inp <-chan *fs.FsFold) {
	defer close(done)
	FsFoldS := []*fs.FsFold{}
	for i := range inp {
		FsFoldS = append(FsFoldS, i)
	}
	done <- FsFoldS
}

// DoneFsFoldSlice returns a channel which will receive a slice
// of all the FsFolds received on inp channel before close.
// Unlike DoneFsFold, a full slice is sent once, not just an event.
func DoneFsFoldSlice(inp <-chan *fs.FsFold) (done <-chan ([]*fs.FsFold)) {
	cha := make(chan ([]*fs.FsFold))
	go doitFsFoldSlice(cha, inp)
	return cha
}

func doitFsFoldFunc(done chan<- struct{}, inp <-chan *fs.FsFold, act func(a *fs.FsFold)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneFsFoldFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneFsFoldFunc(inp <-chan *fs.FsFold, act func(a *fs.FsFold)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a *fs.FsFold) { return }
	}
	go doitFsFoldFunc(cha, inp, act)
	return cha
}

func pipeFsFoldBuffer(out chan<- *fs.FsFold, inp <-chan *fs.FsFold) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeFsFoldBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeFsFoldBuffer(inp <-chan *fs.FsFold, cap int) (out <-chan *fs.FsFold) {
	cha := make(chan *fs.FsFold, cap)
	go pipeFsFoldBuffer(cha, inp)
	return cha
}

func pipeFsFoldFunc(out chan<- *fs.FsFold, inp <-chan *fs.FsFold, act func(a *fs.FsFold) *fs.FsFold) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeFsFoldFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeFsFoldMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeFsFoldFunc(inp <-chan *fs.FsFold, act func(a *fs.FsFold) *fs.FsFold) (out <-chan *fs.FsFold) {
	cha := make(chan *fs.FsFold)
	if act == nil {
		act = func(a *fs.FsFold) *fs.FsFold { return a }
	}
	go pipeFsFoldFunc(cha, inp, act)
	return cha
}

func pipeFsFoldFork(out1, out2 chan<- *fs.FsFold, inp <-chan *fs.FsFold) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeFsFoldFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeFsFoldFork(inp <-chan *fs.FsFold) (out1, out2 <-chan *fs.FsFold) {
	cha1 := make(chan *fs.FsFold)
	cha2 := make(chan *fs.FsFold)
	go pipeFsFoldFork(cha1, cha2, inp)
	return cha1, cha2
}
