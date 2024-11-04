package dispolib

import (
	"strconv"
)

func (s *Session) UpdateUsers(users []User, runHandlers bool) error {
	if len(s.room.users) == 0 || !runHandlers {
		s.room.users = users
		return nil
	}

	var newUsers []User
	var leftUsers []User

	for _, user := range users {
		found := false
		for _, existingUser := range s.room.users {
			if user.ID == existingUser.ID {
				found = true
				break
			}
		}
		if !found {
			newUsers = append(newUsers, user)
		}
	}

	for _, existingUser := range s.room.users {
		found := false
		for _, user := range users {
			if user.ID == existingUser.ID {
				found = true

				if s.handlers.onTyping == nil {
					continue
				}

				userTyping := user.Status == "type" || user.Status == "mesg"
				existingUserTyping := existingUser.Status == "type" || existingUser.Status == "mesg"

				if userTyping != existingUserTyping && s.name != user.Name {
					s.handlers.onTyping(Event{
						user:    user,
						payload: strconv.FormatBool(userTyping),
					})
				}
			}
		}
		if !found {
			leftUsers = append(leftUsers, existingUser)
		}
	}

	for _, user := range newUsers {
		s.handlers.onJoin(Event{user: user})
	}
	for _, user := range leftUsers {
		s.handlers.onLeave(Event{user: user})
	}

	s.room.users = users

	return nil
}

func (s *Session) GetUsers() []User {
	return s.room.users
}

func (s *Session) GetUserByName(name string) User {
	for _, user := range s.room.users {
		if user.Name == name {
			return user
		}
	}
	return User{}
}
