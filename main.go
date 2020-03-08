package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	port := flag.Uint("port", 8081, "port number")
	dry_run := flag.Bool("dry-run", false, "dry run – do not write files")
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
	fmt.Printf("Downloading track metadata from %s port %d..", ip, *port)
	xml := wadm.getTracks()
	fmt.Printf("\nParsing response..")
	num_tracks, albums := parseAlbums(xml)
	count := 0

	cwd, err := filepath.Abs(".")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nDownloading %d tracks\n", num_tracks)
	for dir, tracks := range albums {
		for i := 0; i < len(tracks); i++ {
			count += 1
			fmt.Printf("\n[%d/%d] %s", count, num_tracks, tracks[i].OutputFile())

			if _, err := os.Stat(tracks[i].OutputFile()); err == nil {
				fmt.Printf(" [already exists, skipping]")
			} else {
				if !*dry_run {
					err := os.Mkdir(dir, 0700)
					if err != nil && !os.IsExist(err) {
						log.Fatal(err)
					}

					err, data := wadm.downloadTrack(tracks[i].Url)
					if err != nil {
						log.Fatal(err)
					} else {
						output_file := filepath.Join(cwd, tracks[i].OutputFile())
						err = ioutil.WriteFile(output_file, data, 0400)
						if err != nil {
							log.Fatal(err)
						}
					}
				}
			}
		}
	}
	fmt.Println("\nDone!")
}
