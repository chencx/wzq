package wzq

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var GWinMap map[int][]int = make(map[int][]int)
var GPosMap map[int][]int = make(map[int][]int)
var GW []float64 = []float64{}

func InitWeight() error {
	content, err := ioutil.ReadFile("wzq.sav")
	if err != nil {
		log.Println("打开存档文件失败", err)
		return err
	}

	arr := strings.Split(string(content), "\r\n")
	for _, s := range arr {
		line := strings.Split(s, "=")
		arrw := strings.Split(line[1], "*")
		for _, w := range arrw {
			wi, _ := strconv.ParseFloat(w, 64)
			GW = append(GW, wi)
		}
	}
	//log.Println("载入成功：", GW)
	return nil
}

//初始化赢法数组
func InitWinArr() {
	seq := 0
	//横
	for i := 0; i < 15; i++ {
		for j := 0; j < 11; j++ {
			GWinMap[seq] = []int{i*15 + j, i*15 + j + 1, i*15 + j + 2, i*15 + j + 3, i*15 + j + 4}
			seq++
		}
	}
	//竖
	for i := 0; i < 15; i++ {
		for j := 0; j < 11; j++ {
			GWinMap[seq] = []int{j*15 + i, (j+1)*15 + i, (j+2)*15 + i, (j+3)*15 + i, (j+4)*15 + i}
			seq++
		}
	}
	//斜
	for i := 0; i < 11; i++ {
		for j := 4; j < 15; j++ {
			pos := i*15 + j
			GWinMap[seq] = []int{pos, pos + 14, pos + 28, pos + 42, pos + 56}
			seq++
		}
	}
	//反斜
	for i := 0; i < 11; i++ {
		for j := 0; j < 11; j++ {
			pos := i*15 + j
			GWinMap[seq] = []int{pos, pos + 16, pos + 32, pos + 48, pos + 64}
			seq++
		}
	}
}

//初始化位置对应胜利数组
func InitPosMap() {
	for i := 0; i < 225; i++ {
		for k, v := range GWinMap {
			for _, p := range v {
				if p == i {
					GPosMap[i] = append(GPosMap[i], k)
				}
			}
		}
	}
}
