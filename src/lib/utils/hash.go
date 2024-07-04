package utils

type IHash[K comparable, V any] interface {
	Get(key K) Option[V]
	Insert(key K, value V) Option[V]
	Remove(key K) Option[V]
	Len() int
	Keys() []K
	Values() []V
}

type hash[K comparable, V any] struct {
	storage map[K]V
}

func Hash[K comparable, V any]() IHash[K, V] {
	storage := make(map[K]V)

	return hash[K, V]{storage}
}

func (h hash[K, V]) Len() int {
	return len(h.storage)
}

func (h hash[K, V]) Keys() []K {
	buf := make([]K, 0, len(h.storage))
	for key, _ := range h.storage {
		buf = append(buf, key)
	}
	return buf
}

func (h hash[K, V]) Values() []V {
	buf := make([]V, 0, len(h.storage))
	for _, value := range h.storage {
		buf = append(buf, value)
	}
	return buf
}

func (h hash[K, V]) Get(key K) Option[V] {
	value, ok := h.storage[key]
	if ok {
		// Wrap & return found value
		return Some(value)
	}
	// Otherwise return None
	return None[V]()
}

func (h hash[K, V]) Insert(key K, value V) Option[V] {
	prev, ok := h.storage[key]
	h.storage[key] = value
	if ok {
		// Wrap & return any previous value
		return Some(prev)
	} else {
		// Otherwise return none
		return None[V]()
	}
}

func (h hash[K, V]) Remove(key K) Option[V] {
	prev, ok := h.storage[key]
	delete(h.storage, key)
	if ok {
		return Some(prev)
	} else {
		return None[V]()
	}
}

func OkOr[T any, E error](o Option[T], e E) Result[T, E] {
	if o.IsSome() {
		value := o.Unwrap()
		return Ok[T, E](value)
	} else {
		return Err[T, E](e)
	}
}
