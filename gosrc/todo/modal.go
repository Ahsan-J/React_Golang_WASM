package todo

// Item Item modal
type Item struct {
	Task string `json:"task"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updateAt"`
	Status int `json:"status"`
}