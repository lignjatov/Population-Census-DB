package model

type Population struct {
	OIB                    string `json:"oib"`
	Name                   string `json:"name"`
	Surname                string `json:"surname"`
	Occupation             string `json:"occupation"`
	Date_Of_Birth          string `json:"date_of_birth"`
	Height                 int    `json:"height"`
	Weight                 int    `json:"weight"`
	Blood_Type             string `json:"blood_type"`
	Additional_Information string `json:"additional_information"`
}
