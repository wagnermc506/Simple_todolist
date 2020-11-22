package dbhandle

import (
	"database/sql"
	"fmt"
	"frostwagner/structures"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Task structures.Task

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

func doCall(query string) {
	db, err := sql.Open(driver, getPath())
	if errHandle(err) {
		return
	}
	defer db.Close()

	statement, err := db.Prepare(query)
	if errHandle(err) {
		return
	}
	statement.Exec()
	statement.Close()
}

func CreateTable() {
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (description TEXT, done BOOLEAN)", tableName)
	doCall(query)
}

func CreateTask(description *string) {
	query := fmt.Sprintf("INSERT INTO %s VALUES (\"%s\", %v)", tableName, *description, false)
	doCall(query)
}

func GetList() *[]Task {
	tasks := make([]Task, 0)
	db, err := sql.Open(driver, getPath())
	if err != nil {
		log.Println(err)
		return &tasks
	}
	defer db.Close()
	rows, err := db.Query(fmt.Sprintf("SELECT rowid, description, done FROM %s", tableName))
	if err != nil {
		log.Println(err)
		return &tasks
	}
	defer rows.Close()

	var t Task
	for rows.Next() {
		rows.Scan(&t.Id, &t.Description, &t.Done)
		tasks = append(tasks, t)
	}
	return &tasks
}

func DeleteTask(id int) {
	query := fmt.Sprintf("DELETE FROM %s WHERE rowid = %d", tableName, id)
	doCall(query)
}

func UpdateTaskDone(task *Task) {
	query := fmt.Sprintf("UPDATE %s SET done = %v WHERE rowid = %d", tableName, task.Done, task.Id)
	doCall(query)
}

func UpdateTaskDescription(task *Task) {
	query := fmt.Sprintf("UPDATE %s SET description = \"%s\" WHERE rowid = %d", tableName, task.Description, task.Id)
	doCall(query)
}
