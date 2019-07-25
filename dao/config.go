package dao

import "fmt"

// mysql 连接用户与密码
const (
	USER     = "root"
	PASSWORD = "xxxxx"
)

func getDBMaster() string {
	return fmt.Sprintf("%s:%s@/cloudmusic?charset=utf8mb4&parseTime=True&loc=Local", USER, PASSWORD)
}
