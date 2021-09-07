package commentDb

import (
	"database/sql"
	"fmt"
	"vibraninlyGo/database"
)

func GetNicknameAndColor(postId, token string) (string, string) {
	var nickname, color string
	err := database.Db.QueryRow("select nickname,color from post_user_nickname_table where post_id=$1 and user_id=(select user_id from users where token=$2)", postId, token).Scan(&nickname, &color)
	if err != nil {
		fmt.Println("There is no nickname")
	}
	return nickname, color
}

// InsertNicknameTable if there is no comment belong to the user, nickname and color will be assigned to the user.
func InsertNicknameTable(postId, token, randomNickname, textField, randomColor string) error {
	_, err := database.Db.Exec("insert into  post_user_nickname_table(post_id, user_id, nickname, color) values($1,(select user_id from users where token=$2),$3,$4)", postId, token, randomNickname, randomColor)
	if err != nil {
		return err
	}
	_, err = database.Db.Exec("insert into comment_table( post_id, user_id, text_content, nickname, likes, dislikes, comment_color) values($1,(select user_id from users where token=$2),$3,$4,$5,$5,$6)", postId, token, textField, randomNickname, 0, randomColor)
	if err != nil {
		return err
	}
	return err
}

//InsertComment if the user has already had a nickname into the post, the same nickname will be used and comment will be added to the comment_table
func InsertComment(postId, token, nickname, textField, color string) error {
	_, err := database.Db.Exec("insert into comment_table (post_id, user_id, text_content, nickname, likes, dislikes, comment_color) values($1,(select user_id from users where token=$2),$3,$4,$5,$5,$6)", postId, token, textField, nickname, 0, color)
	if err != nil {
		return err
	}
	return err

}

func GetAllCommentRows(postId string) (*sql.Rows, error) {
	rows, err := database.Db.Query("select comment_id,post_id,text_content,nickname,likes,dislikes,comment_color from comment_table where post_id=$1", postId)
	return rows, err
}

func GetAllNicknames(postId string) (*sql.Rows, error) {
	rows, err := database.Db.Query("select nickname from post_user_nickname_table where post_id=$1", postId)
	return rows, err
}
