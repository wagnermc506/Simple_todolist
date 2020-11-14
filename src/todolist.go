package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type Task struct {
	Id          int
	Description string
	Done        bool
}

const (
	tableName = "tasks"
	driver    = "sqlite3"
	file      = "./db/todolist.db"
)

func CreateTable() {
	db, err := sql.Open(driver, file)
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
	db, err := sql.Open(driver, file)
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

func addTask(w http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)
	bodystr := string(body)
	if bodystr != "" {
		CreateTask(&bodystr)
	}
	fmt.Fprintf(w, "alright")
}

func GetList() *[]Task {
	tasks := make([]Task, 0)
	db, err := sql.Open(driver, file)
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

func ReadList(w http.ResponseWriter, req *http.Request) {
	//returns a json parsed with the id, description, and if it is done
	tasks := *GetList()
	jsonData, err := json.Marshal(tasks)
	if err != nil {
		fmt.Fprintf(w, "null")
		return
	}
	fmt.Fprintf(w, string(jsonData))
}

func DeleteTask(id int) {
	db, err := sql.Open(driver, file)
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

func RemoveTask(w http.ResponseWriter, req *http.Request) {
	// recebe o id da task
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(w, "something went wrong")
		return
	}

	num, err := strconv.Atoi(string(body))
	if err != nil {
		fmt.Fprintf(w, "NotANumber")
		return
	}

	DeleteTask(num)
}

func DBUpdateTaskDone(task *Task) {
	db, err := sql.Open(driver, file)
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	statement, err := db.Prepare(fmt.Sprintf("UPDATE %s SET done = ? WHERE rowid = ?", tableName))
	statement.Exec(task.Done, task.Id)
	statement.Close()
}

func UpdateTaskDone(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
		return
	}
	var task Task
	err = json.Unmarshal(body, &task)
	if err != nil {
		log.Println(err)
		return
	}
	DBUpdateTaskDone(&task)
}

func DBUpdateTaskDescription(task *Task) {
	db, err := sql.Open(driver, file)
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

func UpdateTaskDescription(w http.ResponseWriter, req *http.Request) {
	//recebe um json com os campos id e description
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return
	}
	var task Task
	err = json.Unmarshal(body, &task)
	if err != nil {
		fmt.Println(err)
		return
	}
	DBUpdateTaskDescription(&task)
}

func main() {

	CreateTable()

	http.HandleFunc("/add", addTask)
	http.HandleFunc("/", ReadList)
	http.HandleFunc("/delete", RemoveTask)
	http.HandleFunc("/updateText", UpdateTaskDescription)
	http.HandleFunc("/updateDone", UpdateTaskDone)

	http.ListenAndServe(":8090", nil)
}
