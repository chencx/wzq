package main

import (
	"./wzq"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

const (
	fmt_query  = `{"current":"%s","msg":"%s"}`
	fmt_cookie = `{"cookie":"%s","total":%d,"win":%d}`
	fmt_click  = `{"index":%d,"msg":"%s"}`
	fmt_record = `{"total":%d,"win":%d}`
)

var GMuxClicked sync.Mutex

func Cb_query(rw http.ResponseWriter, req *http.Request) {
	//{"current":"01020102020101021212121"}
	replay, rst := wzq.GChess.GetCurrent()
	var msg string
	if rst == wzq.Color_black {
		msg = "黑胜"
	} else if rst == wzq.Color_white {
		msg = "白胜"
	} else if rst == wzq.Color_eque {
		msg = "和棋"
	} else if rst == -1 {
		msg = "黑棋下禁手，白胜"
	}
	fmt.Fprintf(rw, fmt_query, replay, msg)
}

func Cb_start(rw http.ResponseWriter, req *http.Request) {
	//如果未开始  返回cook {"cookie":"123231231"}
	//如果开始  query()
	gamemod := req.FormValue("gamemod")
	cookie := wzq.GChess.NewGame(gamemod)
	if len(cookie) == 0 {
		Cb_query(rw, req)
		return
	}
	log.Println("开始新游戏", req.RemoteAddr)
	fmt.Fprintf(rw, fmt_cookie, cookie, wzq.GTotal, wzq.GWin)

}
func Cb_click(rw http.ResponseWriter, req *http.Request) {
	//返回  {"index":12,"msg":""}   有输赢{"index":-1,"msg":"谁胜利"}
	GMuxClicked.Lock()
	defer GMuxClicked.Unlock()

	cookie := req.FormValue("cookie")
	pos := req.FormValue("index")
	posInt, _ := strconv.Atoi(pos)
	err, over, rst := wzq.GChess.GetResult(cookie, posInt)
	if !err {
		log.Println("收到错误的指令", cookie, posInt)
		Cb_query(rw, req)
		return
	}
	//游戏结束
	if over > 0 {
		if over == wzq.End_j3 {
			log.Println("游戏结束,黑方三三禁手,白胜")
			fmt.Fprintf(rw, fmt_click, rst, "三三禁手,你输了")
		} else if over == wzq.End_j4 {
			log.Println("游戏结束,黑方四四禁手,白胜")
			fmt.Fprintf(rw, fmt_click, rst, "四四禁手,你输了")
		} else if over == wzq.End_toLong {
			log.Println("游戏结束,黑方长连禁手,白胜")
			fmt.Fprintf(rw, fmt_click, rst, "长连禁手,你输了")
		} else if over == wzq.Color_black {
			log.Println("游戏结束,黑胜")
			fmt.Fprintf(rw, fmt_click, rst, "你赢了")
		} else if over == wzq.Color_white {
			log.Println("游戏结束,白胜")
			fmt.Fprintf(rw, fmt_click, rst, "你输了")
		} else if over == wzq.Color_eque {
			log.Println("游戏结束,和棋")
			fmt.Fprintf(rw, fmt_click, rst, "和棋")
		} else {
			log.Println("黑方超时,游戏结束,白胜")
			fmt.Fprintf(rw, fmt_click, rst, "超时，你输了")
		}
		return
	}
	fmt.Fprintf(rw, fmt_click, rst, "")
}

func Cb_record(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, fmt_record, wzq.GTotal, wzq.GWin)
}

func main() {
	flog, _ := os.OpenFile("wzq.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0x666)
	log.SetOutput(flog)
	log.SetFlags(log.LstdFlags)
	wzq.InitWeight()
	wzq.InitWinArr()
	wzq.InitPosMap()
	wzq.GChess.Start()
	log.Println("ZC五子棋v2.5 服务启动...")
	http.HandleFunc("/start", Cb_start)
	http.HandleFunc("/query", Cb_query)
	http.HandleFunc("/click", Cb_click)
	http.HandleFunc("/queryrecord", Cb_record)
	http.Handle("/", http.FileServer(http.Dir("web")))
	http.ListenAndServe(":8080", nil)
}
