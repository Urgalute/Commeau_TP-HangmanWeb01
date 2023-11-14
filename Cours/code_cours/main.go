package main

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
)

func main() {

	http.HandleFunc("/sayHello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("HELLO WORLD !"))
	})

	temp, err := template.ParseGlob("./templates/*.html")
	if err != nil {
		fmt.Println(fmt.Sprint("ERREUR => %s", err.Error()))
		return
	}

	http.HandleFunc("/template", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "index", nil)
	})

	type PageVariable struct {
		Titre   string
		Intro   string
		Article string
	}

	http.HandleFunc("/var", func(w http.ResponseWriter, r *http.Request) {
		dataPage := PageVariable{"Les bases du WEB",
			"Le web est devenu un incontournable pour....",
			"Pour commencer cet article sur les bases du WEB nous allons aborder la culture du WEB"}
		temp.ExecuteTemplate(w, "var", dataPage)
	})

	type PageCondition struct {
		Check      bool
		CheckOwner bool
		NbrGuest   int
	}

	http.HandleFunc("/cond", func(w http.ResponseWriter, r *http.Request) {
		dataPage := PageCondition{true, false, 12}
		temp.ExecuteTemplate(w, "cond", dataPage)
	})

	type User struct {
		FirstName string
		LastName  string
	}

	type PageLocal struct {
		Titre string
		Users []User
	}

	http.HandleFunc("/local", func(w http.ResponseWriter, r *http.Request) {
		data := PageLocal{"Liste des mentors",
			[]User{{"Cyril", "RODRIGUES"},
				{"Kheir-eddine", "MEDERREG"},
				{"Alan", "PHILIPIERT"}}}
		temp.ExecuteTemplate(w, "local", data)
	})

	type PagePrim struct {
		Titre  string
		ListId []int
	}

	http.HandleFunc("/prim", func(w http.ResponseWriter, r *http.Request) {
		data := PagePrim{"Liste des références des commandes", []int{10, 87, 65, 12, 47, 89, 5874, 247}}
		temp.ExecuteTemplate(w, "prim", data)
	})

	type DataFile struct {
		ImgPath string
		LinkCss string
	}

	http.HandleFunc("/file", func(w http.ResponseWriter, r *http.Request) {
		dataTemp := DataFile{"/static/img/img01.png", "/static/css/main.css"}
		temp.ExecuteTemplate(w, "useFiles", dataTemp)
	})

	http.HandleFunc("/form/get", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "formGet", nil)
	})
	http.HandleFunc("/form/post", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "formPost", nil)
	})

	type DataFormTemp struct {
		CheckData bool
		Data      string
	}

	http.HandleFunc("/form/traitement", func(w http.ResponseWriter, r *http.Request) {
		var data DataFormTemp
		if r.Method == http.MethodGet {
			data = DataFormTemp{
				CheckData: false,
				Data:      r.URL.Query().Get("nom"),
			}
		}

		if r.Method == http.MethodPost {
			data = DataFormTemp{
				CheckData: false,
				Data:      r.FormValue("nom"),
			}
		}

		checkValue, _ := regexp.MatchString("^[a-zA-Z-]{1,64}$", data.Data)
		if !checkValue {
			data.CheckData = true
		}

		temp.ExecuteTemplate(w, "displayForm", data)
	})

	fileServer := http.FileServer(http.Dir("asset"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	http.ListenAndServe("localhost:8080", nil)
}
