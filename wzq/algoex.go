package wzq

import ()

const (
	E_MIN = -10000.0
	E_MAX = 10000.0
)

//计算期望
func GetE(arr []int) float64 {
	val := 0.0
	for i, v := range arr {
		val += float64(v) * GW[i]
	}
	return val
}

//获取黑白期望
func GetWzqEvaluate(arr []int, winArr []*Win) float64 {
	eArr := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < GWinCounts; i++ {
		if winArr[i].Black == 5 {
			eArr[9]++
		}
		if winArr[i].White == 5 {
			eArr[8]++
		}
		if winArr[i].Black == 4 && winArr[i].White == 0 {
			eArr[7]++
		}
		if winArr[i].Black == 0 && winArr[i].White == 4 {
			eArr[6]++
		}
		if winArr[i].Black == 3 && winArr[i].White == 0 {
			eArr[5]++
		}
		if winArr[i].Black == 0 && winArr[i].White == 3 {
			eArr[4]++
		}
		if winArr[i].Black == 2 && winArr[i].White == 0 {
			eArr[3]++
		}
		if winArr[i].Black == 0 && winArr[i].White == 2 {
			eArr[2]++
		}
		if winArr[i].Black == 1 && winArr[i].White == 0 {
			eArr[1]++
		}
		if winArr[i].Black == 0 && winArr[i].White == 1 {
			eArr[0]++
		}
	}
	return GetE(eArr)
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
		e := GetWzqEvaluate(tmp, GChess.WinArr)
		if e >= val {
			return val + 1, i
		}
		//更新最大期望
		if e > e_val {
			e_val = e
			pos = i
		}
	}
	return e_val, pos
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
