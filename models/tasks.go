package models

type Task struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Status      bool   `json:"status"`
	Type        string `json:"type"`
	Priority    int    `json:"priority"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}

var Tasks []Task
