package user

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"vibraninlyGo/helpers"
	"vibraninlyGo/user/userDb"
)

func SetupUser(rg *gin.RouterGroup) {

	rg.POST("/signup", signup)
	rg.POST("/login", login)
}

func signup(c *gin.Context) {
	body := userStruct{}
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
	if len(body.Password) < 6 || len(body.Username) < 6 {
		c.JSON(400, helpers.ErrorStruct{
			Error: "Password and username must be more than 6 character",
		})
		return
	}
	password, _ := hashpassword(body.Password)
	token := tokenGenerator()
	err = userDb.SignUpDb(body.Username, password, token)
	if err != nil {
		c.JSON(400, helpers.ErrorStruct{
			Error: "User is already exist",
		})
		return
	}

	c.JSON(200, "user is successfully created")
}

func login(c *gin.Context) {
	body := userStruct{}
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
			Error: "Bad Input",
		})
		return
	}

	token, password, errLogin := userDb.LoginDb(body.Username)
	if errLogin != nil {
		c.JSON(400, helpers.ErrorStruct{
			Error: "User is not found",
		})
		return
	}
	if checkpassword(body.Password, password) {
		c.JSON(200, token)
	} else {
		c.JSON(400, helpers.ErrorStruct{
			Error: "Check your password",
		})
		return
	}
}
