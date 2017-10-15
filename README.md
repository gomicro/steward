# Steward
[![Build Status](https://travis-ci.org/gomicro/steward.svg)](https://travis-ci.org/gomicro/steward)
[![Go Reportcard](https://goreportcard.com/badge/github.com/gomicro/steward)](https://goreportcard.com/report/github.com/gomicro/steward)
[![GoDoc](https://godoc.org/github.com/gomicro/steward?status.png)](https://godoc.org/github.com/gomicro/steward)
[![License](https://img.shields.io/github/license/gomicro/steward.svg)](https://github.com/gomicro/steward/blob/master/LICENSE.md)
[![Release](https://img.shields.io/github/release/gomicro/steward.svg)](https://github.com/gomicro/steward/releases/latest)

Steward adds a drop in status endpoint with minimal effort for service health checks.

# Requirements
Golang v1.2 or higher

# Versioning
The project will be versioned in accordance with [Semver 2.0.0](http://semver.org).  See the [releases](https://github.com/gomicro/vacay/releases) section for the latest version.  Until version 1.0.0 the SDK is considered to be unstable.

# License
See [LICENSE.md](./LICENSE.md) for more information.

# Usage
To use steward, link this package into your program:
```go
import _ "github.com/gomicro/steward"
```

If your application is not already running an http server, you need to start one. Add "net/http" to your imports and the following code to your main function:
```go
go http.ListenAndServe("localhost:8080", nil)
```

# Examples
1. **[Example1](ext/examples/example1)** - Serve endpoint localhost:8000/status only returning status code
1. **[Example2](ext/examples/example2)** - Serve endpoint localhost:8000/status with custom JSON payload
