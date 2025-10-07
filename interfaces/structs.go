package interfaces

type Key byte

type Value interface {
	string | bool | int64 | float64
}
