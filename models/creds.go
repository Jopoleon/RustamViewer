package models

type CredStruct struct {
	FirstName   string `json:"firstName"`
	SecondName  string `json:"secondName"`
	CompanyName string `json:"companyName"`
	CompanyID   int    `json:"companyID"`
	Profilename string `json:"profilename"`
	Email       string `json:"email"` //login
	Login       string `json:"login"`
	Password    string `json:"password"`
}

func (c *CredStruct) Validate() error {
	return nil
}
