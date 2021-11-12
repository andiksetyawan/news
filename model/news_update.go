package model

type NewsUpdate struct {
	Title     string   `json:"title" bson:"title"`
	Text      string   `json:"text" bson:"text"`
	Tags      []string `json:"tags" bson:"tags"`
	Status    string   `json:"status" bson:"status"`
	UpdatedAt int64    `json:"updated_at" bson:"updated_at"`
}
