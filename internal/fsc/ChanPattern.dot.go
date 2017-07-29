// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsc

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/golangsam/container/ccsafe/fs"
)

// MakePatternChan returns a new open channel
// (simply a 'chan *fs.Pattern' that is).
//
// Note: No 'Pattern-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myPatternPipelineStartsHere := MakePatternChan()
//	// ... lot's of code to design and build Your favourite "myPatternWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myPatternPipelineStartsHere <- drop
//	}
//	close(myPatternPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipePatternBuffer) the channel is unbuffered.
//
func MakePatternChan() (out chan *fs.Pattern) {
	return make(chan *fs.Pattern)
}

func sendPattern(out chan<- *fs.Pattern, inp ...*fs.Pattern) {
	defer close(out)
	for _, i := range inp {
		out <- i
	}
}

// ChanPattern returns a channel to receive all inputs before close.
func ChanPattern(inp ...*fs.Pattern) (out <-chan *fs.Pattern) {
	cha := make(chan *fs.Pattern)
	go sendPattern(cha, inp...)
	return cha
}

func sendPatternSlice(out chan<- *fs.Pattern, inp ...[]*fs.Pattern) {
	defer close(out)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
}

// ChanPatternSlice returns a channel to receive all inputs before close.
func ChanPatternSlice(inp ...[]*fs.Pattern) (out <-chan *fs.Pattern) {
	cha := make(chan *fs.Pattern)
	go sendPatternSlice(cha, inp...)
	return cha
}

func joinPattern(done chan<- struct{}, out chan<- *fs.Pattern, inp ...*fs.Pattern) {
	defer close(done)
	for _, i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinPattern sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinPattern(out chan<- *fs.Pattern, inp ...*fs.Pattern) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinPattern(cha, out, inp...)
	return cha
}

func joinPatternSlice(done chan<- struct{}, out chan<- *fs.Pattern, inp ...[]*fs.Pattern) {
	defer close(done)
	for _, in := range inp {
		for _, i := range in {
			out <- i
		}
	}
	done <- struct{}{}
}

// JoinPatternSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinPatternSlice(out chan<- *fs.Pattern, inp ...[]*fs.Pattern) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinPatternSlice(cha, out, inp...)
	return cha
}

func joinPatternChan(done chan<- struct{}, out chan<- *fs.Pattern, inp <-chan *fs.Pattern) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinPatternChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinPatternChan(out chan<- *fs.Pattern, inp <-chan *fs.Pattern) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinPatternChan(cha, out, inp)
	return cha
}

func doitPattern(done chan<- struct{}, inp <-chan *fs.Pattern) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DonePattern returns a channel to receive one signal before close after inp has been drained.
func DonePattern(inp <-chan *fs.Pattern) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitPattern(cha, inp)
	return cha
}

func doitPatternSlice(done chan<- ([]*fs.Pattern), inp <-chan *fs.Pattern) {
	defer close(done)
	PatternS := []*fs.Pattern{}
	for i := range inp {
		PatternS = append(PatternS, i)
	}
	done <- PatternS
}

// DonePatternSlice returns a channel which will receive a slice
// of all the Patterns received on inp channel before close.
// Unlike DonePattern, a full slice is sent once, not just an event.
func DonePatternSlice(inp <-chan *fs.Pattern) (done <-chan ([]*fs.Pattern)) {
	cha := make(chan ([]*fs.Pattern))
	go doitPatternSlice(cha, inp)
	return cha
}

func doitPatternFunc(done chan<- struct{}, inp <-chan *fs.Pattern, act func(a *fs.Pattern)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DonePatternFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DonePatternFunc(inp <-chan *fs.Pattern, act func(a *fs.Pattern)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a *fs.Pattern) { return }
	}
	go doitPatternFunc(cha, inp, act)
	return cha
}

func pipePatternBuffer(out chan<- *fs.Pattern, inp <-chan *fs.Pattern) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipePatternBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipePatternBuffer(inp <-chan *fs.Pattern, cap int) (out <-chan *fs.Pattern) {
	cha := make(chan *fs.Pattern, cap)
	go pipePatternBuffer(cha, inp)
	return cha
}

func pipePatternFunc(out chan<- *fs.Pattern, inp <-chan *fs.Pattern, act func(a *fs.Pattern) *fs.Pattern) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipePatternFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipePatternMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipePatternFunc(inp <-chan *fs.Pattern, act func(a *fs.Pattern) *fs.Pattern) (out <-chan *fs.Pattern) {
	cha := make(chan *fs.Pattern)
	if act == nil {
		act = func(a *fs.Pattern) *fs.Pattern { return a }
	}
	go pipePatternFunc(cha, inp, act)
	return cha
}

func pipePatternFork(out1, out2 chan<- *fs.Pattern, inp <-chan *fs.Pattern) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipePatternFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipePatternFork(inp <-chan *fs.Pattern) (out1, out2 <-chan *fs.Pattern) {
	cha1 := make(chan *fs.Pattern)
	cha2 := make(chan *fs.Pattern)
	go pipePatternFork(cha1, cha2, inp)
	return cha1, cha2
}
