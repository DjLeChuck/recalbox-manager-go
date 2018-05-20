package recalbox

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/viper"
)

// ProcessRecalboxSettingsForm loops through the form data and update the
// config of Recalbox by calling the config script.
// checkboxes represents checkboxes on the page. When submitted unchecked,
// they have no value, so we force one.
func ProcessRecalboxSettingsForm(data map[string]interface{}, checkboxes []string) (err error) {
	for _, c := range checkboxes {
		if _, ok := data[c]; !ok {
			data[c] = "0"
		}
	}

	pythonFile := viper.GetString("recalbox.pythonSettingsFile")

	for k, v := range data {
		normalizedKey := strings.Replace(k, "-", ".", -1)
		_, err = exec.Command("python", pythonFile, "-command", "save", "-key", normalizedKey, "-value", "'"+v.(string)+"'").CombinedOutput()
		fmt.Println("python", pythonFile, "-command", "save", "-key", normalizedKey, "-value", "'"+v.(string)+"'")
		if err != nil {
			return err
		}
	}

	if data["audio-volume"] != nil {
		configScript := viper.GetString("recalbox.configScript")
		_, err = exec.Command(configScript, "volume", data["audio-volume"].(string)).CombinedOutput()

		if err != nil {
			return err
		}
	}

	return nil
}
