package main

import (
	"fmt"
	"log"
	"unsafe"

	"github.com/dgrr/http2"
	"github.com/valyala/fasthttp"
)

func b2s(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// func s2b(s string) (b []byte) {
// 	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
// 	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
// 	bh.Data = sh.Data
// 	bh.Cap = sh.Len
// 	bh.Len = sh.Len
// 	return b
// }

func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	path := b2s(ctx.Path())

	switch path {
	case "/":
		_, _ = fmt.Fprint(ctx, "FastHttp webserver running...")
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
