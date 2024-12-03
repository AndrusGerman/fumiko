package domain

type RoleID string

const (
	UserRoleID      RoleID = "user"
	AssistantRoleID RoleID = "assistant"
)

type Message struct {
	RoleID  RoleID
	Content string
}

func NewMessage(content string, roleId RoleID) *Message {
	return &Message{
		RoleID:  roleId,
		Content: content,
	}
}
