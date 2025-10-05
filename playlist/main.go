package main

import (
	"fmt"

	"github.com/NoahTek/head-first-go/gadget"
)

type Player interface {
	Play(string)
	Stop()
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func TryOut(player Player) {
	player.Play("Test Track")
	player.Stop()
	recorder, ok := player.(gadget.TapeRecorder)
	if ok {
		recorder.Record()
	} else {
		fmt.Println("player was not TapeRecorder")
	}
}

func main() {
	mixtape := []string{"Jessie's Girl", "Whip it", "9 to 5"}
	var player Player = gadget.TapePlayer{}
	playList(player, mixtape)
	player = gadget.TapeRecorder{}
	playList(player, mixtape)

	TryOut(gadget.TapeRecorder{})
	TryOut(gadget.TapePlayer{})
}
