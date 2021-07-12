package main

import (
	"net/http"
	"fmt"
)
// main function
func main()  {
	http.HandleFunc("/",HttpHandleFunc)
	http.ListenAndServe(":8081",nil)
}

func HttpHandleFunc(response http.ResponseWriter,request *http.Request){
	fmt.Println("server got!")
}