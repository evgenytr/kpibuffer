package memstorage

import (
	"container/list"
	"fmt"
	"sync"

	"github.com/evgenytr/kpibuffer/internal/domain"
)

type memStorage struct {
	storage *list.List
	mutex   *sync.Mutex
}

func NewStorage() *memStorage {
	return &memStorage{
		storage: list.New(),
		mutex:   &sync.Mutex{},
	}
}

func (ms memStorage) PushFact(fact *domain.Fact) {
	ms.mutex.Lock()
	ms.storage.PushBack(fact)
	ms.mutex.Unlock()
}

func (ms memStorage) PopFact() (*domain.Fact, error) {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()

	if ms.storage.Len() > 0 {
		elem := ms.storage.Front()
		fact := ms.storage.Remove(elem).(*domain.Fact)
		return fact, nil
	}

	return nil, fmt.Errorf("no elements in storage")
}

func (ms memStorage) Len() int {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()

	return ms.storage.Len()
}
