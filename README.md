# Logger

A wrapper for [zerolog](https://github.com/rs/zerolog) with a simplified API that encourages
passing the logger through the [context](https://pkg.go.dev/context#Context).

Provides JSON (default) and plain text output.

## grpclog

The library implements [grpclog.LoggerV2](https://pkg.go.dev/google.golang.org/grpc/grpclog#LoggerV2).

```go
package main

import (
	"os"

	"github.com/cdnnow-pro/logger-go"
	"google.golang.org/grpc/grpclog"
)

func init() {
	grpclog.SetLoggerV2(log.NewGrpcLogger(log.InfoLevel,
		log.WithGrpcOutput(os.Stderr),
	))
}
```
