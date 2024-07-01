package driven

import (
	"hexagonal-go/src/lib/core"
)

type FishRepository struct {
	Storage map[string]*core.Fish
}

type NotFound struct{}

func (s NotFound) Error() string {
	return "Not found"
}

func (s FishRepository) Read(id string) (core.Fish, error) {
	fish, ok := s.Storage[id]
	if !ok {
		return *fish, NotFound{}
	}
	return *fish, nil
}

func (s FishRepository) ReadCollection() []core.Fish {
	collection := make([]core.Fish, 0, len(s.Storage))

	for _, value := range s.Storage {
		collection = append(collection, *value)
	}

	return collection
}

func (s FishRepository) Save(fish core.Fish) {
	s.Storage[fish.ID] = &fish
}

func (s FishRepository) Delete(id string) {
	delete(s.Storage, id)
}
