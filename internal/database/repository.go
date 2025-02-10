package database

import (
	"context"
	"fmt"
)

// LinkRepository - Интерфейс для работы с базой данных
type LinkRepository interface {
	SaveLink(ctx context.Context, originalURL, shortenedLink string) error
	GetOriginalLink(ctx context.Context, shortenedLink string) (string, error)
	// Другие методы для работы с данными
}

// InMemoryRepository - Реализация для in-memory
type InMemoryRepository struct {
	data map[string]string
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{make(map[string]string)}
}

func (r *InMemoryRepository) SaveLink(ctx context.Context, originalURL, shortenedLink string) error {
	r.data[shortenedLink] = originalURL
	return nil
}

func (r *InMemoryRepository) GetOriginalLink(ctx context.Context, shortenedLink string) (string, error) {
	url, ok := r.data[shortenedLink]
	if !ok {
		return "", fmt.Errorf("link not found")
	}
	return url, nil
}
