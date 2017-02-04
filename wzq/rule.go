package wzq

import (
	"bytes"
	"fmt"
	"log"
)

//是否获胜
//返回 0继续，1黑胜,2白胜,3和局
func CheckWin(pos int, color int, arr []int, notreg bool) int {
	r := color
	defer func() {
		if r != 0 && !notreg {
			//log.Println("结果", r)
			if r == Color_black || r == Color_eque {
				GTotal++
			}
			if r == Color_white {
				GTotal++
				GWin++
			}
			log.Println("total", GTotal, "win", GWin)
			log.Println("W", GW)
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
			//log.Println("横向5连 ", color, " 胜")
			return r
		}
	}
	//判断竖
	for a := x - 4; a <= x; a++ {
		if stone[a][y] == color && stone[a+1][y] == color && stone[a+2][y] == color && stone[a+3][y] == color && stone[a+4][y] == color {
			//log.Println("竖向5连 ", color, " 胜")
			return r
		}
	}
	//判断斜
	b := y - 4
	for a := x - 4; a <= x; a++ {
		if stone[a][b] == color && stone[a+1][b+1] == color && stone[a+2][b+2] == color && stone[a+3][b+3] == color && stone[a+4][b+4] == color {
			//log.Println("斜5连 ", color, " 胜")
			return r
		}
		b++
	}
	//判断反斜
	b = y + 4
	for a := x - 4; a <= x; a++ {
		if stone[a][b] == color && stone[a+1][b-1] == color && stone[a+2][b-2] == color && stone[a+3][b-3] == color && stone[a+4][b-4] == color {
			//log.Println("反斜5连 ", color, " 胜")
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

//index黑子当前位置   arr当前棋局  返回是否禁手
func xyToIndex(x, y int) int {
	return y*15 + x
}

func indexToXY(index int) (int, int) {
	return index % 15, index / 15
}

//index黑子当前位置   arr当前棋局  返回是否禁手
func Forbid(index int, arr []int) (bool, int) {
	if index > 224 || index < 0 || len(arr) < 225 {
		log.Println("输入参数有误")
		return false, 0
	}
	arr[index] = 1
	x, y := indexToXY(index)
	j4, j3 := 0, 0
	//liveCont := 0
	//当前棋局 一维数换成二维
	stone := make([][]int, 15)
	for i := 0; i < 15; i++ {
		for j := 0; j < 15; j++ {
			stone[i] = append(stone[i], arr[15*i+j])
		}
	}
	//横
	j4 = j4 + rule_Four(x, stone[y])
	j3 = j3 + rule_Three(x, stone[y])
	if rule_TooLong(x, stone[y]) {
		return true, End_toLong
	}
	//竖
	var stone1 []int = make([]int, 15)
	for i := 0; i < 15; i++ {
		stone1[i] = stone[i][x]
	}
	j4 = j4 + rule_Four(y, stone1)
	j3 = j3 + rule_Three(y, stone1)
	if rule_TooLong(y, stone1) {
		return true, End_toLong
	}
	//斜（\）
	var stone2 []int = make([]int, 0)
	xyIndexStone2 := 0
	i := (x - min(x, y))
	j := (y - min(x, y))
	for i < 15 && j < 15 {
		stone2 = append(stone2, stone[j][i])
		if j == y && i == x {
			xyIndexStone2 = len(stone2) - 1
		}
		i++
		j++
	}
	j4 = j4 + rule_Four(xyIndexStone2, stone2)
	j3 = j3 + rule_Three(xyIndexStone2, stone2)
	if rule_TooLong(xyIndexStone2, stone2) {
		return true, End_toLong
	}
	//反斜 （/）先把棋盘上下翻转这样取值方便 上下翻转X不变，Y变化
	var stone3 []int = make([]int, 0)
	fx, fy := x, y-14
	if fy < 0 {
		fy = -fy
	}
	xyIndexStone3 := 0
	i = (fx - min(fx, fy))
	j = (fy - min(fx, fy))
	fstone := make([][]int, 0)
	for i := 14; i >= 0; i-- {
		fstone = append(fstone, stone[i])
	}
	for i < 15 && j < 15 {
		stone3 = append(stone3, fstone[j][i])
		if j == fy && i == fx {
			xyIndexStone3 = len(stone3) - 1
		}
		i++
		j++
	}
	j4 = j4 + rule_Four(xyIndexStone3, stone3)
	j3 = j3 + rule_Three(xyIndexStone3, stone3)
	if rule_TooLong(xyIndexStone3, stone3) {
		return true, End_toLong
	}
	if j4 >= 2 {
		return true, End_j4
	}
	if j3 >= 2 {
		return true, End_j3
	}
	return false, 0
}

//找出两个比较小的数
func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

//长连禁手，规则为出现长于5子的连棋
func rule_TooLong(index int, stone []int) bool {
	num := 1
	for i := index - 1; i >= 0; i-- {
		if stone[i] == 1 {
			num++
		} else {
			break
		}
	}
	for i := index + 1; i < len(stone); i++ {
		if stone[i] == 1 {
			num++
		} else {
			break
		}
	}
	if num > 5 {
		return true
	} else {
		return false
	}
}

//有多少个四子 index当前下的  stone当前某列或某行或某斜
func rule_Four(index int, stone []int) int {
	//当前有几个四子
	four := 0
	start := index - 4
	if start < 0 {
		start = 0
	}
	end := index + 4
	if end > len(stone) {
		end = len(stone)
	}
	for i := start; i < end; i++ {
		hasindex := false //这个次循环中是否包含刚刚所下的字
		blacknum := 0     //黑子数量
		j := i
		for ; j <= i+4 && j < len(stone); j++ {
			if j == index {
				hasindex = true
			}
			if stone[j] == 1 {
				blacknum++
			}
			if stone[i] != 1 || stone[j] == 2 {
				break
			}
			if blacknum == 4 {
				break
			}
		}
		if blacknum == 4 && hasindex {
			if i-1 >= 0 && j+1 < len(stone) {
				if stone[i] == 2 && stone[j+1] == 2 {
					continue
				}
			}
			if i == 0 {
				if (j+1 < len(stone)) && stone[j+1] == 2 {
					continue
				}
			}
			if j == len(stone)-1 {
				if i-1 >= 0 && stone[i] == 2 {
					continue
				}
			}
			four++
		}
	}
	return four
}

//有多少个活三子
func rule_Three(index int, stone []int) int {
	//log.Println(stone)
	start := index - 3
	if start <= 0 {
		start = 1 //原因到最前面了，不可能是活三
	}
	end := index + 3
	if end >= len(stone)-2 {
		end = len(stone) - 2 //到最后面了，也不可能是活三
	}
	//log.Println(start, end)
	for i := start; i < end; i++ {
		hasindex := false
		blacknum := 0
		j := i
		for ; j <= i+3 && j < len(stone)-1; j++ {
			if j == index {
				hasindex = true
			}
			if stone[j] == 1 {
				blacknum++
			}
			if blacknum == 3 {
				break
			}
			if stone[i] != 1 || stone[j] == 2 {
				break
			}
		}
		//log.Println(i, blacknum, hasindex, j+1)
		if stone[i-1] == 0 && blacknum == 3 && hasindex && stone[j+1] == 0 {
			return 1
		}
	}
	return 0
}
