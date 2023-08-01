package role

import "github.com/sandertv/gophertunnel/minecraft/text"

// Plus represents the role structure for the plus role.
type Plus struct{}

// Name returns the name of the role.
func (Plus) Name() string {
	return "+"
}

// Chat returns the formatted chat message using the name and message provided.
func (Plus) Chat(name, message string) string {
	return text.Colourf("<grey>[<black>+</black>]</grey><black> %s</black><grey>:</grey> <black>%s</black>", name, message)
}

// Colour returns the formatted name-Colour using the name provided.
func (Plus) Colour(name string) string {
	return text.Colourf("<black>%s</black>", name)
}
