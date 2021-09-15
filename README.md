# Go

## How to Download Go

You can visit at https://golang.org/dl/ to download go. Don't forget the add it to path.

## Go Gin

I am using go gin framework. To check it please visit https://github.com/gin-gonic/gin".

## Download Needed Modules

#Run this command on terminal 

go module tidy

## Front-End With Angular

Visit https://github.com/akglali/Angular to see Front-End of the app.

## Database Tables (Postgres Sql)
### users Table (name type)
user_id uuid, username text,password text,token text 
### post_table (name type)
post_id uuid, user_id uuid,text_field text, comment_count int, posted_date timestamp, likes int ,dislikes int
### post_user_nickname_table (name type)
post_id uuid, user_id uuid, nickname text,color text
### comment_table (name type)
comment_id uuid,post_id uuid, user_id uuid, text_content text,likes int,dislikes int
### like_dislike_table(name type)
like_dislike_id int(AI),post_id uuid,user_id uuid,comment_id uuid, like_option integer (0 is liked 1 is disliked).
Please check dislikeAndLikeProcedure.txt file for the likes and dislikes procedures. 