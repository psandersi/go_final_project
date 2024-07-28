package main

import (
	"net/http"
)

// Функция TaskDelete обрабатывает DELETE-запросы к /api/task/done. Возвращает {} или, в случае ошибки, JSON с полем error.
func TaskDelete(w http.ResponseWriter, r *http.Request) {
	var err error

	id := r.FormValue("id")

	err = DeleteTask(id)
	if err != nil {
		jsonError(err, w)
		return
	}
	emptyJson(w)

}
