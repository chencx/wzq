package wzq

import (
	"bytes"
	"fmt"
	"log"
)

//是否获胜
//返回0继续，1黑胜、2白胜
func CheckWin(pos int, color int, arr []int) int {
	x := pos/15 + 4
	y := pos%15 + 4

	stone := [23][23]int{}
	for i := 0; i < 15; i++ {
		for j := 0; j < 15; j++ {
			stone[i+4][j+4] = arr[15*i+j]
		}
	}

	//判断横
	for b := y - 4; b <= y; b++ {
		if stone[x][b] == color && stone[x][b+1] == color && stone[x][b+2] == color && stone[x][b+3] == color && stone[x][b+4] == color {
			log.Println("横向5连 ", color, " 胜")
			return color
		}
	}
	//判断竖
	for a := x - 4; a <= x; a++ {
		if stone[a][y] == color && stone[a+1][y] == color && stone[a+2][y] == color && stone[a+3][y] == color && stone[a+4][y] == color {
			log.Println("竖向5连 ", color, " 胜")
			return color
		}
	}
	//判断斜
	b := y - 4
	for a := x - 4; a <= x; a++ {
		if stone[a][b] == color && stone[a+1][b+1] == color && stone[a+2][b+2] == color && stone[a+3][b+3] == color && stone[a+4][b+4] == color {
			log.Println("斜5连 ", color, " 胜")
			return color
		}
		b++
	}
	//判断反斜
	b = y + 4
	for a := x - 4; a <= x; a++ {
		if stone[a][b] == color && stone[a+1][b-1] == color && stone[a+2][b-2] == color && stone[a+3][b-3] == color && stone[a+4][b-4] == color {
			log.Println("反斜5连 ", color, " 胜")
			return color
		}
		b--
	}
	return 0
}

func ArrayToString(arr []int) string {
	buf := bytes.NewBufferString("")
	for i := 0; i < 255; i++ {
		buf.WriteString(fmt.Sprint(arr[i]))
	}
	return buf.String()
}

//传入棋局，传出下棋位置
func ChessNext(arr []int) int {
	for i := 0; i < 255; i++ {
		if arr[i] == 0 {
			return i
		}
	}
	return -1
}
