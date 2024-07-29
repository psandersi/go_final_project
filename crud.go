package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

// Добавляет задачу в таблицу и возвращает id добавленной задачи
func AddTask(task Task) (int64, error) {
	db, _ := sqlx.Connect("sqlite3", dbFile)
	res, err := db.Exec("INSERT INTO scheduler (date, title, comment, repeat) VALUES (:date, :title, :comment, :repeat)",
		sql.Named("date", task.Date), sql.Named("title", task.Title),
		sql.Named("comment", task.Comment), sql.Named("repeat", task.Repeat))
	if err != nil {
		return 0, fmt.Errorf("failed to INSERT a request for database update: %w", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert id: %w", err)
	}
	return id, nil
}

// Функция GetTasksList возвращает список ближайщих задач. Количество задач регулируется в переменной limit
func GetTasksList() ([]Task, error) {
	limit := 10
	db, _ := sqlx.Connect("sqlite3", dbFile)
	var tasks []Task
	var rows *sql.Rows
	var err error
	rows, err = db.Query("SELECT * FROM scheduler ORDER BY id LIMIT :limit", sql.Named("limit", limit))
	if err != nil {
		return []Task{}, err
	}

	defer rows.Close()

	for rows.Next() {
		task := Task{}

		err := rows.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			log.Println(err)
			return []Task{}, err
		}
		tasks = append(tasks, task)

	}
	return tasks, nil
}

// Функция GetTaskByID возвращает задачу по указанному id
func GetTaskByID(id string) (Task, error) {
	var task Task
	db, _ := sqlx.Connect("sqlite3", dbFile)
	row := db.QueryRow("SELECT * FROM scheduler WHERE id = :id", sql.Named("id", id))

	err := row.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		log.Println(err)
		return Task{}, err
	}
	return task, nil

}

// Функция PutTask редактирует задачу
func PutTask(task Task) error {
	db, _ := sqlx.Connect("sqlite3", dbFile)
	res, err := db.Exec("UPDATE scheduler SET date = :date, title = :title, comment = :comment, repeat = :repeat WHERE id = :id",
		sql.Named("date", task.Date),
		sql.Named("title", task.Title),
		sql.Named("comment", task.Comment),
		sql.Named("repeat", task.Repeat),
		sql.Named("id", task.ID))
	if err != nil {
		return err
	}
	if rowsAffected, _ := res.RowsAffected(); rowsAffected != 1 {
		return fmt.Errorf("error while update")
	}
	return nil
}

// Функция DeleteTask удаляет задачу
func DeleteTask(id string) error {
	db, _ := sqlx.Connect("sqlite3", dbFile)
	_, err := GetTaskByID(id)
	if err != nil {
		return err
	}

	res, err := db.Exec("DELETE FROM scheduler WHERE id= :id", sql.Named("id", id))
	if err != nil {
		return err
	}
	smth, _ := res.RowsAffected()
	if smth != 1 {
		return fmt.Errorf("crash while deleting")
	}
	return nil
}
