package dispolib

// AddHandler adds a handler for the given event type
func (s *Session) AddHandler(eventType string, h func(Event)) {
	switch eventType {
	case "message":
		s.handlers.onMessage = h
	case "join":
		s.handlers.onJoin = h
	case "leave":
		s.handlers.onLeave = h
	case "typing":
		s.handlers.onType = h
	case "nameChange":
		s.handlers.onNameChange = h
	case "crash":
		s.handlers.onCrash = h
	default:
		panic("Invalid event type")
	}
}
