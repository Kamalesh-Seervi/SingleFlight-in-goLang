package main

import (
	"Kamalesh-Seervu/singleflight/client"
	"Kamalesh-Seervu/singleflight/server"
	"time"
)

func main() {
	server.ServerSetup()
	time.Sleep(5 * time.Second)
	client.ClientEndPoint()
}
