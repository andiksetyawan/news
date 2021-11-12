package entity

type Tag struct {
	Slug      string `json:"slug" bson:"slug"`
	Name      string `json:"name" bson:"name"`
	CreatedAt int64  `json:"created_at" bson:"created_at"`
	UpdatedAt int64  `json:"updated_at" bson:"updated_at"`
}