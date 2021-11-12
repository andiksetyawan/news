package model

type NewsUpdateRequest struct {
	Title  string   `json:"title" binding:"required"`
	Text   string   `json:"text"`
	Tags   []string `json:"tags"`
	Status string   `json:"status" binding:"required"`
}
