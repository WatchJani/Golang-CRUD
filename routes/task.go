package routes

import (
	"encoding/json"
	"fmt"
	"log"

	"root/db"
	err "root/error"
	"root/model"

	"github.com/valyala/fasthttp"
)

func GetAllTask(ctx *fasthttp.RequestCtx) {
	var task []model.Task
	err.ErrorHandler(db.DB.Select(&task, ALL_TASKS))

	response := struct {
		Length   int
		AllTasks []model.Task
	}{
		Length:   len(task),
		AllTasks: task,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	ctx.SetContentType("application/json")
	ctx.SetBody(jsonResponse)
}

func InsertTask(ctx *fasthttp.RequestCtx) {

	requestBody := ctx.PostBody()

	request := struct {
		Title string
		Task  string
	}{}

	err := json.Unmarshal(requestBody, &request)
	if err != nil {
		fmt.Println("Greška prilikom parsiranja JSON-a:", err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString("Neispravan JSON format")
		return
	}

	_, err = db.DB.Exec(SET_POST, request.Title, request.Task)
	if err != nil {
		fmt.Println("Greška prilikom umetanja podataka u bazu:", err)
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetBodyString("Greška prilikom umetanja podataka u bazu")
		return
	}

	fmt.Println("Primljeni podaci:", request)

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBodyString("Zadatak uspješno umetnut")
}

func DeleteTask(ctx *fasthttp.RequestCtx) {

	taskID := ctx.UserValue("id").(string)

	_, err := db.DB.Exec(DELETE_TASK, taskID)
	if err != nil {
		log.Fatal(err)
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBodyString("Zadatak uspješno obrisan")
}

func UpdateTask(ctx *fasthttp.RequestCtx) {
	taskID := ctx.UserValue("id").(string)

	requestBody := ctx.PostBody()

	request := struct {
		Title string
		Task  string
	}{}

	err := json.Unmarshal(requestBody, &request)
	if err != nil {
		fmt.Println("Greška prilikom parsiranja JSON-a:", err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString("Neispravan JSON format")
		return
	}

	_, err = db.DB.Exec(UPDATE_TASK, request.Title, request.Task, taskID)
	if err != nil {
		log.Fatal(err)
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBodyString("Zadatak uspješno ažuriran")
}
