package yet_another_todo_app

import (
	"fmt"

	"github.com/satori/go.uuid"
)

type Item struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Details string `json:"details"`
	Done	bool `json:"done"`
}

func NewItem(title, details string, done bool) (*Item, error) {
	if len(title) == 0 {
		return nil, fmt.Errorf("title must be a non-empty string")
	}
	if len(details) == 0 {
		return nil, fmt.Errorf("details must be a non-empty string")
	}
	id := uuid.NewV4()
	return &Item{
		ID: id.String(),
		Title: title,
		Details: details,
		Done: done,
	}, nil
}

func (t *Item) GetTitle() string {
	return t.Title
}

func (t *Item) SetTitle(title string) error {
	if len(title) == 0 {
		return fmt.Errorf("title must be a non-empty string")
	}
	return nil
}

func (t *Item) IsDone() bool {
	return t.IsDone()
}

func (t *Item) SetDone() {
	t.Done = true
}

func (t *Item) UnsetDone() {
	t.Done = false
}

type ItemRepository interface {
	List() ([]*Item, error)
	//Delete(id int) error
	Save(t *Item) error
}