package role

import (
	"github.com/moyai-network/carrot"
	"github.com/sandertv/gophertunnel/minecraft/text"
)

// SeniorMod represents the role specification for the senior mod role.
type SeniorMod struct{}

// Name returns the name of the role.
func (SeniorMod) Name() string {
	return "senior_mod"
}

// Chat returns the formatted chat message using the name and message provided.
func (SeniorMod) Chat(name, message string) string {
	return text.Colourf("<grey>[<dark-green>Senior Mod</dark-green>]</grey> <dark-green>%s</dark-green><dark-grey>:</dark-grey> <dark-green>%s</dark-green>", name, message)
}

// Colour returns the formatted name-Colour using the name provided.
func (SeniorMod) Colour(name string) string {
	return text.Colourf("<dark-green>%s</dark-green>", name)
}

// Inherits returns the role that this role inherits from.
func (SeniorMod) Inherits() carrot.Role {
	return Mod{}
}
