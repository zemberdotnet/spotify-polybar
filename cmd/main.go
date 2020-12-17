package main

import (
	"fmt"
	"github.com/zemberdotnet/polybar-spotify-go/pkg/spotifydbus"
)

func main() {
	artist := spotifydbus.GetArtist()
	title := spotifydbus.GetTitle()
	if artist != "" && title != "" {
		fmt.Printf("%v: %v", spotifydbus.GetArtist(), spotifydbus.GetTitle())
	}
}
