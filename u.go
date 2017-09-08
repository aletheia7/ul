// Copyright 2016 aletheia7. All rights reserved. Use of this source code is
// governed by a BSD-2-Clause license that can be found in the LICENSE file.

/*
Package ul provides macOS Sierra/OSX Unified Loggging functionality via cgo.

strings are logged as public. format strings are go fmt format strings.

See man os_log(3).

FYI: go log/syslog no longer works under macOS Sierra (10.12).
*/
package ul

import "fmt"

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
	os_log_t log_t
}

func New() *Logger {
	return new_logger(``, ``)
}

// Must call Release() to free subsystem/category
//
func New_object(subsystem, category string) *Logger {
	return new_logger(subsystem, category)
}

// limit: 1024
//
func (o *Logger) Log(s string) {
	o.log(Default, s)
}

// limit: 1024
//
func (o *Logger) Logf(format string, a ...interface{}) {
	o.log(Default, fmt.Sprintf(format, a...))
}

// limit: 256
//
func (o *Logger) Info(s string) {
	o.log(Info, s)
}

// limit: 256
//
func (o *Logger) Infof(format string, a ...interface{}) {
	o.log(Info, fmt.Sprintf(format, a...))
}

// limit: 256
//
func (o *Logger) Debug(s string) {
	o.log(Debug, s)
}

// limit: 256
//
func (o *Logger) Debugf(format string, a ...interface{}) {
	o.log(Debug, fmt.Sprintf(format, a...))
}

// limit: 256
//
func (o *Logger) Error(s string) {
	o.log(Error, s)
}

// limit: 256
//
func (o *Logger) Errorf(format string, a ...interface{}) {
	o.log(Error, fmt.Sprintf(format, a...))
}

// limit: 256
//
func (o *Logger) Fault(s string) {
	o.log(Fault, s)
}

// limit: 256
//
func (o *Logger) Faultf(format string, a ...interface{}) {
	o.log(Fault, fmt.Sprintf(format, a...))
}

// Satifies io.Writer. Can be used with the log package.
// Set Logger.Level for the default level
// log.SetOutput(*ul.Logger)
//
func (o *Logger) Write(p []byte) (n int, err error) {
	o.log(o.Level, string(p))
	return len(p), nil
}

func (o *Logger) Release() {
	o.release()
}
