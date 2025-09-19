package common

type Person struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Age         int    `json:"age"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email,omitempty"`
}

type Username struct {
	DateTime    string `json:"date_time"`
	Username    string `json:"username"`
	PhoneNumber string `json:"phone_number"`
}

type LogReader struct {
	GET int `json:"GET"`
	POST int `json:"POST"`
	DELETE int `json:"DELETE"`
	PUT int `json:"UPDATE"`
	Error int `json:"error"`
	Alert int `json:"alert"`
	Warn int `json:"warn"`
	Crit int `json:"crit"`
	TotalRequests int `json:"total_requests"`
}	

type MessageWebSocket struct {
	Message  string `json:"message"`
	Username string `json:"username"`
	DateTime string `json:"date_time"`
}