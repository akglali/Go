package post

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
	"vibraninlyGo/database"
	"vibraninlyGo/helpers"
	"vibraninlyGo/post/postDb"
)

func getPostsOwner(c *gin.Context) {
	token := c.GetHeader("Token")
	var allPostId []string
	rows, err := postDb.GetPostsOwnerDb(token)
	if err != nil {
		helpers.MyAbort(c, "Could not reach the user's post")
	}
	for rows.Next() {
		var postId string
		if err := rows.Scan(&postId); err != nil {
			return
		}
		allPostId = append(allPostId, postId)
	}
	c.JSON(200, allPostId)
}

func getSinglePostOwner(c *gin.Context) {
	token := c.GetHeader("Token")
	postId := c.Param("postId")
	trueOrFalse, err := postDb.GetSinglePostOwnerDb(token, postId)
	if err != nil {
		helpers.MyAbort(c, "The post couldn't be found")
		return
	}
	c.JSON(200, trueOrFalse)

}

func editPost(c *gin.Context) {
	body := postStruct{}
	token := c.GetHeader("Token")
	postId := c.Param("postId")
	data, err := c.GetRawData()
	if err != nil {
		helpers.MyAbort(c, "Input format is wrong")
	}
	err = json.Unmarshal(data, &body)
	if err != nil {
		helpers.MyAbort(c, "Bad Input")
		return
	}
	var trueOrFalse bool
	currentTime := time.Now().Format("2006.01.02 15:04:05")
	err = database.Db.QueryRow("select exists(select 1 from post_table where user_id=(select user_id from users where token=$1) and post_id=$2)", token, postId).Scan(&trueOrFalse)
	if err != nil {
		helpers.MyAbort(c, "The post couldn't be found")
		return
	}
	if trueOrFalse {
		_, err := database.Db.Exec("update post_table set text_field=$1,posted_date=$3 where post_id=$2", body.TextField, postId, currentTime)
		if err != nil {
			helpers.MyAbort(c, "Check your input please")
			return
		}
	} else {
		helpers.MyAbort(c, "Are you the owner???")
		return
	}
	c.JSON(200, gin.H{"TextContent": body.TextField, "DateCreated": currentTime})

}

func deletePost(c *gin.Context) {
	token := c.GetHeader("Token")
	postId := c.Param("postId")
	var trueOrFalse bool
	err := database.Db.QueryRow("select exists(select 1 from post_table where user_id=(select user_id from users where token=$1) and post_id=$2)", token, postId).Scan(&trueOrFalse)
	if err != nil {
		helpers.MyAbort(c, "The post could not be found")
	}
	if trueOrFalse {
		_, err := database.Db.Exec("WITH d as (delete from post_table where post_id=$1),cd as ( delete from post_user_nickname_table where post_id = $1) delete from comment_table where  post_id=$1", postId)
		fmt.Println("NEW")
		if err != nil {
			helpers.MyAbort(c, "The post could not be deleted")
			panic(err)
			return
		}
	} else {
		helpers.MyAbort(c, "Are you the owner???")
		return
	}
	c.JSON(200, "Post with postId  "+postId+" is deleted successfully.")

}
