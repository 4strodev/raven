package option

type OptionTypes string

const (
	some OptionTypes = "SOME"
	none OptionTypes = "NONE"
)

type Option[T any] struct {
	value      T
	optionType OptionTypes
}

func (o Option[T]) IsSome() bool {
	return o.optionType == some
}

func (o Option[T]) IsNone() bool {
	return o.optionType == none
}

func (o Option[T]) Unwrap() T {
	if o.IsNone() {
		panic("Trying to get value of an empty option")
	}

	return o.value
}

func Some[T any](value T) Option[T] {
	return Option[T]{
		value:      value,
		optionType: some,
	}
}

func None[T any]() Option[T] {
	return Option[T]{
		optionType: none,
	}
}

