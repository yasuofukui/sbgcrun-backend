package model

// News ... entity for news item
type News struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Category  string `json:"category"`
	Author    string `json:"author"`
	CreatedAt string `json:"created_at"`
}
