package wzq

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var GWinCounts int = 100
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
	log.Println("载入成功：", GW)
	return nil
}
