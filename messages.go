package dispolib

import (
	"errors"
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
		if msg.User == s.Name {
			continue
		}

		systemMessage := msg.User == ""

		if systemMessage {
			s.ParseSystemMessage(msg.Message)
		}

		if s.handlers.onMessage != nil {
			s.handlers.onMessage(Event{
				User:    user{Name: msg.User},
				Payload: msg.Message,
				System:  systemMessage,
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

	endpoint := GetEndpoint(s.Room, "")

	resp, err := s.Client.PostForm(endpoint, url.Values{
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

func (s *Session) ParseSystemMessage(msg string) {
	if s.handlers.onNameChange == nil {
		return
	}

	split := strings.Split(msg, " has changed their user name to ")

	if len(split) > 1 {
		s.handlers.onNameChange(Event{
			User:    user{Name: split[1]},
			Payload: split[0],
		})
	}
}
