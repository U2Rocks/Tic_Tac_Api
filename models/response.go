package models

// this struct is used to send messages from the server
type ResponseObject struct {
	StatusCode uint   `json:"statuscode"`
	Message    string `json:"message"`
}
