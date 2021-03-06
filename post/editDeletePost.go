package post

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"time"
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
	currentTime := time.Now().Format("2006.01.02 15:04:05")
	err, trueOrFalse := postDb.LetOnlyOwner(token, postId)
	if err != nil {
		helpers.MyAbort(c, "The post could not be found")
	}
	if trueOrFalse {
		_, err := postDb.UpdatePost(body.TextField, postId, currentTime)
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
	err, trueOrFalse := postDb.LetOnlyOwner(token, postId)
	if err != nil {
		helpers.MyAbort(c, "The post could not be found")
	}
	if trueOrFalse {
		_, err := postDb.DeletePost(postId)
		if err != nil {
			helpers.MyAbort(c, "The post could not be deleted")
			return
		}
	} else {
		helpers.MyAbort(c, "Are you the owner???")
		return
	}
	c.JSON(200, "Post with postId  "+postId+" is deleted successfully.")

}
