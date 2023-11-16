package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type detail struct {
	Nom    string
	Prenom string
	Birth  string
	Sexe   string
}

var myuser detail

func main() {

	temp, err := template.ParseGlob("./Templates/*.html")
	if err != nil {
		fmt.Println(fmt.Sprint("ERREUR => %s", err.Error()))
		return
	}
	type Etudiant struct {
		A      string
		B      string
		C      string
		Age1   string
		Age2   string
		Age3   string
		ImgSrc string
	}

	http.HandleFunc("/Promo", func(w http.ResponseWriter, r *http.Request) {
		dataPage := Etudiant{"Cyril RODRIGUES", "Kheir-eddine MEDERREG", "Alan PHILIPIERT",
			"22", "22", "26",
			"static/img/OIP.jpg"}
		temp.ExecuteTemplate(w, "Promo", dataPage)
	})

	type CheckView struct {
		Check bool
		Nbr   int
	}

	var even bool
	var view int
	http.HandleFunc("/Change", func(w http.ResponseWriter, r *http.Request) {
		view++
		if view%2 == 0 {
			even = true
		} else {
			even = false
		}
		dataPage := CheckView{even, view}
		temp.ExecuteTemplate(w, "Change", dataPage)
	})

	http.HandleFunc("/Init", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "Init", nil)
	})

	http.HandleFunc("/Display", func(w http.ResponseWriter, r *http.Request) {
		user := detail{
			Nom:    r.FormValue("nom"),
			Prenom: r.FormValue("prenom"),
			Birth:  r.FormValue("birthday"),
			Sexe:   r.FormValue("sexe"),
		}
		temp.ExecuteTemplate(w, "Display", user)
	})

	fileServer := http.FileServer(http.Dir("asset"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	http.ListenAndServe("localhost:8080", nil)
}
