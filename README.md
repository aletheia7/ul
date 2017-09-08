[![](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/aletheia7/ul) 

#### Install 
```bash
go get github.com/aletheia7/ul
go test -v
```

- Requires OSX Xcode compiler
- Made with go 1.9
- Supports macOS os_log subsystem/category logging
- Implements io.Writer and can be used with the go log package

#### Example

```go
package main

import (
	"github.com/aletheia7/ul"
	"log"
)

func main() {
	l := ul.New()
	l.Log("Hello")

	// To see subsystem/category messages:
	// # log stream --level debug --predicate 'subsystem == "com.example.myapp"'
	// Filtering the log data using "subsystem == "com.example.myapp""
	// Timestamp                       Thread     Type        Activity             PID
	// 2017-09-07 17:17:01.680996-0700 0x20870    Default     0x0                  1842   t: [com.example.myapp.whatever] hi

	lo := ul.New_object("com.example.myapp", "whatever")
	// Must call Release() for subsystem/category logger only
	defer lo.Release()
	lo.Log("hi")
	
	// golang log package

	mylogger := log.New(lo, "stuff ", log.Lshortfile|log.Ltime)
	mylogger.Println("wow") 
}
```

#### License 

Use of this source code is governed by a BSD-2-Clause license that can be found
in the LICENSE file.

[![BSD-2-Clause License](img/osi_logo_100X133_90ppi_0.png)](https://opensource.org/)
