package test

type APIServiceX interface {
	SendMessage(message string) error
	IncrementCounterBy(incr int) error
}
