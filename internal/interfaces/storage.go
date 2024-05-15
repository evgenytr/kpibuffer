package interfaces

import "github.com/evgenytr/kpibuffer/internal/domain"

// Storage interface should be implemented by any used storage
type Storage interface {
	PushFact(*domain.Fact)
	PopFact() (*domain.Fact, error)
	Len() int
}
