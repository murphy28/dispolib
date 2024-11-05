package dispolib

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

// Create a new session with the given username
func New(name string) (*Session, error) {
	jar, _ := cookiejar.New(nil)

	return &Session{
		Name:   name,
		Room:   Room{},
		Client: http.Client{Jar: jar},
	}, nil
}

func (s *Session) UpdateName(name string) error {
	if !s.IsJoined() {
		return errors.New("bot is not joined to a room")
	}

	endpoint := GetEndpoint(s.Room, "/change-user")

	resp, err := s.Client.PostForm(endpoint, url.Values{
		"new_name": {name},
		"noRender": {"true"},
	})

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("unable to change name: status " + resp.Status)
	}

	s.Name = name

	return nil
}

func (s *Session) SetTyping(isTyping bool) error {
	if !s.IsJoined() {
		return errors.New("bot is not joined to a room")
	}

	s.Typing = isTyping

	return nil
}

func (s *Session) Listen() {
	fmt.Println("Dispolib listening...")
	for {
		s.GetMessages()
		time.Sleep(1 * time.Second)
	}
}
