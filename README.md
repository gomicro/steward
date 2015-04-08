# Steward
Steward adds a status endpoint for service health checks.

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
**[Example1](ext/examples/example1)** - Serve endpoint localhost:8000/status only returning status code
**[Example2](ext/examples/example2)** - Serve endpoint localhost:8000/status with custom JSON payload
