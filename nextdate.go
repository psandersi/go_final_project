package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Функция NextDate рассчитвает дату следующего выполнения задачи с учетом правила повторения и возращает ее в формате 20060102
func NextDate(now time.Time, date string, repeat string) (string, error) {
	if repeat == "" {
		return "", fmt.Errorf("repeat cannot be empty: %w", errors.New("not found"))
	}
	dateTask, err := time.Parse("20060102", date)
	if err != nil {
		return "", fmt.Errorf("failed to convert string to date: %w", errors.New("not found"))
	}
	timeRepeat := strings.Split(repeat, " ")

	switch strings.ToLower(timeRepeat[0]) {
	case "y":
		dateTask = addDateTask(now, dateTask, 1, 0, 0)
	case "d":
		if len(timeRepeat) == 1 {
			return "", fmt.Errorf("d cannot be empty: %w", errors.New("not found"))
		}

		digit, err := convert(timeRepeat[1])
		if err != nil {
			return "", err
		}
		dateTask = addDateTask(now, dateTask, 0, 0, digit)
	default:
		return "", fmt.Errorf("unknown key in repeat: %w", errors.New("not found"))

	}
	return dateTask.Format("20060102"), nil
}

// Функция convert переводит строку в тип int, а также проверяет находится ли это число в нужном диапазоне
func convert(n string) (int, error) {
	d, err := strconv.Atoi(n)
	if err != nil {
		return 0, err
	}
	if d >= 400 || d < 0 {
		return 0, fmt.Errorf("the number of days is greater than 400 or a negative number is specified: %w", errors.New("not found"))
	}
	return d, nil
}

// Функция addDateTask добавляет ко времени указанное количество лет, месяцев и дней
func addDateTask(now time.Time, dateTask time.Time, y int, m int, d int) time.Time {
	dateTask = dateTask.AddDate(y, m, d)

	for dateTask.Before(now) {

		dateTask = dateTask.AddDate(y, m, d)
	}
	return dateTask
}
