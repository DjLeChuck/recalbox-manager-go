package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

// GetFileMd5 calculates the MD5 of a file and returns it.
func GetFileMd5(filePath string) string {
	file, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	hasher := md5.New()
	_, err = io.Copy(hasher, file)

	if err != nil {
		panic(err)
	}

	return hex.EncodeToString(hasher.Sum(nil))
}
