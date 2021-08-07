package userDb

import (
	"vibraninlyGo/database"
)

func SignUpDb(username, password, token string) error {
	_, err := database.Db.Query("insert into  users(username, password, token) values($1,$2,$3)", username, password, token)
	if err != nil {
		return err
	}
	return err
}

func LoginDb(username string) (string, string, error) {
	var token, password string
	err := database.Db.QueryRow("select token,password from users where username=$1", username).Scan(&token, &password)
	return token, password, err
}
