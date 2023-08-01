package role

import (
	"github.com/moyai-network/carrot"
	"github.com/sandertv/gophertunnel/minecraft/text"
)

// Media represents the role structure for the Media role.
type Media struct{}

// Name returns the name of the role.
func (Media) Name() string {
	return "media"
}

// Chat returns the formatted chat message using the name and message provided.
func (Media) Chat(name, message string) string {
	return text.Colourf("<grey><i>[<aqua>Media</aqua>]</grey></i><aqua> %s</aqua><grey>:</grey> <white>%s</white>", name, message)
}

// Colour returns the formatted name-Colour using the name provided.
func (Media) Colour(name string) string {
	return text.Colourf("<i><aqua>%s</aqua></i>", name)
}

// Inherits returns the role that this role inherits from.
func (Media) Inherits() carrot.Role {
	return Plus{}
}
