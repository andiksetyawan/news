package model

type TagCreateRequest struct {
	Name string `json:"name" binding:"required"`
}
