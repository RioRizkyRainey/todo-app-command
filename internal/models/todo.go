package models

type Todo struct {
	ID        string `json:"_id"`
	Rev       string `json:"_rev,omitempty"`
	Text      string `json:"text"`
	IsDone    bool   `json:"is_done"`
	CreatedAt string `json:"createdAt"`
}
