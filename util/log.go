package util

import "fmt"

func Log(info ...interface{}) {
	fmt.Println()
	for _, v := range info {
		fmt.Printf("\n%v", v)
	}
}

func Help() {
	Log("请仔细阅读以下步骤: ")
	Log("0. 登录Mysql,执行`source 项目根目录/create.sql`命令", "1. 修改`dao/config.go`中`USER`与`PASSWORD`常量","2. 登录网易云音乐官网(https://music.163.com).", "3. 在官网右上角搜索框,键入音乐名,并搜索.", "4. 在搜索结果列表中,点击目标音乐.", "5. 复制页面地址栏中的值(https://music.163.com/#/song?id=复制的值).", "6. 将复制的值粘贴到命令行,等待获取结果.")
	Log("输入 `exit` ,退出程序.", "输入 `help` ,打印帮助信息.")
}

func Poster() {
	fmt.Println(`+-------------------------+-----+--+--+--+--+--+--+
|                         |     |  |  |  |  |  |  |
|                         |     |  |  |  |  |  |  |
|    Hot Comments Spider  |     +--+--+--+--+--+--+
|                         v                       |
|                             -- music.163.com    |
+-->   X      X    XXXXXXXXX                      |
|     XX     XX        X                          |
|    X X    X X        X       <------------------+
|   X   X  X   X       X                          |
|  X     XX     X      X       Anthor: SimonMa    |
| X      X       X   XXX                          |
+-------------------------------------------------+

`)
}

func End() {
	fmt.Println("程序退出...")
}
