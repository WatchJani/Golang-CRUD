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

	log.Fatal(fasthttp.ListenAndServe(":5000", r.Handler))
}
