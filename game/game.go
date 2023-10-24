package game

import (
	"fmt"

	"github.com/wwi23ama-prog/aufgaben_tictactoe/board"
)

func Run() {

	// Spielfeld initialisieren.

	/* Hinweis:
	   Erzeugen Sie hier ein leeres Spielfeld.
	   Leer heißt in diesem FAll, dass kein Feld "X" oder "O" enthält.
	   D.h. Sie haben die Wahl, ob Sie ein Feld mit Leerzeichen
	   initialisieren oder es z.B. mit den Zahlen 1-9 initialisieren,
	   um den Spielern die Eingabe zu erleichtern.
	*/
	board := board.MakeBoardFromStrings(
		"123",
		"456",
		"789",
	)

	// Spieler und Zug-Zähler initialisieren.
	/* Hinweis:
	   Die Spieler sind X und O, X beginnt.
	   Wir legen hier als Startspieler aber O fest,
	   damit der erste SwitchPlayer-Aufruf X liefert.
	   Unsere Hauptschleife führt zuerst einen
	   SwitchPlayer-Aufruf aus, bevor der erste Zug
	   gemacht wird.
	*/
	currentPlayer := "O"

	// Zug-Zähler initialisieren.
	/* Hinweis:
	   Wir zählen die Züge, um zu erkennen,
	   wann das Spiel unentschieden ist.
	*/
	moveCount := 0

	// Hauptschleife.
	for moveCount < 9 {
		// Prüfen, ob der aktuelle Spieler gewonnen hat.
		if PlayerWins(currentPlayer, board) {
			fmt.Println("Spieler", currentPlayer, "gewinnt!")
			return
		}

		// Spieler wechseln und Zug machen.
		currentPlayer = SwitchPlayer(currentPlayer)

		// Zug machen und Zug-Zähler erhöhen.
		MakeMove(currentPlayer, board)
		moveCount++
	}

	fmt.Println("Unentschieden!")
}
