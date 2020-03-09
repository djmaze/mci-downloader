package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/kennygrant/sanitize"
)

type Tracks struct {
	XMLName xml.Name `xml:"contentdataset"`
	Tracks  []Track  `xml:"contentdata"`
}

type Track struct {
	XMLName xml.Name `xml:"contentdata"`
	Artist  string   `xml:"artist"`
	Album   string   `xml:"album"`
	Trackno uint16   `xml:"trackno"`
	Name    string   `xml:"name"`
	Url     string   `xml:"url"`
}

type Albums map[string][]Track

func (track Track) OutputDir() string {
	return sanitizeFilename(
		fmt.Sprintf("%s - %s", track.Artist, track.Album))
}

func (track Track) OutputFile() string {
	var file = sanitizeFilename(
		fmt.Sprintf("%02d - %s.mp3", track.Trackno, track.Name))
	return filepath.Join(track.OutputDir(), file)
}

func parseAlbums(data []byte) (int, Albums) {
	var tracks Tracks

	err := xml.Unmarshal(data, &tracks)
	if err != nil {
		log.Fatal(err)
	}

	albums := make(Albums)
	num_tracks := len(tracks.Tracks)
	for i := 0; i < num_tracks; i++ {
		track := tracks.Tracks[i]
		if album, ok := albums[track.OutputDir()]; ok {
			album = append(album, track)
			albums[track.OutputDir()] = album
		} else {
			albums[track.OutputDir()] = []Track{track}
		}
	}

	return num_tracks, albums
}

func sanitizeFilename(name string) string {
	replaceFunc := func(r rune) rune {
		switch r {
		case '/':
			return '-'
		case ',':
			return ' '
		case '?':
			return -1
		case ':':
			return '.'
		default:
			return r
		}
	}

	return strings.Map(replaceFunc, sanitize.Accents(name))
}
