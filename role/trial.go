package role

import (
	"github.com/sandertv/gophertunnel/minecraft/text"
)

// Trial represents the role specification for the trial role.
type Trial struct{}

// Name returns the name of the role.
func (Trial) Name() string {
	return "Trial"
}

// Chat returns the formatted chat message using the name and message provided.
func (Trial) Chat(name, message string) string {
	return text.Colourf("<grey>[<yellow>Trial</yellow>]</grey> <yellow>%s</yellow><dark-grey>:</dark-grey> <yellow>%s</yellow>", name, message)
}

// Colour returns the formatted name-Colour using the name provided.
func (Trial) Colour(name string) string {
	return text.Colourf("<yellow>%s</yellow>", name)
}
