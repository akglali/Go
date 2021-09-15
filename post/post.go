package post

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
	"vibraninlyGo/helpers"
	"vibraninlyGo/post/postDb"
)

func SetupPost(rg *gin.RouterGroup) {
	rg.POST("/newpost", postSinglePost)
	rg.GET("/getallpost", getAllPost)
	rg.GET("/getpost/:postId", getSinglePost)
	rg.GET("/getowner", getPostsOwner)
	rg.GET("/getowner/:postId", getSinglePostOwner)
	rg.PUT("/edit/:postId", editPost)
	rg.DELETE("/delete/:postId", deletePost)
	rg.POST("/like", likeAndDislike)
}

func postSinglePost(c *gin.Context) {
	body := postStruct{}
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
	nickname := RandomNickname()
	currentTime := time.Now().Format("2006.01.02 15:04:05")
	err, row := postDb.PostSinglePostDb(token, nickname, body.TextField, currentTime, RandomColor())

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
