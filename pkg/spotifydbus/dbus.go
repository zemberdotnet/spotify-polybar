package spotifydbus

import (
	"fmt"
	"github.com/godbus/dbus"
)

const SPOT_PATH = "/org/mpris/MediaPlayer2"
const SPOT_MEMB = "org.mpris.MediaPlayer2.Player"
const SPOT_DEST = "org.mpris.MediaPlayer2.spotify"

const DBUS_PROP_GET = "org.freedesktop.DBus.Properties.Get"

func GetArtist() string {
	metadata := GetMetadata()
	if metadata == nil {
		return ""
	}
	return metadata["xesam:artist"].Value().([]string)[0]
}

func GetAlbum() string {
	metadata := GetMetadata()
	if metadata == nil {
		return ""
	}
	return metadata["xesam:album"].Value().(string)
}

func GetTitle() string {
	metadata := GetMetadata()
	if metadata == nil {
		return ""
	}
	return metadata["xesam:title"].Value().(string)
}

func GetMetadata() map[string]dbus.Variant {
	conn, err := dbus.SessionBus()
	if err != nil {
		fmt.Println(err)
	}

	spotifyBus := conn.Object(SPOT_DEST, SPOT_PATH)

	call := spotifyBus.Call(DBUS_PROP_GET, 0, SPOT_MEMB, "Metadata")
	if call.Err != nil {
		//
	}
	if len(call.Body) == 0 {
		return nil
	}

	return call.Body[0].(dbus.Variant).Value().(map[string]dbus.Variant)
}
