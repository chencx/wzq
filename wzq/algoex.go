package wzq

import (
	"log"
	"math/rand"
	"time"
)

const (
	E_MIN = -52000.0
	E_MAX = 52000.0
)

var GRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

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
	//log.Println("更新前:", GW)
	//log.Println("棋盘状态:", x)
	for i := 0; i < 16; i++ {
		GW[i] += 0.000001 * float64(x[i]) * (E - Enow)
		//log.Println("权:", i, x[i], E, Enow, GW[i])
	}
	log.Println("更新后:", GW)
}

//计算期望
func GetE(arr []int) float64 {
	val := 0.0
	for i, v := range arr {
		val += float64(v) * GW[i]
	}
	return val
}

//计算期望
func GetEX(arr []int) float64 {
	val := 0.0
	for i, v := range arr {
		val += float64(v) * GW[i]
		//log.Println(v, "*", GW[i], val)
	}
	return val
}

//计算周围是否空
func IsEmpty(x, y, pos, color int, arr []int) int {
	//四角
	if x == 0 && y == 0 {
		if (arr[pos+1] != 0 && arr[pos+1] != color) || (arr[pos+15] != 0 && arr[pos+15] != color) || (arr[pos+16] != 0 && arr[pos+16] != color) {
			return 0
		}
		return 1
	}
	if x == 0 && y == 14 {
		if (arr[pos-1] != 0 && arr[pos-1] != color) || (arr[pos+15] != 0 && arr[pos+15] != color) || (arr[pos+14] != 0 && arr[pos+14] != color) {
			return 0
		}
		return 1
	}
	if x == 14 && y == 0 {
		if (arr[pos+1] != 0 && arr[pos+1] != color) || (arr[pos-15] != 0 && arr[pos-15] != color) || (arr[pos-14] != 0 && arr[pos-14] != color) {
			return 0
		}
		return 1
	}
	if x == 14 && y == 14 {
		if (arr[pos-1] != 0 && arr[pos-1] != color) || (arr[pos-15] != 0 && arr[pos-15] != color) || (arr[pos-16] != 0 && arr[pos-16] != color) {
			return 0
		}
		return 1
	}
	//四边
	if x == 0 && y != 0 && y != 14 {
		n := 0
		if arr[pos-1] != 0 && arr[pos-1] != color {
			n++
		}
		if arr[pos+1] != 0 && arr[pos+1] != color {
			n++
		}
		if arr[pos+14] != 0 && arr[pos+14] != color {
			n++
		}
		if arr[pos+15] != 0 && arr[pos+15] != color {
			n++
		}
		if arr[pos+16] != 0 && arr[pos+16] != color {
			n++
		}
		if n > 1 {
			return 0
		}
		return 1
	}
	if x == 14 && y != 0 && y != 14 {
		n := 0
		if arr[pos-1] != 0 && arr[pos-1] != color {
			n++
		}
		if arr[pos+1] != 0 && arr[pos+1] != color {
			n++
		}
		if arr[pos-14] != 0 && arr[pos-14] != color {
			n++
		}
		if arr[pos-15] != 0 && arr[pos-15] != color {
			n++
		}
		if arr[pos-16] != 0 && arr[pos-16] != color {
			n++
		}
		if n > 1 {
			return 0
		}
		return 1
	}
	if y == 0 && x != 0 && x != 14 {
		n := 0
		if arr[pos-15] != 0 && arr[pos-15] != color {
			n++
		}
		if arr[pos+15] != 0 && arr[pos+15] != color {
			n++
		}
		if arr[pos+1] != 0 && arr[pos+1] != color {
			n++
		}
		if arr[pos-14] != 0 && arr[pos-14] != color {
			n++
		}
		if arr[pos+16] != 0 && arr[pos+16] != color {
			n++
		}
		if n > 1 {
			return 0
		}
		return 1
	}
	if y == 14 && x != 0 && x != 14 {
		n := 0
		if arr[pos-1] != 0 && arr[pos-1] != color {
			n++
		}
		if arr[pos+15] != 0 && arr[pos+15] != color {
			n++
		}
		if arr[pos-15] != 0 && arr[pos-15] != color {
			n++
		}
		if arr[pos+14] != 0 && arr[pos+14] != color {
			n++
		}
		if arr[pos-16] != 0 && arr[pos-16] != color {
			n++
		}
		if n > 1 {
			return 0
		}
		return 1
	}
	//其他
	n := 0
	if arr[pos-1] != 0 && arr[pos-1] != color {
		n++
	}
	if arr[pos+1] != 0 && arr[pos+1] != color {
		n++
	}
	if arr[pos-14] != 0 && arr[pos-14] != color {
		n++
	}
	if arr[pos-15] != 0 && arr[pos-15] != color {
		n++
	}
	if arr[pos-16] != 0 && arr[pos-16] != color {
		n++
	}
	if arr[pos+14] != 0 && arr[pos+14] != color {
		n++
	}
	if arr[pos+15] != 0 && arr[pos+15] != color {
		n++
	}
	if arr[pos+16] != 0 && arr[pos+16] != color {
		n++
	}
	if n > 2 {
		return 0
	}
	return 1

}

//获取棋盘状态
func GetXVlues(arr []int, winArr map[int]Win) []int {
	eArr := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	//大局观
	for i, p := range arr {
		if p == Color_black {
			x := i / 15
			y := i % 15
			eArr[15] += IsEmpty(x, y, i, Color_black, arr)
			eArr[13] += (7-x)*(7-x) + (7-y)*(7-y)
			eArr[11] += len(GPosMap[i])
		}
		if p == Color_white {
			x := i / 15
			y := i % 15
			eArr[14] += IsEmpty(x, y, i, Color_white, arr)
			eArr[12] += (7-x)*(7-x) + (7-y)*(7-y)
			eArr[10] += len(GPosMap[i])
		}
	}
	if eArr[13] > eArr[12] {
		eArr[13] = eArr[13] - eArr[12]
		eArr[12] = 0
	} else {
		eArr[12] = eArr[12] - eArr[13]
		eArr[13] = 0
	}
	if eArr[11] > eArr[10] {
		eArr[11] = eArr[11] - eArr[10]
		eArr[10] = 0
	} else {
		eArr[10] = eArr[10] - eArr[11]
		eArr[11] = 0
	}

	//连子
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
		// if v.Black == 0 && v.White == 1 {
		// 	eArr[0]++
		// }
	}
	//log.Println("棋盘状态", eArr)
	return eArr
}

func GetCanPut(arr []int) []int {
	r := []int{}
	for i, p := range arr {
		if p == 0 {
			r = append(r, i)
		}
	}
	l := len(r)
	for i := 0; i < l; i++ {
		n := GRand.Intn(l - 1)
		r[i], r[n] = r[n], r[i]
	}
	return r
}

//输入棋盘，期望，输出期望，位置
func Moni_Put(arr []int, val float64, wArr map[int]Win) float64 {
	e_val := -1000000.0
	//pos := 0
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
		// if CheckWin(i, 1, tmp) == 1 {
		// 	tmp[i] = 0
		// 	log.Println("人下", i, "赢")
		// 	return 1000000.0
		// }
		e := GetE(GetXVlues(tmp, tmpwArr))
		if e >= val {
			tmp[i] = 0
			//log.Println("期望大于传入", e)
			return e
		}
		//更新最大期望
		if e > e_val {
			//log.Println("更新期望", e, "位置", i)
			e_val = e
			//pos = i
		}
		tmp[i] = 0
	}
	//log.Println("人期望：", e_val, pos)
	return e_val
}

//下棋，返回最大优势的点
func Put(arr []int, wArr map[int]*Win) int {
	e_val := 1000000.0
	pos := 0
	tmp := make([]int, 225)
	copy(tmp, arr)
	rarr := GetCanPut(tmp)
	for _, i := range rarr {
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
		if CheckWin(i, 2, tmp, true) == 2 {
			return i
		}
		e := Moni_Put(tmp, e_val, tmpwArr)
		if e < e_val {
			e_val = e
			pos = i
			//log.Println("机器更新期望", e_val, pos)
		}
		tmp[i] = 0
	}
	//log.Println("最终期望", e_val, pos)
	return pos
}
