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
		Persons := CreatePerson()
		common.PrintByTime(Persons)

		Books := CreateBook(Persons)
		common.PrintByTime(Books)
	case 2:
		countries := CreateCountries()
		common.PrintMapStringJustKey(countries)
		capital := FindInMap(countries)
		common.PrintByTime(capital)
	case 3:
		username := GetInfo()
		err := CheckUniqUsername(username)
		if err != nil {
			// fmt.Println(err) 
			return
		}
		err = SaveInFile(username)
		if err != nil {
			// fmt.Println(err) 
			return
		}
		

	case 4:
		AllWeek2()
	}
}

func AllWeek2() {
	Persons := CreatePerson()
	common.PrintByTime(Persons)

	Books := CreateBook(Persons)
	common.PrintByTime(Books)

	countries := CreateCountries()
	// common.PrintMapString(Capital)
	capital := FindInMap(countries)
	common.PrintByTime(capital)

	username := GetInfo()
		err := CheckUniqUsername(username)
		if err != nil {
			// fmt.Println(err) 
			return
		}
		err = SaveInFile(username)
		if err != nil {
			// fmt.Println(err) 
			return
		}
	

}
