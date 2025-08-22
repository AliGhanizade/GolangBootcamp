package week_2

type person struct {
	ID          int
	Name        string
	Age         int
	PhoneNumber string
	Email       string
}

type book struct {
	ID     uint
	Title  string
	Author person
	Year   int
}

func createPerson() []person {
	var Persons []person
	Persons = append(Persons, person{
		ID:          1,
		Name:        "ali",
		Age:         17,
		PhoneNumber: "+989052315520",
		Email:       "MraliGhanizade@gmail.com",
	})
	Persons = append(Persons, person{
		ID:          1,
		Name:        "Reza",
		Age:         25,
		PhoneNumber: "+989000000000",
		Email:       "MrReza@gmail.com",
	})
	return Persons

}

func createBook(Authors []person) []book {
	var Books []book
	Books = append(Books, book{
		ID:     1,
		Title:  "Book1",
		Author: Authors[0],
		Year:   2024,
	})
	Books = append(Books, book{
		ID:     2,
		Title:  "Book2",
		Author: Authors[1],
		Year:   2024,
	})
	return Books
}
