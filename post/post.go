package post

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"time"
	"vibraninlyGo/database"
	"vibraninlyGo/helpers"
)

func SetupPost(rg *gin.RouterGroup) {
	rg.POST("/newpost", postSinglePost)
}

func postSinglePost(c *gin.Context) {
	body := postStruct{}
	data, err := c.GetRawData()
	if err != nil {
		c.JSON(400, helpers.ErrorStruct{
			Error: "Input format is wrong",
		})
		return
	}
	err = json.Unmarshal(data, &body)
	if err != nil {
		c.JSON(400, helpers.ErrorStruct{
			Error: "Bad input",
		})
	}
	token := c.GetHeader("token")
	nickname := randomNickname()
	currentTime := time.Now().Format("2006.01.02 15:04:05")
	_, err = database.Db.Query("insert into  post_table( user_id, nickname, text_field, comment_count, comment_date, likes, dislikes) values((select user_id from users where token=$1),$2,$3,$4,$5,$4,$4)", token, nickname, body.TextField, 0, currentTime)

	if err != nil {
		c.JSON(400, helpers.ErrorStruct{
			Error: "Something went wrong",
		})
		return
	}

	c.String(200, "Post Is created")

}
