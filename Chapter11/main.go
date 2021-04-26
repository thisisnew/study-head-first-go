package main

import (
	gadget "study-head-first-go/Chapter11/gadget"
)

func playList(device gadget.TapePlayer, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}

	device.Stop()
}

func main() {
	player := gadget.TapePlayer{}
	mixtape := []string{"Jessie's Girl", "Whip it", "9 to 5"}
	playList(player, mixtape)
}
