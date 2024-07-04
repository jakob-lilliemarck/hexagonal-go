package utils

type result[T any, E error] struct {
	ok  *T
	err *E
}

type Result[T any, E error] interface {
	IsOk() bool
	IsErr() bool
	Unwrap() T
	UnwrapErr() E
}

func Ok[T any, E error](value T) result[T, E] {
	return result[T, E]{
		ok:  &value,
		err: nil,
	}
}

func Err[T any, E error](err E) result[T, E] {
	return result[T, E]{
		ok:  nil,
		err: &err,
	}
}

func (r result[T, E]) IsOk() bool {
	return r.ok != nil
}

func (r result[T, E]) IsErr() bool {
	return r.err != nil
}

func (r result[T, E]) Unwrap() T {
	if r.ok != nil {
		return *r.ok
	}
	panic("called Unwrap on an Err value")
}

func (r result[T, E]) UnwrapErr() E {
	if r.err != nil {
		return *r.err
	}
	panic("called UnwrapErr on an Ok value")
}

func Map[T any, U any, E error](
	callback func(T) U,
	result Result[T, E],
) Result[U, E] {
	if result.IsOk() {
		value := result.Unwrap()
		mapped := callback(value)
		return Ok[U, E](mapped)
	}
	return Err[U](result.UnwrapErr())
}

func MapError[T any, E1 error, E2 error](
	callback func(E1) E2,
	result Result[T, E1],
) Result[T, E2] {
	if result.IsErr() {
		err := result.UnwrapErr()
		mapped := callback(err)
		return Err[T, E2](mapped)
	}
	return Ok[T, E2](result.Unwrap())
}

func AndThen[T any, U any, E error](
	callback func(T) Result[U, E],
	result Result[T, E],
) Result[U, E] {
	if result.IsOk() {
		value := result.Unwrap()
		mapped := callback(value)
		return mapped
	}
	return Err[U](result.UnwrapErr())
}
