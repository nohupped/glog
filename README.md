# glog
Go's modified log package with an option to set and filter loglevels.
(This is modified code from GO's `log` package at https://github.com/golang/go/tree/master/src/log (version 1.6.2))

Newly added functions like `(Error|Warn|Info|Debug)ln(), (Error|Warn|Info|Debug)f() and (Error|Warn|Info|Debug)()` will only output logs to output based on the configured loglevel with the helper function `SetLogLevel()`. Accepted log levels are `ErrorLevel, WarnLevel, InfoLevel and DebugLevel` which are of values `int 0, 1, 2, 3` respectively.

##### Example:

```
package main

import (
	"bytes"
	"fmt"
	log "https://github.com/nohupped/glog"
)

func main() {
	var buf bytes.Buffer
	logger := log.New(&buf, "logger: ", log.Lshortfile)
	// Below two lines to set logging to file
	//file, _ := os.OpenFile("/tmp/testlog.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//logger.SetOutput(file)
	strfn := "Errorf()!"
	logger.Errorf("Hello, this is %d Error from %s", 1, strfn)
	logger.Error("Hello, this is 1 Error from Error()!")
	logger.Warnln("Hello, This is from Warnln!") // This will not print because the default loglevel when initiating logger is set to ErrorLevel.
	logger.SetLogLevel(log.InfoLevel)
	logger.Errorln("Hello, This is Error from Errorln!")
	logger.Warnln("Hello, This is Warn from Warnln!")
	logger.Infoln("Hello, This is Info from Infoln!")
	logger.Debugln("Hello, This is Debug from Debugln!") // This will not print because loglevel is set to InfoLevel

	fmt.Print(&buf)

	// Output:
	// logger: example_test.go:19: ERROR: Hello, this is 1 Error from Errorf()!
	// logger: example_test.go:20: ERROR: Hello, this is 1 Error from Error()!
	// logger: example_test.go:23: ERROR: Hello, This is Error from Errorln!
	// logger: example_test.go:24: WARN: Hello, This is Warn from Warnln!
	// logger: example_test.go:25: INFO: Hello, This is Info from Infoln!

```
