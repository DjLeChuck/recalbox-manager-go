package recalbox

import (
	"crypto/rand"
	"fmt"
	"os/exec"
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

// convertToSimpleType takes a reflect Value and convert it into a
// "simple type" like string or int.
func convertToSimpleType(v reflect.Value) interface{} {
	switch v.Interface().(type) {
	case string:
		return v.String()
	case int:
		return v.Int()
	case bool:
		if v.Bool() == true {
			return 1
		}

		return 0
	default:
		return ""
	}
}

// FormatFormData takes a form struct and convert it into an iterable map of
// key: value for making others processings.
func FormatFormData(form interface{}) (data map[string]interface{}) {
	data = make(map[string]interface{})
	s := reflect.ValueOf(form).Elem()
	typeOfT := s.Type()

	for i := 0; i < s.NumField(); i++ {
		v := s.Field(i)
		data[typeOfT.Field(i).Tag.Get("form")] = convertToSimpleType(v)
	}

	return data
}

// ProcessRecalboxSettingsForm loops through the form data and update the
// config of Recalbox by calling the config script.
func ProcessRecalboxSettingsForm(data map[string]interface{}) (err error) {
	pythonFile := viper.GetString("recalbox.pythonSettingsFile")

	for k, v := range data {
		value := fmt.Sprintf("%v", v)

		if _, ok := v.(string); ok {
			value = "'" + v.(string) + "'"
		}

		normalizedKey := strings.Replace(k, "-", ".", -1)
		_, err = exec.Command("python", pythonFile, "-command", "save", "-key", normalizedKey, "-value", value).CombinedOutput()

		if err != nil {
			return err
		}
	}

	if data["audio-volume"] != nil {
		configScript := viper.GetString("recalbox.configScript")
		_, err = exec.Command(configScript, "volume", fmt.Sprintf("%v", data["audio-volume"])).CombinedOutput()

		if err != nil {
			return err
		}
	}

	return nil
}

// PseudoUuid generates a sort of UUID.
func PseudoUuid() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:]), nil
}
