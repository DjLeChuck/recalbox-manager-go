package structs

// Credentials represents the login/password used to access to the manager.
type Credentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// Authentication represents the state of the application authentication.
type Authentication struct {
	Enabled         bool
	Credentials     Credentials
	IsAuthenticated bool
}

// Reset Authentication with new Credentials.
func (a *Authentication) Reset(cred Credentials) {
	a.Enabled = true
	a.Credentials = cred
	a.IsAuthenticated = false
}

// AvailableLanguage represents an available language in the menu.
type AvailableLanguage struct {
	Locale, Name string
}

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

// HelpLink represents an entry of the help links list.
type HelpLink struct {
	Link, Label string
}

// SelectList represents an entry of a <select> list.
type SelectList struct {
	Value, Label string
}

// SmartFileLink represents a Link object on SmartFile.
type SmartFileLink struct {
	Href string `json:"href"`
}

// CPU represents a CPU on the system.
type CPU struct {
	Number int
	Value  string
}

// Disk represents a disk mounted on the system.
type Disk struct {
	Device, Path, UsedPercent, Used, Free, Total string
}

// RecalboxConfValue represents a value of recalbox.conf file.
type RecalboxConfValue struct {
	Value    string
	Disabled bool
}
