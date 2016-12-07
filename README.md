# Steward
[![Build Status](https://travis-ci.org/gomicro/steward.svg)](https://travis-ci.org/gomicro/steward)
[![Coverage](http://gocover.io/_badge/github.com/gomicro/steward)](http://gocover.io/github.com/gomicro/steward)
[![GoDoc](https://godoc.org/github.com/gomicro/steward?status.png)](https://godoc.org/github.com/gomicro/steward)

Steward adds a drop in status endpoint with minimal effort for service health checks.

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
