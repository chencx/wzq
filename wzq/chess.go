package wzq

import (
	"crypto/md5"
	"fmt"
	"log"
	"sync"
	"time"
)

const (
	Color_black = 1
	Color_white = 2
)

type Chess struct {
	Started       bool
	MuxChess      sync.Mutex
	cookie        string
	Current       []int
	CurrentString string
	lastResult    int
	lastTime      int64
}

var GChess *Chess = &Chess{}

func (c *Chess) GetCurrent() (string, int) {
	c.MuxChess.Lock()
	defer c.MuxChess.Unlock()
	r := 0
	if !c.Started {
		r = c.lastResult
	}
	return c.CurrentString, r
}

func (c *Chess) NewGame() string {
	c.MuxChess.Lock()
	defer c.MuxChess.Unlock()
	if c.Started {
		return ""
	}
	log.Println("新游戏开始")
	c.Started = true
	c.Current = make([]int, 255)
	c.CurrentString = ArrayToString(c.Current)
	c.cookie = fmt.Sprintf("%x", md5.Sum([]byte(time.Now().String())))
	c.lastTime = time.Now().Unix()
	return c.cookie
}

//返回值  格式是否正确,谁赢)，下的点
func (c *Chess) GetResult(cookie string, pos int) (bool, int, int) {
	c.MuxChess.Lock()
	defer c.MuxChess.Unlock()
	if !c.Started || c.cookie != cookie || pos < 0 || pos > 224 || c.Current[pos] != 0 {
		return false, 0, 0
	}
	c.lastTime = time.Now().Unix()
	c.Current[pos] = Color_black
	c.CurrentString = ArrayToString(c.Current)
	r := CheckWin(pos, Color_black, c.Current)
	//未分胜负，机器出牌
	if r == 0 {
		time.Sleep(time.Second * 1)
		posWhite := ChessNext(c.Current)
		c.Current[posWhite] = Color_white
		c.CurrentString = ArrayToString(c.Current)
		rw := CheckWin(posWhite, Color_white, c.Current)
		if rw == 0 {
			return true, 0, posWhite
		} else {
			c.Started = false
			c.lastResult = rw
			return true, rw, posWhite
		}
	} else {
		c.Started = false
		c.lastResult = r
		return true, r, -1
	}
}

func (c *Chess) Start() {
	go c.CheckCookie()
}

func (c *Chess) CheckCookie() {
	chCheckTimes := time.NewTicker(time.Second * 2)

	for {
		select {
		case <-chCheckTimes.C:
			c.MuxChess.Lock()
			if c.Started {
				if time.Now().Unix()-c.lastTime > 60 {
					log.Println("客户端超时, 游戏结束.")
					c.Started = false
					c.lastResult = Color_white
				}
			}
			c.MuxChess.Unlock()
		}
	}
}
