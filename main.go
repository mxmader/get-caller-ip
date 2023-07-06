package main

import (
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"
)

func main(){

    fmt.Println("Initializing API")
    router := gin.Default()
    router.GET("/", getCallerIp)

    // NOTE: never run this service on "standard" HTTP (tcp/80) or HTTPS (tcp/443)
    // ports, as this will cause clients using some telco hotspot networks to receive
    // the outer IP of a "transparent proxy" imposed on such requests.
    // TODO: add link to source describing the above phenomenon.
    fmt.Println("Starting webserver")
    router.Run("0.0.0.0:9999")
}

func getCallerIp(c *gin.Context){

    // here, we want the "remote IP" of the client since we're not concerned
    // with HTTP semantics; clients are expected to use the same comms path
    // as their SSH client (i.e. bypass any and all HTTP proxies)
    c.String(http.StatusOK, c.RemoteIP())
}

