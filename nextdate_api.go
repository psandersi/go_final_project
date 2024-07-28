package main

import (
	"log"
	"net/http"
	"time"
)

func GetNextDate(w http.ResponseWriter, r *http.Request) {
	now := r.FormValue("now")
	nowTime, err := time.Parse("20060102", now)
	if err != nil {
		log.Panic(err)
	}
	date := r.FormValue("date")
	repeat := r.FormValue("repeat")
	d, err := NextDate(nowTime, date, repeat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ans := []byte(d)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(ans)
	if err != nil {
		log.Println(err)
	}
}
