package game

import "github.com/wwi23ama-prog/aufgaben_tictactoe/board"

// SwitchPlayer erwartet einen Spieler-String und liefert den anderen Spieler zurück.
func SwitchPlayer(currentPlayer string) string {
	// TODO
	return "X"
}

// PlayerWins erwartet einen Spieler-String und ein Board und liefert true zurück,
// wenn der Spieler gewonnen hat.
func PlayerWins(player string, b board.Board) bool {
	// TODO
	return false
}

// MoveAllowed erwartet ein Board und eine Zahl und liefert true zurück,
// wenn das Feld noch nicht belegt ist.
//
// Die Zahl muss zwischen 1 und 9 liegen und gibt das Feld an,
// das geprüft werden soll. Die Zählung beginnt links oben bei 1
// und geht nach rechts und dann nach unten.
//
// Liegt die Zahl außerhalb des Bereichs 1-9, wird false zurückgegeben.
func MoveAllowed(b board.Board, move int) bool {
	// TODO
	return false
}
