package dispolib

import (
	"net/http"
)

type Room struct {
	Name     string
	Password string
	Index    int
	Users    []User
}

type Session struct {
	Name     string
	Room     Room
	Client   http.Client
	Handlers Handlers
	Typing   bool
}

type Handlers struct {
	OnMessage    func(Event)
	OnJoin       func(Event)
	OnLeave      func(Event)
	OnTyping     func(Event)
	OnNameChange func(Event)
	OnCrash      func(Event)
}

type Event struct {
	User    User
	Payload string
	System  bool
}

type Message struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
	User    string `json:"user"`
}

type User struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type RoomState struct {
	Chats  []Message `json:"chats"`
	LastID int       `json:"lastId"`
	Users  []User    `json:"users"`
}
