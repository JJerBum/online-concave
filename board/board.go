package board

import (
	"errors"
	"fmt"
)

var isNotValidError = errors.New("This position is not valid")

type PieceType uint8

const (
	None PieceType = iota
	White
	Black
)

const (
	// 가로
	BoardLength int = 15

	// 세로
	BoardWidth int = 15

	BlackPieceUnicode = '●'
	WhitePieceUnicode = '○'
)

type Board [BoardWidth][BoardLength]PieceType

func New() *Board {
	return new(Board)
}

func (b *Board) Initialize() {
	for outer := 0; outer < BoardLength; outer++ {
		for inner := 0; inner < BoardWidth; inner++ {
			b[outer][inner] = None
		}
	}
}

func (b *Board) Render() {
	var charIndex = 'a'
	var numIndex = 0

	// 윗줄의 번호 영어 출력
	for outer := 0; outer < BoardLength; outer++ {
		fmt.Printf("%2c", charIndex)
		charIndex++
	}

	// 개행 출력
	fmt.Println()

	for outer := 0; outer < BoardLength; outer++ {
		for inner := 0; inner < BoardWidth; inner++ {
			var want rune
			switch b[outer][inner] {
			case White:
				want = WhitePieceUnicode
			case Black:
				want = BlackPieceUnicode
			case None:
				want = rune(' ')
			}
			fmt.Printf("%2c", want)
		}
		// 마지막 줄에 숫자 출력
		fmt.Printf("%3d", numIndex)
		numIndex++

		// 개행
		fmt.Println()
	}
}

func (b *Board) FindWinner() PieceType {
	// TODO: MinMax 알고리즘을 사용하여 속도 향상

	// 현재 사용할 알고리즘은, 15 * 15 = 255 개의 모든 경우의 수를 탐색하여 결과값을 내게 함.
	for outer := 0; outer < BoardLength; outer++ {
		for inner := 0; inner < BoardWidth; inner++ {
			if b[outer][inner] != None && b.isWin(outer, inner) {
				switch b[outer][inner] {
				case White:
					return White
				case Black:
					return Black
				}
			}
		}
	}
	return None
}

func (b *Board) PutPiece(row, col int, piece PieceType) error {
	if b[row][col] != None {
		return isNotValidError
	}

	b[row][col] = piece
	return nil
}

// NOT EXPORT, using from FindWinner func
func (b *Board) isWin(row, col int) bool {

	currentPiece := b[row][col]

	horizontal := func() bool {
		startRow := row - 2
		endRow := row + 2

		if b.isInBoard(startRow, col) == false {
			return false
		}

		if b.isInBoard(endRow, col) == false {
			return false
		}

		// 왼쪽끝에서 오른쪽 끝으로
		for outer := startRow; outer <= endRow; outer++ {
			if b[outer][col] != currentPiece {
				return false
			}
		}

		return true
	}

	vertical := func() bool {
		startCol := col - 2
		endCol := col + 2

		if b.isInBoard(row, startCol) == false {
			return false
		}

		if b.isInBoard(row, endCol) == false {
			return false
		}

		// 위쪽 끝에서 맨 아래 끝까지
		for outer := startCol; outer <= endCol; outer++ {
			if b[row][outer] != currentPiece {
				return false
			}
		}

		return true
	}

	upperLeftDiagonal := func() bool {
		startPoint := [2]int{row - 2, col - 2}
		endPoint := [2]int{row + 2, col + 2}

		if b.isInBoard(startPoint[0], startPoint[1]) == false {
			return false
		}

		if b.isInBoard(endPoint[0], endPoint[1]) == false {
			return false
		}

		startRow := startPoint[0]
		startCol := startPoint[1]
		for startRow <= endPoint[0] && startCol <= endPoint[1] {

			if b[startRow][startCol] != None {
				return false
			}

			startRow++
			startCol++
		}

		return true
	}

	upperRightDiagonal := func() bool {
		startPoint := [2]int{row + 2, col + 2}
		endPoint := [2]int{row - 2, col - 2}

		if b.isInBoard(startPoint[0], startPoint[1]) == false {
			return false
		}

		if b.isInBoard(endPoint[0], endPoint[1]) == false {
			return false
		}

		startRow := startPoint[0]
		startCol := startPoint[1]
		for startRow <= endPoint[0] && startCol <= endPoint[1] {

			if b[startRow][startCol] != None {
				return false
			}

			startRow--
			startCol--
		}

		return true
	}

	return horizontal() || vertical() || upperLeftDiagonal() || upperRightDiagonal()
}

// NOT EXPORT
func (b *Board) isInBoard(row, col int) bool {
	return (row >= 0 && row <= BoardLength-1) && (col >= 0 && col <= BoardWidth-1)
}
