package core

import (
	"hexagonal-go/src/lib/utils"
)

type Fish struct {
	ID      string
	Species string
	Gender  string
	Age     uint32
}

func (s Fish) UpdateAge(Age uint32) Fish {
	s.Age = Age
	return s
}

/*
On implicit inteface implementations:

1. When the interface is refactored, it doesn't throw the error in the implementing function, as it doesn't know which function that would be. Instead, the error will occur in the caller, since the provided argument no longer implements the required interface.

2. When refactoring the implementing function, there is no error at all if the caller expects the new signature. The implementing function is then not implementing the interface any longer, but has drifted away from the definition without any errors. A struct can't enforce it's own implementations, the caller must enfore them.

--- --- ---
On error handling:
1. Errors is just another type and unhandled errors do not panic the code. Golang doesn't require the developer to handle errors, and won't tell you if they occured, that depends upon the developer.

--- --- ---
On zero-type initialization:
1. Go has something called zero-type, which is used as a fallback initialization for instance when accessing a map-key that holds no value, a zero-initialized struct is returned along with a boolean. So the struct is in fact initialized *even when it doesn't exist!*

--- --- ---
On maps, variables and move semantics:

Accessing a value in a map returns a copy of that value. Modifying values returned from a map will have unexpected effects while the overhead of accessing items will be larger than neccessary. Maps should really only be used for pointer storage, but that is pretty impractical.

Go lacks move-semantics, but is still passed by value. That means the values passed to a function is still accessible after passing them in the caller, modifying the value in the called function will not modify the value in the caller, since these are now two different instances of the passed type. For the beginner, this creates some pretty confusing behaviour where there are no compiler or runtime errors.
utils
	func (lhs: MyEnum) Test(rhs MyEnum) bool {
		return lhs == rhs
	}

The above example doesn't enforce type safety at compile time. Test() can be called with any two strings!

Since there are no real union-types, one can't model enums using structs either, which could otherwise be a way to provide type saftey.

*/

type FishDrivingPort interface {
	Read(id string) utils.Result[Fish, error]
	ReadCollection() utils.Result[[]Fish, error]
	Create(species string, age uint32) utils.Result[Fish, error]
	Update(id string, age uint32) utils.Result[Fish, error]
	Delete(id string) utils.Result[string, error]
}

// TODO - Rename ReadCollection to Load, and remove Read - there should only be a single function needed to load shit.
type FishDrivenPort interface {
	Read(id string) utils.Result[Fish, error]
	ReadCollection() utils.Result[[]Fish, error]
	Save(fish Fish) utils.Result[Fish, error]
	Delete(id string) utils.Result[string, error]
}
