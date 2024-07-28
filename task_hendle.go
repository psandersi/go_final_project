package main

import "net/http"

//Функция TaskHandler обрабатывает запросы к /api/task и в зависимости от метода вызывает нужные функции
func TaskHandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	switch method {
	case http.MethodGet:
		TaskGet(w, r)
	case http.MethodPost:
		TaskPost(w, r)
	case http.MethodPut:
		TaskPut(w, r)
	case http.MethodDelete:
		TaskDelete(w, r)
	}
}
