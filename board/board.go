package board

import "fmt"

type printer interface {
	print()
}

type Board struct {
}

func (b *Board) print() {
	fmt.Println("   A B C D E F G H I J K L M N O")
	fmt.Println("   ------------------------------")
	for i := 0; i < b.BoardWidth; i++ {
		// 세로 좌표를 출력할 때, 일정한 폭을 맞추어 출력합니다.
		fmt.Printf("%2d|", i+1)
		for j := 0; j < b.BoardLength; j++ {
			fmt.Printf("%s ", b[i][j]) // 각 교점의 돌 또는 점 출력
		}
		// 세로 좌표를 출력할 때, 일정한 폭을 맞추어 출력합니다.
		fmt.Printf("|%2d\n", i+1)
	}
	fmt.Println("   ------------------------------")
	fmt.Println("   A B C D E F G H I J K L M N O")
}

type BoardPrintr struct {
}

func (b BoardPrintr) Print(p printer) {
	p.print()
}
