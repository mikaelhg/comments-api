package domain

type Comment struct {
	CommentId int64  `json:"id"`
	ThreadId  int64  `json:"thread_id"`
	Sender    string `json:"sender"`
	Text      string `json:"text"`
}
