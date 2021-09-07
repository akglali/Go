package comment

type postCommentStruct struct {
	TextField string
	PostId    string
}
type commentStruct struct {
	CommentId          string
	PostId             string
	TextField          string
	Nickname           string
	Likes              uint
	Dislikes           uint
	CommentColor       string
	CommentDateCreated string
}
