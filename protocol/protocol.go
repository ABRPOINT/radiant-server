package protocol

type MessageType string

const (
	MessageTypeSendMessage  MessageType = "send_message"
	MessageTypeCreateGroup  MessageType = "create_group"
	MessageTypeJoinGroup    MessageType = "join_group"
	MessageTypeLeaveGroup   MessageType = "leave_group"
	MessageTypeGroupMessage MessageType = "group_message"
	// Add more message types as needed.
)

type Message struct {
	Type    MessageType `json:"type"`
	Payload string      `json:"payload"` // Depending on the message type, this can be further unmarshaled to a more specific type.
}
