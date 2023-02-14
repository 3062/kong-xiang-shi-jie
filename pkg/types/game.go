package types

type GameID int

type Game interface {
	Update() error
	Load() error
}
