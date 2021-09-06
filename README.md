# Go

## How to Download Go

You can visit at https://golang.org/dl/ to download go. Don't forget the add it to path.

## Go Gin

I am using go gin framework. To check it please visit https://github.com/gin-gonic/gin".

## Download Needed Modules

Run this command on terminal 
#go module tidy

## Database Tables (Postgres Sql)
### users Table
user_id uuid, username text,password text,token text 
### post_table 
post_id uuid, user_id uuid, nickname text ,text_field text, comment_count int, posted_date timestamp, likes int ,dislikes int,color text
### post_user_nickname_table
post_id uuid, user_id uuid, nickname text,color text