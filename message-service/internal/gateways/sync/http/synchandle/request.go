package synchandle

type addMessageRequest struct {
	Content string `json:"content" binding:"required"`
}

const (
	AddMessageRequest = iota
)

const (
	AddMessageRequestPartition = iota
)
