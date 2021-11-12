package model

type NewsCreateRequest struct {
	Title string   `json:"title" binding:"required"`
	Text  string   `json:"text"`
	Tags  []string `json:"tags"`
}
