package main

import (
	"net/http"
)

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
