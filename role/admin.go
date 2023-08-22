package role

import (
	"github.com/moyai-network/carrot"
	"github.com/sandertv/gophertunnel/minecraft/text"
)

// Admin represents the role specification for the admin role.
type Admin struct{}

// Name returns the name of the role.
func (Admin) Name() string {
	return "admin"
}

// Chat returns the formatted chat message using the name and message provided.
func (Admin) Chat(name, message string) string {
	return text.Colourf("<grey>[<red>Admin</red>]</grey> <red>%s</red><dark-grey>:</dark-grey> <red>%s</red>", name, message)
}

// Colour returns the formatted name-Colour using the name provided.
func (Admin) Colour(name string) string {
	return text.Colourf("<red>%s</red>", name)
}

// Inherits returns the role that this role inherits from.
func (Admin) Inherits() carrot.Role {
	return SeniorMod{}
}
