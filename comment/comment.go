package comment

import (
	"database/sql"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"vibraninlyGo/database/commentDb"
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
		return
	}
	token := c.GetHeader("token")
	nickname, color := commentDb.GetNicknameAndColor(body.PostId, token)

	if nickname == "" {
		randomNickname := post.RandomNickname()
		randomColor := post.RandomColor()
		nicknames, _ := getAllNicknames(body.PostId)
		for checkNickname(&randomNickname, nicknames) {
		}
		err := commentDb.InsertNicknameTable(body.PostId, token, randomNickname, body.TextField, randomColor)
		if err != nil {
			c.JSON(400, helpers.ErrorStruct{
				Error: "Something went wrong with nickname",
			})
			return
		}
	} else {
		err := commentDb.InsertComment(body.PostId, token, nickname, body.TextField, color)
		if err != nil {
			c.JSON(400, helpers.ErrorStruct{
				Error: "Comment couldn't be added",
			})
			return
		}
	}
	c.JSON(200, "comment is added")
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
		if err := rows.Scan(&comment.CommentId, &comment.PostId, &comment.TextField, &comment.Nickname, &comment.Likes, &comment.Dislikes, &comment.CommentColor); err != nil {
			return comments, err
		}
		comments = append(comments, comment)
	}
	return comments, err
}