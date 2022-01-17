package Model

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
)

func DownloadFile(URL,filename string) error {
	response, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("Received non 200 response code")
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()


	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}
	log.Println("Downloaded cover")
	return nil
}
