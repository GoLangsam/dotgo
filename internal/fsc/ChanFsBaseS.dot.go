// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
)

// MakeFsBaseSChan returns a new open channel
// (simply a 'chan fs.FsBaseS' that is).
//
// Note: No 'FsBaseS-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myFsBaseSPipelineStartsHere := MakeFsBaseSChan()
//	// ... lot's of code to design and build Your favourite "myFsBaseSWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myFsBaseSPipelineStartsHere <- drop
//	}
//	close(myFsBaseSPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeFsBaseSBuffer) the channel is unbuffered.
//
func MakeFsBaseSChan() (out chan fs.FsBaseS) {
	return make(chan fs.FsBaseS)
}

func sendFsBaseS(out chan<- fs.FsBaseS, inp ...fs.FsBaseS) {
	defer close(out)
	for _, i := range inp {
		out <- i
	}
}

// ChanFsBaseS returns a channel to receive all inputs before close.
func ChanFsBaseS(inp ...fs.FsBaseS) (out <-chan fs.FsBaseS) {
	cha := make(chan fs.FsBaseS)
	go sendFsBaseS(cha, inp...)
	return cha
}

func sendFsBaseSSlice(out chan<- fs.FsBaseS, inp ...[]fs.FsBaseS) {
	defer close(out)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
}

// ChanFsBaseSSlice returns a channel to receive all inputs before close.
func ChanFsBaseSSlice(inp ...[]fs.FsBaseS) (out <-chan fs.FsBaseS) {
	cha := make(chan fs.FsBaseS)
	go sendFsBaseSSlice(cha, inp...)
	return cha
}

func joinFsBaseS(done chan<- struct{}, out chan<- fs.FsBaseS, inp ...fs.FsBaseS) {
	defer close(done)
	for _, i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinFsBaseS
func JoinFsBaseS(out chan<- fs.FsBaseS, inp ...fs.FsBaseS) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinFsBaseS(cha, out, inp...)
	return cha
}

func joinFsBaseSSlice(done chan<- struct{}, out chan<- fs.FsBaseS, inp ...[]fs.FsBaseS) {
	defer close(done)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
	done <- struct{}{}
}

// JoinFsBaseSSlice
func JoinFsBaseSSlice(out chan<- fs.FsBaseS, inp ...[]fs.FsBaseS) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinFsBaseSSlice(cha, out, inp...)
	return cha
}

func joinFsBaseSChan(done chan<- struct{}, out chan<- fs.FsBaseS, inp <-chan fs.FsBaseS) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinFsBaseSChan
func JoinFsBaseSChan(out chan<- fs.FsBaseS, inp <-chan fs.FsBaseS) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinFsBaseSChan(cha, out, inp)
	return cha
}

func doitFsBaseS(done chan<- struct{}, inp <-chan fs.FsBaseS) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneFsBaseS returns a channel to receive one signal before close after inp has been drained.
func DoneFsBaseS(inp <-chan fs.FsBaseS) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitFsBaseS(cha, inp)
	return cha
}

func doitFsBaseSSlice(done chan<- ([]fs.FsBaseS), inp <-chan fs.FsBaseS) {
	defer close(done)
	FsBaseSS := []fs.FsBaseS{}
	for i := range inp {
		FsBaseSS = append(FsBaseSS, i)
	}
	done <- FsBaseSS
}

// DoneFsBaseSSlice returns a channel which will receive a slice
// of all the FsBaseSs received on inp channel before close.
// Unlike DoneFsBaseS, a full slice is sent once, not just an event.
func DoneFsBaseSSlice(inp <-chan fs.FsBaseS) (done <-chan ([]fs.FsBaseS)) {
	cha := make(chan ([]fs.FsBaseS))
	go doitFsBaseSSlice(cha, inp)
	return cha
}

func doitFsBaseSFunc(done chan<- struct{}, inp <-chan fs.FsBaseS, act func(a fs.FsBaseS)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneFsBaseSFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneFsBaseSFunc(inp <-chan fs.FsBaseS, act func(a fs.FsBaseS)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a fs.FsBaseS) { return }
	}
	go doitFsBaseSFunc(cha, inp, act)
	return cha
}

func pipeFsBaseSBuffer(out chan<- fs.FsBaseS, inp <-chan fs.FsBaseS) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeFsBaseSBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeFsBaseSBuffer(inp <-chan fs.FsBaseS, cap int) (out <-chan fs.FsBaseS) {
	cha := make(chan fs.FsBaseS, cap)
	go pipeFsBaseSBuffer(cha, inp)
	return cha
}

func pipeFsBaseSFunc(out chan<- fs.FsBaseS, inp <-chan fs.FsBaseS, act func(a fs.FsBaseS) fs.FsBaseS) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeFsBaseSFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeFsBaseSMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeFsBaseSFunc(inp <-chan fs.FsBaseS, act func(a fs.FsBaseS) fs.FsBaseS) (out <-chan fs.FsBaseS) {
	cha := make(chan fs.FsBaseS)
	if act == nil {
		act = func(a fs.FsBaseS) fs.FsBaseS { return a }
	}
	go pipeFsBaseSFunc(cha, inp, act)
	return cha
}

func pipeFsBaseSFork(out1, out2 chan<- fs.FsBaseS, inp <-chan fs.FsBaseS) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeFsBaseSFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeFsBaseSFork(inp <-chan fs.FsBaseS) (out1, out2 <-chan fs.FsBaseS) {
	cha1 := make(chan fs.FsBaseS)
	cha2 := make(chan fs.FsBaseS)
	go pipeFsBaseSFork(cha1, cha2, inp)
	return cha1, cha2
}
