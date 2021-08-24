package main

import (
	"github.com/gin-gonic/gin"
	"vibraninlyGo/database"
	"vibraninlyGo/post"
	user "vibraninlyGo/user"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	database.SetupDatabase()

	r.Use(func(context *gin.Context) {
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "*")
		if context.Request.Method == "OPTIONS" {
			context.Status(200)
			context.Abort()
		}
	})
	users := r.Group("/user")
	user.SetupUser(users)

	posts := r.Group("/post")
	post.SetupPost(posts)
	err := r.Run(":8000")
	if err != nil {
		return
	}

}
