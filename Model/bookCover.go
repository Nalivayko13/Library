package Model

import (
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
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
func SaveCoverFile(cover []byte) {
	name:=strconv.Itoa(rand.Int())
	file, err := os.Create(fmt.Sprintf("C:/GO/booksCover/%s.jpg",name))
	if err!=nil{
		log.Println(err)
	}
	defer file.Close()
	file.Write(cover)
	log.Println("Downloaded cover")

}