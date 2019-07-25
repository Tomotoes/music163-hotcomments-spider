package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	. "music163-hotcomments-spider/model"
	. "music163-hotcomments-spider/util"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	db, err = gorm.Open("mysql", getDBMaster())
	Check(err)
}

func Save(hotComments <-chan []HotComments) {
	for comments := range hotComments {
		for _, comment := range comments {
			if err := db.Create(&comment).Error; err != nil {
				Check(err)
			}
		}
	}
	db.Close()
}
