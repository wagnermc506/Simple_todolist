package route

import (
	"encoding/json"
	"fmt"
	"frostwagner/dbhandle"
	"io/ioutil"
	"log"
	"net/http"
)

func errHandle(err error) bool {
	if err != nil {
		log.Println(err)
		return true
	}
	return false
}

func AddTask(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if errHandle(err) {
		return
	}

	task, err := decodeJson(&body)
	if errHandle(err) {
		return
	}
	bodystr := &task.Description

	if *bodystr != "" {
		dbhandle.CreateTask(bodystr)
	}
	fmt.Fprintf(w, "alright")
}

func ReadList(w http.ResponseWriter, req *http.Request) {
	//returns a json parsed with the id, description, and if it is done
	tasks := *dbhandle.GetList()
	jsonData, err := json.Marshal(tasks)
	if errHandle(err) {
		return
	}
	// fmt.Fprintf(w, string(jsonData))
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func RemoveTask(w http.ResponseWriter, req *http.Request) {
	// recebe o id da task
	body, err := ioutil.ReadAll(req.Body)
	if errHandle(err) {
		return
	}

	task, err := decodeJson(&body)
	if errHandle(err) {
		return
	}

	dbhandle.DeleteTask(task.Id)
}

func UpdateTaskDone(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if errHandle(err) {
		return
	}

	task, err := decodeJson(&body)
	if errHandle(err) {
		return
	}

	dbhandle.UpdateTaskDone(task)
}

func UpdateTaskDescription(w http.ResponseWriter, req *http.Request) {
	//recebe um json com os campos id e description
	body, err := ioutil.ReadAll(req.Body)
	if errHandle(err) {
		return
	}

	task, err := decodeJson(&body)
	if errHandle(err) {
		return
	}

	dbhandle.UpdateTaskDescription(task)
}

func decodeJson(body *[]byte) (*dbhandle.Task, error) {
	var task dbhandle.Task
	err := json.Unmarshal(*body, &task)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &task, nil
}
