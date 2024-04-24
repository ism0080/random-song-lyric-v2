package geniusapi

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	baseUrl string = "https://api.genius.com"
)

func GetRandomLyrics(artist string) RandomLyric {
	songsRes := searchSongs(url.QueryEscape(artist))
	songIndex := rand.Intn(len(songsRes.Response.Hits))
	lyrics := songsRes.getLyrics(songIndex)
	var arrLyrics []string
	for _, line := range strings.Split(lyrics, "\n") {
		if len(line) > 0 && line[0] != '[' {
			arrLyrics = append(arrLyrics, line)
		}
	}

	randLyric := arrLyrics[rand.Intn(len(arrLyrics))]
	title := songsRes.Response.Hits[songIndex].Result.FullTitle
	image := songsRes.Response.Hits[songIndex].Result.AlbumImageUrl

	return NewRandomLyrics(randLyric, title, image)
}

func searchSongs(artist string) searchResponse {
	res, err := makeRequest(fmt.Sprintf("%s/search?q=%s", baseUrl, artist))
	if err != nil {
		fmt.Println(err)
	}

	var data searchResponse
	err = json.Unmarshal([]byte(res), &data)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
	}

	fmt.Printf("%+v", data)
	return data
}

func makeRequest(url string) (string, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("GENIUS_API_KEY")))
	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
