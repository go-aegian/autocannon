package main

import (
	"log"
	"unsafe"

	"github.com/dgrr/http2"
	"github.com/sainzg/autocannon/webservers/go"
	"github.com/valyala/fasthttp"
)

func b2s(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	path := b2s(ctx.Path())

	switch path {
	case "/":
		_, _ = ctx.Write(weather.Predict(5))
	}
}

func main() {
	s := &fasthttp.Server{
		Handler: fastHTTPHandler,
		Name:    "HTTP2 test",
	}

	http2.ConfigureServer(s, http2.ServerConfig{})

	log.Fatal(s.ListenAndServe(":3006"))
}
