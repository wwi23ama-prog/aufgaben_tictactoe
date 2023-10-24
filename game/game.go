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
	// tag::initBoard[]
	board := board.MakeBoardFromStrings(
		"123",
		"456",
		"789",
	)
	// end::initBoard[]

	// Spieler und Zug-Zähler initialisieren.
	/* Hinweis:
	   Die Spieler sind X und O, X beginnt.
	   Wir legen hier als Startspieler aber O fest,
	   damit der erste SwitchPlayer-Aufruf X liefert.
	   Unsere Hauptschleife führt zuerst einen
	   SwitchPlayer-Aufruf aus, bevor der erste Zug
	   gemacht wird.
	*/
	// tag::initPlayer[]
	currentPlayer := "O"
	// end::initPlayer[]

	// Zug-Zähler initialisieren.
	/* Hinweis:
	   Wir zählen die Züge, um zu erkennen,
	   wann das Spiel unentschieden ist.
	*/
	// tag::initMoveCount[]
	moveCount := 0
	// end::initMoveCount[]

	// Hauptschleife.
	for moveCount < 9 {
		// Prüfen, ob der aktuelle Spieler gewonnen hat.
		// tag::playerWins[]
		if PlayerWins(currentPlayer, board) {
			fmt.Println("Spieler", currentPlayer, "gewinnt!")
			return
		}
		// end::playerWins[]

		// Spieler wechseln und Zug machen.
		// tag::switchPlayer[]
		currentPlayer = SwitchPlayer(currentPlayer)
		// end::switchPlayer[]

		// Zug machen und Zug-Zähler erhöhen.
		// tag::makeMove[]
		MakeMove(currentPlayer, board)
		moveCount++
		// end::makeMove[]
	}

	fmt.Println("Unentschieden!")
}
