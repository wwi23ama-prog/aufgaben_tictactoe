package botgame

import (
	"github.com/wwi23ama-prog/aufgaben_tictactoe/board"
)

// CopyBoard erwartet ein Board und liefert eine Kopie des Boards zurück.
func CopyBoard(b board.Board) board.Board {
	newBoard := board.MakeBoard(" ", len(b), len(b[0]))
	// TODO
	return newBoard
}

// WinningMove erwartet einen Spieler-String und ein Board und liefert
// die Zahl für einen Zug zurück, der zum Gewinn führt.
// Ist kein solcher Zug möglich, wird -1 zurückgegeben.
//
// Die Zahl liegt zwischen 1 und 9 und gibt das Feld an,
// auf dem der Zug gemacht werden muss, um zu gewinnen.
// Die Zählung beginnt links oben bei 1 und geht nach rechts und dann nach unten.
// Falls mehrere Züge zum Gewinn führen, wird einer davon zurückgegeben.
func WinningMove(player string, b board.Board) int {
	/*
		boardcopy := CopyBoard(b)
		field := 1
	*/
	// TODO (Kommentieren Sie die Zeilen am Anfang der Funktion ein.)
	return -1
}

// BotMove erwartet einen Spieler-String und ein Board und liefert
// die Zahl für einen Zug zurück, den der Bot machen soll.
// Ist kein solcher Zug möglich, wird -1 zurückgegeben.
//
// Die Zahl liegt zwischen 1 und 9 und gibt das Feld an,
// auf dem der Zug gemacht werden soll.
//
// Der Bot versucht zu gewinnen, indem er die Funktion WinningMove verwendet.
// Falls kein Gewinn möglich ist, versucht er, den Gegner am Gewinnen zu hindern.
// Weitere Intelligenz ist nicht gefordert. Der Bot muss also nicht perfekt spielen.
func BotMove(player string, b board.Board) int {
	// TODO
	return -1
}
