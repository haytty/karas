package event

type Event interface {
	Act() error
}
