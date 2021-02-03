// Copyright (c) 2020 github.com/nyantlet
// Use of this source code is governed by a zlib-style
// license that can be found in the LICENSE file.

// +build !debugnya

package nya

// DebugTrue is debug-only version of True.
func DebugTrue(cond bool, msg string) {}

// DebugTruef is debug-only version of Truef.
func DebugTruef(cond bool, format string, v ...interface{}) {}

// DebugFalse is debug-only version of False.
func DebugFalse(cond bool, msg string) {}

// DebugFalsef is debug-only version of Falsef.
func DebugFalsef(cond bool, format string, v ...interface{}) {}

// DebugEqual is debug-only version of Equal.
func DebugEqual(expected, actual interface{}, msg string) {}

// DebugEqualf is debug-only version of Equalf.
func DebugEqualf(expected, actual interface{}, format string, v ...interface{}) {}

// DebugDeepEqual is debug-only version of DeepEqual.
func DebugDeepEqual(expected, actual interface{}, msg string) {}

// DebugDeepEqualf is debug-only version of DeepEqualf.
func DebugDeepEqualf(expected, actual interface{}, format string, v ...interface{}) {}
