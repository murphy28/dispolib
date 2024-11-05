package dispolib

import (
	"strconv"
)

func (s *Session) UpdateUsers(users []User, runHandlers bool) error {
	if len(s.Room.Users) == 0 || !runHandlers {
		s.Room.Users = users
		return nil
	}

	var newUsers []User
	var leftUsers []User

	for _, user := range users {
		found := false
		for _, existingUser := range s.Room.Users {
			if user.ID == existingUser.ID {
				found = true
				break
			}
		}
		if !found {
			newUsers = append(newUsers, user)
		}
	}

	for _, existingUser := range s.Room.Users {
		found := false
		for _, user := range users {
			if user.ID == existingUser.ID {
				found = true

				if s.Handlers.OnTyping == nil {
					continue
				}

				userTyping := user.Status == "type" || user.Status == "mesg"
				existingUserTyping := existingUser.Status == "type" || existingUser.Status == "mesg"

				if userTyping != existingUserTyping && s.Name != user.Name {
					s.Handlers.OnTyping(Event{
						User:    user,
						Payload: strconv.FormatBool(userTyping),
					})
				}
			}
		}
		if !found {
			leftUsers = append(leftUsers, existingUser)
		}
	}

	for _, user := range newUsers {
		s.Handlers.OnJoin(Event{User: user})
	}
	for _, user := range leftUsers {
		s.Handlers.OnLeave(Event{User: user})
	}

	s.Room.Users = users

	return nil
}

func (s *Session) GetUsers() []User {
	return s.Room.Users
}

func (s *Session) GetUserByName(name string) User {
	for _, user := range s.Room.Users {
		if user.Name == name {
			return user
		}
	}
	return User{}
}
