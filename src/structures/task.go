package structures

import (
	"encoding/json"
	"log"

	"frostwagner/dbhandle"
)

const NOT_EXISTS = 0

type Task struct {
	ID           int64
	Description  string
	Done         bool
	ErrorMessage string
}

func (task *Task) EncodeToJson() []byte {
	jsonData, err := json.Marshal(task)
	if err != nil {
		log.Fatalln(err)
	}
	return jsonData
}

func (task *Task) DecodeFromJson(jsonData []byte) {
	err := json.Unmarshal(jsonData, &task)
	if err != nil {
		log.Fatalln(err)
	}
}

func (task *Task) Save() {
	if task.descriptionIsEmpty() {
		task.ErrorMessage = "Description must not be empty"
	} else if task.ID == NOT_EXISTS {
		newTaskId := dbhandle.CreateTask(&task.Description)
		task.ID = newTaskId
	} else {
		dbhandle.UpdateTask(task.ID, task.Description, task.Done)
	}
}

func (task *Task) descriptionIsEmpty() bool {
	if task.Description == "" {
		return true
	}
	return false
}

func (task *Task) Delete() {
	dbhandle.DeleteTask(task.ID)
}

// func (task Task) overwrite(newTask Task) {
// 	task = newTask
// }

// func (task Task) Fetch(ID int64) {
// 	task = dbhandle.GetTask(ID)
// }

func FetchAll() []Task {
	rows := dbhandle.GetList()
	defer rows.Close()
	tasks := make([]Task, 0)
	var t Task
	for rows.Next() {
		rows.Scan(&t.ID, &t.Description, &t.Done)
		tasks = append(tasks, t)
	}
	return tasks
}
