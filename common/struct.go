package common

type Person struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Age         int    `json:"age"`
	PhoneNumber string `json:"phone_number"`
	Email       string `jsom:"email,omitempty"`
}

type Username struct {
	DateTime    string `json:"date_time"`
	Username    string `json:"username"`
	PhoneNumber string `json:"phone_number"`
}
