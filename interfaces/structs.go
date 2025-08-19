package interfaces

type Key string

type Value interface {
	string | bool | int64 | float64
}
