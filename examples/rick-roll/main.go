package main

import (
	"fmt"
	"log"

	"github.com/fhs/gompd/mpd"
)

func main() {
	conn, err := mpd.Dial("tcp", "localhost:6600")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	status, err := conn.Status()
	if err != nil {
		log.Fatal(err)
	}
	if status["state"] == "play" {
		//conn.Stop()
	} else {

		results, err := conn.Search("artist", "\"Rick Astley\"", "title", "\"Never Gonna Give You Up\"")
		if err != nil {
			log.Fatal(err)
		}
		if err = conn.Clear(); err != nil {
			log.Fatal(err)
		}
		if err = conn.Add(results[0]["file"]); err != nil {
			log.Fatal(err)
		}
		if err = conn.Play(-1); err != nil {
			log.Fatal(err)
		}
	}

	song, err := conn.CurrentSong()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(song)
	state, err := conn.Status()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(state)

	attrs, err := conn.PlaylistInfo(-1, -1)
	fmt.Println(attrs)

	stats, err := conn.Stats()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(stats)

	fmt.Printf("Now playing %s by %s\n", song["Title"], song["Artist"])
}
