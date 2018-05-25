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

// Configuration represents the form on /configuration page.
type Configuration struct {
	SystemLanguage                 string `form:"system-language"`
	SystemKblayout                 string `form:"system-kblayout"`
	SystemTimezone                 string `form:"system-timezone"`
	SystemHostname                 string `form:"system-hostname"`
	WifiEnabled                    bool   `form:"wifi.enabled"`
	WifiSsid                       string `form:"wifi-ssid"`
	WifiKey                        string `form:"wifi-key"`
	WifiSsid2                      string `form:"wifi-ssid2"`
	WifiKey2                       string `form:"wifi-key2"`
	WifiSsid3                      string `form:"wifi-ssid3"`
	WifiKey3                       string `form:"wifi-key3"`
	KodiEnabled                    bool   `form:"kodi-enabled"`
	KodiAtStartup                  bool   `form:"kodi-atstartup"`
	KodiXButton                    bool   `form:"kodi-xbutton"`
	SystemEsMenu                   string `form:"system-es-menu"`
	EmulationStationSelectedSystem string `form:"emulationstation-selectedsystem"`
	EmulationStationBootOnGamelist bool   `form:"emulationstation-bootongamelist"`
	EmulationStationHideSystemView bool   `form:"emulationstation-hidesystemview"`
	EmulationStationGamelistOnly   bool   `form:"emulationstation-gamelistonly"`
	SystemEmulatorsSpecialKeys     string `form:"system-emulators-specialkeys"`
	SystemAPIEnabled               bool   `form:"system-api-enabled"`
	UpdatesEnabled                 bool   `form:"updates-enabled"`
	UpdatesType                    string `form:"updates-type"`
}

// Logs represents the form on /logs page.
type Logs struct {
	File string `form:"log-file"`
}

// RecalboxConf represents the form on /recalbox-conf page.
type RecalboxConf struct {
	Content string `form:"content"`
}
