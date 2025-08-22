package db

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

var db *sql.DB
var Dbpath string = "scheduler.db"

const (
	schema = `
CREATE TABLE IF NOT EXISTS scheduler (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    date CHAR(8) NOT NULL DEFAULT "",
    title VARCHAR(255) NOT NULL,
    comment TEXT,
    repeat VARCHAR(128)
);
CREATE INDEX IF NOT EXISTS idx_date ON scheduler(date);
`
)

func Init(dbFile string) error {
	var err error
	db, err = sql.Open("sqlite", dbFile)

	if err != nil {
		return fmt.Errorf("ошибка открытия базы данных: %v", err)
	}

	_, err = db.Exec(schema)
	if err != nil {
		return fmt.Errorf("ошибка создания таблицы: %v", err)
	}

	return nil
}

func Close() {
	db.Close()
}
