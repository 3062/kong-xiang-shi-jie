package audio

func init() {

}

type Music []byte

var MusicConf = map[string]string{}
var MusicSet = map[string]Music{}

var CurrentMusic Music

func (c *Music) Play() {

}

func (c *Music) Pause() {

}

func (c *Music) Continue() {

}

func SetMusic(name string) (err error) {
	music, exist := MusicSet[name]
	if !exist {
		music, err = load(name)
		if err != nil {
			return err
		}
		MusicSet[name] = music
	}
	CurrentMusic = music
	CurrentMusic.Play()
	return nil
}

func load(path string) (Music, error) {
	return Music{}, nil
}

func Pause() {
	if CurrentMusic == nil {
		panic("CurrentMusic is nil")
	}
	CurrentMusic.Pause()
}

func Continue() {
	if CurrentMusic == nil {
		panic("CurrentMusic is nil")
	}
	CurrentMusic.Continue()
}
