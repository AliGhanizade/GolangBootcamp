package week_2

import common "GolangBootcamp/common"

func AllWeek2()  {
	Persons := CreatePerson()
	common.PrintByTime(Persons)
	Books := CreateBook(Persons)
	common.PrintByTime(Books)
}