package models

type List struct {
	ID        string   `json:"id"`
	UserID    string   `json:"user_id"`
	FolderID  string   `json:"folder_id"`
	Title     string   `json:"title"`
	List      []string `json:"list" `
	DoneItems []string `json:"doneItems" `
}
