package commentDb

import (
	"database/sql"
	"fmt"
	"vibraninlyGo/database"
)

func GetNicknameAndColor(postId, token string) string {
	var nickname, color string
	err := database.Db.QueryRow("select nickname,color from post_user_nickname_table where post_id=$1 and user_id=(select user_id from users where token=$2)", postId, token).Scan(&nickname, &color)
	if err != nil {
		fmt.Println("There is no nickname")
	}
	return nickname
}

// InsertNicknameTable if there is no comment belong to the user, nickname and color will be assigned to the user.
func InsertNicknameTable(postId, token, randomNickname, textField, currentTime, randomColor string) (error, *sql.Row) {
	var commentId string
	_, err := database.Db.Exec("insert into  post_user_nickname_table(post_id, user_id, nickname, color) values($1,(select user_id from users where token=$2),$3,$4)", postId, token, randomNickname, randomColor)
	if err != nil {
		return err, nil
	}
	err = database.Db.QueryRow("insert into comment_table( post_id, user_id, text_content, likes, dislikes,comment_date_created) values($1,(select user_id from users where token=$2),$3,$4,$4,$5)  returning comment_id ", postId, token, textField, 0, currentTime).Scan(&commentId)
	row := database.Db.QueryRow("select comment_id,comment_table.post_id,text_content, post_user_nickname_table.nickname, likes, dislikes, post_user_nickname_table.color,comment_date_created from comment_table left join post_user_nickname_table on comment_table.user_id= post_user_nickname_table.user_id where comment_id=$1", commentId)
	if err != nil {
		return err, nil
	}
	err = incrementCommentCount(postId)
	if err != nil {
		return err, nil
	}
	return err, row
}

//InsertComment if the user has already had a nickname into the post, the same nickname will be used and comment will be added to the comment_table
func InsertComment(postId, token, textField, currentTime string) (error, *sql.Row) {
	var commentId string
	err := database.Db.QueryRow("insert into comment_table (post_id, user_id, text_content, likes, dislikes,comment_date_created) values($1,(select user_id from users where token=$2),$3,$4,$4,$5) returning comment_id", postId, token, textField, 0, currentTime).Scan(&commentId)
	row := database.Db.QueryRow("select comment_id,comment_table.post_id,text_content, post_user_nickname_table.nickname, likes, dislikes, post_user_nickname_table.color,comment_date_created from comment_table left join post_user_nickname_table on comment_table.user_id = post_user_nickname_table.user_id where comment_table.comment_id=$1", commentId)
	if err != nil {
		return err, nil
	}
	err = incrementCommentCount(postId)
	if err != nil {
		return err, nil
	}
	return err, row
}

func incrementCommentCount(postId string) error {
	_, err := database.Db.Exec("update post_table set comment_count=comment_count+1 where post_id=$1", postId)
	return err
}

func GetAllCommentRows(postId string) (*sql.Rows, error) {
	rows, err := database.Db.Query("select comment_id,comment_table.post_id,text_content,post_user_nickname_table.nickname,likes,dislikes,post_user_nickname_table.color,comment_date_created from comment_table left join post_user_nickname_table on comment_table.post_id = post_user_nickname_table.post_id and comment_table.user_id=post_user_nickname_table.user_id where comment_table.post_id=$1 order by comment_date_created desc ", postId)
	return rows, err
}

func GetAllNicknames(postId string) (*sql.Rows, error) {
	rows, err := database.Db.Query("select nickname from post_user_nickname_table where post_id=$1", postId)
	return rows, err
}
