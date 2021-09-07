package postDb

import (
	"database/sql"
	"vibraninlyGo/database"
)

func PostSinglePostDb(token, nickname, textField, currentTime, color string) (error, *sql.Row) {

	var postId string
	err := database.Db.QueryRow("insert into  post_table( user_id, nickname, text_field, comment_count, posted_date, likes, dislikes,color) values((select user_id from users where token=$1),$2,$3,$4,$5,$4,$4,$6) returning post_id", token, nickname, textField, 0, currentTime, color).Scan(&postId)
	row := database.Db.QueryRow("select post_id,nickname,text_field,comment_count,color,posted_date,likes,dislikes from post_table where post_id=$1", postId)
	_, err = database.Db.Exec("insert into post_user_nickname_table (post_id, user_id, nickname, color) values ($1,(select user_id from users where token=$2),$3,$4)", postId, token, nickname, color)
	if err != nil {
		return err, nil
	}
	return err, row
}

func GetSinglePostDb(postId string) *sql.Row {
	row := database.Db.QueryRow("select post_id,nickname,text_field,comment_count,posted_date,likes,dislikes,color from post_table where post_id=$1", postId)
	return row
}

func GetAllPostDb() (*sql.Rows, error) {
	rows, err := database.Db.Query("select post_id,nickname,text_field,comment_count,posted_date,likes,dislikes,color from post_table order by posted_date desc")
	return rows, err

}
