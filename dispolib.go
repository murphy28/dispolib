package dispolib

// Disposable Chat Endpoint
const BaseURL = "http://www.disposablechat.com/chat"

// Construct a URL for a chat room
func GetEndpoint(room Room, path string) string {
	return BaseURL + "/" + room.Name + path
}

func GetTypingParam(isTyping bool) string {
	if isTyping {
		return "&status=type"
	}
	return "&status=null"
}
