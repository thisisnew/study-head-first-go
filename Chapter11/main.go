package main

import "study-head-first-go/Chapter11/gadget"

type Player interface {
	Play(s string)
	Stop()
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}

	device.Stop()
}

func main() {
	mixtape := []string{"Jessie's Girl", "Whip it", "9 to 5"}
	var player Player = gadget.TapePlayer{}
	playList(player, mixtape)

	player = gadget.TapeRecorder{}
	playList(player, mixtape)

}
