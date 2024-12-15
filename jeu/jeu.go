package HangmanWeb

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type HangManData struct {
	Level           string   // Niveau de difficulté (Facile, Moyen, Difficile)
	ToFind          string   // Mot à deviner
	Word            string   // Mot masqué (avec des underscores)
	RevealedLetters []bool   // Lettres révélées (true si la lettre à cet index a été trouvée)
	Attempts        int      // Nombre de tentatives restantes
	UsedLetters     []string // Liste des lettres déjà utilisées
}

func (h *HangManData) ChoisirMotAleatoire(fichier string) error {
	file, err := os.Open(fichier)
	if err != nil {
		return fmt.Errorf("erreur d'ouverture du fichier %s: %v", fichier, err)
	}
	defer file.Close()

	var mots []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mots = append(mots, scanner.Text())
	}

	if len(mots) == 0 {
		return fmt.Errorf("le fichier %s est vide", fichier)
	}

	rand.Seed(time.Now().UnixNano())
	indexAleatoire := rand.Intn(len(mots))

	h.ToFind = mots[indexAleatoire]
	h.RevealedLetters = make([]bool, len(h.ToFind))
	h.UsedLetters = []string{}

	h.InitAttempts()

	h.InitWordMask()

	return nil
}

func contains(slice []string, letter string) bool {
	for _, v := range slice {
		if v == letter {
			return true
		}
	}
	return false
}

func (h *HangManData) InitAttempts() {
	h.Attempts = 10
}

func (h *HangManData) InitWordMask() {
	h.Word = strings.Repeat("_ ", len(h.ToFind))
}

func (h *HangManData) GetMaskedWordWithSpaces() string {

	return strings.Join(strings.Fields(h.Word), "  ")
}

func (h *HangManData) APerdu() bool {
	return h.Attempts <= 0
}

func (h *HangManData) IsWinner() bool {
	return h.MotEstRevele()
}

func (h *HangManData) MotEstRevele() bool {
	for _, revealed := range h.RevealedLetters {
		if !revealed {
			return false
		}
	}
	return true
}

func (h *HangManData) DevinerLettre(lettre string) {
	if !contains(h.UsedLetters, lettre) {
		h.UsedLetters = append(h.UsedLetters, lettre)
		if strings.Contains(h.ToFind, lettre) {

			for i, char := range h.ToFind {
				if string(char) == lettre {
					h.RevealedLetters[i] = true
				}
			}
			h.updateMaskedWord()
		} else {
			if h.Attempts > 0 {
				h.Attempts--
			}
		}
	}
}

func (h *HangManData) updateMaskedWord() {
	var masked []rune
	for i, char := range h.ToFind {
		if h.RevealedLetters[i] {
			masked = append(masked, char)
		} else {
			masked = append(masked, '_')
		}
	}
	h.Word = string(masked)
}

func (h *HangManData) LettreDejaUtilisee(lettre string) bool {
	for _, used := range h.UsedLetters {
		if used == lettre {
			return true
		}
	}
	return false
}
