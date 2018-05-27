package smartfile

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/djlechuck/recalbox-manager/structs"
	"github.com/spf13/viper"
)

// UploadArchive uploads a support archive.
func UploadArchive(file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()

	// Prepare the file to upload by copying local archive in buffer.
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	filename := filepath.Base(f.Name())
	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return "", err
	}
	io.Copy(part, f)
	writer.Close()

	// Prepare the request (URL, auth, headers)
	url := viper.GetString("smartFile.url") + viper.GetString("smartFile.api.upload") + viper.GetString("smartFile.folderName")
	r, err := http.NewRequest("POST", url, body)
	if err != nil {
		return "", err
	}
	r.Close = true
	r.SetBasicAuth(viper.GetString("smartFile.keys.public"), viper.GetString("smartFile.keys.private"))
	r.Header.Add("Content-Type", writer.FormDataContentType())

	// Execute the upload
	c := &http.Client{Timeout: time.Second * 20}
	resp, err := c.Do(r)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return viper.GetString("smartFile.folderName") + "/" + filename, nil
}

// GetLink generates a public link to access the given path.
func GetLink(path string) (string, error) {
	// Prepare data
	data := url.Values{}
	data.Set("path", path)
	data.Add("read", "1")
	data.Add("list", "1")

	// Prepare the request (URL, auth, headers)
	url := viper.GetString("smartFile.url") + viper.GetString("smartFile.api.link")
	body := strings.NewReader(data.Encode())
	r, err := http.NewRequest("POST", url, body)
	if err != nil {
		return "", err
	}
	r.Close = true
	r.SetBasicAuth(viper.GetString("smartFile.keys.public"), viper.GetString("smartFile.keys.private"))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	// Execute the upload
	c := &http.Client{Timeout: time.Second * 20}
	resp, err := c.Do(r)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	l := structs.SmartFileLink{}
	err = json.Unmarshal(b, &l)
	if err != nil {
		return "", err
	}

	return l.Href, nil
}
