package main

import (
	"./wzq"
	"fmt"
	//"io"
	"log"
	"net/http"
	"strconv"
	"sync"
)

const (
	fmt_query  = `{"current":"%s","msg":"%s"}`
	fmt_cookie = `{"cookie":"%s"}`
	fmt_click  = `{"index":%d,"msg":"%s"}`
)

var GMuxClicked sync.Mutex

func Cb_query(rw http.ResponseWriter, req *http.Request) {
	//{"current":"01020102020101021212121"}
	replay, rst := wzq.GChess.GetCurrent()
	var msg string
	if rst == wzq.Color_black {
		msg = "黑胜"
	}
	if rst == wzq.Color_white {
		msg = "白胜"
	}
	if rst == wzq.Color_eque {
		msg = "和棋"
	}
	fmt.Fprintf(rw, fmt_query, replay, msg)
}

func Cb_start(rw http.ResponseWriter, req *http.Request) {
	//如果未开始  返回cook {"cookie":"123231231"}
	//如果开始  query()
	cookie := wzq.GChess.NewGame()
	if len(cookie) == 0 {
		Cb_query(rw, req)
		return
	}
	log.Println("开始新游戏", req.RemoteAddr)
	fmt.Fprintf(rw, fmt_cookie, cookie)

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
		if over == wzq.Color_black {
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

func main() {
	wzq.GChess.Start()
	log.Println("五子棋服务启动...")
	http.HandleFunc("/start", Cb_start)
	http.HandleFunc("/query", Cb_query)
	http.HandleFunc("/click", Cb_click)
	http.Handle("/", http.FileServer(http.Dir("web")))
	http.ListenAndServe(":8080", nil)
}
