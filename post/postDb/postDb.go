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

	err := database.Db.QueryRow("with pInfo as (insert into post_table (user_id, text_field, comment_count, posted_date, likes, dislikes) values ((select user_id from  users where  token=$1),$2,$3,$4,$3,$3)returning post_id,user_id) insert into post_user_nickname_table(post_id, user_id, nickname, color) values((SELECT pInfo.post_id from pInfo),(SELECT pInfo.user_id from pInfo),$6,$5)  returning post_id", token, textField, 0, currentTime, color, nickname).Scan(&postId)
	row := database.Db.QueryRow("select post_table.post_id,nickname,text_field,comment_count,color,posted_date,likes,dislikes from post_user_nickname_table  left join post_table  on post_table.post_id = post_user_nickname_table.post_id where post_table.post_id=$1", postId)
	if err != nil {
		return err, nil
	}
	return err, row
}

func GetSinglePostDb(postId string) *sql.Row {
	row := database.Db.QueryRow("select post_table.post_id,post_user_nickname_table.nickname,text_field,comment_count,posted_date,likes,dislikes,post_user_nickname_table.color from post_table left join post_user_nickname_table on post_table.post_id = post_user_nickname_table.post_id where post_table.post_id=$1 ", postId)
	return row
}

func GetAllPostDb() (*sql.Rows, error) {
	rows, err := database.Db.Query("select post_table.post_id,post_user_nickname_table.nickname,text_field,comment_count,posted_date,likes,dislikes,post_user_nickname_table.color from post_table left join post_user_nickname_table  on post_table.post_id = post_user_nickname_table.post_id and post_table.user_id=post_user_nickname_table.user_id order by post_table.posted_date desc")
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

func LetOnlyOwner(token, postId string) (error, bool) {
	var trueFalse bool
	err := database.Db.QueryRow("select exists(select 1 from post_table where user_id=(select user_id from users where token=$1) and post_id=$2)", token, postId).Scan(&trueFalse)
	return err, trueFalse
}

func UpdatePost(textField, postId, currentTime string) (sql.Result, error) {
	result, err := database.Db.Exec("update post_table set text_field=$1,posted_date=$3 where post_id=$2", textField, postId, currentTime)
	return result, err
}

func DeletePost(postId string) (sql.Result, error) {
	result, err := database.Db.Exec("WITH d as (delete from post_table where post_id=$1),cd as ( delete from post_user_nickname_table where post_id = $1) ,del as (delete from comment_table where  post_id=$1) delete from like_dislike_table where post_id=$1", postId)
	return result, err
}
