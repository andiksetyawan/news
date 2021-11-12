package entity

type News struct {
	ID        string   `json:"id" bson:"id"`
	Title     string   `json:"title" bson:"title"`
	Slug      string   `json:"slug" bson:"slug"`
	Text      string   `json:"text" bson:"text"`
	Tags      []string `json:"tags" bson:"tags"`
	Status    string   `json:"status" bson:"status"`
	CreatedAt int64    `json:"created_at" bson:"created_at"`
	UpdatedAt int64    `json:"updated_at" bson:"updated_at"`
	DeletedAt int64    `json:"deleted_at" bson:"deleted_at"`
}
