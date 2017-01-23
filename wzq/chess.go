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
	Color_eque  = 3
)

type Win struct {
	id    int
	White int
	Black int
}

type Chess struct {
	Started       bool
	MuxChess      sync.Mutex
	cookie        string
	Current       []int
	CurrentString string
	lastResult    int
	lastTime      int64
	lastE         float64
	MapWinArr     map[int]*Win ///当前棋局（经过转换）[赢法序号](状态)
}

var GChess *Chess = &Chess{}

//获取当前状态，如果未开始，返回上一局结果，否则返回当前棋盘
func (c *Chess) GetCurrent() (string, int) {
	c.MuxChess.Lock()
	defer c.MuxChess.Unlock()
	r := 0
	if !c.Started {
		r = c.lastResult
	}
	return c.CurrentString, r
}

//请求新游戏，如果已经开始，返回空，否则初始化游戏，返回cookie
func (c *Chess) NewGame() string {
	c.MuxChess.Lock()
	defer c.MuxChess.Unlock()
	if c.Started {
		return ""
	}
	log.Println("新游戏开始")
	c.Started = true
	c.Current = make([]int, 225)
	c.CurrentString = ArrayToString(c.Current)
	c.cookie = fmt.Sprintf("%x", md5.Sum([]byte(time.Now().String())))
	c.lastTime = time.Now().Unix()
	c.MapWinArr = make(map[int]*Win)
	//log.Println(c.MapWinArr)
	return c.cookie
}

//返回值  格式是否正确,谁赢，下的点
func (c *Chess) GetResult(cookie string, pos int) (bool, int, int) {
	c.MuxChess.Lock()
	defer c.MuxChess.Unlock()
	if c.cookie != cookie || pos < 0 || pos > 224 || c.Current[pos] != 0 {
		return false, 0, 0
	}
	//超时结束
	if !c.Started {
		return true, 4, -1
	}
	//人类下棋
	c.lastTime = time.Now().Unix()
	c.Current[pos] = Color_black
	//log.Println(c.MapWinArr)
	UpdateWinMap(c.MapWinArr, pos, Color_black)
	c.CurrentString = ArrayToString(c.Current)
	r := CheckWin(pos, Color_black, c.Current, false)

	//未分胜负，机器出牌
	if r == 0 {
		tmpwArr := make(map[int]Win)
		for k, v := range c.MapWinArr {
			tmpwArr[k] = *v
		}
		x := GetXVlues(c.Current, tmpwArr)
		enow := GetE(x)
		log.Println("当前期望:", enow)
		if c.lastE != 0 {
			UpdateW(x, enow, c.lastE)
		}
		c.lastE = enow

		time.Sleep(time.Millisecond * 100)
		posWhite := Put(c.Current, c.MapWinArr)
		c.Current[posWhite] = Color_white
		UpdateWinMap(c.MapWinArr, posWhite, Color_white)
		c.CurrentString = ArrayToString(c.Current)
		rw := CheckWin(posWhite, Color_white, c.Current, false)
		if rw == 0 {
			return true, rw, posWhite
		} else {
			//机器赢或平，更新期望
			tmpwArr := make(map[int]Win)
			for k, v := range c.MapWinArr {
				tmpwArr[k] = *v
			}
			x := GetXVlues(c.Current, tmpwArr)
			enow := GetE(x)
			if rw == Color_white {
				log.Println("机器赢,期望", enow)
				UpdateW(x, E_MIN, enow)
			} else {
				log.Println("平局,期望", enow)
				UpdateW(x, 0, enow)
			}
			c.Started = false
			c.lastResult = rw
			return true, rw, posWhite
		}
	} else {
		//人赢或平,更新期望
		tmpwArr := make(map[int]Win)
		for k, v := range c.MapWinArr {
			tmpwArr[k] = *v
		}
		x := GetXVlues(c.Current, tmpwArr)
		enow := GetE(x)
		if r == Color_black {
			log.Println("人赢,期望", enow)
			UpdateW(x, E_MAX, enow)
		} else {
			log.Println("平局,期望", enow)
			UpdateW(x, 0, enow)
		}

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
