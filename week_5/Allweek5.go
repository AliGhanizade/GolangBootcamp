package week_5

import (
	"GolangBootcamp/common"
	"fmt"
)

func SelectResult()  {
	var practice int
	msg := `Write Number of practice
1-Run server on http://127.0.0.1:8080
	1-1 on route /struct -> convert struct to json	
	1-2 on route /csv -> convert csv to json
2-print csv file
3-convert csv to json and save in file data.json
4-log reader and count total request and successful request (status 200)
`
	fmt.Print(msg)
	fmt.Scan(&practice)
	switch practice {
	case 1:
		Run()
	case 2:
		readCSV()
	case 3:
		jsonData := csvToJson()
		saveInJsonFile(jsonData)
	case 4:
		countLogAccess()
		countLogError()
		logA := countLogAccess()
		logE := countLogError()
		saveData(&common.LogReader{
			TotalRequests: logA.TotalRequests,
			GET:           logA.GET,
			POST:          logA.POST,
			PUT:           logA.PUT,
			DELETE:        logA.DELETE,
			Error:         logE.Error,
			Alert:         logE.Alert,
			Warn:          logE.Warn,
			Crit:          logE.Crit,
		})
	default:
		fmt.Println("Invalid practice number. Please enter a valid number.")
	}
}