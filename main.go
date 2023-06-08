package main

import (
	"log"
	"root/db"

	handler "root/routes"

	"github.com/fasthttp/router"
	_ "github.com/lib/pq"
	"github.com/valyala/fasthttp"
)

func main() {
	db.Connect()

	r := router.New()

	r.GET("/", handler.GetAllTask)
	r.POST("/", handler.InsertTask)
	r.PUT("/{id}", handler.UpdateTask)
	r.DELETE("/{id}", handler.DeleteTask)

	log.Fatal(fasthttp.ListenAndServe(":5000", handleCORS(r.Handler)))
}

func handleCORS(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		// Postavi zaglavlja za podršku CORS
		ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
		ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		ctx.Response.Header.Set("Access-Control-Allow-Headers", "Content-Type")

		// Provjeri zahtjev OPTIONS za CORS pre-flight
		if string(ctx.Method()) == fasthttp.MethodOptions {
			ctx.SetStatusCode(fasthttp.StatusOK)
			return
		}

		// Pozovi sljedeći rukovatelj zahtjeva
		h(ctx)
	})
}
