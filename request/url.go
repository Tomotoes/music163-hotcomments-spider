package request

import (
	"fmt"
)

func generatorUrl(id int) string {
	return fmt.Sprintf("http://music.163.com/api/v1/resource/comments/R_SO_4_%d?limit=0&offset=0", id)
}
