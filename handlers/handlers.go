package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"rpsweb/rps"
	"strconv"
)

type Player struct {
	Name string
}

var player Player

const (
	templateDir   = "templates/"
	temaplateBase = templateDir + "base.html"
)

func Index(w http.ResponseWriter, r *http.Request) {
	restarValue()
	RenderTemplate(w, "index.html", nil)
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	restarValue()
	RenderTemplate(w, "new-game.html", nil)
}

func Game(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		err := r.ParseForm()

		if err != nil {
			http.Error(w, "Error al procesar el formulario", http.StatusBadRequest)
			return
		}
		player.Name = r.Form.Get("name")

		if player.Name == "" {
			http.Redirect(w, r, "/NewGame", http.StatusFound)
			return
		}

		RenderTemplate(w, "game.html", player)
	}
}

func Play(w http.ResponseWriter, r *http.Request) {
	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("c"))

	result := rps.PlayRound(playerChoice)

	fmt.Println(result)

	out, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func About(w http.ResponseWriter, r *http.Request) {
	restarValue()
	RenderTemplate(w, "about.html", nil)
}

func RenderTemplate(w http.ResponseWriter, page string, data any) {
	tpl := template.Must(template.ParseFiles(temaplateBase, templateDir+page))

	err := tpl.ExecuteTemplate(w, "base", data)

	if err != nil {
		http.Error(w, "No se pudo renderizar la plantilla", http.StatusInternalServerError)
		return
	}
}

func restarValue() {
	player.Name = ""
	rps.ComputerScore = 0
	rps.PlayerScore = 0
}
