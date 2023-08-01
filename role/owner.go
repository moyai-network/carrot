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
	return text.Colourf("<grey>[<lapis>Owner</lapis>]</grey> <lapis>%s</lapis><dark-grey>:</dark-grey> <lapis>%s</lapis>", name, message)
}

// Colour returns the formatted name-Colour using the name provided.
func (Owner) Colour(name string) string {
	return text.Colourf("<lapis>%s</lapis>", name)
}

// Inherits returns the role that this role inherits from.
func (Owner) Inherits() carrot.Role {
	return Manager{}
}
