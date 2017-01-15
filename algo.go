package main

import (
	"log"
	"math/rand"
	"time"
)

////////////////////////
//<参数1：棋盘
//<返回1：总评价
//<返回2：分评价
////////////////////////
func GetEvaluate(arr []int) (int, []int) {
	tmp := make([][]int, 200, 200)
	seq := 0
	//横向
	// for i := 0; i < 15; i++ {
	// 	tmp[seq] = arr[15*i : 15*(i+1)]
	// 	seq++
	// }
	// //纵向
	// for i := 0; i < 15; i++ {
	// 	t := make([]int, 15)
	// 	for j := 0; j < 15; j++ {
	// 		t[j] = arr[i+15*j]
	// 	}
	// 	tmp[seq] = t
	// 	seq++
	// }
	//斜
	// for i := 4; i < 14; i++ {
	// 	t := make([]int, i+1)
	// 	for j := 0; j <= i; j++ {
	// 		t[j] = arr[i+14*j]
	// 	}
	// 	tmp[seq] = t
	// 	seq++
	// }
	j := 4
	for i := 220; i > 210; i-- {
		t := make([]int, j+1)
		for k := 0; k <= j; k++ {
			t[k] = arr[i-14*k]
		}
		j++
		log.Println(t)
		tmp[seq] = t
		seq++
	}
	log.Println(tmp)
	return 0, []int{}
}

/////////////////////////
//<参数1：当前棋盘
//<参数2：搜索深度
//<返回1：权重/下棋位置
/////////////////////////
func AlphaSearch(arr []int, deap int) (int, int) {
	return 0, 0
}

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	arr := make([]int, 225)
	for i := 0; i < 225; i++ {
		arr[i] = r.Int() % 10
	}
	log.Println(arr)
	log.Println(GetEvaluate(arr))
}
