package role

import (
	"github.com/moyai-network/carrot"
	"github.com/sandertv/gophertunnel/minecraft/text"
)

// Famous represents the role structure for the famous role.
type Famous struct{}

// Name returns the name of the role.
func (Famous) Name() string {
	return "famous"
}

// Chat returns the formatted chat message using the name and message provided.
func (Famous) Chat(name, message string) string {
	return text.Colourf("<grey><i>[<purple>Famous</purple>]</grey></i><purple> %s</purple><grey>:</grey> <white>%s</white>", name, message)
}

// Colour returns the formatted name-Colour using the name provided.
func (Famous) Colour(name string) string {
	return text.Colourf("<i><purple>%s</purple></i>", name)
}

// Inherits returns the role that this role inherits from.
func (Famous) Inherits() carrot.Role {
	return Media{}
}
