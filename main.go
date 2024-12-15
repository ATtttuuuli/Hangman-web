package main

import (
	HangmanWeb "HangmanWeb/jeu" 
	"fmt"
	"html/template"
	"net/http"
)

var currentGame HangmanWeb.HangManData

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl := template.Must(template.ParseFiles("templetes/index.html"))
		err := tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func hangman(w http.ResponseWriter, r *http.Request) {
	reset := r.URL.Query().Get("reset")
	if reset == "true" {
		currentGame = HangmanWeb.HangManData{
			Level:       r.URL.Query().Get("level"),
			Attempts:    10,
			UsedLetters: []string{},
			ToFind:      "",
		}
	}

	level := r.URL.Query().Get("level")
	if level != "" && currentGame.ToFind == "" {
		currentGame.Level = level
	}

	if currentGame.ToFind == "" {
		var fichier string
		switch currentGame.Level {
		case "easy":
			fichier = "jeu/words1.txt"
		case "medium":
			fichier = "jeu/words2.txt"
		case "hard":
			fichier = "jeu/words3.txt"
		}

		if err := currentGame.ChoisirMotAleatoire(fichier); err != nil {
			http.Error(w, "Erreur lors du choix du mot: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if r.Method == http.MethodPost {
		r.ParseForm()
		lettre := r.Form.Get("letter")
		if lettre != "" && len(lettre) == 1 && lettre[0] >= 'a' && lettre[0] <= 'z' {

			currentGame.DevinerLettre(lettre)
		} else {
			// Lettre invalide ou vide
			fmt.Println("Lettre invalide") // Utilisé pour déboguer
		}
	}

	if currentGame.IsWinner() {

		currentGame.Attempts = 0
	}

	var imagePath string
	if currentGame.Attempts == 10 {
		imagePath = "/images/hangman10.png"
	} else {
		// Sinon, utiliser les images hangman0.png, hangman1.png, etc.
		imagePath = fmt.Sprintf("/images/hangman%d.png", currentGame.Attempts)
	}

	// Préparer les données à passer au template
	tmpl := template.Must(template.ParseFiles("templetes/hangman.html"))

	data := struct {
		Word        string
		UsedLetters []string
		Attempts    int
		Level       string
		IsGameOver  bool
		IsWinner    bool
		ImagePath   string
	}{
		Word:        currentGame.GetMaskedWordWithSpaces(),
		UsedLetters: currentGame.UsedLetters,
		Attempts:    currentGame.Attempts,
		Level:       currentGame.Level,
		IsWinner:    currentGame.IsWinner(),
		IsGameOver:  currentGame.APerdu(),
		ImagePath:   imagePath,
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("templetes/css"))))
	http.Handle("/images/", http.StripPrefix("/images", http.FileServer(http.Dir("templetes/images"))))

	// Définir les routes
	http.HandleFunc("/", index)          // Route d'accueil
	http.HandleFunc("/hangman", hangman) // Route du jeu
	// j'ai fais 8081 pca j'avais un probleme AVEC LE 8080
	err := http.ListenAndServe(":8081", nil)
	fmt.Println("le serveur est demarer sur le port 8081")
	if err != nil {
		fmt.Println("Erreur de démarrage du serveur : ", err)
	}
}
