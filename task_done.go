package main

import (
	"net/http"
	"time"
)

// Функция PostTaskDone обрабатывает POST-запросы к /api/task/done, который делает задачу выполненной.
// Одноразовая задача с пустым полем repeat удаляется. Для периодической задачи рассчитывает и поменяет дату следующего выполнения
// В случае успешного удаления возвращается пустой JSON, а в случае ошибки - JSON с полем error.
func PostTaskDone(w http.ResponseWriter, r *http.Request) {
	var err error

	id := r.FormValue("id")

	task, err := GetTaskByID(id)
	if err != nil {
		jsonError(err, http.StatusInternalServerError, w)
		return
	}
	if len(task.Repeat) == 0 {
		err = DeleteTask(id)
		if err != nil {
			jsonError(err, http.StatusInternalServerError, w)
			return
		}
		emptyJson(w)
		return
	} else {
		nextDate, err := NextDate(time.Now(), task.Date, task.Repeat)
		if err != nil {
			jsonError(err, http.StatusInternalServerError, w)
			return
		}
		task.Date = nextDate
	}
	err = PutTask(task)
	if err != nil {
		jsonError(err, http.StatusInternalServerError, w)
		return
	}
	emptyJson(w)

}
