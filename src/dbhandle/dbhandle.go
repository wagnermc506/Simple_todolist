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

func CreateTable() {
	db, err := sql.Open(driver, getPath())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	statement, _ := db.Prepare(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (description TEXT, done BOOLEAN)", tableName))
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()
	statement.Close()
}

func CreateTask(description *string) {
	db, err := sql.Open(driver, getPath())
	if err != nil {
		log.Print(err)
		return
	}
	defer db.Close()

	insert, err := db.Prepare(fmt.Sprintf("INSERT INTO %s VALUES (?, ?)", tableName))
	if err != nil {
		log.Println(err)
		return
	}
	insert.Exec(description, false)
	insert.Close()
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
	db, err := sql.Open(driver, getPath())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	statement, err := db.Prepare(fmt.Sprintf("DELETE FROM %s WHERE rowid = ?", tableName))
	if err != nil {
		fmt.Println(err)
		return
	}
	statement.Exec(id)
	statement.Close()
}

func UpdateTaskDone(task *Task) {
	db, err := sql.Open(driver, getPath())
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	statement, err := db.Prepare(fmt.Sprintf("UPDATE %s SET done = ? WHERE rowid = ?", tableName))
	statement.Exec(task.Done, task.Id)
	statement.Close()
}

func UpdateTaskDescription(task *Task) {
	db, err := sql.Open(driver, getPath())
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	statement, err := db.Prepare(fmt.Sprintf("UPDATE %s SET description = ? WHERE rowid = ?", tableName))
	if err != nil {
		log.Println(err)
		return
	}
	statement.Exec(task.Description, task.Id)
	statement.Close()
}
