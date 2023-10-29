package main

import "fmt"

const (
	// 가로
	BoardLength int = 15

	// 세로
	BoardWidth int = 15

	// 흑 돌 외의 출력할 점
	BoardPoint = "·"
)

type board [BoardWidth][BoardLength]string

func (b *board) init(defaultPoint string) {
	for i := 0; i < BoardWidth; i++ {
		for j := 0; j < BoardWidth; j++ {
			b[i][j] = BoardPoint
		}
	}
}

func (b *board) show() {
	fmt.Println("   A B C D E F G H I J K L M N O")
	fmt.Println("  ------------------------------")
	for i := 0; i < BoardWidth; i++ {
		// 세로 좌표를 출력할 때, 일정한 폭을 맞추어 출력합니다.
		fmt.Printf("%2d|", i+1)
		for j := 0; j < BoardLength; j++ {
			fmt.Printf("%s ", b[i][j]) // 각 교점의 돌 또는 점 출력
		}
		// 세로 좌표를 출력할 때, 일정한 폭을 맞추어 출력합니다.
		fmt.Printf("|%2d\n", i+1)
	}
	fmt.Println("  ------------------------------")
	fmt.Println("   A B C D E F G H I J K L M N O")
}

func main() {
	var board board
	board.init(BoardPoint)
	board.show()

}
