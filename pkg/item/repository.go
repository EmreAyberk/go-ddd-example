package item

import (
	"errors"
	"sync"
)

type IRepository interface {
	Create(item Item) error
	GetAll() ([]Item, error)
	GetOne(key string) (Item, error)
	DeleteAll() error
	DeleteOne(key string) error
}

type Repository struct {
	Data map[string]string
}

func NewRepository() Repository {
	return Repository{Data: map[string]string{}}
}

var mu sync.Mutex

func init() {
	mu = sync.Mutex{}
}

// Create returns error because if any db action happens, it may cause error
func (r *Repository) Create(item Item) error {
	mu.Lock()
	defer mu.Unlock()
	r.Data[item.Key] = item.Value

	return nil
}

// GetAll returns error because if any db action happens, it may cause error
func (r *Repository) GetAll() (map[string]string, error) {
	mu.Lock()
	defer mu.Unlock()
	return r.Data, nil
}

// GetOne returns error because if any db action happens, it may cause error
func (r *Repository) GetOne(key string) (Item, error) {
	mu.Lock()
	defer mu.Unlock()
	if _, exist := r.Data[key]; exist {
		return Item{
			Key:   key,
			Value: r.Data[key],
		}, nil
	} else {
		return Item{}, errors.New("cannot find the selected key")
	}
}

// DeleteAll returns error because if any db action happens, it may cause error
func (r *Repository) DeleteAll() error {
	mu.Lock()
	defer mu.Unlock()
	r.Data = map[string]string{}
	return nil
}

// DeleteOne returns error because if any db action happens, it may cause error
func (r *Repository) DeleteOne(key string) error {
	mu.Lock()
	defer mu.Unlock()
	if _, exist := r.Data[key]; exist {
		delete(r.Data, key)
		return nil
	} else {
		return errors.New("cannot find the selected key")
	}
}
