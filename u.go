// Copyright 2016 aletheia7. All rights reserved. Use of this source code is
// governed by a BSD-2-Clause license that can be found in the LICENSE file.

/*
Package ul provides macOS Sierra/OSX Unified Loggging functionality via cgo.

strings are logged as public. format strings are go fmt format strings.

See man os_log(3).

FYI: go log/syslog no longer works under macOS Sierra (10.12).
*/

package ul

/*
#include <stdlib.h>
#include <os/log.h>

const os_log_t os_log_default = OS_LOG_DEFAULT;

void ul_log(unsigned char level, os_log_t log, const char* const s) {
	if(level == 0) {
		os_log(log, "%{public}s", s);
	}
	else if(level == 1) {
		os_log_info(log, "%{public}s", s);
	}
	else if(level == 2) {
		os_log_debug(log, "%{public}s", s);
	}
	else if(level == 3) {
		os_log_error(log, "%{public}s", s);
	}
	else if(level == 4) {
		os_log_fault(log, "%{public}s", s);
	}
}

void release(os_log_t t) {
	os_release(t);
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type level uint8

const (
	Default level = iota
	Info
	Debug
	Error
	Fault
)

type Logger struct {
	Subsystem string
	Category  string
	// io.Writer Default level
	Level    level
	os_log_t C.os_log_t
}

func New() *Logger {
	return New_object(``, ``)
}

// Must call Release() to free subsystem/category
//
func New_object(subsystem, category string) *Logger {
	r := &Logger{
		Subsystem: subsystem,
		Category:  category,
	}
	if 0 < len(subsystem) || 0 < len(category) {
		s := C.CString(subsystem)
		c := C.CString(category)
		r.os_log_t = C.os_log_create(s, c)
		C.free(unsafe.Pointer(s))
		C.free(unsafe.Pointer(c))
	} else {
		r.os_log_t = C.os_log_default
	}
	return r
}

// limit: 1024
//
func (o *Logger) Log(s string) {
	cs := C.CString(s)
	C.ul_log((C.uchar)(Default), o.os_log_t, cs)
	C.free(unsafe.Pointer(cs))
}

// limit: 1024
//
func (o *Logger) Logf(format string, a ...interface{}) {
	cs := C.CString(fmt.Sprintf(format, a...))
	C.ul_log((C.uchar)(Default), o.os_log_t, cs)
	C.free(unsafe.Pointer(cs))
}

// limit: 256
//
func (o *Logger) Info(s string) {
	cs := C.CString(s)
	C.ul_log((C.uchar)(Info), o.os_log_t, cs)
	C.free(unsafe.Pointer(cs))
}

// limit: 256
//
func (o *Logger) Infof(format string, a ...interface{}) {
	cs := C.CString(fmt.Sprintf(format, a...))
	C.ul_log((C.uchar)(Info), o.os_log_t, cs)
	C.free(unsafe.Pointer(cs))
}

// limit: 256
//
func (o *Logger) Debug(s string) {
	cs := C.CString(s)
	C.ul_log((C.uchar)(Debug), o.os_log_t, cs)
	C.free(unsafe.Pointer(cs))
}

// limit: 256
//
func (o *Logger) Debugf(format string, a ...interface{}) {
	cs := C.CString(fmt.Sprintf(format, a...))
	C.ul_log((C.uchar)(Debug), o.os_log_t, cs)
	C.free(unsafe.Pointer(cs))
}

// limit: 256
//
func (o *Logger) Error(s string) {
	cs := C.CString(s)
	C.ul_log((C.uchar)(Error), o.os_log_t, cs)
	C.free(unsafe.Pointer(cs))
}

// limit: 256
//
func (o *Logger) Errorf(format string, a ...interface{}) {
	cs := C.CString(fmt.Sprintf(format, a...))
	C.ul_log((C.uchar)(Error), o.os_log_t, cs)
	C.free(unsafe.Pointer(cs))
}

// limit: 256
//
func (o *Logger) Fault(s string) {
	cs := C.CString(s)
	C.ul_log((C.uchar)(Fault), o.os_log_t, cs)
	C.free(unsafe.Pointer(cs))
}

// limit: 256
//
func (o *Logger) Faultf(format string, a ...interface{}) {
	cs := C.CString(fmt.Sprintf(format, a...))
	C.ul_log((C.uchar)(Fault), o.os_log_t, cs)
	C.free(unsafe.Pointer(cs))
}

// Satifies io.Writer. Can be used with the log package.
// Set Logger.Level for the default level
// log.SetOutput(*ul.Logger)
//
func (o *Logger) Write(p []byte) (n int, err error) {
	b := C.CBytes(p)
	C.ul_log((C.uchar)(o.Level), o.os_log_t, (*C.char)(b))
	C.free(b)
	return len(p), nil
}

func (o *Logger) Release() {
	C.os_release(unsafe.Pointer(o.os_log_t))
}
