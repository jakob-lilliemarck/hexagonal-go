package driven

import (
	"errors"
	"hexagonal-go/src/lib/core"
	"hexagonal-go/src/lib/utils"
)

type FishRepository struct {
	Storage map[string]*core.Fish
}

type NotFound struct{}

func (s NotFound) Error() string {
	return "Not found"
}

func (s FishRepository) Read(id string) utils.Result[core.Fish, error] {
	fish, ok := s.Storage[id]
	if !ok {
		return utils.Err[core.Fish, error](errors.New("something went wrong"))
	}
	return utils.Ok[core.Fish, error](*fish)
}

func (s FishRepository) ReadCollection() utils.Result[[]core.Fish, error] {
	collection := make([]core.Fish, 0, len(s.Storage))

	for _, value := range s.Storage {
		collection = append(collection, *value)
	}

	return utils.Ok[[]core.Fish, error](collection)
}

func (s FishRepository) Save(fish core.Fish) utils.Result[core.Fish, error] {
	s.Storage[fish.ID] = &fish
	return utils.Ok[core.Fish, error](fish)
}

func (s FishRepository) Delete(id string) utils.Result[string, error] {
	delete(s.Storage, id)
	return utils.Ok[string, error](id)
}
