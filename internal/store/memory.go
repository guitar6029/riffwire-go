package store

import "riffwire/internal/models"

type InMemoryStore struct {
	items []models.Item
}

func NewInMemoryStore(items []models.Item) *InMemoryStore {
	return &InMemoryStore{items: items}
}

func (s *InMemoryStore) ListItems() []models.Item {
	return s.items
}

func (s *InMemoryStore) GetItemByID(id int) (*models.Item, bool) {
	for i := range s.items {
		if s.items[i].ID == id {
			return &s.items[i], true
		}
	}
	return nil, false
}
