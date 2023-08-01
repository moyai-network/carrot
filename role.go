package carrot

type Role interface {
	// Name returns the name of the role
	Name() string
	// Colour returns the provided name with the colour associated with the role
	Colour(name string) string
	// Chat returns the formatted chat message using the provided name and message
	Chat(name, message string) string
}

// HeirRole represents a role that inherits from another role.
type HeirRole interface {
	// Inherits returns the role that this role inherits from.
	Inherits() Role
}
