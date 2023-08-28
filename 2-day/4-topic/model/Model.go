package model

type Topic struct {
	Id         int64  `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CreateTime string `json:"create_time"`
}

type Post struct {
	Id         int64  `json:"id"`
	Content    string `json:"content"`
	CreateTime string `json:"create_time"`
}

type ListPosts struct {
	TopicId int64   `json:"topic_id"`
	Posts   []*Post `json:"posts"`
}
