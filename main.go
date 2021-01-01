package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//4つの運勢を設定
	result := []string{"大吉", "中吉", "小吉", "凶"}

	//乱数を設定
	rand.Seed(time.Now().UnixNano())

	//4つの運勢でどれかがランダムで選ばれる
	n := rand.Intn(len(result))
	res := result[n]

	//go run main.go -myname 〇〇　を入れると運勢が出る
	var stringFlag string
	flag.StringVar(&stringFlag, "myname", "", "this is stringflag.")
	flag.Parse()

	fmt.Println(res, stringFlag)

}
