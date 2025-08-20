package week_2

type Person struct {
	ID          int
	Name        string
	Age         int
	PhoneNumber string
	Email       string
}

type Book struct {
	ID     uint
	Title  string
	Author Person
	Year   int
}

func CreatePerson() []Person {
	var Persons []Person
	Persons = append(Persons, Person{
		ID:          1,
		Name:        "ali",
		Age:         17,
		PhoneNumber: "+989052315520",
		Email:       "MraliGhanizade@gmail.com",
	})
	Persons = append(Persons, Person{
		ID:          1,
		Name:        "Reza",
		Age:         25,
		PhoneNumber: "+989000000000",
		Email:       "MrReza@gmail.com",
	})
	return Persons

}

func CreateBook(Authors []Person) []Book {
	var Books []Book
	Books = append(Books, Book{
		ID:     1,
		Title:  "Book1",
		Author: Authors[0],
		Year:   2024,
	})
	Books = append(Books, Book{
		ID:     2,
		Title:  "Book2",
		Author: Authors[1],
		Year:   2024,
	})
	return Books
}
