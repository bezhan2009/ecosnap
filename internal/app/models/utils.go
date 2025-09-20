package models

type FileRequest struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"name,omitempty"`
}
