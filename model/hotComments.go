package model

import (
	"encoding/json"
	"fmt"
	. "music163-hotcomments-spider/util"
)

type HotComments struct {
	Content    string `gorm:"type:text;not null"`
	LikedCount int    `gorm:"type:int;not null"`
	Nickname   string `gorm:"type:varchar(50);not null"`
	ID         int    `gorm:"primary_key;type:int;AUTO_INCREMENT"`
}

func (h HotComments) String() string {
	return fmt.Sprintf("%s: %s (%d 赞)\n", h.Nickname, h.Content, h.LikedCount)
}

func Parse(body <-chan string, hotComments chan<- []HotComments) {
	for b := range body {
		var data map[string]interface{}
		if err := json.Unmarshal([]byte(b), &data); err != nil {
			Check(err)
		}
		m := data["hotComments"].([]interface{})

		var comments = make([]HotComments, len(m))

		for i, v := range m {
			o := v.(map[string]interface{})
			comments[i] = HotComments{
				Content:    o["content"].(string),
				LikedCount: ToInt(o["likedCount"].(float64)),
				Nickname:   o["user"].(map[string]interface{})["nickname"].(string),
			}
		}

		hotComments <- comments
	}

}
