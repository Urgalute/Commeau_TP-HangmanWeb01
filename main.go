package main

import (
	"fmt"
	"html/template"
	"net/http"
)

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
		Css string
	}

	dataPage := Etudiant{"Cyril RODRIGUES", "Kheir-eddine MEDERREG", "Alan PHILIPIERT",
		"22", "22", "26",
		"static/img/OIP.jpg", "static/css/main.css"}
	http.HandleFunc("/Promo", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "Promo", dataPage)
	})

	type CheckView struct {
		Check bool
		Nbr   int
		Css string
	}

	even := false
	view := 0
	http.HandleFunc("/Change", func(w http.ResponseWriter, r *http.Request) {
		view++
		if view%2 == 0 {
			even = true
		} else {
			even = false
		}
		dataPage := CheckView{even, view, "static/css/main.css"}
		temp.ExecuteTemplate(w, "Change", dataPage)
	})

	fileServer := http.FileServer(http.Dir("asset"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	http.ListenAndServe("localhost:8080", nil)
}
