package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	port := flag.Uint("port", 8081, "port number")
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("Usage: mci-downloader [--port <PORT>] <IP or hostname>")
		os.Exit(1)
	}
	ip := args[0]
	if ip == "" {
		fmt.Println("IP must be given")
		os.Exit(1)
	}

	wadm := WADM(ip, *port)
	xml := wadm.getTracks()
	num_tracks, albums := parseAlbums(xml)
	count := 0

	for dir, tracks := range albums {
		for i := 0; i < len(tracks); i++ {
			count += 1
			fmt.Printf("[%d/%d] %s\n", count, num_tracks, tracks[i].OutputFile())

			err := os.Mkdir(dir, 0700)
			if err != nil && !os.IsExist(err) {
				log.Fatal(err)
			}

			err, data := wadm.downloadTrack(tracks[i].Url)
			if err != nil {
				log.Fatal(err)
			} else {
				err = ioutil.WriteFile(tracks[i].OutputFile(), data, 0400)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}
