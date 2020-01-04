# glog

Go's modified log package with an option to set a loglevel, which will filter the output of logs based on that.
(This is a lame attempt to modify GO's original [`log` package](https://github.com/golang/go/tree/master/src/log) (version 1.6.2)), and the logic borrowed from [Sirupsen](https://github.com/Sirupsen/logrus) package.

## Why

I was trying to find a way to log `filename` and `line number` using logrus for obvious debugging purpose, but I failed to get a workaround. I am not using advanced features of `logrus` like `SetFormatter` or `hooks`. `Go`'s native `log` package already provides a feature to log filename and line numbers but it lacks setting a loglevel for filtering logs. I just borrowed and incorporated the idea of using a `loglevel`.
Newly added methods for `*Logger` like `(Error|Warn|Info|Debug)ln(), (Error|Warn|Info|Debug)f() and (Error|Warn|Info|Debug)()` will only output logs based on the configured loglevel with the helper function `SetLogLevel()`. Accepted log levels are `ErrorLevel, WarnLevel, InfoLevel and DebugLevel` which are of values `int 0, 1, 2, 3` respectively.

Default value of `loglevel` at the time of initing the `*Logger` is set to the lowest level, which is `ErrorLevel`.

## Breaking change because of using pointer to loglevel

Use

```bash
require github.com/nohupped/glog v0.0.0-20200102070319-ef78151f855d
```

in the go.mod file to continue using the older version that uses `int` as `loglevel` instead of `*uint`.

* why? To have the convenience of using this logger in `init()` functions and have a global variable that can be modified without calling the `SetStandardLogLevel` or similar functions in every importing sub-packages.

Eg: Package `foo/blah`

```golang
package blah
import (
     // "flag"
    log "github.com/nohupped/glog"
)

// Loglevel is the logging handler
var Loglevel *uint

func init() {
    // loglevel := flag.Uint("loglevel", uint(glog.InfoLevel), "Loglevel.")
    // flag.Parse()
    // Don't do the above 2 lines in init functions. This will mess with go's testing because of https://github.com/golang/go/issues/33774.
    // flag.Parse() and init() functions doesn't go well.
    if Loglevel == nil {
        l := log.ErrorLevel
        Loglevel = &l
    }
    log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
    log.SetStandardLogLevel(Loglevel)
    log.Infof("Setting loglevel to %d", log.InfoLevel)
}

```

and in package foo,

```golang
package main

import(
    "foo/blah"
)
func main(){
    loglevel := flag.Uint("loglevel", log.InfoLevel, "Loglevel.")
    flag.Parse()
    *blah.Loglevel = *loglevel // Dereference and modify the pointer value to update the loglevel globally on all the packages that imports foo/blah.
}
```

Now, for another sub-package `foo/extreme_computation`,

```golang
package extreme_computation
import(
    _ "foo/blah"
    log "github.com/nohupped/glog"
)

func PvsNP() {

}

```

it will see the updated loglevel from the `main` although the initial loglevel was set in `blah`, and this is because of the `var Loglevel *uint`

### Reference when setting loglevel

`logger.SetLogLevel(log.ErrorLevel)` will log only Errors using the method `log.Error*()`. Any calls to `log.Warn*()`, `log.Info*()` and `log.Debug*()` will not be logged unless the `loglevel` is changed.

`logger.SetLogLevel(log.WarnLevel)` will log Errors and Warnings. Rest are not logged.

`logger.SetLogLevel(log.InfoLevel)` will log Errors, Warnings and Info. Rest are not logged.

`logger.SetLogLevel(log.DebugLevel)` will log Errors, Warnings, Info and Debug.

Note: To disable logging line numbers and file names, set the logger's flag similar to `logger := log.New(&buf, "logger: ", log.Ldate)` instead of `log.Lshortfile` or similar flags.

#### Example

```golang
package main

import (
    "bytes"
    "fmt"
    log "github.com/nohupped/glog"
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
    il := log.InfoLevel
    logger.SetLogLevel(&iL)
    logger.Errorln("Hello, This is Error from Errorln!")
    logger.Warnln("Hello, This is Warn from Warnln!")
    logger.Infoln("Hello, This is Info from Infoln!")
    logger.Debugln("Hello, This is Debug from Debugln!") // This will not print because loglevel is set to InfoLevel
    dl := log.DebugLevel
    log.SetStandardLogLevel(&dl) // set std loglevel. This will NOT set the *Logger struct's loglevel.
    log.SetFlags(log.Lshortfile) // set flags to log. This will add short filename and line number.
    log.Debugf("This is from %s", strfn)
    log.Printf("%d error from %s", 1, "Error function")
    log.Warnf("%d error from %s", 1, "Error function")
    fmt.Println(&buf)

    fmt.Print(&buf)

    // Output:
    // logger: main.go:19: ERROR: Hello, this is 1 Error from Errorf()!
    // logger: main.go:20: ERROR: Hello, this is 1 Error from Error()!
    // logger: main.go:23: ERROR: Hello, This is Error from Errorln!
    // logger: main.go:24: WARN: Hello, This is Warn from Warnln!
    // logger: main.go:25: INFO: Hello, This is Info from Infoln!

```
`* Line numbers may vary.`
