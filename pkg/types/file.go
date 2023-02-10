package types

type File interface {
	Save(path string) error
	Load(path string) error
}

type GameFile interface {
}
