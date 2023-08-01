package role

import (
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
	return text.Colourf("<grey>[<dark-green>Mod</dark-green>]</grey> <dark-green>%s</dark-green><dark-grey>:</dark-grey> <dark-green>%s</dark-green>", name, message)
}

// Colour returns the formatted name-Colour using the name provided.
func (Mod) Colour(name string) string {
	return text.Colourf("<dark-green>%s</dark-green>", name)
}
