package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Film struct {
  Title string
  Director string
}

func main(){
  var h1 http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("index.html")) 
    films := map[string][]Film{
      "Films": {
        {Title: "Black Panther", Director: "Ryan Coogler"},
        {Title: "Fight Club", Director: "David Fincher"},
        {Title: "13 Bom Di Jakarta", Director: "Angga Dimas Sasongko"},
      },
    }

    if err := tmpl.Execute(w,films); err != nil {
      http.Error(w,err.Error(),http.StatusInternalServerError)
    }
  }
  var h2 http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
    title := r.PostFormValue("title")
    director := r.PostFormValue("director")
    htmlStrx := fmt.Sprintf("<p class='px-3 py-2'>%s - %s</p>", title,director)
    tmepx,_ := template.New("res").Parse(htmlStrx)
    tmepx.Execute(w,nil)
  } 

  http.HandleFunc("/",h1)
  http.HandleFunc("/add-film",h2)

  log.Fatal(http.ListenAndServe(":8080",nil))
}
