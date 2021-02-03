// Copyright (c) 2020 github.com/nyantlet
// Use of this source code is governed by a zlib-style
// license that can be found in the LICENSE file.

/*
Package nya provides a set of assertion functions.

By default, debug assertions do nothing.
You can turn on debug assertions with build tag "debugnya":

	go build -tags debugnya PACKAGE
*/
package nya

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
)

func outputAndDie(depth int, msg string) {
	var fn string
	pc, file, line, ok := runtime.Caller(depth)
	if ok {
		fn = runtime.FuncForPC(pc).Name()
	} else {
		fn = "???"
		file = "???"
		line = 0
	}

	// nya: file.go:5: package.function: MESSAGE
	fmt.Fprint(os.Stderr,
		"nya: ", file, ":", line, ": ", fn, ": ", msg, "\n")
	os.Exit(2)
}

// True expects cond is true.
func True(cond bool, msg string) {
	if cond {
		return
	}
	outputAndDie(2, msg)
}

// Truef is True with format.
// Arguments are handled in the manner of fmt.Printf.
func Truef(cond bool, format string, v ...interface{}) {
	if cond {
		return
	}
	msg := fmt.Sprintf(format, v...)
	outputAndDie(2, msg)
}

// False expects cond is false.
func False(cond bool, msg string) {
	if !cond {
		return
	}
	outputAndDie(2, msg)
}

// Falsef is False with format.
// Arguments are handled in the manner of fmt.Printf.
func Falsef(cond bool, format string, v ...interface{}) {
	if !cond {
		return
	}
	msg := fmt.Sprintf(format, v...)
	outputAndDie(2, msg)
}

// Equal expects expected == acutual.
// You cannot use non-comparable types (slice, map, ...).
func Equal(expected, actual interface{}, msg string) {
	if expected == actual {
		return
	}
	msg2 := fmt.Sprintf("%s\nexpected %v (type %T), found %v (type %T)",
		msg, expected, expected, actual, actual)
	outputAndDie(2, msg2)
}

// Equalf is Equal with format.
// Arguments are handled in the manner of fmt.Printf.
func Equalf(expected, actual interface{}, format string, v ...interface{}) {
	if expected == actual {
		return
	}
	msg := fmt.Sprintf(format, v...)
	msg2 := fmt.Sprintf("%s\nexpected %v (type %T), found %v (type %T)",
		msg, expected, expected, actual, actual)
	outputAndDie(2, msg2)
}

// DeepEqual expects reflect.DeepEqual(expected, actual) is true.
func DeepEqual(expected, actual interface{}, msg string) {
	if reflect.DeepEqual(expected, actual) {
		return
	}
	msg2 := fmt.Sprintf("%s\nexpected %v (type %T), found %v (type %T)",
		msg, expected, expected, actual, actual)
	outputAndDie(2, msg2)
}

// DeepEqualf is DeepEqual with format.
func DeepEqualf(expected, actual interface{}, format string, v ...interface{}) {
	if reflect.DeepEqual(expected, actual) {
		return
	}
	msg := fmt.Sprintf(format, v...)
	msg2 := fmt.Sprintf("%s\nexpected %v (type %T), found %v (type %T)",
		msg, expected, expected, actual, actual)
	outputAndDie(2, msg2)
}
