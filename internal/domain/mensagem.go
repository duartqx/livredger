package domain

type Event string

type Message string

type Command interface {
	Validate() error
}
