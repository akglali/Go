package post

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
	"vibraninlyGo/database"
	"vibraninlyGo/helpers"
)

func SetupPost(rg *gin.RouterGroup) {
	rg.POST("/newpost", postSinglePost)
	rg.GET("/getallpost", getAllPost)
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
	_, err = database.Db.Query("insert into  post_table( user_id, nickname, text_field, comment_count, posted_date, likes, dislikes,color) values((select user_id from users where token=$1),$2,$3,$4,$5,$4,$4,$6)", token, nickname, body.TextField, 0, currentTime, randomColor())

	if err != nil {
		c.JSON(400, helpers.ErrorStruct{
			Error: "Something went wrong",
		})
		return
	}

	c.String(200, "Post Is created")

}

func getAllPost(c *gin.Context) {
	allRows, _ := getAllRows()
	c.JSON(200, allRows)

}

func getAllRows() ([]post, error) {
	rows, err := database.Db.Query("select * from post_table order by posted_date desc")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Rows couldn't be read")
		}
	}(rows)

	var posts []post

	for rows.Next() {
		var pst post
		if err := rows.Scan(&pst.PostId, &pst.UserId, &pst.VirtualName, &pst.TextContent, &pst.CommentCount, &pst.DateCreated, &pst.Likes, &pst.Dislikes, &pst.Color); err != nil {
			return posts, err
		}
		posts = append(posts, pst)
	}
	return posts, err
}
