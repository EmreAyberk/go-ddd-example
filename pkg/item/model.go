package item

type Item struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type CreateRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
