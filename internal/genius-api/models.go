package geniusapi

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type RandomLyric struct {
	Lyrics string
	Title  string
	Image  string
}

func NewRandomLyrics(lyrics, title, image string) RandomLyric {
	return RandomLyric{
		Lyrics: lyrics,
		Title:  title,
		Image:  image,
	}
}

type searchResponse struct {
	Meta struct {
		Status int `json:"status"`
	} `json:"meta"`
	Response struct {
		Hits []struct {
			Result struct {
				ArtistNames            string `json:"artist_names"`
				FullTitle              string `json:"full_title"`
				Title                  string `json:"title"`
				AlbumImageUrl          string `json:"song_art_image_url"`
				AlbumImageThumbnailUrl string `json:"song_art_image_thumbnail_art"`
				Url                    string `json:"url"`
			} `json:"result"`
		} `json:"hits"`
	}
}

func (s *searchResponse) getLyrics(index int) string {
	res, err := makeRequest(s.Response.Hits[index].Result.Url)
	if err != nil {
		return fmt.Sprint("Request err:", err)
	}

	document, err := goquery.NewDocumentFromReader(strings.NewReader(res))
	if err != nil {
		return fmt.Sprint("Document err:", err)
	}

	lyricsRoot := document.Find("#lyrics-root")

	lyrics := ""
	lyricsRoot.Find("[data-lyrics-container='true']").Each(func(i int, sel *goquery.Selection) {
		// Replace <br> with newline character
		sel.Find("br").Each(func(j int, br *goquery.Selection) {
			br.ReplaceWithHtml("\n")
		})
		// Append text to lyrics string
		lyrics += sel.Text() + "\n"
	})
	return lyrics
}
