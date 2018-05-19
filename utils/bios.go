package utils

import (
	"bufio"
	"os"
	"regexp"
)

// GetBiosMD5List returns list of all BIOS file with their valid MD5.
func GetBiosMD5List(md5File string) map[string][]string {
	re := regexp.MustCompile("^([a-f0-9]{32})[ ]+(.*)$")
	file, err := os.Open(md5File)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	list := map[string][]string{}

	for scanner.Scan() {
		m := re.FindStringSubmatch(scanner.Text())

		if 3 != len(m) {
			continue
		}

		// m[1] -> MD5 ; m[2] -> BIOS name
		list[m[2]] = append(list[m[2]], m[1])
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return list
}
