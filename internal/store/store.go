package store

import "riffwire/internal/models"

type Store interface {
	ListItems() []models.Item
	GetItemByID(id int) (*models.Item, bool)
}
