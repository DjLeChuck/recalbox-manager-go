package structs

// HomeTile represents a tile on the homepage.
type HomeTile struct {
	Link, Image, Label string
}

// Language represents an available language for the interface.
type Language struct {
	Locale, Name, File string
}

// MenuItem represents an entry of the menu.
type MenuItem struct {
	Link, Glyph, Label string
	Children           []MenuItem
}
