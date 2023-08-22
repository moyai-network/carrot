package role

import "github.com/sandertv/gophertunnel/minecraft/text"

// Dev represents the role structure for the dev role.
type Dev struct{}

// Name returns the name of the role.
func (Dev) Name() string {
	return "Dev"
}

// Chat returns the formatted chat message using the name and message provided.
func (Dev) Chat(name, message string) string {
	return text.Colourf("<grey>[<dark-aqua>Dev</dark-aqua>]</grey><dark-aqua> %s</dark-aqua><grey>:</grey> <dark-aqua>%s</dark-aqua>", name, message)
}

// Colour returns the formatted name-Colour using the name provided.
func (Dev) Colour(name string) string {
	return text.Colourf("<dark-aqua>%s</dark-aqua>", name)
}
