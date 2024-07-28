package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Функция writeTasksGet возвращает ответ обработчика TasksGet: список ближайщих задач или ошибку
func writeTasksGet(tasks []Task, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var res []byte
	if err != nil {
		jsonError(err, w)
		return
	} else {
		if len(tasks) == 0 {
			tasksResp := map[string][]Task{
				"tasks": {},
			}
			res, err = json.Marshal(tasksResp)
		} else {
			tasksResp := map[string][]Task{
				"tasks": tasks,
			}
			res, err = json.Marshal(tasksResp)

		}

		if err != nil {
			log.Println(err)
		}
		w.WriteHeader(http.StatusCreated)
		_, err = w.Write(res)
		if err != nil {
			log.Println(err)
		}
		return
	}
}

// Функция TasksGet обрабатывает GET-запросы к /api/tasks, вовзращает список ближайших задач в формате JSON в виде списка в поле tasks
// При ошибке возвращается JSON с полем error.
func TasksGet(w http.ResponseWriter, r *http.Request) {
	var tasks []Task
	var err error

	tasks, err = GetTasksList()
	if err != nil {
		log.Println(err)
	}

	writeTasksGet(tasks, err, w)
}
