package dispolib

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// GetMessages retrieves the current state of the chat room and runs any
// registered message handlers.
func (s *Session) GetMessages() error {
	state, err := s.RoomState()
	if err != nil {
		return err
	}

	for _, msg := range state.Chats {
		// Ignore messages from self
		if msg.User == s.name {
			continue
		}

		systemMessage := msg.User == ""

		if systemMessage {
			s.ParseSystemMessage(msg.Message)
			fmt.Print("Hi")
		}

		if s.handlers.onMessage != nil {
			s.handlers.onMessage(Event{
				user:    User{Name: msg.User},
				payload: msg.Message,
				system:  systemMessage,
			})
		}
	}

	return nil
}

// SendMessage sends a message to the current chat room.
func (s *Session) SendMessage(msg string) error {
	if !s.IsJoined() {
		return errors.New("bot is not joined to a room")
	}

	endpoint := GetEndpoint(s.room, "")

	resp, err := s.client.PostForm(endpoint, url.Values{
		"message_input_window": {msg},
		"noRender":             {"false"},
	})

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("unable to send message: status " + resp.Status)
	}

	return nil
}

func (s *Session) ParseSystemMessage(msg string) bool {
	split := strings.Split(msg, " has changed their user name to ")
	fmt.Println("hi")
	if len(split) > 1 {
		s.handlers.onNameChange(Event{
			user:    User{Name: split[1]},
			payload: split[0],
		})
	}

	return true
}
