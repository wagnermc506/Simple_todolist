package main

import (
	"frostwagner/dbhandle"
	"frostwagner/route"
	"net/http"
)

func main() {

	dbhandle.CreateTable()

	http.HandleFunc("/add", route.AddTask)
	http.HandleFunc("/", route.ReadList)
	http.HandleFunc("/delete", route.RemoveTask)
	// http.HandleFunc("/updateText", route.UpdateTaskDescription)
	// http.HandleFunc("/updateDone", route.UpdateTaskDone)
	http.HandleFunc("/update", route.UpdateTask)

	http.ListenAndServe(":8090", nil)
}
