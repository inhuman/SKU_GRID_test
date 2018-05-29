package main

import (
	"server"
	"router"
	"flag"
)
var port = flag.Int(
	"port",
	80,
	"tmt serve port (default: 80)")

func main() {
	flag.Parse()
	server.Start(router.GetRouter(), port)
}