// Copyright (c) 2020 github.com/nyantlet
// Use of this source code is governed by a zlib-style
// license that can be found in the LICENSE file.

// +build debugnya

package nya

import (
	"fmt"
	"reflect"
)

func DebugTrue(cond bool, msg string) {
	if cond {
		return
	}
	outputAndDie(2, msg)
}

func DebugTruef(cond bool, format string, v ...interface{}) {
	if cond {
		return
	}
	msg := fmt.Sprintf(format, v...)
	outputAndDie(2, msg)
}

func DebugFalse(cond bool, msg string) {
	if !cond {
		return
	}
	outputAndDie(2, msg)
}

func DebugFalsef(cond bool, format string, v ...interface{}) {
	if !cond {
		return
	}
	msg := fmt.Sprintf(format, v...)
	outputAndDie(2, msg)
}

func DebugEqual(expected, actual interface{}, msg string) {
	if expected == actual {
		return
	}
	msg2 := fmt.Sprintf("%s\nexpected %v (type %T), found %v (type %T)",
		msg, expected, expected, actual, actual)
	outputAndDie(2, msg2)
}

func DebugEqualf(expected, actual interface{}, format string, v ...interface{}) {
	if expected == actual {
		return
	}
	msg := fmt.Sprintf(format, v...)
	msg2 := fmt.Sprintf("%s\nexpected %v (type %T), found %v (type %T)",
		msg, expected, expected, actual, actual)
	outputAndDie(2, msg2)
}

func DebugDeepEqual(expected, actual interface{}) {
	if reflect.DeepEqual(expected, actual) {
		return
	}
	msg2 := fmt.Sprintf("%s\nexpected %v (type %T), found %v (type %T)",
		msg, expected, expected, actual, actual)
	outputAndDie(2, msg2)
}

func DebugDeepEqualf(expected, actual interface{}, format string, v ...interface{}) {
	if reflect.DeepEqual(expected, actual) {
		return
	}
	msg := fmt.Sprintf(format, v...)
	msg2 := fmt.Sprintf("%s\nexpected %v (type %T), found %v (type %T)",
		msg, expected, expected, actual, actual)
	outputAndDie(2, msg2)
}
