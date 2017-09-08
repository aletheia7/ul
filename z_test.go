// Copyright 2016 aletheia7. All rights reserved. Use of this source code is
// governed by a BSD-2-Clause license that can be found in the LICENSE file.

package ul_test

import (
	. "github.com/aletheia7/ul"
	"log"
	"strconv"
	"testing"
)

var i = 0
var l = New()
var lo = New_object("com.github/aletheia/ul", "test")

func Test_log(t *testing.T) {
	const pre = "ul log "
	t.Log(`To see resuls exec: log stream --level debug --predicate 'processImagePath endswith "ul.test"'`)
	l.Log(pre + next())
	l.Logf(pre+"%v", next())
	lo.Log(pre + next())
	lo.Logf(pre+"%v", next())
	lo.Release()
}

func Test_info(t *testing.T) {
	const pre = "ul info "
	l.Info(pre + next())
}

func Test_debug(t *testing.T) {
	const pre = "ul debug "
	l.Debug(pre + next())
}

func Test_error(t *testing.T) {
	const pre = "ul error "
	l.Error(pre + next())
}

func Test_fault(t *testing.T) {
	const pre = "ul fault "
	l.Fault(pre + next())
}

func Test_writer(t *testing.T) {
	const pre = "ul io.writer "
	go_log := log.New(l, pre, log.Llongfile)
	go_log.Println(next())
	l.Level = Error
	go_log.Println("error " + next())

}

func next() string {
	i++
	return strconv.Itoa(i)
}
