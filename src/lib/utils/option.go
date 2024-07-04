package utils

type option[T any] struct {
	some *T
	none bool
}

type Option[T any] interface {
	IsSome() bool
	IsNone() bool
	Unwrap() T
}

func Some[T any](value T) option[T] {
	return option[T]{
		some: &value,
		none: false,
	}
}

func None[T any]() option[T] {
	return option[T]{
		some: nil,
		none: true,
	}
}

func (o option[T]) IsSome() bool {
	return o.some != nil
}

func (o option[T]) IsNone() bool {
	return !o.none
}

func (o option[T]) Unwrap() T {
	if o.some != nil {
		return *o.some
	}
	panic("called Unwrap on an None value")
}
