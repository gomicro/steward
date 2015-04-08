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
