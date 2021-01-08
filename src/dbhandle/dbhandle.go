package dbhandle

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const (
	tableName = "tasks"
	driver    = "sqlite3"
)

func getPath() string {
	home := fmt.Sprintf("%s/todolist", os.Getenv("HOME"))
	return home
}

func errHandle(err error) bool {
	if err != nil {
		log.Println(err)
		return true
	}
	return false
}

func doCall(query string) int64 {
	db, err := sql.Open(driver, getPath())
	if errHandle(err) {
		return -1
	}
	defer db.Close()

	statement, err := db.Prepare(query)
	if errHandle(err) {
		return -1
	}
	result, err := statement.Exec()
	if errHandle(err) {
		return -1
	}

	newID, _ := result.LastInsertId()
	statement.Close()
	return newID
}

func CreateTable() {
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (description TEXT, done BOOLEAN)", tableName)
	doCall(query)
}

func CreateTask(description *string) int64 {
	query := fmt.Sprintf("INSERT INTO %s VALUES (\"%s\", %v)", tableName, *description, false)
	id := doCall(query)
	return id
}

func DeleteTask(id int64) {
	query := fmt.Sprintf("DELETE FROM %s WHERE rowid = %d", tableName, id)
	doCall(query)
}

func UpdateTask(id int64, description string, done bool) {
	query := fmt.Sprintf("UPDATE %s SET description = \"%s\", done = %v WHERE rowid = %d", tableName, description, done, id)
	doCall(query)
}

// func UpdateTaskDone(task *Task) {
// 	query := fmt.Sprintf("UPDATE %s SET done = %v WHERE rowid = %d", tableName, task.Done, task.Id)
// 	doCall(query)
// }

// func UpdateTaskDescription(task *Task) {
// 	query := fmt.Sprintf("UPDATE %s SET description = \"%s\" WHERE rowid = %d", tableName, task.Description, task.Id)
// 	doCall(query)
// }

func GetList() *sql.Rows {
	// tasks := make([]Task, 0)
	db, err := sql.Open(driver, getPath())
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	rows, err := db.Query(fmt.Sprintf("SELECT rowid, description, done FROM %s", tableName))
	if err != nil {
		log.Fatal(err)
	}
	// defer rows.Close()

	// var t Task

	return rows
}
