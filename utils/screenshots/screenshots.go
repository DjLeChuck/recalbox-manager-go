package screenshots

import (
	"io/ioutil"
	"path/filepath"
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
