package model

// TodoModel todo model
type TodoModel struct {
	ID          int    `json:"id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
