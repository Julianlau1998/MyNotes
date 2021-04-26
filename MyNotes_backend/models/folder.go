package models

type Folder struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	Title  string `json:"title"`
	Color  string `json:"color" `
	Type   string `json:"type" `
}
