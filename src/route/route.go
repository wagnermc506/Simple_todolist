package route

import (
	"encoding/json"
	"fmt"
	"frostwagner/dbhandle"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// type Task structures.Task

func AddTask(w http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)
	bodystr := string(body)
	if bodystr != "" {
		dbhandle.CreateTask(&bodystr)
	}
	fmt.Fprintf(w, "alright")
}

func ReadList(w http.ResponseWriter, req *http.Request) {
	//returns a json parsed with the id, description, and if it is done
	tasks := *dbhandle.GetList()
	jsonData, err := json.Marshal(tasks)
	if err != nil {
		fmt.Fprintf(w, "null")
		return
	}
	fmt.Fprintf(w, string(jsonData))
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

	dbhandle.DeleteTask(num)
}

func UpdateTaskDone(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
		return
	}
	var task dbhandle.Task
	err = json.Unmarshal(body, &task)
	if err != nil {
		log.Println(err)
		return
	}
	dbhandle.UpdateTaskDone(&task)
}

func UpdateTaskDescription(w http.ResponseWriter, req *http.Request) {
	//recebe um json com os campos id e description
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return
	}
	var task dbhandle.Task
	err = json.Unmarshal(body, &task)
	if err != nil {
		fmt.Println(err)
		return
	}
	dbhandle.UpdateTaskDescription(&task)
}
