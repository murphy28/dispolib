package dispolib

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

// JoinRoom attempts to join a chat room with the given room details.
func (s *Session) JoinRoom(room Room) error {
	s.room = room

	resp, err := s.client.PostForm(BaseURL, url.Values{
		"user-name": {s.name},
		"room-name": {s.room.name},
		"pass":      {s.room.password},
	})

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("unable to join room: status " + resp.Status)
	}

	return s.UpdateIndex()
}

// RoomState retrieves the current state of the chat room.
func (s *Session) RoomState() (RoomState, error) {
	endpoint := GetEndpoint(s.room, "/ajax?lastId="+strconv.Itoa(s.room.index)+GetTypingParam(s.typing))

	resp, err := s.client.Get(endpoint)
	if err != nil {
		return RoomState{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if s.handlers.onCrash != nil {
			s.handlers.onCrash(Event{})
		}

		return RoomState{}, errors.New("cannot get room state: status " + resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return RoomState{}, err
	}

	var state RoomState
	err = json.Unmarshal(body, &state)
	if err != nil {
		return RoomState{}, err
	}

	s.room.index = state.LastID

	s.UpdateUsers(state.Users, true)

	return state, nil
}

// UpdateIndex updates the room index to the latest state.
func (s *Session) UpdateIndex() error {
	state, err := s.RoomState()
	if err != nil {
		return err
	}

	s.UpdateUsers(state.Users, false)

	s.room.index = state.LastID
	return nil
}

// IsJoined checks if the session is joined to a room.
func (s *Session) IsJoined() bool {
	return s.room.name != "" || s.room.index != 0
}
