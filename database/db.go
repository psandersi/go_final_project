package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func DoNewTable(dbPath string) (Storage, error) {
	//проверяем существует ли файл
	_, err := os.Stat(dbPath)
	var install bool
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

	err = CreateTable(dbPath)
	if err != nil {
		log.Println(err)
	}

	return dbStorage, nil
}

func CreateTable(dbPath string) error {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec("CREATE TABLE scheduler (id  INTEGER PRIMARY KEY AUTOINCREMENT, date VARCHAR, title VARCHAR(128) NOT NULL, comment VARCHAR, repeat VARCHAR(128) )")
	if err != nil {
		log.Panic(err)
		return err
	}
	_, err = db.Exec("CREATE INDEX ID_Date ON scheduler (date)")
	if err != nil {
		log.Panic(err)
		return err
	}
	return nil
}

func (s *Storage) dbClose() {
	s.db.Close()
}
