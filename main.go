package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	dbPath := "../scheduler.db"
	//Создаем или открываем базу данных
	dbStorage, err := DoNewTable(dbPath)
	if err != nil {
		log.Println(err)
	}
	defer dbStorage.dbClose()
	//настраиваем сервер
	r := chi.NewRouter()
	r.Handle("/*", http.FileServer(http.Dir("web")))
	r.Get("/api/nextdate", GetNextDate)
	r.Get("/api/tasks", TasksGet)
	r.HandleFunc("/api/task", TaskHandler)
	r.Post("/api/task/done", PostTaskDone)
	if err := http.ListenAndServe(":7540", r); err != nil {
		fmt.Printf("Error while opening server: %s", err.Error())
		return
	}
}
