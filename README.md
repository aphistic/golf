golf
====

[![GoDoc](https://godoc.org/github.com/aphistic/golf?status.svg)](https://godoc.org/github.com/aphistic/golf)
[![Build Status](https://travis-ci.org/aphistic/golf.svg?branch=master)](https://travis-ci.org/aphistic/golf)
[![codecov.io](http://codecov.io/github/aphistic/golf/coverage.svg?branch=master)](http://codecov.io/github/aphistic/golf?branch=master)

Golf is an MIT-licensed Go client library for servers supporting the Graylog
Extended Log Format (GELF, https://www.graylog.org/resources/gelf-2/).

If you run into any issues with the library or have any feature requests, please create issues for them!

As this library is very new, the API could still change.  I don't expect it to change much because I'm pretty happy with how it is now but if anyone has suggestions easier ways to use the library via API changes I would be open to it.

Test coverage is an ongoing process!

Features
========

* GELF 1.1 support
* Native Go implementation
* Supports Logger-level and Message-level attributes

Installation
============

The recommended way to install is via http://gopkg.in

    go get gopkg.in/aphistic/golf.v0
    ...
    import "gopkg.in/aphistic/golf.v0"

Golf can also be installed the standard way as well

    go get github.com/aphistic/golf
    ...
    import "github.com/aphistic/golf"

Examples
========

For brevity a lot of error checking has been omitted from these examples, be sure you do your checks!

The standard way to implement the golf library is by creating a Client, connecting to a server and creating Loggers off that Client:

```go
package main

import (
    "gopkg.in/aphistic/golf.v0"
)

func main() {
    c, _ := golf.NewClient()
    c.Dial("udp://192.168.30.150")

    l, _ := c.NewLogger()
    // Attributes set at the Logger level will automatically be included
    // on each message sent from that Logger.  This is helpful if there's
    // any consistent information you don't want to include every time you
    // log a message.
    l.SetAttr("facility", "golf.example")
    l.SetAttr("instance_id", 12345)

    for idx := 1; idx <= 10; idx++ {
        l.Dbgm(map[string]interface{}{
            "msg_attr1": 1234,
        }, "Test message %v", idx)
    }

    c.Close()
}
```

It is also possible to set a Logger as the default for the golf library so you don't need to keep track of a main Logger manually:

```go
package main

import (
    "gopkg.in/aphistic/golf.v0"
)

func main() {
    c, _ := golf.NewClient()
    c.Dial("udp://192.168.30.150")

    l, _ := c.NewLogger()
    // Set l as the default logger
    golf.DefaultLogger(l)
    // Attributes set at the Logger level will automatically be included
    // on each message sent from that Logger.  This is helpful if there's
    // any consistent information you don't want to include every time you
    // log a message.
    l.SetAttr("facility", "golf.example")
    l.SetAttr("instance_id", 12345)

    for idx := 1; idx <= 10; idx++ {
        // Use the default logger to send the message
        golf.Dbgm(map[string]interface{}{
            "msg_attr1": 1234,
        }, "Test message %v", idx)
    }

    c.Close()
}
```

You can use the query parameter "compress" in the Dial URL, with one of the following value:

* none
* zlib
* gzip

```
udp://192.168.30.150?compress=none
```

Default is gzip compression.
