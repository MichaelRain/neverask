package handlers

import (
	"github.com/fate-lovely/phi"
	"github.com/valyala/fasthttp"
)

func API() phi.Router {
	r := phi.NewRouter()

	r.Get("/messages/{articleID}", fetch)
	r.Post("/messages/", create)

	return r
}

func fetch(ctx *fasthttp.RequestCtx) {

}

func create(ctx *fasthttp.RequestCtx) {

}
