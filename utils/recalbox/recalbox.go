package recalbox

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"strings"

	"github.com/djlechuck/recalbox-manager/structs"
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

// PseudoUUID generates a sort of UUID.
func PseudoUUID() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:]), nil
}

// GetConfValue returns the value of the given key of recalbox.conf file.
func GetConfValue(keys []string) (map[string]*structs.RecalboxConfValue, error) {
	file := viper.GetString("recalbox.confPath")
	// keysStr := strings.Replace(strings.Join(keys, "|"), ".", "\\.", -1)
	res := make(map[string]*structs.RecalboxConfValue)

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for len(keys) > 0 && scanner.Scan() {
		l := scanner.Text()

		for i, t := range keys {
			if strings.Contains(l, t) {
				parts := strings.Split(l, "=")
				disabled := false

				name := parts[0]
				if name[0:1] == ";" {
					name = name[1:]
					disabled = true
				}

				res[name] = &structs.RecalboxConfValue{
					Value:    parts[1],
					Disabled: disabled,
				}

				// Remove element from slice
				keys[i] = keys[len(keys)-1]
				keys = keys[:len(keys)-1]
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return res, nil
}
