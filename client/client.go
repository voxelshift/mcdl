package client

import (
	"encoding/json"
	"io"
	"mime"
	"net/http"
	"os"
	"path"

	log "github.com/sirupsen/logrus"
)

// Make a GET request to the given URL and parse its response into the provided interface
func GetJson(url string, v interface{}) error {
	log.Debugf("making http request to %s", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Errorf("http request to %s failed: %v", url, err)
		return err
	}
	defer resp.Body.Close()
	log.Debugf("successfully made http request to %s", url)

	log.Debugf("decoding json from http request to %s", url)
	if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
		log.Errorf("failed to decode json from %s: %v", url, err)
		return err
	}
	log.Debugf("successfully decoded json from http request to %s", url)

	return nil
}

func Download(url string, filename *string) (*string, error) {
	log.Debugf("initializing http request to %s", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Errorf("http request to %s failed: %v", url, err)
		return filename, err
	}
	defer resp.Body.Close()
	log.Debugf("successfully made http request to %s", url)

	// If a filename isn't provided, base the filename on the download itself
	if filename == nil {
		log.Debug("filename parameter not set; getting filename based on request")

		contentDisposition := resp.Header.Get("Content-Disposition")
		_, params, err := mime.ParseMediaType(contentDisposition)
		if err != nil {
			// If the Content-Disposition header isn't set, default to the base path
			basePath := path.Base(url)
			log.Debugf("Content-Disposition header missing; defaulting to base path name: %s", basePath)
			filename = &basePath
		} else {
			// Set the filename based on the Content-Disposition header
			contentPath := params["filename"]
			filename = &contentPath
		}
	}

	log.Debugf("creating output file %s", *filename)
	out, err := os.Create(*filename)
	if err != nil {
		log.Errorf("failed to create file %s: %v", *filename, err)
		return filename, err
	}
	defer out.Close()
	log.Debugf("successfully created output file %s", *filename)

	log.Debug("writing to output file")
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Errorf("failed to write to file %s: %v", *filename, err)
		return filename, err
	}
	log.Debug("successfully wrote to output file")

	return filename, nil
}
