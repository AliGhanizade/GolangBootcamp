package week_2

import (
	common "GolangBootcamp/common"
	"fmt"
)

func SelectResult() {
	var practice int
	fmt.Print("Write Number of practice 1-struct 2-map 3-SaveInFile 0-All:")
	fmt.Scan(&practice)
	switch practice {

	case 1:
		Persons := createPerson()
		common.PrintByTime(Persons)

		Books := createBook(Persons)
		common.PrintByTime(Books)

	case 2:
		countries := createCountries()
		common.PrintMapStringJustKey(countries)
		capital := findInMap(countries)
		common.PrintByTime(capital)
		
	case 3:
		username := getInfo()
		err := checkUniqUsername(username)
		if err != nil {
			// fmt.Println(err) 
			return
		}
		err = saveInFile(username)
		if err != nil {
			// fmt.Println(err) 
			return
		}
	
	case 0:
		AllWeek2()
	}
}

func AllWeek2() {
	Persons := createPerson()
	common.PrintByTime(Persons)

	Books := createBook(Persons)
	common.PrintByTime(Books)

	countries := createCountries()
	// common.PrintMapString(Capital)
	capital := findInMap(countries)
	common.PrintByTime(capital)

	username := getInfo()
		err := checkUniqUsername(username)
		if err != nil {
			// fmt.Println(err) 
			return
		}
		err = saveInFile(username)
		if err != nil {
			// fmt.Println(err) 
			return
		}
	

}
