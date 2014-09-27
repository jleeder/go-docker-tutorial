package routes

import (
	//"fmt"
	"net/http"
	"html/template"
)

type Home struct{}

func (h *Home) Index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/home/index.html")
    data := struct {
    	title string
    }{
    	title : "Boardy",
    }
    t.Execute(w, data)		
}