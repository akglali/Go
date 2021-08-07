package main

import (
	"github.com/gin-gonic/gin"
	"vibraninlyGo/database"
	user "vibraninlyGo/user"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	database.SetupDatabase()

	users := r.Group("/user")
	user.SetupUser(users)

	err := r.Run(":8000")
	if err != nil {
		return
	}

}
