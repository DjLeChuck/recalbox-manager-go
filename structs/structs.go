package structs

// BiosFile represents a BIOS file.
type BiosFile struct {
	Name               string
	Md5                []string
	IsPresent, IsValid bool
}

// CheckValidity checks if the given MD5 is correct.
func (b BiosFile) CheckValidity(md5 string) bool {
	for _, m := range b.Md5 {
		if m == md5 {
			return true
		}
	}

	return false
}

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
	Link, Glyph, Label, ActiveClass string
	Children                        []MenuItem
	IsActive                        bool
}
