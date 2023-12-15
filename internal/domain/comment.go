package domain

type Comment struct {
	Id     int64  `json:"id"`
	Sender string `json:"sender"`
	Text   string `json:"text"`
}
