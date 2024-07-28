package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Функция writeTaskGet возвращает ответ обработчика TaskGet или ошибку
func writeTaskGet(task Task, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var res []byte
	if err != nil {
		jsonError(err, w)
		return
	} else {
		res, err = json.Marshal(task)
	}

	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)
	if err != nil {
		log.Println(err)
	}
}

// Функция TaskGet обрабатывает GET-запросы к /api/task, а именно возвращает задачу по ее id.
func TaskGet(w http.ResponseWriter, r *http.Request) {
	var err error
	var task Task

	id := r.FormValue("id")
	task, err = GetTaskByID(id)
	if err != nil {
		log.Println(err)
	}
	writeTaskGet(task, err, w)
}
