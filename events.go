package dispolib

// AddHandler adds a handler for the given event type
func (s *Session) AddHandler(eventType string, h func(Event)) {
	switch eventType {
	case "message":
		s.Handlers.OnMessage = h
	case "join":
		s.Handlers.OnJoin = h
	case "leave":
		s.Handlers.OnLeave = h
	case "typing":
		s.Handlers.OnTyping = h
	case "nameChange":
		s.Handlers.OnNameChange = h
	case "crash":
		s.Handlers.OnCrash = h
	default:
		panic("Invalid event type")
	}
}
