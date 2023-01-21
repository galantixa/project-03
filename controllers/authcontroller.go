package controllers

import (
	"net/http"
	"html/template"
)

func Index (w http.ResponseWriter, r *http.Request) {
	temp, _ := template.ParseFiles("views/index.js")
	temp.Execute(w, nil)
}

func Login (w http.ResponseWriter, r *http.Response) {
	temp, _ := template.ParseFiles("views/login.js")
	temp.Execute(w, nil)
}