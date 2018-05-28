package screenshots

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
)

// ListImages lists all images (PNG files) in the given directory.
func ListImages(directory string) (list []string, err error) {
	files, err := ioutil.ReadDir(directory)

	if err != nil {
		return nil, err
	}

	for _, file := range files {
		ext := filepath.Ext(file.Name())

		if ext != ".png" {
			continue
		}

		list = append(list, file.Name())
	}

	return list, nil
}

// CanTakeScreenshot indicates if the system can take a screenshot.
func CanTakeScreenshot() bool {
	f := viper.GetString("recalbox.arch")
	out, err := exec.Command("cat", f).CombinedOutput()
	if err != nil {
		return false
	}

	return "rpi" == string(out)[:3]
}

// TakeScreenshot calls rapis2png and take a screenshot of the actual recalbox screen.
func TakeScreenshot(destination string) error {
	if info, err := os.Stat(destination); err != nil || !info.IsDir() {
		return err
	}

	r := viper.GetString("recalbox.raspi2png")
	date := time.Now().Local().Format("2006-01-02-15-04-05")
	name := "screenshot-" + date + ".png"
	path := destination + name
	_, err := exec.Command(r, "-p", path).CombinedOutput()

	return err
}
