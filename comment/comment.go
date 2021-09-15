package comment

import (
	"database/sql"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"time"
	"vibraninlyGo/comment/commentDb"
	"vibraninlyGo/helpers"
	"vibraninlyGo/post"
)

func SetupComment(rg *gin.RouterGroup) {
	rg.POST("/post", postComment)
	rg.GET("/get_comment/:postId", getAllComment)
}

// post a new comment.
func postComment(c *gin.Context) {
	body := postCommentStruct{}
	data, err := c.GetRawData()
	if err != nil {
		helpers.MyAbort(c, "Input format is wrong")
		return
	}
	err = json.Unmarshal(data, &body)
	if err != nil {
		helpers.MyAbort(c, "Bad input")
		return
	}
	token := c.GetHeader("token")
	nickname := commentDb.GetNicknameAndColor(body.PostId, token)
	currentTime := time.Now().Format("2006.01.02 15:04:05")
	var row *sql.Row
	if nickname == "" {
		var randomNickname string
		randomColor := post.RandomColor()
		nicknames, err := getAllNicknames(body.PostId)
		if err != nil {
			c.AbortWithStatusJSON(400, helpers.ErrorStruct{Error: "wops"})
		}
		for checkNickname(&randomNickname, nicknames) {
		}
		err, row = commentDb.InsertNicknameTable(body.PostId, token, randomNickname, body.TextField, currentTime, randomColor)
		if err != nil {
			c.AbortWithStatusJSON(400, helpers.ErrorStruct{
				Error: "Something went wrong with nickname",
			})
			return
		}
	} else {
		err, row = commentDb.InsertComment(body.PostId, token, body.TextField, currentTime)
		if err != nil {
			c.JSON(400, helpers.ErrorStruct{
				Error: "Comment couldn't be added",
			})
			return
		}
	}
	var comment commentStruct
	err = row.Scan(&comment.CommentId, &comment.PostId, &comment.TextField, &comment.Nickname, &comment.Likes, &comment.Dislikes, &comment.CommentColor, &comment.CommentDateCreated)
	if err != nil {
		c.JSON(400, helpers.ErrorStruct{
			Error: "We could not reach the comment!",
		})
		return
	}

	c.JSON(200, comment)
}

// getting all nicknames since every user has unique nickname for single post. So need to check to prevent duplicate nickname
func getAllNicknames(postId string) ([]string, error) {
	rows, err := commentDb.GetAllNicknames(postId)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var nicknames []string

	for rows.Next() {
		var nickname string
		if err := rows.Scan(&nickname); err != nil {
			return nicknames, err
		}
		nicknames = append(nicknames, nickname)
	}
	return nicknames, err

}

// it checks the nickname whether it is unique or not
func checkNickname(a *string, list []string) bool {
	virtualName := post.RandomNickname()
	*a = virtualName
	for _, b := range list {
		if b == *a {
			return true
		}
	}
	return false
}

//sending all comments.
func getAllComment(c *gin.Context) {
	postId := c.Param("postId")
	allCommentRows, err := getAllCommentRows(postId)
	if err != nil {
		c.JSON(400, helpers.ErrorStruct{
			Error: "The post's comment couldn't be found",
		})
		return
	}
	c.JSON(200, allCommentRows)
}

//getting all comment belongs to the post.
func getAllCommentRows(postId string) ([]commentStruct, error) {

	rows, err := commentDb.GetAllCommentRows(postId)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
		}
	}(rows)

	var comments []commentStruct

	for rows.Next() {
		var comment commentStruct
		if err := rows.Scan(&comment.CommentId, &comment.PostId, &comment.TextField, &comment.Nickname, &comment.Likes, &comment.Dislikes, &comment.CommentColor, &comment.CommentDateCreated); err != nil {
			return comments, err
		}
		comments = append(comments, comment)
	}
	return comments, err
}
