package wzq

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Weight struct {
	Wi []float64
}

func (w *Weight) Start() error {
	content, err := ioutil.ReadFile("wxz.sav")
	if err != nil {
		log.Println("打开存档文件失败")
		return err
	}

	w.Wi = make([]float64)
	arr := strings.Split(string(content), "*")
	for _, s := range arr {
		wi, _ := strconv.ParseFloat(s, 64)
		w.Wi = append(w.Wi, wi)
	}
	log.Println("载入成功：", w.Wi)
	return nil
}
