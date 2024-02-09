package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type Film struct {
	FilmID   uuid.UUID
	Title    string
	Director string
}

type Films struct {
	Filems []Film
}

var Filmx Films = Films{
	Filems: []Film{
		{FilmID: uuid.New(), Title: "13 Bom Di jakarta", Director: "Angga Dimas Sasongko"},
	},
}

func (f *Films) addfilm(tit, dir string) uuid.UUID {
	newFilm := Film{FilmID: uuid.New(), Title: tit, Director: dir}
	f.Filems = append(f.Filems, newFilm)
	return newFilm.FilmID
}

func (f *Films) DeleteFilm(FilmID uuid.UUID) {
	for i, v := range f.Filems {
		if v.FilmID == FilmID {
			f.Filems = append(f.Filems[:i], f.Filems[i+1:]...)
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
		res := Films{
			Filems: []Film{
				{Title: title, Director: director, FilmID: Flz},
			},
		}

		if err := childe.ExecuteTemplate(w, "A", res); err != nil {
			log.Panic(err)
		}
	}

	var h3 http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		aidi := r.URL.Query().Get("id")
		idd, _ := uuid.Parse(aidi)
		Filmx.DeleteFilm(idd)

		t, _ := template.New("foo").ParseFiles("comps/child.html")
		if err := t.ExecuteTemplate(w, "B", Filmx); err != nil {
			panic(err)
		}
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film", h2)
	http.HandleFunc("/del", h3)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
