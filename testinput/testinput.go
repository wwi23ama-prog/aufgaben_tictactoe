package main

import (
	"fmt"

	"github.com/wwi23ama-prog/aufgaben_tictactoe/board"
	"github.com/wwi23ama-prog/aufgaben_tictactoe/game"
)

func main() {
	// Je nach Bedarf ein- oder auskommentieren.

	TestAskForMove()
	// TestMakeMove()
}

func TestAskForMove() {
	result := game.AskForMove("X")
	fmt.Printf("AskForMove(\"X\") == %d\n", result)
}

func TestMakeMove() {
	b := board.MakeBoardFromStrings(
		"OX ",
		" O ",
		"  X",
	)

	game.MakeMove("X", b)
	fmt.Printf("MakeMove(\"X\", b) == \n%s\n", b)
}
