package models

import geniusapi "github.com/ism0080/random-song-lyric-v2/internal/genius-api"

type Data struct {
	RandomLyric geniusapi.RandomLyric
}

func NewData() Data {
	return Data{
		RandomLyric: geniusapi.RandomLyric{},
	}
}

type Page struct {
	Data Data
}

func NewPage() Page {
	return Page{
		Data: NewData(),
	}
}
