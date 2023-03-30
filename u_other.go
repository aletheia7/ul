// Copyright 2022 aletheia7. All rights reserved. Use of this source code is
// governed by a BSD-2-Clause license that can be found in the LICENSE file.

//go:build !apple

/*
Package ul provides macOS Sierra/OSX Unified Loggging functionality via cgo.

strings are logged as public. format strings are go fmt format strings.

See man os_log(3).

FYI: go log/syslog no longer works under macOS Sierra (10.12).
*/
package ul

import "C"
import (
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
func (o *Logger) Log(s string) {
	//pass
}

// limit: 1024
func (o *Logger) Logf(format string, a ...interface{}) {
	//pass
}

// limit: 256
func (o *Logger) Info(s string) {
	//pass
}

// limit: 256
func (o *Logger) Infof(format string, a ...interface{}) {
	//pass
}

// limit: 256
func (o *Logger) Debug(s string) {
	//pass
}

// limit: 256
func (o *Logger) Debugf(format string, a ...interface{}) {
	//pass
}

// limit: 256
func (o *Logger) Error(s string) {
	//pass
}

// limit: 256
func (o *Logger) Errorf(format string, a ...interface{}) {
	//pass
}

// limit: 256
func (o *Logger) Fault(s string) {
	//pass
}

// limit: 256
func (o *Logger) Faultf(format string, a ...interface{}) {
	//pass
}

// Satifies io.Writer. Can be used with the log package.
// Set Logger.Level for the default level
// log.SetOutput(*ul.Logger)
func (o *Logger) Write(p []byte) (n int, err error) {
	//pass
	return len(p), nil
}

func (o *Logger) Release() {
	//pass
}
