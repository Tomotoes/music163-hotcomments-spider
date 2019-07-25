package main

import (
	"fmt"
	. "music163-hotcomments-spider/dao"
	. "music163-hotcomments-spider/model"
	. "music163-hotcomments-spider/request"
	. "music163-hotcomments-spider/util"
	"runtime"
	"strings"
)

var (
	workerChanCap = func() int {
		if runtime.GOMAXPROCS(0) == 1 {
			return 0
		}
		return 1
	}()
	id          = make(chan int, workerChanCap)
	body        = make(chan string, workerChanCap)
	hotComments = make(chan []HotComments, workerChanCap)
)

func init() {
	Poster()
	Help()
}

func main() {
	for {
		fmt.Printf("\n\n请输入指令(数字ID|exit|help): ")
		command := strings.ToLower(strings.TrimSpace(GetInput()))
		switch {
		case command == "exit":
			close(id)
			close(body)
			close(hotComments)
			End()
			return
		case command == "help":
			Help()
		default:
			id <- ToInt(command)
			go Request(id, body)
			go Parse(body, hotComments)
			go Save(hotComments)
			fmt.Println("正在爬取热评...")
		}
	}
}
