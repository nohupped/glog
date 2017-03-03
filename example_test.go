// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package glog_test

import (
	"bytes"
	"fmt"
	log "glog"
)

func ExampleLogger() {
	var buf bytes.Buffer
	logger := log.New(&buf, "logger: ", log.Lshortfile)
	logger.Warnln("Hello, This is Error!") // This will not print because the default loglevel is set to ErrorLevel.
	logger.SetLogLevel(log.WarnLevel)
	logger.Errorln("Hello, This is Error!")
	logger.Warnln("Hello, This is Warn!")
	logger.Infoln("Hello, This is Info!") // This will not print because loglevel is set to WarnLevel
	logger.Debugln("Hello, This is Debug!") // This will not print because loglevel is set to WarnLevel
	fmt.Print(&buf)
	// Output:
	// logger: example_test.go:18: ERROR: Hello, This is Error!
	// logger: example_test.go:19: WARN: Hello, This is Warn!
}
