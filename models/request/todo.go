package request

type CreateTodoRequestModel struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdateTodoRequestModel struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  bool   `json:"status"`
}
