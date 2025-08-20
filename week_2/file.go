package week_2

import (
	"GolangBootcamp/common"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

func GetInfo() common.Username {
	var username string
	var PhoneNumber string
	dateTime := time.Now().Format(time.DateTime)
	fmt.Print("Write your name : ")
	fmt.Scan(&username)
	fmt.Print("Write your PhoneNumber : ")
	fmt.Scan(&PhoneNumber)
	un := common.Username{
		DateTime:    dateTime,
		Username:    username,
		PhoneNumber: PhoneNumber,
	}
	return un

}
//No need to use it, but maybe later.
// func CreateUsernames(username common.Username) []common.Username {
// 	var Usernames []common.Username
// 	Usernames = append(Usernames, common.Username{
// 		DateTime:    username.DateTime,
// 		Username:    username.Username,
// 		PhoneNumber: username.PhoneNumber,
// 	})
// 	return Usernames
// }

func readFile() ([]string, error) {
	fileName := "week_2/UserName.txt"

	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	stringData := string(data)
	splitStringData := strings.Split(stringData, "|")

	return splitStringData, nil
}

func CheckUniqUsername(unInfo common.Username) error {
	data, err := readFile()
	if err != nil {
		return err
	}
	for i := 1; i < len(data); i += 3 {
		// fmt.Println(i ,data[i])
		if strings.TrimSpace(data[i]) == unInfo.Username {
			return errors.New("username used")
		}
	}
	

	for i := 2; i < len(data); i += 3 {
		// fmt.Println(i ,data[i])
		if strings.TrimSpace(data[i]) == unInfo.PhoneNumber {
			return errors.New("PhoneNumber used")
		}
	}
	return nil
}

func SaveInFile(username common.Username) error {
	fileName := "week_2/UserName.txt"

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return errors.Join(errors.New("we have error in open or create file"), err)

	}
	defer file.Close()

	usernameString := fmt.Sprintf("%s|%s|%s|\n",
		username.DateTime,
		username.Username,
		username.PhoneNumber)

	_, err = file.WriteString(usernameString)
	if err != nil {
		return errors.Join(errors.New("we have error in Write file"), err)

	}

	return nil
}
