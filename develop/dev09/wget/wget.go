package wget

import (
	"dev09/utils"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

func WGet(url string, dirName string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("Status code is not 200, but %v", res.StatusCode)
	}

	if strings.HasPrefix(res.Header.Get("Content-Type"), "text/html") {
		fmt.Println("hello")
		filename := path.Join(dirName, strings.TrimLeft(url, "http://"))
		filename = strings.TrimSuffix(filename, "/")
		filename = strings.TrimSuffix(filename, ".html")
		filename = filename + "/index.html"
		filename = strings.ReplaceAll(filename, "/", "_")
		filename = strings.ReplaceAll(filename, ":", "_")
		file, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer file.Close()

		if _, err := io.Copy(file, res.Body); err != nil {
			return err
		}

		links, err := utils.GetLinks(url, res.Body)
		if err != nil {
			return fmt.Errorf("Error parsing HTML: %v", err)
		}

		for _, link := range links {
			fullLink := url + "/" + link
			if utils.FileExists(fullLink) {
				continue
			}

			err := WGet(fullLink, filename[:len(filename)-10])
			if err != nil {
				return fmt.Errorf("Error downloading file. %s:%v", fullLink, err)
			}
		}
	} else {
		filename := path.Join(dirName, strings.TrimLeft(url, "http://"))
		filename = strings.ReplaceAll(filename, "/", "_")
		filename = strings.ReplaceAll(filename, ":", "_")
		file, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer file.Close()

		if _, err := io.Copy(file, res.Body); err != nil {
			return err
		}
	}
	return nil
}
