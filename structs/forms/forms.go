package forms

// Audio represents the form on /audio page.
type Audio struct {
	BackgroundMusic bool   `form:"audio-bgmusic"`
	Volume          int    `form:"audio-volume"`
	Device          string `form:"audio-device"`
}

// Controllers represents the form on /controllers page.
type Controllers struct {
	Db9Enabled     bool   `form:"controllers-db9-enabled"`
	Db9Args        string `form:"controllers-db9-args"`
	GameconEnabled bool   `form:"controllers-gamecon-enabled"`
	GameconArgs    string `form:"controllers-gamecon-args"`
	GpioEnabled    bool   `form:"controllers-gpio-enabled"`
	GpioArgs       string `form:"controllers-gpio-args"`
	Ps3Enabled     bool   `form:"controllers-ps3-enabled"`
	Ps3Driver      string `form:"controllers-ps3-driver"`
}

// Systems represents the form on /systems page.
type Systems struct {
	Ratio                            string `form:"global-ratio"`
	Shaderset                        string `form:"global-shaderset"`
	SmoothEnabled                    bool   `form:"global-smooth"`
	RewindEnabled                    bool   `form:"global-rewind"`
	AutosaveEnabled                  bool   `form:"global-autosave"`
	IntegerscaleEnabled              bool   `form:"global-integerscale"`
	RetroachievementsEnabled         bool   `form:"global-retroachievements"`
	RetroachievementsUsername        string `form:"global-retroachievements-username"`
	RetroachievementsPassword        string `form:"global-retroachievements-password"`
	RetroachievementsHardcoreEnabled bool   `form:"global-retroachievements-hardcore"`
}
