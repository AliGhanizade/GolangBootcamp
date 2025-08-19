package week_2

import (
	"fmt"
)

func CreateCountries() map[string]string {
	Countries := map[string]string{
		"France":       "Paris",
		"Germany":      "Berlin",
		"Japan":        "Tokyo",
		"Brazil":       "Brasilia",
		"Australia":    "Canberra",
		"Canada":       "Ottawa",
		"China":        "Beijing",
		"India":        "New Delhi",
		"Egypt":        "Cairo",
		"Nigeria":      "Abuja",
		"Mexico":       "Mexico City",
		"Russia":       "Moscow",
		"South Africa": "Pretoria",
		"South Korea":  "Seoul",
		"Italy":        "Rome",
		"Spain":        "Madrid",
		"Argentina":    "Buenos Aires",
		"Saudi Arabia": "Riyadh",
		// "ژاپن":   "توکیو",
		// "چین":   "پکن",
		// "فرانسه": "پاریس",
		// "ایتالیا": "رم",
		// "آلمان":  "برلین",
		// "اسپانیا": "مادرید",
		// "برزیل":  "برازیلیا",
		// "استرالیا": "کانبرا",
		// "هند":   "دهلی نو",
		// "کانادا":  "اتاوا",
		// "ایران" : "تهران",
	}
	Countries["Iran"] = "Tehran"

	return Countries
}

func FindInMap(Countries map[string]string) string {
	var country string
	fmt.Print("Write your Country :")
	fmt.Scan(&country)

	return Countries[country]
}
