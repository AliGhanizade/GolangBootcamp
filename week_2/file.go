package week_2

import (
	"fmt"
	"os"
	"time"
)

func GetInfo() string {
	var name string
	var phoneNumber string

	fmt.Print("Write your name : ")
	fmt.Scan(&name)
	fmt.Print("Write your PhoneNumber : ")
	fmt.Scan(&phoneNumber)

	output := fmt.Sprintf("%s|%s|%s\n", time.Now().Format(time.DateTime), name, phoneNumber)
	return output
}

func SaveInFile(username []byte) error {
	fileName := "week_2/UserName.txt"

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(username)
	if err != nil {
		return err
	}

	_, err = file.WriteString("\n")
	if err != nil {
		return err
	}

	return nil
}


