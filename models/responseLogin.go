package models

/* ResponseLogin ave a token that return with the login */
type ResponseLogin struct {
	Token 	string 	`json:"token,omitempty"`
}
