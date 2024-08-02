package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

//Функция CreateTable создает новую базу данных c полями:
/*id — автоинкрементный идентификатор;
date — дата задачи;
title — заголовок задачи;
comment — комментарий к задаче;
repeat — строковое поле; ID_Date - индекс по дате */
var install bool

func CreateTable(dbPath string) error {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		if os.IsNotExist(err) {
			install = true
		} else {
			log.Panic(err)
		}
	}
	defer db.Close()
	_, err = db.Exec("CREATE TABLE scheduler (id  INTEGER PRIMARY KEY AUTOINCREMENT, date VARCHAR, title VARCHAR(128) NOT NULL, comment VARCHAR, repeat VARCHAR(128) )")
	if err != nil {
		return err
	}
	_, err = db.Exec("CREATE INDEX ID_Date ON scheduler (date)")
	if err != nil {
		return err
	}
	return nil
}

/*
Функция DoNewTable работет с базой данных: проверяет существует ли нужный файл, если нет, то создает новый.
На вход принимает строку с нужным расположением файла, на выходе выводит структуру Storage
*/
func DoNewTable(dbPath string) (Storage, error) {
	//проверяем существует ли файл
	_, err := os.Stat(dbPath)
	if err != nil {
		install = true
	}
	//если файла нет, то создаем новый
	if install {
		_, err = os.Create(dbPath)
		if err != nil {
			log.Fatal(err)
		}
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return Storage{}, err
	}
	dbStorage := Storage{}
	dbStorage.db = db
	if install {
		err = CreateTable(dbPath)
		if err != nil {
			log.Fatal(err)
		}
	}

	return dbStorage, nil
}

func (s *Storage) dbClose() {
	s.db.Close()
}
