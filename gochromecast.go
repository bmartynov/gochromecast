package main

import (
	"github.com/zenazn/goji"
	"github.com/bmartynov/gochromecast/http_server"
)



func main() {
	mux := http_server.New()
	goji.Handle("/*", mux)
	goji.Serve()

}