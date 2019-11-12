package main

import (
    "os"
    "./server"
    "github.com/spf13/pflag"
)

func main() {
    port := os.Getenv("PORT")
    // set default port
    if (len(port) == 0) {
        port = "1234"
    }
    // parse port 
    currentPort := pflag.StringP("port", "p", "1234", "Port for http listening")
    pflag.Parse()
    if (len(*currentPort) != 0) {
        port = *currentPort
    }
    // run server
    server.Run(port)
}