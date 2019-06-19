# slf4go-logrus [![Build Status](https://travis-ci.org/go-eden/slf4go-logrus.svg?branch=master)](https://travis-ci.org/go-eden/slf4go-logrus)

This is a default [`Driver`](https://github.com/go-eden/slf4go#driver) implementation for bridging [`slf4go`](https://github.com/go-eden/slf4go) and [`logrus`](https://travis-ci.org/sirupsen/logrus).

# Install

`slf4go-logrus` dependents on `slf4go` and `logrus`, you don't need to dependent them directly: 

```
go get github.com/go-eden/slf4go-logrus
```

# Usage

`slf4go-logrus` focuses on bridging logs, you should configure `logrus` according to your needs.  

This is a simple example of `slf4go-logrus`:

```go
package main

import (
	slog "github.com/go-eden/slf4go"
	sl "github.com/go-eden/slf4go-logrus"
	"github.com/sirupsen/logrus"
	"os"
)

func init() {
	sl.Init()

	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.InfoLevel)
}

func main() {
	slog.Debug("hello")
	slog.Info("what???")
	slog.Warnf("warnning: %v", "surrender")

	log := slog.GetLogger()
	log.BindFields(slog.Fields{"logger": "test"})
	log.Errorf("error!!! %v", 100)
}
```

# Account

I haven't used `logrus` too much, so this library don't have lots of features currently.

Hope you can help me improve this library, any `Pull Request` will be very welcomed.

# License

MIT