package driving

import (
	"hexagonal-go/src/lib/core"

	"github.com/oklog/ulid/v2"
)

type FishService struct {
	Repository core.FishDrivenPort
}

func (adapter FishService) Read(id string) (core.Fish, error) {
	return adapter.Repository.Read(id)
}

func (s FishService) ReadCollection() []core.Fish {
	return s.Repository.ReadCollection()
}

func (s FishService) Create(
	species string,
	age uint32,
) core.Fish {
	id := ulid.Make()
	fish := core.Fish{
		ID:      id.String(),
		Species: species,
		Age:     age,
	}

	s.Repository.Save(fish)
	return fish
}

func (s FishService) Update(id string, age uint32) (core.Fish, error) {
	fish, err := s.Repository.Read(id)

	if err != nil {
		return fish, err
	}

	fish.UpdateAge(age)

	s.Repository.Save(fish)

	return fish, nil
}

func (s FishService) Delete(id string) {
	s.Repository.Delete(id)
}
