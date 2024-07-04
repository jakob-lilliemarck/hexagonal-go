package driven

import (
	"errors"
	"hexagonal-go/src/lib/core"
	"hexagonal-go/src/lib/utils"
)

type drivenAdapter struct {
	hash utils.IHash[string, *core.Fish]
}

func DrivenAdapter(hash utils.IHash[string, *core.Fish]) core.Driven {
	return drivenAdapter{hash}
}

func (adapter drivenAdapter) Read(id string) utils.Result[*core.Fish, error] {
	option := adapter.hash.Get(id)
	result := utils.OkOr(option, errors.New("blabla"))
	return result
}

func (adapter drivenAdapter) ReadCollection() utils.Result[[]*core.Fish, error] {
	values := adapter.hash.Values()
	return utils.Ok[[]*core.Fish, error](values)
}

func (adapter drivenAdapter) Save(fish *core.Fish) utils.Result[*core.Fish, error] {
	adapter.hash.Insert(fish.ID, fish)
	return utils.Ok[*core.Fish, error](fish)
}

func (adapter drivenAdapter) Delete(id string) utils.Result[string, error] {
	adapter.hash.Remove(id)
	return utils.Ok[string, error](id)
}
