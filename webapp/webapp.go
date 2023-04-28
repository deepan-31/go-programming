package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

var pl = fmt.Println

type ToDolist struct {
	ToDoCount int
	ToDos     []string
}

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func write(writer http.ResponseWriter, msg string) {
	_, err := writer.Write([]byte(msg))
	errorCheck(err)
}
func getStrings(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	errorCheck(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	errorCheck(scanner.Err())
	return lines
}

func englishHandler(writer http.ResponseWriter, req *http.Request) {
	write(writer, "Hello Internet")
}
func spanishHandler(writer http.ResponseWriter, req *http.Request) {
	write(writer, "Hola Internet")
}
func interactHanler(writer http.ResponseWriter, req *http.Request) {
	todoVals := getStrings("todos.txt")
	fmt.Printf("%#v\n", todoVals)
	tmp1, err := template.ParseFiles("view.html")
	errorCheck(err)
	todos := ToDolist{
		ToDoCount: len(todoVals),
		ToDos:     todoVals,
	}
	err = tmp1.Execute(writer, todos)
}

func newhandler(writer http.ResponseWriter, req *http.Request) {
	tmp1, err := template.ParseFiles("new.html")
	errorCheck(err)
	err = tmp1.Execute(writer, nil)
}

func createHandler(writer http.ResponseWriter, req *http.Request) {
	todo := req.FormValue("todo")
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("todos.txt", options, os.FileMode(0600))
	errorCheck(err)
	_, err = fmt.Fprintln(file, todo)
	err = file.Close()
	errorCheck(err)
	http.Redirect(writer, req, "/interact", http.StatusFound)
}

func main() {
	http.HandleFunc("/hello", englishHandler)
	http.HandleFunc("/hola", spanishHandler)
	http.HandleFunc("/interact", interactHanler)
	http.HandleFunc("/new", newhandler)
	http.HandleFunc("/create", createHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	errorCheck(err)
}
