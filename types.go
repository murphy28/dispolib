package dispolib

import (
	"net/http"
)

type Room struct {
	Name     string
	Password string
	index    int
	users    []user
}

type Session struct {
	Name     string
	Room     Room
	Client   http.Client
	handlers handlers
	typing   bool
}

type handlers struct {
	onMessage    func(Event)
	onJoin       func(Event)
	onLeave      func(Event)
	onType       func(Event)
	onNameChange func(Event)
	onCrash      func(Event)
}

type Event struct {
	User    user
	Payload string
	System  bool
}

type message struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
	User    string `json:"user"`
}

type user struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type roomState struct {
	Chats  []message `json:"chats"`
	LastID int       `json:"lastId"`
	Users  []user    `json:"users"`
}
