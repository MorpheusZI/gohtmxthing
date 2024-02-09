package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type Todo struct {
	TodoID uuid.UUID
	Title  string
	Todo   string
}

type TodoList struct {
	TodoList []Todo
}

var Filmx TodoList = TodoList{
	TodoList: []Todo{
		{TodoID: uuid.New(), Title: "Wash Da Dishes", Todo: "Wash Da fucking Dishes"},
	},
}

func (f *TodoList) addfilm(tit, dir string) uuid.UUID {
	newFilm := Todo{TodoID: uuid.New(), Title: tit, Todo: dir}
	f.TodoList = append(f.TodoList, newFilm)
	return newFilm.TodoID
}

func (f *TodoList) DeleteFilm(FilmID uuid.UUID) {
	for i, v := range f.TodoList {
		if v.TodoID == FilmID {
			f.TodoList = append(f.TodoList[:i], f.TodoList[i+1:]...)
		}
	}
}

func main() {
	var h1 http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		if err := tmpl.Execute(w, Filmx); err != nil {
			panic(err)
		}
	}

	var h2 http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")

		Flz := Filmx.addfilm(title, director)
		childe, _ := template.ParseFiles("comps/child.html")
		res := TodoList{
			TodoList: []Todo{
				{Title: title, Todo: director, TodoID: Flz},
			},
		}

		if err := childe.ExecuteTemplate(w, "A", res); err != nil {
			log.Panic(err)
		} else {
			log.Printf("Todo item created. Values : \n{ Title: %v, Todo: %v, TodoID: %v }\n\n", title, director, Flz)
		}
	}

	var h3 http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		aidi := r.URL.Query().Get("id")
		idd, _ := uuid.Parse(aidi)
		Filmx.DeleteFilm(idd)

		t, _ := template.New("foo").ParseFiles("comps/child.html")
		if err := t.ExecuteTemplate(w, "B", Filmx); err != nil {
			log.Panic(err)
		} else {
			log.Printf("Todo item '%v' was Deleted. \n\n", aidi)
		}
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film", h2)
	http.HandleFunc("/del", h3)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
