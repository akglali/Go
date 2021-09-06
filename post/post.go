package post

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
	"vibraninlyGo/database/postDb"
	"vibraninlyGo/helpers"
)

func SetupPost(rg *gin.RouterGroup) {
	rg.POST("/newpost", postSinglePost)
	rg.GET("/getallpost", getAllPost)
	rg.GET("/getpost/:postId", getSinglePost)
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
		return
	}

	token := c.GetHeader("token")
	nickname := randomNickname()
	currentTime := time.Now().Format("2006.01.02 15:04:05")
	err, row := postDb.PostSinglePostDb(token, nickname, body.TextField, currentTime, randomColor())

	if err != nil {
		c.JSON(400, helpers.ErrorStruct{
			Error: "Something went wrong please try again later!",
		})
		return
	}
	var pst post
	err = row.Scan(&pst.PostId, &pst.VirtualName, &pst.TextContent, &pst.CommentCount, &pst.Color, &pst.DateCreated, &pst.Likes, &pst.Dislikes)
	if err != nil {
		c.JSON(400, helpers.ErrorStruct{
			Error: "We could not reach the post!",
		})
		return
	}
	c.JSON(200, pst)

}

func getAllPost(c *gin.Context) {
	allRows, _ := getAllRows()
	c.JSON(200, allRows)
}

func getSinglePost(c *gin.Context) {
	postId := c.Param("postId")
	row := postDb.GetSinglePostDb(postId)
	var pst post
	err := row.Scan(&pst.PostId, &pst.VirtualName, &pst.TextContent, &pst.CommentCount, &pst.DateCreated, &pst.Likes, &pst.Dislikes, &pst.Color)
	if err != nil {
		c.JSON(400, helpers.ErrorStruct{
			Error: "We could not reach the post",
		})
		return
	}
	c.JSON(200, pst)

}

func getAllRows() ([]post, error) {
	rows, err := postDb.GetAllPostDb()
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Rows couldn't be read")
		}
		return
	}(rows)

	var posts []post
	for rows.Next() {
		var pst post
		if err := rows.Scan(&pst.PostId, &pst.VirtualName, &pst.TextContent, &pst.CommentCount, &pst.DateCreated, &pst.Likes, &pst.Dislikes, &pst.Color); err != nil {
			return posts, err
		}
		posts = append(posts, pst)
	}
	return posts, err
}
