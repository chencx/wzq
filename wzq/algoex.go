package wzq

import (
	"log"
	//"time"
)

const (
	E_MIN = -100000.0
	E_MAX = 100000.0
)

//更新棋盘状态
func UpdateWinMap(winArr map[int]*Win, pos int, color int) {
	//log.Println("更新前", winArr)
	//panic("1111111")
	arr := GPosMap[pos]
	for _, v := range arr {
		_, ok := winArr[v]
		if !ok {
			st := &Win{}
			st.White = 0
			st.Black = 0
			winArr[v] = st
		}
		if color == Color_black {
			winArr[v].Black += 1
		} else {
			winArr[v].White += 1
		}
	}
	//log.Println("更新后", winArr)
}

//调整权
func UpdateW(x []int, E, Enow float64) {
	for i, w := range GW {
		w += 0.2 * float64(x[i]) * (E - Enow)
	}
}

//计算期望
func GetE(arr []int) float64 {
	val := 0.0
	for i, v := range arr {
		val += float64(v) * GW[i]
	}
	return val
}

//获取棋盘状态
func GetXVlues(arr []int, winArr map[int]Win, pos, color int) []int {
	eArr := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	leng := len(GPosMap[pos])
	if color == Color_black {
		eArr[11] = leng
	}
	if color == Color_white {
		eArr[10] = -leng
	}
	for _, v := range winArr {
		if v.Black == 5 {
			eArr[9]++
		}
		if v.White == 5 {
			eArr[8]++
		}
		if v.Black == 4 && v.White == 0 {
			eArr[7]++
		}
		if v.Black == 0 && v.White == 4 {
			eArr[6]++
		}
		if v.Black == 3 && v.White == 0 {
			eArr[5]++
		}
		if v.Black == 0 && v.White == 3 {
			eArr[4]++
		}
		if v.Black == 2 && v.White == 0 {
			eArr[3]++
		}
		if v.Black == 0 && v.White == 2 {
			eArr[2]++
		}
		if v.Black == 1 && v.White == 0 {
			eArr[1]++
		}
		if v.Black == 0 && v.White == 1 {
			eArr[0]++
		}
	}
	//log.Println("棋盘状态", eArr)
	return eArr
}

//输入棋盘，期望，输出期望，位置
func Moni_Put(arr []int, val float64, wArr map[int]Win) float64 {
	e_val := E_MIN
	pos := 0
	tmp := make([]int, 225)
	copy(tmp, arr)
	//log.Println("当前期望", tmp, val)
	for i, p := range tmp {
		if p != 0 {
			continue
		}
		tmpwArr := make(map[int]Win)
		for k, v := range wArr {
			tmpwArr[k] = v
		}
		//不能比输入的期望高
		tmp[i] = 1
		//log.Println("下", i)
		arrpos := GPosMap[i]
		for _, v := range arrpos {
			w, ok := tmpwArr[v]
			if !ok {
				st := Win{}
				st.White = 0
				st.Black = 1
				tmpwArr[v] = st
			} else {
				w.Black += 1
				tmpwArr[v] = w
			}
		}
		//UpdateWinMap(tmpwArr, i, Color_black)
		if CheckWin(p, 1, tmp) == 1 {
			tmp[i] = 0
			log.Println("人下", i, "赢")
			return val + 1
		}
		e := GetE(GetXVlues(tmp, tmpwArr, i, Color_black))
		if e >= val {
			tmp[i] = 0
			//log.Println("期望大于传入", e)
			return e
		}
		//更新最大期望
		if e > e_val {
			//log.Println("更新期望", e, "位置", i)
			e_val = e
			pos = i
		}
		tmp[i] = 0
	}
	//log.Println("人期望：", e_val, pos)
	return e_val
}

//下棋，返回最大优势的点
func Put(arr []int, wArr map[int]*Win) int {
	e_val := E_MAX
	pos := 0
	tmp := make([]int, 225)
	copy(tmp, arr)
	for i, p := range tmp {
		if p != 0 {
			continue
		}
		tmpwArr := make(map[int]Win)
		for k, v := range wArr {
			tmpwArr[k] = *v
		}
		tmp[i] = 2
		//log.Println("机下", i)
		arrpos := GPosMap[i]
		for _, v := range arrpos {
			w, ok := tmpwArr[v]
			if !ok {
				st := Win{}
				st.White = 1
				st.Black = 0
				tmpwArr[v] = st
			} else {
				w.White++
				tmpwArr[v] = w
			}
		}

		//UpdateWinMap(tmpwArr, i, Color_white)
		if CheckWin(p, 2, tmp) == 2 {
			return p
		}
		e := Moni_Put(tmp, e_val, tmpwArr)
		if e < e_val {
			e_val = e
			pos = i
			//log.Println("机器更新期望", e_val, pos)
		}
		tmp[i] = 0
	}
	log.Println("最终期望", e_val, pos)
	return pos
}
