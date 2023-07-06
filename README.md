# About

This is a simple HTTP server that will return the caller (client) IP address without any regard for proxy semantics.

It listens on port `9999` to avoid telco "transparent proxies" which munges server-side overvability.

This is a primitive way of determining what "WAN IP" an SSH server hosted on a public cloud server will see from the given client.

# Running locally

```
➜  get-caller-ip git:(main) ✗ go run .
Initializing API
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /                         --> main.getCallerIp (3 handlers)
Starting webserver
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on 0.0.0.0:9999
[GIN] 2023/07/06 - 10:05:09 | 200 |       66.86µs |       127.0.0.1 | GET      "/"
[GIN] 2023/07/06 - 10:05:11 | 200 |       37.66µs |       127.0.0.1 | GET      "/"
[GIN] 2023/07/06 - 10:05:12 | 200 |      16.796µs |       127.0.0.1 | GET      "/"
```

