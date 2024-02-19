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

var Todoz TodoList = TodoList{
	TodoList: []Todo{
		{TodoID: uuid.New(), Title: "Wash Da Dishes", Todo: "Wash Da fucking Dishes"},
	},
}

func (f *TodoList) addTodo(tit, dir string) uuid.UUID {
	newTodo := Todo{TodoID: uuid.New(), Title: tit, Todo: dir}
	f.TodoList = append(f.TodoList, newTodo)
	return newTodo.TodoID
}

func (f *TodoList) DeleteTodo(FilmID uuid.UUID) {
	for i, v := range f.TodoList {
		if v.TodoID == FilmID {
			f.TodoList = append(f.TodoList[:i], f.TodoList[i+1:]...)
		}
	}
}

func (f *TodoList) EditTodo(eTodoID uuid.UUID, tit, tod string) {
	EditedTodo := Todo{TodoID: eTodoID, Title: tit, Todo: tod}
	for a, v := range f.TodoList {
		if v.TodoID == eTodoID {
			f.TodoList[a] = EditedTodo
		}
	}
}

func main() {
	var h1 http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		if err := tmpl.Execute(w, Todoz); err != nil {
			panic(err)
		}
	}

	var h2 http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		title := r.PostFormValue("title")
		todo := r.PostFormValue("todo")

		Flz := Todoz.addTodo(title, todo)
		childe, _ := template.ParseFiles("comps/child.html")
		res := TodoList{
			TodoList: []Todo{
				{Title: title, Todo: todo, TodoID: Flz},
			},
		}

		if err := childe.ExecuteTemplate(w, "A", res); err != nil {
			log.Panic(err)
		} else {
			log.Printf("Todo item created.Values : \n{ Title: %v, Todo: %v, TodoID: %v }\n\n", title, todo, Flz)
		}
	}

	var h3 http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		aidi := r.URL.Query().Get("id")
		idd, _ := uuid.Parse(aidi)
		Todoz.DeleteTodo(idd)

		t, _ := template.New("grahh").ParseFiles("comps/child.html")
		if err := t.ExecuteTemplate(w, "B", Todoz); err != nil {
			log.Panic(err)
		} else {
			log.Printf("Todo item '%v' was Deleted. \n\n", aidi)
		}
	}

	var h4 http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		eid := r.URL.Query().Get("eid")
		idd, _ := uuid.Parse(eid)
		tit := r.URL.Query().Get("title")
		tod := r.URL.Query().Get("todo")

		Todoz.EditTodo(idd, tit, tod)

		t, _ := template.New("grahh").ParseFiles("comps/child.html")
		if err := t.ExecuteTemplate(w, "B", Todoz); err != nil {
			log.Panic(err)
		} else {
			log.Printf("Todo item '%v' was Edited. \n\n", eid)
		}
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film", h2)
	http.HandleFunc("/del", h3)
	http.HandleFunc("/edt", h4)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
