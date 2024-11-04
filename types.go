package dispolib

import (
	"net/http"
)

type Room struct {
	name     string
	password string
	index    int
	users    []User
}

type Session struct {
	name     string
	room     Room
	client   http.Client
	handlers Handlers
	typing   bool
}

type Handlers struct {
	onMessage    func(Event)
	onJoin       func(Event)
	onLeave      func(Event)
	onTyping     func(Event)
	onNameChange func(Event)
	onCrash      func(Event)
}

type Event struct {
	user    User
	payload string
	system  bool
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
