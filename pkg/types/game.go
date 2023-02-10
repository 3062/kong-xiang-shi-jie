package types

type GameIden int

type Game interface {
	Update() error
	Load() error
}
