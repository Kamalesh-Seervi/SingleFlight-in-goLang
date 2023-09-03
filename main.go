package main

import (
	"Kamalesh-Seervu/singleflight/client"
	"Kamalesh-Seervu/singleflight/server"
	"runtime"
)

func main() {
	server.ServerSetup()
	runtime.GOMAXPROCS(runtime.NumCPU())
	client.ClientEndPoint()
}
