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
type Films struct {
  Filems []Film
}

var Filmx Films

func (f *Films) addfilm(tit,dir string) *Films {
  f.Filems = append(f.Filems, []Film{
    {Title: tit,Director: dir},
  }...)
  return f
}

func main(){
  var h1 http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("index.html")) 
    if err:= tmpl.Execute(w,Filmx); err != nil {
      panic(err)
    }
  }

  var h2 http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
    title := r.PostFormValue("title")
    director := r.PostFormValue("director")
    Filmx.addfilm(title,director)

    htmlStrx := fmt.Sprintf("<p class='px-3 py-2'>%s - %s</p>", title,director)
    tmepx,_ := template.New("res").Parse(htmlStrx)
    tmepx.Execute(w,nil)
  } 

  http.HandleFunc("/",h1)
  http.HandleFunc("/add-film",h2)

  log.Fatal(http.ListenAndServe(":8080",nil))
}
