package server

import (
	"strings"

	"github.com/nosborn/federation-1999/internal/model"
)

type Storage struct {
	Warehouse [model.MAX_STORES]*model.Warehouse
}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) FindWarehouse(planet string) *model.Warehouse {
	for i := range s.Warehouse {
		if s.Warehouse[i] == nil {
			continue
		}
		if !strings.EqualFold(s.Warehouse[i].Planet, planet) {
			continue
		}
		return s.Warehouse[i]
	}
	return nil
}
