// Copyright 2016 aletheia7. All rights reserved. Use of this source code is
// governed by a BSD-2-Clause license that can be found in the LICENSE file.

package ul

// This file has the C package and allows the go tools/guru package to complete
// u.go with vim-go

/*
#include <stdlib.h>
#include <os/log.h>

const os_log_t os_log_default = OS_LOG_DEFAULT;

void ul_log(os_log_t log, const char* const s) {
	os_log(log, "%{public}s", s);
}

void ul_log_info(os_log_t log, const char* const s) {
	os_log_info(log, "%{public}s", s);
}

void ul_log_debug(os_log_t log, const char* const s) {
	os_log_debug(log, "%{public}s", s);
}

void ul_log_error(os_log_t log, const char* const s) {
	os_log_error(log, "%{public}s", s);
}

void ul_log_fault(os_log_t log, const char* const s) {
	os_log_fault(log, "%{public}s", s);
}

void release(os_log_t t) {
	os_release(t);
}
*/
import "C"
import "unsafe"

type log_t C.os_log_t

func new_logger(subsystem, category string) *Logger {
	r := &Logger{
		Subsystem: subsystem,
		Category:  category,
	}
	if 0 < len(subsystem) || 0 < len(category) {
		s := C.CString(subsystem)
		c := C.CString(category)
		r.os_log_t = log_t(C.os_log_create(s, c))
		C.free(unsafe.Pointer(s))
		C.free(unsafe.Pointer(c))
	} else {
		r.os_log_t = log_t(C.os_log_default)
	}
	return r
}

func (o *Logger) log(l level, s string) {
	cs := C.CString(s)
	switch l {
	case Default:
		C.ul_log(o.os_log_t, cs)
	case Info:
		C.ul_log_info(o.os_log_t, cs)
	case Debug:
		C.ul_log_debug(o.os_log_t, cs)
	case Error:
		C.ul_log_error(o.os_log_t, cs)
	case Fault:
		C.ul_log_fault(o.os_log_t, cs)
	}
	C.free(unsafe.Pointer(cs))
}

func (o *Logger) release() {
	C.release(o.os_log_t)
}
