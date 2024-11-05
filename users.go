package dispolib

import (
	"strconv"
)

func (s *Session) UpdateUsers(users []user, runHandlers bool) error {
	if len(s.Room.users) == 0 || !runHandlers {
		s.Room.users = users
		return nil
	}

	var newUsers []user
	var leftUsers []user

	if s.handlers.onJoin != nil {

		for _, user := range users {
			found := false
			for _, existingUser := range s.Room.users {
				if user.ID == existingUser.ID {
					found = true
					break
				}
			}
			if !found {
				newUsers = append(newUsers, user)
			}
		}

		for _, user := range newUsers {
			s.handlers.onJoin(Event{User: user})
		}

	}

	for _, existingUser := range s.Room.users {
		found := false
		for _, user := range users {
			if user.ID == existingUser.ID {
				found = true

				if s.handlers.onType == nil {
					continue
				}

				userTyping := user.Status == "type" || user.Status == "mesg"
				existingUserTyping := existingUser.Status == "type" || existingUser.Status == "mesg"

				if userTyping != existingUserTyping && s.Name != user.Name {
					s.handlers.onType(Event{
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

	if s.handlers.onLeave != nil {
		for _, user := range leftUsers {
			s.handlers.onLeave(Event{User: user})
		}
	}

	s.Room.users = users

	return nil
}

func (s *Session) GetUsers() []user {
	return s.Room.users
}

func (s *Session) GetUserByName(name string) user {
	for _, user := range s.Room.users {
		if user.Name == name {
			return user
		}
	}
	return user{}
}
