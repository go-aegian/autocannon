package main

import (
	"fmt"

	"github.com/dgrr/http2"
	"github.com/valyala/fasthttp"
)

// this is not as router but just a simple straight handler

func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "FastHttp webserver running...")
}

func main() {
	s := &fasthttp.Server{
		Handler: fastHTTPHandler,
		Name:    "HTTP2 test",
	}

	http2.ConfigureServer(s, http2.ServerConfig{})

	s.ListenAndServe(":3006")
}
