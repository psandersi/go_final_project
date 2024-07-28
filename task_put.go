package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// Функция emptyJson делает пустой json
func emptyJson(w http.ResponseWriter) {
	ok := map[string]string{}
	res, err := json.Marshal(ok)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)
	if err != nil {
		log.Println(err)
	}
}

// Функция writeTaskPut возвращает ошибку в виде json или пустой json
func writeTaskPut(task Task, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		jsonError(err, w)
		return
	} else {
		emptyJson(w)
		return
	}
}

// Функция TaskPut обрабатывает PUT-запросы к /api/task, а именно редактирует задачу.
// В случае успешного изменения возвращается пустой JSON {}, а в случае ошибки, она записывается в поле error.
func TaskPut(w http.ResponseWriter, r *http.Request) {
	var task Task
	var buf bytes.Buffer
	var err error

	_, err = buf.ReadFrom(r.Body)
	if err != nil {
		writeTaskPut(task, err, w)
		return
	}

	if err = json.Unmarshal(buf.Bytes(), &task); err != nil {
		writeTaskPut(task, err, w)
		return
	}

	task, err = task.checkData()
	if err != nil {
		writeTaskPut(task, err, w)
		return
	}

	err = PutTask(task)
	writeTaskPut(task, err, w)
}
