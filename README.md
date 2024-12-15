Hangman - Jeu de Pendu 🎮
Hangman est une version web du célèbre jeu de Pendu. Ce projet a été développé en Go, HTML et CSS.

Installation 🔧
Prérequis
Avant de commencer, assurez-vous d'avoir Go installé sur votre machine. Vous pouvez le télécharger et l'installer à partir de ce lien.

Étapes d'installation
Clonez le repository Git sur votre machine :

bash
Copier le code
git clone https://github.com/7n4xt/Hangman-Web-LesSwappers.git
Ouvrez un terminal dans le dossier du projet cloné et exécutez la commande suivante pour démarrer l'application :

bash
Copier le code
go run main.go
Une fois l'application lancée, ouvrez votre navigateur et accédez à l'adresse suivante : http://localhost:8080 pour commencer à jouer !

Fonctionnalités 🎮
Page d'Accueil (/index)
Description : La page d'accueil avec les règles du jeu et un bouton pour démarrer une nouvelle session.
Action : Choisir le niveau de difficulté (Facile, Moyen, Difficile).
Jeu (/hangman)
Description : L'interface principale du jeu où le joueur devine les lettres d'un mot.
Logique : Le joueur commence avec un nombre de tentatives (10). À chaque mauvaise lettre, il perd une tentative.
Difficulté : Le mot à deviner change en fonction du niveau choisi (Facile, Moyen, Difficile).
Résultats (/result)
Description : Affiche les résultats du jeu une fois que le joueur a gagné ou perdu. Le joueur peut alors choisir de recommencer une partie ou de quitter.
Logique du Jeu 🤔
Le joueur commence avec 10 tentatives et doit deviner le mot en proposant des lettres.
Si le joueur devine toutes les lettres avant d'épuiser ses tentatives, il gagne et le jeu s'arrête.
Si le joueur n'a plus de tentatives restantes, il perd et le jeu s'arrête.
Une fois le jeu terminé, le mot complet est révélé et le message "Félicitations, vous avez gagné !" ou "Dommage, vous avez perdu..." est affiché.
