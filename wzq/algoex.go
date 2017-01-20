package wzq

import (
	"log"
)

const (
	E_MIN = -10000.0
	E_MAX = 10000.0
)

var GWinCounts int = 100
var GW []float64 = []float64{}

//计算期望
func GetE(arr []int) float64 {
	val := 0
	for i, v := range arr {
		val += v * GW[i]
	}
	return val
}

//获取黑白期望
func GetWzqEvaluate(arr []int, winArr []*Win) (float64, float64) {
	blackArr := []int{0, 0, 0, 0}
	whiteArr := []int{0, 0, 0, 0}

	for i := 0; i < GWinCounts; i++ {
		if winArr[i].Black == 5 {
			return E_MAX
		}
		if winArr[i].white {
			return E_MIN
		}
		if winArr[i].Black == 4 && winArr[i].white == 0 {
			blackArr[3]++
		}
		if winArr[i].Black == 0 && winArr[i].white == 4 {
			whiteArr[3]++
		}
		if winArr[i].Black == 3 && winArr[i].white == 0 {
			blackArr[2]++
		}
		if winArr[i].Black == 0 && winArr[i].white == 3 {
			whiteArr[2]++
		}
		if winArr[i].Black == 2 && winArr[i].white == 0 {
			blackArr[1]++
		}
		if winArr[i].Black == 0 && winArr[i].white == 2 {
			whiteArr[1]++
		}
		if winArr[i].Black == 1 && winArr[i].white == 0 {
			blackArr[0]++
		}
		if winArr[i].Black == 0 && winArr[i].white == 1 {
			whiteArr[0]++
		}
	}
	return GetE(blackArr), GetE(whiteArr)
}

//输入棋盘，期望，输出期望，位置
func Moni_Put(arr []int, val float64) (float64, int) {
	e_val := E_MIN
	pos := 0
	tmp := arr
	for i, p := range tmp {
		if p != 0 {
			continue
		}
		//不能比输入的期望高
		p = 1
		if CheckWin(p, 1, tmp) == 1 {
			return val + 1, i
		}
		eb, ew := GetWzqEvaluate(tmp)
		e = eb - ew
		if e >= val {
			return val + 1, i
		}
		//更新最大期望
		if e > e_val {
			e_val = e
			pos = i
		}
	}
	return e, i
}

//下棋，返回最大优势的点
func Put(arr []int) int {
	e_val := E_MAX
	pos := 0
	tmp := arr
	for _, p := range tmp {
		if p != 0 {
			continue
		}
		p = 2
		if CheckWin(p, 2, tmp) == 2 {
			return p
		}
		e, ep := Moni_Put(tmp, e_val)
		if e < e_val {
			e_val = e
			pos = ep
		}
	}
	return pos
}
