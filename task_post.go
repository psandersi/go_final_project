package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title" binding:"required"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

// Функция jsonError записывает и возвращает ошибку в виде json
func jsonError(err error, st int, w http.ResponseWriter) {
	log.Println(err)
	resErr := map[string]string{
		"error": err.Error(),
	}
	res, err := json.Marshal(resErr)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(st)
	_, err = w.Write(res)
	if err != nil {
		log.Println(err)
	}
}

// Функция jsonAns записвает ответ обработчика TaskPost в виде JSON
func jsonAns(id int64, st int, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		jsonError(err, st, w)
	}
	idres := map[string]string{"id": strconv.Itoa(int(id))}
	res, err := json.Marshal(idres)
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

// Функция checkData проверяет указанные даты на соответсвие формату
func (task Task) checkData() (Task, error) {
	var err error
	if task.Title == "" {
		err = errors.New("Title cannot be empty")
		return Task{}, err
	}
	if len(task.Date) == 0 || strings.ToLower(task.Date) == "today" {
		task.Date = time.Now().Format("20060102")
	} else {
		_, err := time.Parse("20060102", task.Date)
		if err != nil {
			log.Println(err)
			return Task{}, err
		}
	}
	if task.Date < time.Now().Format("20060102") {
		if len(task.Repeat) == 0 {
			task.Date = time.Now().Format("20060102")
		} else {
			task.Date, err = NextDate(time.Now(), task.Date, task.Repeat)
			if err != nil {
				log.Println(err)
				return Task{}, err
			}
		}
	}
	return task, nil
}

// Функция TaskPost обрабатывает POST-запросы к /api/task и добавляет задачу в базу данных
func TaskPost(w http.ResponseWriter, r *http.Request) {
	var task Task
	var buf bytes.Buffer
	var err error
	var id int64
	_, err = buf.ReadFrom(r.Body)
	if err != nil {
		jsonAns(0, http.StatusBadRequest, err, w)
		return
	}

	if err = json.Unmarshal(buf.Bytes(), &task); err != nil {
		jsonAns(0, http.StatusBadRequest, err, w)
		return
	}
	task, err = task.checkData()
	if err != nil {
		jsonAns(0, http.StatusBadRequest, err, w)
		return
	}

	id, err = AddTask(task)
	jsonAns(id, http.StatusBadRequest, err, w)
}
