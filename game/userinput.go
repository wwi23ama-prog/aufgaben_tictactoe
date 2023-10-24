package game

/******************************************************************
 *                                                                *
 * Anmerkung zu den Funktionen in dieser Datei:                   *
 *                                                                *
 * Die Funktionen in dieser Datei beinhalten Benutzereingaben.    *
 * Dies kann nicht auf einfache Weise simuliert werden und daher  *
 * gibt es auch keine automatisierten Tests für diese Funktionen. *
 *                                                                *
 * Diese Funktionen sollten daher manuell getestet werden.        *
 * Dies kann z.B. in einem eigenen `main`-Package geschehen.      *
 * Zu diesem Zweck gibt es ein separates Package `testinput`.     *
 *                                                                *
 * Es kann auch sinnvoll sein, zuerst die Funktionen in           *
 * gamelogic_test.go` zu implementieren                           *
 *                                                                *
 ******************************************************************/

import (
	"github.com/wwi23ama-prog/aufgaben_tictactoe/board"
)

// MakeMove erwartet einen Spieler-String und ein Board und lässt den Spieler
// einen Zug machen.
// Dazu zeigt die Funktion dem Spieler das Spielfeld und fragt den Spieler
// nach einer Zahl für das Feld. Anschließend setzt sie das entsprechende Feld
// auf den Spieler-String.
//
// Ist der Zug ungültig, wird der Spieler erneut nach einer Zahl gefragt.
func MakeMove(player string, b board.Board) {
	// TODO

	// TODO
}

// AskForMove fragt den Spieler nach einer Zahl für ein Feld
// auf dem Spielfeld und liefert diese zurück.
// Ist die Eingabe ungültig, wird -1 zurückgegeben.
func AskForMove(player string) int {
	// TODO
	return -1 // ungültige Eingabe
}
