package post

import "time"

type postStruct struct {
	TextField string
}
type post struct {
	PostId       string
	UserId       string
	VirtualName  string
	TextContent  string
	CommentCount *uint
	DateCreated  time.Time
	Likes        uint
	Dislikes     uint
	Color        string
}
