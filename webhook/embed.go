package webhook

type Field struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}

type Embed struct {
	Title       string  `json:"title,omitempty"`
	Type        string  `json:"type,omitempty"`
	Description string  `json:"description,omitempty"`
	URL         string  `json:"url,omitempty"`
	Color       int     `json:"color,omitempty"`
	Footer      string  `json:"footer,omitempty"`
	Fields      []Field `json:"fields,omitempty"`
}
