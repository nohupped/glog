// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package glog_test

import (
	"bytes"
	log"github.com/nohupped/glog"
	"fmt"
)

func ExampleLogger() {
	var buf bytes.Buffer
	log.Println()
	logger := log.New(&buf, "logger: ", log.Lshortfile)
	//file, _ := os.OpenFile("/tmp/testlog.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//logger.SetOutput(file)
	strfn := "Errorf()!"
	logger.Errorf("Hello, this is %d Error from %s", 1, strfn)
	logger.Error("Hello, this is 1 Error from Error()!")
	logger.Warnln("Hello, This is from Warnln!") // This will not print because the default loglevel is set to ErrorLevel.
	logger.SetLogLevel(log.InfoLevel)
	logger.Errorln("Hello, This is Error from Errorln!")
	logger.Warnln("Hello, This is Warn from Warnln!")
	logger.Infoln("Hello, This is Info from Infoln!")
	logger.Debugln("Hello, This is Debug from Debugln!") // This will not print because loglevel is set to InfoLevel
	log.SetStandardLogLevel(log.DebugLevel)
	log.SetFlags(log.Lshortfile)
	log.Debugf("This is from %s", strfn)
	log.Printf("%d error from %s", 1, "Error function")
	log.Warnf("%d error from %s", 1, "Error function")
	fmt.Println(&buf)


	// Output:
	// logger: example_test.go:20: ERROR: Hello, this is 1 Error from Errorf()!
	// logger: example_test.go:21: ERROR: Hello, this is 1 Error from Error()!
	// logger: example_test.go:24: ERROR: Hello, This is Error from Errorln!
	// logger: example_test.go:25: WARN: Hello, This is Warn from Warnln!
	// logger: example_test.go:26: INFO: Hello, This is Info from Infoln!
	

}
