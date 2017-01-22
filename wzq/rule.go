package wzq

import (
	"bytes"
	"fmt"
	"log"
)

//是否获胜
//返回 0继续，1黑胜,2白胜,3和局
func CheckWin(pos int, color int, arr []int) int {
	r := color
	defer func() {
		if r != 0 {
			log.Println("结果", r)
			if r == Color_black || r == Color_eque {
				GTotal++
			}
			if r == Color_white {
				GTotal++
				GWin++
			}
			SaveResult()
		}
	}()
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
			return r
		}
	}
	//判断竖
	for a := x - 4; a <= x; a++ {
		if stone[a][y] == color && stone[a+1][y] == color && stone[a+2][y] == color && stone[a+3][y] == color && stone[a+4][y] == color {
			log.Println("竖向5连 ", color, " 胜")
			return r
		}
	}
	//判断斜
	b := y - 4
	for a := x - 4; a <= x; a++ {
		if stone[a][b] == color && stone[a+1][b+1] == color && stone[a+2][b+2] == color && stone[a+3][b+3] == color && stone[a+4][b+4] == color {
			log.Println("斜5连 ", color, " 胜")
			return r
		}
		b++
	}
	//判断反斜
	b = y + 4
	for a := x - 4; a <= x; a++ {
		if stone[a][b] == color && stone[a+1][b-1] == color && stone[a+2][b-2] == color && stone[a+3][b-3] == color && stone[a+4][b-4] == color {
			log.Println("反斜5连 ", color, " 胜")
			return r
		}
		b--
	}
	//判断是否和局
	for _, v := range arr {
		if v == 0 {
			r = 0
			return r
		}
	}
	r = 3
	return r
}

//数组转string
func ArrayToString(arr []int) string {
	buf := bytes.NewBufferString("")
	for i := 0; i < 225; i++ {
		buf.WriteString(fmt.Sprint(arr[i]))
	}
	return buf.String()
}
