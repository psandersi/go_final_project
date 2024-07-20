package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/psandersi/go_final_project/database"
)

func main() {
	dbPath := "C:/Users/dashu/dev/go_final_project/scheduler.db"
	dbStorage, err := database.DoNewTable(dbPath)
	if err != nil {
		log.Println(err)
	}
	err = dbStorage.dbClose()
	if err != nil {
		log.Println(err)
	}

	r := chi.NewRouter()
	r.Handle("/*", http.FileServer(http.Dir("web")))
	if err := http.ListenAndServe(":7540", r); err != nil {
		fmt.Printf("Ошибка при запуске сервера: %s", err.Error())
		return
	}
}
