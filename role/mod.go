package role

import (
	"github.com/moyai-network/carrot"
	"github.com/sandertv/gophertunnel/minecraft/text"
)

// Mod represents the role specification for the mod role.
type Mod struct{}

// Name returns the name of the role.
func (Mod) Name() string {
	return "mod"
}

// Chat returns the formatted chat message using the name and message provided.
func (Mod) Chat(name, message string) string {
	return text.Colourf("<grey>[<green>Mod</green>]</grey> <green>%s</green><dark-grey>:</dark-grey> <green>%s</green>", name, message)
}

// Colour returns the formatted name-Colour using the name provided.
func (Mod) Colour(name string) string {
	return text.Colourf("<green>%s</green>", name)
}

// Inherits returns the role that this role inherits from.
func (Mod) Inherits() carrot.Role {
	return Trial{}
}
