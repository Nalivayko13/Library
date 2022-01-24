package Model

import (
	"fmt"
	"library/dao"
	"log"
	"net/smtp"
	"os"
	"time"
)

func FindDebtorAndSendEmail() {
	var rent []dao.Rent
	dao.Get_AllRent_fromDB(&rent)
	log.Println(rent)
	for _, r := range rent {
		t1,_:=time.Parse("01-02-2006", r.Last_date)
		t2:=time.Now()
		if t2.After(t1) && r.Comlete=="0" {
			readersEmail:=dao.Get_email_byId(r.Id_reeder)
			//email(readersEmail)
			fmt.Println(readersEmail)
		}
	}


}


func email(email string) {
	from := os.Getenv("John.Doe@gmail.com")
	password := os.Getenv("ieiemcjdkejspqz")
	toEmail := os.Getenv(email) // ex: "Jane.Smith@yahoo.com"
	to := []string{toEmail}
	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port
	subject := "Subject: Library email\n"
	body := "bring the book back!"
	message := []byte(subject + body)
	auth := smtp.PlainAuth("", from, password, host)
	// func SendMail(addr string, a Auth, from string, to []string, msg []byte) error
	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
}
// Вызов переданной функции раз в сутки в указанное время.
func CallAt(hour, min, sec int, f func()) error {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return err
	}

	// Вычисляем время первого запуска.
	now := time.Now().Local()
	firstCallTime := time.Date(
		now.Year(), now.Month(), now.Day(), hour, min, sec, 0, loc)
	if firstCallTime.Before(now) {
		// Если получилось время раньше текущего, прибавляем сутки.
		firstCallTime = firstCallTime.Add(time.Hour * 24)
	}

	// Вычисляем временной промежуток до запуска.
	duration := firstCallTime.Sub(time.Now().Local())

	go func() {
		time.Sleep(duration)
		for {
			f()
			// Следующий запуск через сутки.
			time.Sleep(time.Hour * 24)
		}
	}()

	return nil
}

