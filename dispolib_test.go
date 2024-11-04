package dispolib

import (
	"testing"
	"time"
)

func TestJoinRoom(t *testing.T) {
	dispo, err := New("Henry")
	if err != nil {
		return
	}

	dispo.AddHandler("message", func(e Event) {
		if dispo.typing {
			dispo.SetTyping(false)
		} else {
			dispo.SetTyping(true)
		}
	})

	dispo.AddHandler("crash", func(e Event) {
		println("Crash!")
	})

	dispo.JoinRoom(Room{name: "Test"})

	go dispo.Listen()

	time.Sleep(600 * time.Second)
}
