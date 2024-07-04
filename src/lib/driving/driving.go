package driving

import (
	"hexagonal-go/src/lib/core"
	"hexagonal-go/src/lib/utils"

	"github.com/oklog/ulid/v2"
)

type FishService struct {
	Repository core.Driven
}

func (adapter FishService) Read(id string) utils.Result[*core.Fish, error] {
	return adapter.Repository.Read(id)
}

func (s FishService) ReadCollection() utils.Result[[]*core.Fish, error] {
	return s.Repository.ReadCollection()
}

func (s FishService) Create(
	species string,
	age uint32,
) utils.Result[*core.Fish, error] {
	id := ulid.Make()
	fish := core.Fish{
		ID:      id.String(),
		Species: species,
		Age:     age,
	}
	return s.Repository.Save(&fish)
}

func (s FishService) Update(id string, age uint32) utils.Result[*core.Fish, error] {
	r := s.Repository.Read(id)

	// Map over result and update the age of the fish
	r = utils.Map(
		func(f *core.Fish) *core.Fish { return f.UpdateAge(age) },
		r,
	)

	// Map over result and save to repository
	return utils.AndThen(
		func(f *core.Fish) utils.Result[*core.Fish, error] { return s.Repository.Save(f) },
		r,
	)
}

func (s FishService) Delete(id string) utils.Result[string, error] {
	return s.Repository.Delete(id)
}
