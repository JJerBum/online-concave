package main

import (
	"fmt"
	"online-concave/board"
)

// TODO: test 코드 작성

func main() {
	// 보드 판 출력
	myBoard := board.New()
	myBoard.Initialize()
	isWhiteTurn := true

	var inputRune rune
	var inputNum int

	for {
		myBoard.Render()

		if _, err := fmt.Scanf("%c %d", &inputRune, &inputNum); err != nil {
			fmt.Println("Retry Enter, ", err)
			continue
		}

		if isWhiteTurn {
			if err := myBoard.PutPiece(inputNum, int(inputRune)-'a', board.White); err != nil {
				fmt.Println(err)
				continue
			}

		} else {
			if err := myBoard.PutPiece(inputNum, int(inputRune)-'a', board.Black); err != nil {
				fmt.Println(err)
				continue
			}
		}

		if winner := myBoard.FindWinner(); winner != board.None {
			switch winner {
			case board.White:
				fmt.Println("Winner is White Player!")
			case board.Black:
				fmt.Println("Winner is Black Player!")
			}
			break
		}

		isWhiteTurn = !isWhiteTurn
	}

	myBoard.Render()

}
