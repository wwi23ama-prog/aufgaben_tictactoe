package botgame

import (
	"fmt"

	"github.com/wwi23ama-prog/aufgaben_tictactoe/board"
	"github.com/wwi23ama-prog/aufgaben_tictactoe/game"
)

func Run() {

	// Spielfeld initialisieren.
	board := board.MakeBoardFromStrings(
		"123",
		"456",
		"789",
	)

	// Spieler initialisieren.
	currentPlayer := "O"

	// Zug-Zähler initialisieren.
	moveCount := 0

	// Hauptschleife.
	for moveCount < 9 {
		// Prüfen, ob der aktuelle Spieler gewonnen hat.
		if game.PlayerWins(currentPlayer, board) {
			fmt.Println("Spieler", currentPlayer, "gewinnt!")
			return
		}

		// Spieler wechseln und Zug machen.
		currentPlayer = game.SwitchPlayer(currentPlayer)

		// Zug machen und Zug-Zähler erhöhen.
		/* Hinweis:
		   Entscheiden Sie je nach aktuellem Spieler,
		   ob der Zug vom Menschen oder vom Bot gemacht wird.
		   Verwenden Sie dazu die Funktionen game.MakeMove und BotMove.
		   Falls der Bot am Zug ist, führen Sie seinen Zug aus.
		*/
		// TODO
		moveCount++
	}

	fmt.Println("Unentschieden!")
}
