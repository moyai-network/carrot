package role

import (
	"github.com/moyai-network/carrot"
	"github.com/sandertv/gophertunnel/minecraft/text"
)

// Owner represents the role specification for the owner role.
type Owner struct{}

// Name returns the name of the role.
func (Owner) Name() string {
	return "owner"
}

// Chat returns the formatted chat message using the name and message provided.
func (Owner) Chat(name, message string) string {
	return text.Colourf("<grey>[<orange>Owner</orange>]</grey> <orange>%s</orange><dark-grey>:</dark-grey> <orange>%s</orange>", name, message)
}

// Colour returns the formatted name-Colour using the name provided.
func (Owner) Colour(name string) string {
	return text.Colourf("<orange>%s</orange>", name)
}

// Inherits returns the role that this role inherits from.
func (Owner) Inherits() carrot.Role {
	return Manager{}
}
