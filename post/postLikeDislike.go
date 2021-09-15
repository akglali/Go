package post

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"vibraninlyGo/database"
	"vibraninlyGo/helpers"
)

func likeAndDislike(c *gin.Context) {
	body := likeDislike{}
	token := c.GetHeader("token")
	data, err := c.GetRawData()
	if err != nil {
		helpers.MyAbort(c, "Input format is wrong")
		return
	}
	err = json.Unmarshal(data, &body)
	if err != nil {
		fmt.Println(err)
		helpers.MyAbort(c, "Bad input")
		return
	}
	var userId string
	err = database.Db.QueryRow("select user_id from users where token=$1", token).Scan(&userId)
	if err != nil {
		helpers.MyAbort(c, "The User could not be found")
		return
	}
	if body.LikeOrDislike == "like" {
		var postId string
		err := database.Db.QueryRow("SELECT * FROM likes($1,$2)", body.PostId, userId).Scan(&postId)

		if err != nil {
			helpers.MyAbort(c, "Something went wrong while liking the post.")
			return
		}
		row := database.Db.QueryRow("select post_table.post_id,post_user_nickname_table.nickname,text_field,comment_count,posted_date,likes,dislikes,post_user_nickname_table.color from post_table left join post_user_nickname_table on post_table.post_id = post_user_nickname_table.post_id where post_table.post_id=$1 ", postId)
		var pst post
		err = row.Scan(&pst.PostId, &pst.VirtualName, &pst.TextContent, &pst.CommentCount, &pst.DateCreated, &pst.Likes, &pst.Dislikes, &pst.Color)
		if err != nil {
			c.JSON(400, helpers.ErrorStruct{
				Error: "We could not reach the post",
			})
			return
		}
		c.JSON(200, pst)

	} else if body.LikeOrDislike == "dislike" {
		var postId string
		err := database.Db.QueryRow("SELECT * FROM dislike($1,$2)", body.PostId, userId).Scan(&postId)
		if err != nil {
			helpers.MyAbort(c, "Something went wrong while disliking the post.")
			return
		}
		row := database.Db.QueryRow("select post_table.post_id,post_user_nickname_table.nickname,text_field,comment_count,posted_date,likes,dislikes,post_user_nickname_table.color from post_table left join post_user_nickname_table on post_table.post_id = post_user_nickname_table.post_id where post_table.post_id=$1 ", postId)
		var pst post
		err = row.Scan(&pst.PostId, &pst.VirtualName, &pst.TextContent, &pst.CommentCount, &pst.DateCreated, &pst.Likes, &pst.Dislikes, &pst.Color)
		if err != nil {
			c.JSON(400, helpers.ErrorStruct{
				Error: "We could not reach the post",
			})
			return
		}
		c.JSON(200, pst)
	} else {
		c.JSON(200, "There is no such an option!!")
	}

}
