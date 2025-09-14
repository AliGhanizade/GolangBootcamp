package week_5

import (
	"GolangBootcamp/common"
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func countLogAccess() *common.LogReader {

	f, err := os.Open("../week_5/log/access.log")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	re := regexp.MustCompile(`^.+?"([A-Z]+)\s.+?"\s(\d{3})`)

	var logreader common.LogReader

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		logreader.TotalRequests++
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		if len(matches) > 1 {
			switch matches[1] {
			case "GET":
				logreader.GET++
			case "POST":
				logreader.POST++
			case "PUT":
				logreader.PUT++
			case "DELETE":
				logreader.DELETE++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return &logreader

}

func countLogError() *common.LogReader {
	f, err := os.Open("../week_5/log/error.log")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	re := regexp.MustCompile(`\[(.*?)\]`)
	var logreader common.LogReader

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		if len(matches) > 1 {
			switch matches[1] {
			case "error":
				logreader.Error++
			case "alert":
				logreader.Alert++
			case "warn":
				logreader.Warn++
			case "crit":
				logreader.Crit++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return &logreader

}

func saveData(logR *common.LogReader) {
	file, err := os.OpenFile("../week_5/log/analyze.json", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data := fmt.Sprintf(`{
	"GET": %d,
	"POST": %d,
	"DELETE": %d,
	"UPDATE": %d,
	"error": %d,
	"alert": %d,
	"warn": %d,
	"crit": %d,
	"total_requests": %d
}`,logR.GET,
 	logR.POST,
 	logR.DELETE,
	logR.PUT,
 	logR.Error,
 	logR.Alert,
	logR.Warn,
	logR.Crit,
 	logR.TotalRequests)
	_, err = file.WriteString(data)
	if err != nil {
		panic(err)
	}
	fmt.Println("JSON data saved to analyze.json")
}
