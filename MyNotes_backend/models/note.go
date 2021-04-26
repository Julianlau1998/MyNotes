package models

type Note struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	FolderID string `json:"folder_id"`
	Title    string `json:"title"`
	Note     string `json:"body" `
}
