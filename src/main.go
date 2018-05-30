package main

import (
	"server"
	"router"
	"flag"
	"url_processor"
)
var port = flag.Int(
	"port",
	80,
	"tmt serve port (default: 80)")

func main() {
	flag.Parse()
	go url_processor.Start()
	server.Start(router.GetRouter(), port)
}