package postDb

import (
	"database/sql"
	"vibraninlyGo/database"
)

func PostSinglePostDb(token, nickname, textField, currentTime, color string) (error, *sql.Row) {

	var postId string
	//with b as (insert into post_table (user_id, nickname, text_field, comment_count, posted_date, likes, dislikes, color)
	//values ('56b81bf5-82bd-49ee-af95-6f00ac55b0f3', 'whatever', 'text', 0, '2021-09-07 16:24:45.000000',
	//	0, 0, 'green')
	//returning post_id,user_id, nickname, color)
	//INSERT
	//INTO post_user_nickname_table (post_id, user_id, nickname, color) SELECT post_id, user_id, nickname, color FROM b

	//insert into  post_table( user_id, nickname, text_field, comment_count, posted_date, likes, dislikes,color) values((select user_id from users where token=$1),$2,$3,$4,$5,$4,$4,$6) returning post_id", token, nickname, textField, 0, currentTime, color).Scan(&postId)
	err := database.Db.QueryRow("with pInfo as (insert into post_table (user_id, nickname, text_field, comment_count, posted_date, likes, dislikes, color) values ((select user_id from  users where  token=$1),$2,$3,$4,$5,$4,$4,$6)returning post_id,user_id,nickname,color) insert into post_user_nickname_table(post_id, user_id, nickname, color) SELECT post_id,user_id,nickname,color from pInfo returning post_id", token, nickname, textField, 0, currentTime, color).Scan(&postId)
	row := database.Db.QueryRow("select post_id,nickname,text_field,comment_count,color,posted_date,likes,dislikes from post_table where post_id=$1", postId)
	if err != nil {
		panic(err)
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

func GetPostsOwnerDb(token string) (*sql.Rows, error) {
	rows, err := database.Db.Query("select post_id from post_table where user_id=(select user_id from users where token=$1)", token)
	return rows, err
}

func GetSinglePostOwnerDb(token, postId string) (bool, error) {
	var trueOrFalse bool
	err := database.Db.QueryRow("select exists(select 1 from post_table where user_id=(select user_id from users where token=$1) and post_id=$2)", token, postId).Scan(&trueOrFalse)
	return trueOrFalse, err

}
