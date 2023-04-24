// Реализовать утилиту wget с возможностью скачивать сайты целиком.

package main

import (
	"dev09/utils"
	"dev09/wget"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: wget <url>")
		os.Exit(1)
	}

	URLStr := os.Args[1]

	parsedURL, err := url.Parse(URLStr)
	if err != nil {
		log.Fatal("Error parsing URL", err)
	}

	// создаём директорию для скачивания
	baseURL := parsedURL.Scheme + "://" + parsedURL.Host
	dirName := parsedURL.Host
	if len(parsedURL.Path) > 1 {
		dirName = path.Join(dirName, parsedURL.Path[1:])
	}

	if !utils.FileExists(dirName) {
		os.MkdirAll(dirName, os.ModePerm)
	}

	// Получаем HTML-код страницы и сохраняем его в файл index.html
	res, err := http.Get(URLStr)
	if err != nil {
		log.Fatal("Error on GET", err)
	}

	defer res.Body.Close()

	file, err := os.Create(path.Join(dirName, "index.html"))
	if err != nil {
		log.Fatal("Error creating file: ", err)
	}

	defer file.Close()

	if _, err := io.Copy(file, res.Body); err != nil {
		log.Fatal("Error reading body: ", err)
	}

	res, err = http.Get(URLStr)
	if err != nil {
		log.Fatal("Error on GET: ", err)
	}

	defer res.Body.Close()

	// Парсим HTML-код и ищем ссылки на другие файлы
	links, err := utils.GetLinks(baseURL, res.Body)
	if err != nil {
		log.Fatal("Error getting links: ", err)
	}

	// Рекурсивно скачиваем найденные файлы
	for idx, link := range links {
		fullLink := baseURL + link
		// Ссылка уже скачана, пропускаем
		if utils.FileExists(link) {
			continue
		}
		// Абсолютная ссылка, добавляем домен
		if strings.HasPrefix(link, "/") {
			fullLink = baseURL + link
		} else if strings.HasPrefix(link, "http") {
			fullLink = link
		} else {
			// Относительная ссылка, добавляем директорию
			fullLink = baseURL + "/" + filepath.Join(filepath.Dir(parsedURL.Path), link)
		}

		err := wget.WGet(fullLink, dirName)
		if err != nil {
			log.Println(idx, err)
		}
	}
}
